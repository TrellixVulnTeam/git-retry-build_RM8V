# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import copy
import logging
import time

from common import time_util
from common.http_client_appengine import HttpClientAppengine as HttpClient
from common.pipeline_wrapper import BasePipeline
from model import analysis_status
from waterfall import swarming_util
from waterfall import waterfall_config


class TriggerBaseSwarmingTaskPipeline(BasePipeline):  # pragma: no cover.
  """A pipeline to trigger a Swarming task to re-run selected tests of a step.

  This pipeline only supports test steps that run on Swarming and support the
  gtest filter.
  """

  def _GetSwarmingTaskName(self, ref_task_id):  # pragma: no cover.
    return 'findit/deflake/ref_task_id/%s/%s' % (
        ref_task_id, time_util.GetUTCNow().strftime('%Y-%m-%d %H:%M:%S %f'))

  def _CreateNewSwarmingTaskRequest(self, ref_task_id, ref_request, master_name,
                                    builder_name, build_number, step_name,
                                    tests, iterations):
    """Returns a SwarmingTaskRequest instance to run the given tests only."""

    # Make a copy of the referred request and drop or overwrite some fields.
    new_request = copy.deepcopy(ref_request)
    new_request.name = self._GetSwarmingTaskName(ref_task_id)
    new_request.parent_task_id = ''
    new_request.user = ''

    # To force a fresh re-run and ignore cached result of any equivalent run.
    new_request.idempotent = False

    # Set the gtest_filter to run the given tests only.
    new_request.extra_args.append('--gtest_repeat=%s' % iterations)
    new_request.extra_args.append('--test-launcher-retry-limit=0')
    new_request.extra_args = [
        a for a in new_request.extra_args if not a.startswith('--gtest_filter')
    ]
    new_request.extra_args.append('--gtest_filter=%s' % ':'.join(tests))

    # Remove the env setting for sharding.
    sharding_settings = ['GTEST_SHARD_INDEX', 'GTEST_TOTAL_SHARDS']
    new_request.env = [
        e for e in new_request.env if e['key'] not in sharding_settings
    ]

    # Reset tags for searching and monitoring.
    ref_name = swarming_util.GetTagValue(ref_request.tags, 'name')
    new_request.tags = []
    new_request.tags.append('purpose:deflake')
    new_request.tags.append('ref_master:%s' % master_name)
    new_request.tags.append('ref_buildername:%s' % builder_name)
    new_request.tags.append('ref_buildnumber:%s' % build_number)
    new_request.tags.append('ref_stepname:%s' % step_name)
    new_request.tags.append('ref_task_id:%s' % ref_task_id)
    new_request.tags.append('ref_name:%s' % ref_name)

    return new_request

  def _GetArgs(self, master_name, builder_name, build_number, step_name, tests):
    # Returns an array you can pass into _GetSwarmingTask, _CreateSwarmingTask,
    # _NeedANewSwarmingTask as the arguments.

    # Should be overwritten in child method.
    raise NotImplementedError(
        '_GetArgs should be implemented in child class')

  def _GetSwarmingTask(self):
    # Get the appropriate kind of Swarming Task (Wf or Flake).

    # Should be overwritten in child method.
    raise NotImplementedError(
        '_GetSwarmingTask should be implemented in child class')

  def _CreateSwarmingTask(self):
    # Create the appropriate kind of Swarming Task (Wf or Flake)

    # Should be overwritten in child method.
    raise NotImplementedError(
        '_CreateSwarmingTask should be implemented in child class')

  def _NeedANewSwarmingTask(self, *args):
    swarming_task = self._GetSwarmingTask(*args)
    if not swarming_task:
      swarming_task = self._CreateSwarmingTask(*args)
      swarming_task.status = analysis_status.PENDING
      swarming_task.put()
      return True
    else:
      # TODO(http://crbug.com/585676): Rerun the Swarming task if it runs into
      # unexpected infra errors.
      return False

  def _GetSwarmingTaskId(self, *args):
    swarming_settings = waterfall_config.GetSwarmingSettings()
    wait_seconds = swarming_settings.get('get_swarming_task_id_wait_seconds')
    timeout_seconds = swarming_settings.get(
        'get_swarming_task_id_timeout_seconds')
    deadline = time.time() + timeout_seconds

    while time.time() < deadline:
      swarming_task = self._GetSwarmingTask(*args)

      if not swarming_task:  # pragma: no cover. Pipeline will retry.
        raise Exception('Swarming task was deleted unexpectedly!')

      if swarming_task.task_id:
        return swarming_task.task_id

      # Wait for the existing pipeline to start the Swarming task.
      time.sleep(wait_seconds)

    raise Exception('Time out!')  # pragma: no cover. Pipeline will retry.

  def _GetIterationsToRerun(self):
    # How many times we want to run the swarming rerun
    # By default, it's what's in wf_config
    raise NotImplementedError(
        '_GetIterationsToRerun should be implemented in child class')

  # Arguments number differs from overridden method - pylint: disable=W0221
  def run(self, master_name, builder_name, build_number, step_name, tests):
    """Triggers a new Swarming task to run the given tests.

    Args:
      master_name (str): The master name.
      builder_name (str): The builder name.
      build_number (str): The build number.
      step_name (str): The failed test step name.
      tests (list): A list of test cases, eg: ['suite1.test1', 'suite2.testw2']

    Returns:
      task_id (str): The new Swarming task that re-run the given tests.
    """
    call_args = self._GetArgs(master_name, builder_name,
                              build_number, step_name, tests)
    # Check if a new Swarming Task is really needed.
    if not self._NeedANewSwarmingTask(*call_args):
      return self._GetSwarmingTaskId(*call_args)
    assert tests
    http_client = HttpClient()

    # 0. Retrieve existing Swarming task ids for the given step.
    swarming_task_items = swarming_util.ListSwarmingTasksDataByTags(
        master_name, builder_name, build_number, http_client, step_name)
    assert len(swarming_task_items) > 0, 'No Swarming task was run.'
    ref_task_id = swarming_task_items[0]['task_id']

    # 1. Retrieve Swarming task parameters from a given Swarming task id.
    ref_request = swarming_util.GetSwarmingTaskRequest(
        ref_task_id, http_client)

    # 2. Update/Overwrite parameters for the re-run.
    iterations_to_rerun = self._GetIterationsToRerun()

    new_request = self._CreateNewSwarmingTaskRequest(
        ref_task_id, ref_request, master_name, builder_name, build_number,
        step_name, tests, iterations_to_rerun)

    # 3. Trigger a new Swarming task to re-run the failed tests.
    task_id = swarming_util.TriggerSwarmingTask(new_request, http_client)

    # Save the task id.
    swarming_task = self._GetSwarmingTask(*call_args)
    swarming_task.task_id = task_id
    swarming_task.parameters['tests'] = tests
    swarming_task.parameters['iterations_to_rerun'] = iterations_to_rerun
    swarming_task.parameters['ref_name'] = swarming_util.GetTagValue(
        new_request.tags, 'ref_name')
    swarming_task.put()

    logging.info('A Swarming task was triggered:%s', task_id)
    return task_id
