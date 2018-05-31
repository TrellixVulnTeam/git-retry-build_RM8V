# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import mock

from dto.int_range import IntRange
from gae_libs import pipelines
from gae_libs.pipeline_wrapper import pipeline_handlers
from model.flake.flake_culprit import FlakeCulprit
from model.flake.master_flake_analysis import DataPoint
from model.flake.master_flake_analysis import MasterFlakeAnalysis
from pipelines.flake_failure.next_commit_position_pipeline import (
    NextCommitPositionInput)
from pipelines.flake_failure.next_commit_position_pipeline import (
    NextCommitPositionPipeline)
from services import step_util
from services.flake_failure import heuristic_analysis
from services.flake_failure import lookback_algorithm
from services.flake_failure import next_commit_position_utils
from waterfall.build_info import BuildInfo
from waterfall.test.wf_testcase import WaterfallTestCase


class NextCommitPositionPipelineTest(WaterfallTestCase):
  app_module = pipeline_handlers._APP

  @mock.patch.object(lookback_algorithm, 'GetNextCommitPosition')
  def testNextCommitPositionPipelineFoundCulprit(self, mock_next_commit):
    master_name = 'm'
    builder_name = 'b'
    build_number = 100
    step_name = 's'
    test_name = 't'
    start_commit_position = 1000

    calculated_next_commit_position = None
    culprit_commit_position = 1000
    mock_next_commit.return_value = (calculated_next_commit_position,
                                     culprit_commit_position)

    analysis = MasterFlakeAnalysis.Create(master_name, builder_name,
                                          build_number, step_name, test_name)
    analysis.Save()

    next_commit_position_input = NextCommitPositionInput(
        analysis_urlsafe_key=analysis.key.urlsafe(),
        commit_position_range=IntRange(lower=None, upper=start_commit_position))

    pipeline_job = NextCommitPositionPipeline(next_commit_position_input)
    pipeline_job.start()
    self.execute_queued_tasks()

    pipeline_job = pipelines.pipeline.Pipeline.from_id(pipeline_job.pipeline_id)
    next_commit_position_output = pipeline_job.outputs.default.value

    self.assertFalse(pipeline_job.was_aborted)
    self.assertEqual(culprit_commit_position,
                     next_commit_position_output['culprit_commit_position'])
    self.assertIsNone(next_commit_position_output['next_commit_position'])

  @mock.patch.object(lookback_algorithm, 'GetNextCommitPosition')
  def testNextCommitPositionPipelineNotReproducible(self, mock_next_commit):
    master_name = 'm'
    builder_name = 'b'
    build_number = 100
    step_name = 's'
    test_name = 't'
    start_commit_position = 1000

    mock_next_commit.return_value = (None, None)

    analysis = MasterFlakeAnalysis.Create(master_name, builder_name,
                                          build_number, step_name, test_name)
    analysis.Save()

    next_commit_position_input = NextCommitPositionInput(
        analysis_urlsafe_key=analysis.key.urlsafe(),
        commit_position_range=IntRange(lower=None, upper=start_commit_position))

    pipeline_job = NextCommitPositionPipeline(next_commit_position_input)
    pipeline_job.start()
    self.execute_queued_tasks()

    pipeline_job = pipelines.pipeline.Pipeline.from_id(pipeline_job.pipeline_id)
    next_commit_position_output = pipeline_job.outputs.default.value

    self.assertFalse(pipeline_job.was_aborted)
    self.assertIsNone(next_commit_position_output['culprit_commit_position'])
    self.assertIsNone(next_commit_position_output['next_commit_position'])

  @mock.patch.object(lookback_algorithm, 'GetNextCommitPosition')
  @mock.patch.object(MasterFlakeAnalysis, 'CanRunHeuristicAnalysis')
  @mock.patch.object(next_commit_position_utils,
                     'GetNextCommitPositionFromHeuristicResults')
  def testNextCommitPositionPipelineWithHeuristicResults(
      self, mock_heuristic_result, mock_run_heuristic, mock_next_commit):
    master_name = 'm'
    builder_name = 'b'
    build_number = 105
    step_name = 's'
    test_name = 't'
    start_commit_position = 1000
    suspect_commit_position = 95
    expected_next_commit_position = 94

    suspect = FlakeCulprit.Create('repo', 'revision', suspect_commit_position)
    suspect.commit_position = suspect_commit_position
    suspect.put()

    analysis = MasterFlakeAnalysis.Create(master_name, builder_name,
                                          build_number, step_name, test_name)
    analysis.suspect_urlsafe_keys.append(suspect.key.urlsafe())
    analysis.put()

    mock_run_heuristic.return_value = False
    mock_heuristic_result.return_value = expected_next_commit_position

    calculated_next_commit_position = 999
    culprit_commit_position = None
    mock_next_commit.return_value = (calculated_next_commit_position,
                                     culprit_commit_position)

    next_commit_position_input = NextCommitPositionInput(
        analysis_urlsafe_key=analysis.key.urlsafe(),
        commit_position_range=IntRange(lower=None, upper=start_commit_position))

    pipeline_job = NextCommitPositionPipeline(next_commit_position_input)
    pipeline_job.start()
    self.execute_queued_tasks()

    pipeline_job = pipelines.pipeline.Pipeline.from_id(pipeline_job.pipeline_id)
    next_commit_position_output = pipeline_job.outputs.default.value

    self.assertFalse(pipeline_job.was_aborted)
    self.assertIsNone(next_commit_position_output['culprit_commit_position'])
    self.assertEqual(expected_next_commit_position,
                     next_commit_position_output['next_commit_position'])
    mock_heuristic_result.assert_called_once_with(analysis.key.urlsafe())

  @mock.patch.object(lookback_algorithm, 'GetNextCommitPosition')
  @mock.patch.object(MasterFlakeAnalysis, 'CanRunHeuristicAnalysis')
  @mock.patch.object(next_commit_position_utils,
                     'GetNextCommitPositionFromHeuristicResults')
  @mock.patch.object(step_util, 'GetValidBoundingBuildsForStep')
  @mock.patch.object(heuristic_analysis, 'RunHeuristicAnalysis')
  def testNextCommitPositionPipelineRunHeuristicResults(
      self, _, mock_bounding_builds, mock_heuristic_result,
      mock_can_run_heuristic, mock_next_commit):
    master_name = 'm'
    builder_name = 'b'
    build_number = 105
    step_name = 's'
    test_name = 't'
    start_commit_position = 1000
    expected_next_commit_position = 94

    analysis = MasterFlakeAnalysis.Create(master_name, builder_name,
                                          build_number, step_name, test_name)
    analysis.put()

    mock_can_run_heuristic.return_value = True
    mock_heuristic_result.side_effect = [None, expected_next_commit_position]

    calculated_next_commit_position = 999
    culprit_commit_position = None
    mock_next_commit.return_value = (calculated_next_commit_position,
                                     culprit_commit_position)

    lower_bound_build = BuildInfo(master_name, builder_name, build_number - 1)
    lower_bound_build.commit_position = expected_next_commit_position
    upper_bound_build = BuildInfo(master_name, builder_name, build_number)
    upper_bound_build.commit_position = start_commit_position
    mock_bounding_builds.return_value = (lower_bound_build, upper_bound_build)

    next_commit_position_input = NextCommitPositionInput(
        analysis_urlsafe_key=analysis.key.urlsafe(),
        commit_position_range=IntRange(lower=None, upper=start_commit_position))

    pipeline_job = NextCommitPositionPipeline(next_commit_position_input)
    pipeline_job.start()
    self.execute_queued_tasks()

    pipeline_job = pipelines.pipeline.Pipeline.from_id(pipeline_job.pipeline_id)
    next_commit_position_output = pipeline_job.outputs.default.value

    self.assertFalse(pipeline_job.was_aborted)
    self.assertIsNone(next_commit_position_output['culprit_commit_position'])
    self.assertEqual(expected_next_commit_position,
                     next_commit_position_output['next_commit_position'])

  @mock.patch.object(lookback_algorithm, 'GetNextCommitPosition')
  @mock.patch.object(MasterFlakeAnalysis, 'CanRunHeuristicAnalysis')
  @mock.patch.object(next_commit_position_utils,
                     'GetNextCommitPositionFromHeuristicResults')
  @mock.patch.object(step_util, 'GetValidBoundingBuildsForStep')
  @mock.patch.object(heuristic_analysis, 'RunHeuristicAnalysis')
  def testNextCommitPositionPipelineRunHeuristicResultsNoResults(
      self, _, mock_bounding_builds, mock_heuristic_result,
      mock_can_run_heuristic, mock_next_commit):
    master_name = 'm'
    builder_name = 'b'
    build_number = 105
    step_name = 's'
    test_name = 't'
    start_commit_position = 1000
    lower_bound_commit_position = 990

    analysis = MasterFlakeAnalysis.Create(master_name, builder_name,
                                          build_number, step_name, test_name)
    analysis.data_points = [
        DataPoint.Create(commit_position=start_commit_position)
    ]
    analysis.put()

    mock_can_run_heuristic.return_value = True
    mock_heuristic_result.return_value = None

    mock_next_commit.return_value = (999, None)

    lower_bound_build = BuildInfo(master_name, builder_name, build_number - 1)
    lower_bound_build.commit_position = lower_bound_commit_position
    upper_bound_build = BuildInfo(master_name, builder_name, build_number)
    upper_bound_build.commit_position = start_commit_position
    mock_bounding_builds.return_value = (lower_bound_build, upper_bound_build)

    next_commit_position_input = NextCommitPositionInput(
        analysis_urlsafe_key=analysis.key.urlsafe(),
        commit_position_range=IntRange(lower=None, upper=start_commit_position))

    pipeline_job = NextCommitPositionPipeline(next_commit_position_input)
    pipeline_job.start()
    self.execute_queued_tasks()

    pipeline_job = pipelines.pipeline.Pipeline.from_id(pipeline_job.pipeline_id)
    next_commit_position_output = pipeline_job.outputs.default.value

    self.assertFalse(pipeline_job.was_aborted)
    self.assertIsNone(next_commit_position_output['culprit_commit_position'])
    self.assertEqual(lower_bound_commit_position,
                     next_commit_position_output['next_commit_position'])

  @mock.patch.object(lookback_algorithm, 'GetNextCommitPosition')
  @mock.patch.object(next_commit_position_utils, 'GetEarliestCommitPosition')
  def testNextCommitPositionPipelineLongStandingFlake(
      self, mock_earliest_commit, mock_next_commit):
    master_name = 'm'
    builder_name = 'b'
    build_number = 100
    step_name = 's'
    test_name = 't'
    start_commit_position = 1000
    cutoff_commit_position = 500

    mock_earliest_commit.return_value = cutoff_commit_position
    mock_next_commit.return_value = (cutoff_commit_position - 1, None)

    analysis = MasterFlakeAnalysis.Create(master_name, builder_name,
                                          build_number, step_name, test_name)
    analysis.Save()

    next_commit_position_input = NextCommitPositionInput(
        analysis_urlsafe_key=analysis.key.urlsafe(),
        commit_position_range=IntRange(lower=None, upper=start_commit_position))

    pipeline_job = NextCommitPositionPipeline(next_commit_position_input)
    pipeline_job.start()
    self.execute_queued_tasks()

    pipeline_job = pipelines.pipeline.Pipeline.from_id(pipeline_job.pipeline_id)
    next_commit_position_output = pipeline_job.outputs.default.value

    self.assertFalse(pipeline_job.was_aborted)
    self.assertIsNone(next_commit_position_output['culprit_commit_position'])
    self.assertIsNone(next_commit_position_output['next_commit_position'])

  @mock.patch.object(lookback_algorithm, 'GetNextCommitPosition')
  @mock.patch.object(step_util, 'GetValidBoundingBuildsForStep')
  @mock.patch.object(MasterFlakeAnalysis, 'CanRunHeuristicAnalysis')
  def testNextCommitPositionPipelineContinueAnalysis(
      self, mock_heuristic, mock_bounding_builds, mock_next_commit):
    master_name = 'm'
    builder_name = 'b'
    build_number = 100
    step_name = 's'
    test_name = 't'
    start_commit_position = 1000
    expected_next_commit_position = 990

    mock_heuristic.return_value = False

    calculated_next_commit_position = 999
    culprit_commit_position = None
    mock_next_commit.return_value = (calculated_next_commit_position,
                                     culprit_commit_position)

    lower_bound_build = BuildInfo(master_name, builder_name, build_number - 1)
    lower_bound_build.commit_position = expected_next_commit_position
    upper_bound_build = BuildInfo(master_name, builder_name, build_number)
    upper_bound_build.commit_position = start_commit_position
    mock_bounding_builds.return_value = (lower_bound_build, upper_bound_build)

    analysis = MasterFlakeAnalysis.Create(master_name, builder_name,
                                          build_number, step_name, test_name)
    analysis.data_points = [
        DataPoint.Create(commit_position=start_commit_position)
    ]
    analysis.Save()

    next_commit_position_input = NextCommitPositionInput(
        analysis_urlsafe_key=analysis.key.urlsafe(),
        commit_position_range=IntRange(lower=None, upper=start_commit_position))

    pipeline_job = NextCommitPositionPipeline(next_commit_position_input)
    pipeline_job.start()
    self.execute_queued_tasks()

    pipeline_job = pipelines.pipeline.Pipeline.from_id(pipeline_job.pipeline_id)
    next_commit_position_output = pipeline_job.outputs.default.value

    self.assertFalse(pipeline_job.was_aborted)
    self.assertIsNone(next_commit_position_output['culprit_commit_position'])
    self.assertEqual(expected_next_commit_position,
                     next_commit_position_output['next_commit_position'])
