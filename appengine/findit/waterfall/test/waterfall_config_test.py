# Copyright 2015 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from waterfall import waterfall_config
from waterfall.test import wf_testcase


class MastersTest(wf_testcase.WaterfallTestCase):

  def testConvertOldMastersFormatToNew(self):
    self.assertEqual(
        {
            'supported_masters': {
                'master1': {
                    'unsupported_steps': ['1', '2']
                },
                'master2': {}
            },
            'global': {}
        },
        waterfall_config._ConvertOldMastersFormatToNew(
            {
                'master1': ['1', '2'],
                'master2': {}
            }))

  def testGetStepsForMastersRulesWithSettingsProvided(self):
    class MockSettings():
      steps_for_masters_rules = {'blabla': 'blabla'}

    self.assertEqual(waterfall_config.GetStepsForMastersRules(MockSettings()),
                     MockSettings().steps_for_masters_rules)

  def testMasterIsSupported(self):
    self.assertTrue(waterfall_config.MasterIsSupported('master1'))
    self.assertFalse(waterfall_config.MasterIsSupported('blabla'))

  def testStepIsSupportedForMaster(self):
    self.assertFalse(
        waterfall_config.StepIsSupportedForMaster('unsupported_step1',
                                                  'master1'))
    self.assertTrue(
        waterfall_config.StepIsSupportedForMaster('step4', 'master1'))
    self.assertTrue(
        waterfall_config.StepIsSupportedForMaster('step4', 'master2'))
    self.assertFalse(
        waterfall_config.StepIsSupportedForMaster('blabla', 'blabla'))
    self.assertTrue(
        waterfall_config.StepIsSupportedForMaster('step4', 'master2'))
    self.assertTrue(
        waterfall_config.StepIsSupportedForMaster('blabla', 'master3'))
    self.assertTrue(
        waterfall_config.StepIsSupportedForMaster('step5', 'master1'))
    self.assertTrue(
        waterfall_config.StepIsSupportedForMaster('step5', 'master2'))
    self.assertFalse(
        waterfall_config.StepIsSupportedForMaster('unsupported_step7',
                                                  'master2'))
    self.assertTrue(
        waterfall_config.StepIsSupportedForMaster('unsupported_step6',
                                                  'master1'))
    self.assertFalse(
        waterfall_config.StepIsSupportedForMaster('unsupported_step6',
                                                  'master2'))
    self.assertFalse(
        waterfall_config.StepIsSupportedForMaster('unsupported_step6',
                                                  'master3'))
    self.assertFalse(
        waterfall_config.StepIsSupportedForMaster('unsupported_step7',
                                                  'master3'))

  def testGetTrybotForWaterfallBuilder(self):
    self.assertEqual(
        ('tryserver1', 'trybot1'),
        waterfall_config.GetTrybotForWaterfallBuilder('master1', 'builder1'))
    self.assertEqual(
        (None, None),
        waterfall_config.GetTrybotForWaterfallBuilder('master3', 'builder3'))

  def testGetTryJobSettings(self):
    self.assertEqual(
        {
            'server_query_interval_seconds': 60,
            'job_timeout_hours': 5,
            'allowed_response_error_times': 5,
            'max_seconds_look_back_for_group': 86400
        },
        waterfall_config.GetTryJobSettings())

  def testGetSwarmingSettings(self):
    self.assertEqual(
        {
            'server_host': 'chromium-swarm.appspot.com',
            'default_request_priority': 150,
            'request_expiration_hours': 20,
            'server_query_interval_seconds': 60,
            'task_timeout_hours': 23,
            'isolated_server': 'https://isolateserver.appspot.com',
            'isolated_storage_url': 'isolateserver.storage.googleapis.com',
            'iterations_to_rerun': 10,
            'get_swarming_task_id_timeout_seconds': 300,
            'get_swarming_task_id_wait_seconds': 10
        },
        waterfall_config.GetSwarmingSettings())

  def testGetDownloadBuildDataSettings(self):
    self.assertEqual(
        {
            'download_interval_seconds': 10,
            'memcache_master_download_expiration_seconds': 3600,
            'use_chrome_build_extract': True
        },
        waterfall_config.GetDownloadBuildDataSettings())

  def testEnableStrictRegexForCompileLinkFailures(self):
    self.assertFalse(
        waterfall_config.EnableStrictRegexForCompileLinkFailures('m', 'b'))
    self.assertTrue(
        waterfall_config.EnableStrictRegexForCompileLinkFailures(
            'master1', 'builder1'))

  def testShouldSkipTestTryJobs(self):
    self.assertFalse(
        waterfall_config.ShouldSkipTestTryJobs('master1', 'builder1'))
    self.assertFalse(
        waterfall_config.ShouldSkipTestTryJobs('master2', 'builder3'))
    self.assertTrue(
        waterfall_config.ShouldSkipTestTryJobs('master2', 'builder2'))

  def testGetActionSettings(self):
    self.assertEqual(
        {
            'cr_notification_build_threshold': 2,
            'cr_notification_latency_limit_minutes': 30,
        },
        waterfall_config.GetActionSettings())

  def testGetCheckFlakeSettings(self):
    self.assertEqual(
        {
            'lower_flake_threshold': 0.02,
            'upper_flake_threshold': 0.98,
            'max_flake_in_a_row': 4,
            'max_stable_in_a_row': 4,
            'iterations_to_rerun': 100
        },
        waterfall_config.GetCheckFlakeSettings())
