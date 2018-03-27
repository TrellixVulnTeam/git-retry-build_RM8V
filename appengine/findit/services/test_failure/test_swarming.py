# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging

from google.appengine.ext import ndb

from common.findit_http_client import FinditHttpClient
from dto import swarming_task_error
from dto.swarming_task_error import SwarmingTaskError
from infra_api_clients.swarming import swarming_util
from libs import analysis_status
from model.wf_swarming_task import WfSwarmingTask
from services import constants
from services import monitoring
from services import swarmed_test_util
from services import swarming
from services import test_results
from waterfall import waterfall_config


def NeedANewSwarmingTask(master_name, builder_name, build_number, step_name,
                         force):
  """Checks if a WfSwarmingTask for the given params exists, or creates it."""
  swarming_task = WfSwarmingTask.Get(master_name, builder_name, build_number,
                                     step_name)

  if not swarming_task:
    swarming_task = WfSwarmingTask.Create(master_name, builder_name,
                                          build_number, step_name)
    swarming_task.status = analysis_status.PENDING
    swarming_task.put()
    return True, swarming_task.key.urlsafe()

  if force:
    swarming_task.Reset()
    swarming_task.put()
    return True, swarming_task.key.urlsafe()

  # TODO(http://crbug.com/585676): Rerun the Swarming task if it runs into
  # unexpected infra errors.
  return False, swarming_task.key.urlsafe()


def CreateNewSwarmingTaskRequest(runner_id, ref_task_id, ref_request,
                                 master_name, builder_name, build_number,
                                 step_name, tests, iterations):
  new_request = swarming.CreateNewSwarmingTaskRequestTemplate(
      runner_id,
      ref_task_id,
      ref_request,
      master_name,
      builder_name,
      step_name,
      tests,
      iterations,
      use_new_pubsub=True)

  # Add additional tags.
  new_request.tags.append('ref_task_id:%s' % ref_task_id)
  new_request.tags.append('purpose:identify-flake')
  new_request.tags.append('ref_buildnumber:%s' % build_number)
  return new_request


def TriggerSwarmingTask(run_swarming_task_parameters, runner_id):
  """Triggers a swarming rerun for the given tests in a given build."""

  master_name, builder_name, build_number = (
      run_swarming_task_parameters.build_key.GetParts())
  step_name = run_swarming_task_parameters.step_name
  tests = run_swarming_task_parameters.tests

  assert tests, 'No tests to trigger swarming task for.'
  http_client = FinditHttpClient()

  # 1. Retrieve Swarming task parameters from a given Swarming task id.
  ref_task_id, ref_request = swarming.GetReferredSwarmingTaskRequestInfo(
      master_name, builder_name, build_number, step_name, http_client)

  # 2. Update/Overwrite parameters for the re-run.
  iterations_to_rerun = waterfall_config.GetSwarmingSettings().get(
      'iterations_to_rerun')
  new_request = CreateNewSwarmingTaskRequest(
      runner_id, ref_task_id, ref_request, master_name, builder_name,
      build_number, step_name, tests, iterations_to_rerun)

  # 3. Trigger a new Swarming task to re-run the failed tests.
  task_id, _ = swarming_util.TriggerSwarmingTask(swarming.SwarmingHost(),
                                                 new_request, http_client)

  if task_id:  # pragma: no branch.
    # 4. Update swarming task.
    OnSwarmingTaskTriggered(master_name, builder_name, build_number, step_name,
                            tests, task_id, iterations_to_rerun, new_request)

  return task_id


@ndb.transactional
def _UpdateSwarmingTaskEntity(master_name,
                              builder_name,
                              build_number,
                              step_name,
                              status=None,
                              task_id=None,
                              error=None,
                              tests_statuses=None,
                              parameters=None,
                              canonical_step_name=None):
  task = WfSwarmingTask.Get(master_name, builder_name, build_number, step_name)
  assert task
  task.status = status or task.status
  task.task_id = task_id or task.task_id
  task.error = error.ToSerializable() if error else task.error
  task.tests_statuses = tests_statuses or task.tests_statuses
  task.parameters = task.parameters or {}
  task.parameters.update(parameters or {})
  task.canonical_step_name = canonical_step_name or task.canonical_step_name
  task.put()


def OnSwarmingTaskTriggered(master_name, builder_name, build_number, step_name,
                            tests, task_id, iterations_to_rerun, new_request):
  canonical_step_name = swarming_util.GetTagValue(new_request.tags, 'ref_name')
  parameters = {
      'tests': tests,
      'iterations_to_rerun': iterations_to_rerun,
      'ref_name': canonical_step_name,
      'priority': new_request.priority
  }
  _UpdateSwarmingTaskEntity(
      master_name,
      builder_name,
      build_number,
      step_name,
      task_id=task_id,
      parameters=parameters,
      canonical_step_name=canonical_step_name)
  monitoring.OnSwarmingTaskStatusChange('trigger', 'identify-flake')


def OnSwarmingTaskTimeout(run_swarming_task_params, task_id):
  master_name, builder_name, build_number = (
      run_swarming_task_params.build_key.GetParts())
  step_name = run_swarming_task_params.step_name

  error = SwarmingTaskError.GenerateError(swarming_task_error.RUNNER_TIMEOUT)

  _state, output_json, _error = swarmed_test_util.GetSwarmingTaskData(task_id)
  if output_json:
    tests_statuses = test_results.GetTestsRunStatuses(output_json)
    _UpdateSwarmingTaskEntity(
        master_name,
        builder_name,
        build_number,
        step_name,
        status=analysis_status.COMPLETED,
        error=error,
        tests_statuses=tests_statuses)
  else:
    _UpdateSwarmingTaskEntity(
        master_name,
        builder_name,
        build_number,
        step_name,
        status=analysis_status.ERROR,
        error=error)


def OnSwarmingTaskError(master_name,
                        builder_name,
                        build_number,
                        step_name,
                        error,
                        should_complete_pipeline=True):
  logging.error('Error %s when processing a swarming task %s/%s/%d/%s',
                error.message, master_name, builder_name, build_number,
                step_name)

  if should_complete_pipeline:
    _UpdateSwarmingTaskEntity(
        master_name,
        builder_name,
        build_number,
        step_name,
        status=analysis_status.ERROR,
        error=error)
    return False
  else:
    _UpdateSwarmingTaskEntity(
        master_name, builder_name, build_number, step_name, error=error)
    return


def OnSwarmingTaskCompleted(master_name, builder_name, build_number, step_name,
                            output_json):
  tests_statuses = test_results.GetTestsRunStatuses(output_json)
  _UpdateSwarmingTaskEntity(
      master_name,
      builder_name,
      build_number,
      step_name,
      status=analysis_status.COMPLETED,
      tests_statuses=tests_statuses)
  return True


def OnSwarmingTaskStateChanged(run_swarming_task_parameters, task_id):
  master_name, builder_name, build_number = (
      run_swarming_task_parameters.build_key.GetParts())
  step_name = run_swarming_task_parameters.step_name

  task_state, output_json, error = swarmed_test_util.GetSwarmingTaskData(
      task_id)
  if not task_state:
    # Error when get task state.
    OnSwarmingTaskError(master_name, builder_name, build_number, step_name,
                        error, False)
    return None
  elif task_state == constants.STATE_COMPLETED and output_json:
    return OnSwarmingTaskCompleted(master_name, builder_name, build_number,
                                   step_name, output_json)
  elif task_state in constants.STATE_NOT_STOP:
    if task_state == constants.STATE_RUNNING:  # pragma: no branch
      _UpdateSwarmingTaskEntity(
          master_name,
          builder_name,
          build_number,
          step_name,
          status=analysis_status.RUNNING)
    return None
  else:
    # Swarming task finished with error.
    return OnSwarmingTaskError(master_name, builder_name, build_number,
                               step_name, error)
