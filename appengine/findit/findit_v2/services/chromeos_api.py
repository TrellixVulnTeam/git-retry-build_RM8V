# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
"""Defines the chromeos-specific APIs required by Findit."""

from collections import defaultdict
import logging

from google.protobuf import json_format

from findit_v2.model.compile_failure import CompileFailureGroup
from findit_v2.services.failure_type import StepTypeEnum
from findit_v2.services.project_api import ProjectAPI

_COMPILE_FAILURE_OUTPUT_NAME = 'compile_failure'
_TEST_FAILURE_OUTPUT_NAME = 'test_failure'


class ChromeOSProjectAPI(ProjectAPI):

  def _GetFailureOutput(self, build, output_name):
    """Gets failure output from the build's output property."""
    # Converts the Struct to standard dict, to use .get, .iteritems etc.
    build_failure_output = json_format.MessageToDict(
        build.output.properties).get(output_name)

    return build_failure_output

  def ClassifyStepType(self, build, step):
    """Returns the failure type of the given build step.

    In ChromeOS builds,
    - if they have compile failures, they will produce an
      output property called 'compile_failure', which includes the
      failed step name. So that step will be classified as the compile step.
    - if they have test failures, they will produce an
      output property called 'test_failure', which includes the
      failed step name. So that step will be classified as the test step.

    Args:
      build (buildbucket build.proto): ALL info about the build.
      step (buildbucket step.proto): ALL info about the build step.
    """

    def classify_compile_step(compile_failure_output):
      # Format of compile_failure_output:
      # {
      #   'failures': [{
      #     'output_targets': ['target'],
      #     'rule': 'emerge',
      #   }, ],
      #   'failed_step': 'step'
      # }
      failed_compile_step = compile_failure_output.get('failed_step')
      if not failed_compile_step:
        logging.error(
            'No failed_step in compile_failure property of ChromeOS'
            ' build %d.', build.id)
        return StepTypeEnum.INFRA

      if step.name == failed_compile_step:
        # Noted for ChromeOS the current supported compile step is nested.
        # To be consistent with sheriff-o-matic, the matching step name is a
        # leaf step. Although in reality the parent step also has 'FAILIURE'
        # state and is a compile step, Findit still returns a StepTypeEnum.INFRA
        # to intentionality ignore it.
        return StepTypeEnum.COMPILE
      return StepTypeEnum.INFRA

    def classify_test_step(test_failure_output):
      # Format of test_failure_output:
      # {
      #   'xx_test_failures': [  # failure type
      #     {
      #       'failed_step': 'step',
      #       'test_spec': 'test_spec'
      #     }
      #   ]
      # }
      for failures in test_failure_output.itervalues():
        for failure in failures:
          if step.name == failure.get('failed_step'):
            return StepTypeEnum.TEST
      return StepTypeEnum.INFRA

    compile_failure_output = self._GetFailureOutput(
        build, _COMPILE_FAILURE_OUTPUT_NAME)
    if compile_failure_output:
      return classify_compile_step(compile_failure_output)

    test_failure_output = self._GetFailureOutput(build,
                                                 _TEST_FAILURE_OUTPUT_NAME)
    if test_failure_output:
      return classify_test_step(test_failure_output)

    # No compile/test failure output, classifies step as INFRA failure.
    return StepTypeEnum.INFRA

  def GetCompileFailures(self, build, compile_steps):
    """Returns the detailed compile failures from a failed build.

    For ChromeOS builds, the failures are stored in the build's output
    property 'compile_failure'.

    Args:
      build (buildbucket build.proto): ALL info about the build.
      compile_steps (list of buildbucket step.proto): The failed compile steps.

    Returns:
      (dict): Information about detailed compile failures.
      {
        'step_name': {
          'failures': {
            frozenset(['target1', 'target2']): {
              'first_failed_build': {
                'id': 8765432109,
                'number': 123,
                'commit_id': 654321
              },
              'last_passed_build': None,
              'properties': {
                # Arbitrary information about the failure if exists.
              }
            },
          'first_failed_build': {
            'id': 8765432109,
            'number': 123,
            'commit_id': 654321
          },
          'last_passed_build': None,
          'properties': {
            # Arbitrary information about the failure if exists.
          }
        },
      }
    """
    # pylint: disable=unused-argument
    build_info = {
        'id': build.id,
        'number': build.number,
        'commit_id': build.input.gitiles_commit.id
    }

    build_compile_failure_output = self._GetFailureOutput(
        build, _COMPILE_FAILURE_OUTPUT_NAME)

    if not build_compile_failure_output:
      logging.error('No %s for ChromeOS build %d.',
                    _COMPILE_FAILURE_OUTPUT_NAME, build.id)
      return {}

    failed_step = build_compile_failure_output.get('failed_step')
    if not failed_step:
      logging.error(
          'No failed_step in compile_failure property of ChromeOS'
          ' build %d.', build.id)
      return {}

    detailed_compile_failures = {
        failed_step: {
            'failures': {},
            'first_failed_build': build_info,
            'last_passed_build': None,
        }
    }
    failures_dict = detailed_compile_failures[failed_step]['failures']
    for failure in build_compile_failure_output.get('failures', []):
      # In ChromeOS build, output_target is a json string like
      # "{\"category\": \"chromeos-base\", \"packageName\": \"cryptohome\"}"
      output_targets = frozenset(failure['output_targets'])
      failures_dict[output_targets] = {
          'properties': {
              'rule': failure.get('rule')
          },
          'first_failed_build': build_info,
          'last_passed_build': None,
      }

    return detailed_compile_failures

  def GetTestFailures(self, build, test_steps):
    """Returns the detailed test failures from a failed build.

    For ChromeOS builds, the failures are stored in the build's output
    property 'build_test_failure_output'.

    Args:
      build (buildbucket build.proto): ALL info about the build.
      test_steps (list of buildbucket step.proto): The failed test steps.

    Returns:
      (dict): Information about detailed test failures.
      {
        'step_name1': {
          'failures': {},
          'first_failed_build': {
            'id': 8765432109,
            'number': 123,
            'commit_id': 654321
          },
          'last_passed_build': None,
          'properties': {
            'failure_type': 'xx_test_failures',
            'test_spec': 'test_spec'
          }
        },
        ...
      }
    """
    # pylint: disable=unused-argument
    build_info = {
        'id': build.id,
        'number': build.number,
        'commit_id': build.input.gitiles_commit.id
    }

    # The format of test_failure property is like:
    # 'test_failure': {
    #   'xx_test_failures': [
    #     {
    #       'failed_step': 'results|xx test results|[FAILED] <suite1>',
    #       'test_spec': 'json serialized proto for suite1'
    #     }
    #   ],
    #   'yy_test_failures': [
    #     {
    #       'failed_step': 'results|yy test results|[FAILED] <suite2>',
    #       'test_spec': 'json serialized proto for suite2'
    #     }
    #   ],
    # }
    build_test_failure_output = self._GetFailureOutput(
        build, _TEST_FAILURE_OUTPUT_NAME)

    if not build_test_failure_output:
      logging.error('No %s for ChromeOS build %d.', _TEST_FAILURE_OUTPUT_NAME,
                    build.id)
      return {}

    detailed_test_failures = {}
    for failure_type, failures in build_test_failure_output.iteritems():
      for failure in failures:
        failed_step = failure.get('failed_step')
        test_spec = failure.get('test_spec')
        if not failed_step or not test_spec:
          logging.error(
              'Malformed %s for ChromeOs build %d - failure_type: %s,'
              ' failure_info: %r.', _TEST_FAILURE_OUTPUT_NAME, build.id,
              failure_type, failure)
          continue

        detailed_test_failures[failed_step] = {
            'failures': {},
            'first_failed_build': build_info,
            'last_passed_build': None,
            'properties': {
                'failure_type': failure_type,
                'test_spec': test_spec
            }
        }

    return detailed_test_failures

  def GetFailuresWithMatchingCompileFailureGroups(
      self, context, build, first_failures_in_current_build):
    """Gets reusable failure groups for given compile failure(s).

    Each failure in detailed_compile_failures will be updated with the failure
    group it belongs to if an existing failure group is found for it.

    Criteria for a matching group:
      + same project
      + group contains exactly the same failed targets
      + same regression range

    Here are some special cases:
    1. compile result is unknown if a build ends with infra_failure. For Findit,
      the regression range is from the commit/build a compile target actually
      passed, to the commit/build a compile target actually failed. So it's
      possible for build(s) in between with infra_failure status.
      a. If 2 builds(with gitiles_commit 100 and 110 respectively) in a row on
        builder A failed at compile. And only 1 build(with gitiles_commit 110)
        on builder B failed at compile, though the build failed with commit 100
        ended with some infra_failures. In current criteria Findit will group
        those failures together. No matter builds on which builder are analyzed
        first.
      b. But if on builder A only build with gitiles_commit 110 failed at
        compile and build with gitiles_commit 100 passed (same build result for
        builder B). In current criteria Findit will NOT group
        those failures together. No matter builds on which builder are analyzed
        first.
    2. If build C failed to compile target 1 and target 2 and build D failed to
      compile target 1 only, these 2 builds will not be grouped.

    Args:
      context (findit_v2.services.context.Context): Scope of the analysis.
      build (buildbucket build.proto): ALL info about the build.
      first_failures_in_current_build (dict): A dict for failures that happened
      the first time in current build.
      {
        'failures': {
          'install packages': {
            'atomic_failures': [
              frozenset(['target4']),
              frozenset(['target1', 'target2'])],
            'last_passed_build': {
              'id': 8765432109,
              'number': 122,
              'commit_id': 'git_sha1'
            },
          },
        },
        'last_passed_build': {
          'id': 8765432109,
          'number': 122,
          'commit_id': 'git_sha1'
        }
      }

    Returns:
      failures_with_existing_group (dict): A dict of failures from
        first_failures_in_current_build that found a matching group.
        {
          'install packages': {
            frozenset(['target4']):  8765432000,
          ]
        }
    """
    groups = CompileFailureGroup.query(
        CompileFailureGroup.luci_project == build.builder.project).filter(
            CompileFailureGroup.first_failed_commit.gitiles_id == build.input
            .gitiles_commit.id).fetch()

    failures_with_existing_group = defaultdict(dict)

    # Looks for existing groups to reuse.
    for group in groups:
      group_last_passed_commit = group.last_passed_commit

      if (context.gitiles_host != group_last_passed_commit.gitiles_host or
          context.gitiles_project != group_last_passed_commit.gitiles_project or
          context.gitiles_ref !=
          group_last_passed_commit.gitiles_ref):  # pragma: no cover
        logging.debug(
            'Group %d and build %d have commits from different repo'
            ' or branch.', group.key.id(), build.id)
        continue

      failures_in_group = group.failed_targets
      for step_ui_name, step_failure in first_failures_in_current_build[
          'failures'].iteritems():
        if (step_ui_name not in failures_in_group or
            step_failure['last_passed_build']['commit_id'] !=
            group_last_passed_commit.gitiles_id):
          # The group doesn't have failures in the step or the group has a
          # different regression range.
          continue

        failed_targets_in_current_build = frozenset.union(
            *step_failure['atomic_failures'])
        if not failed_targets_in_current_build == set(
            failures_in_group[step_ui_name]):
          continue
        # Matching failure found in the group. Should reuse this group.
        for output_target_frozenset in step_failure['atomic_failures']:
          failures_with_existing_group[step_ui_name][
              output_target_frozenset] = group.key.id()

    return failures_with_existing_group

  def GetRerunBuilderId(self, build):
    """Gets builder id to run the rerun builds.

    Args:
      build (buildbucket build.proto): ALL info about the build.

    Returns:
      (str): Builder id in the format luci_project/luci_bucket/luci_builder
    """
    rerun_builder = json_format.MessageToDict(
        build.output.properties).get('BISECT_BUILDER')

    assert rerun_builder, 'Failed to find rerun builder for build {}'.format(
        build.id)

    return '{project}/{bucket}/{builder}'.format(
        project=build.builder.project,
        bucket=build.builder.bucket,
        builder=rerun_builder)

  def GetCompileRerunBuildInputProperties(self, failed_targets):
    """Gets input properties of a rerun build for compile failures.

    Args:
      failed_targets (dict): Targets Findit wants to rerun in the build.

    Returns:
      (dict): input properties of the rerun build."""
    targets = []
    for step_targets in failed_targets.itervalues():
      targets.extend(step_targets)
    if not targets:
      return None

    return {
        '$chromeos/cros_bisect': {
            'targets': list(set(targets))
        },
    }
