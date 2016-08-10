# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import webapp2

from waterfall.test import wf_testcase
from handlers.flake import filter_flake
from handlers.flake.filter_flake import FilterMasterFlakeAnalysis
from model.flake.master_flake_analysis import MasterFlakeAnalysis


class FilterFlakeTest(wf_testcase.WaterfallTestCase):
  app_module = webapp2.WSGIApplication([
      ('/waterfall/filter-flake', filter_flake.FilterFlake),
  ], debug=True)

  def _CreateAndSaveMasterFlakeAnalysis(
      self, master_name, builder_name, build_number,
      step_name, test_name):
    analysis = MasterFlakeAnalysis.Create(
        master_name, builder_name, build_number, step_name, test_name)
    analysis.put()
    return analysis

  def setUp(self):
    super(FilterFlakeTest, self).setUp()
    self.mock_current_user(user_email='test@chromium.org', is_admin=True)
    self.master_name1 = 'm1'
    self.master_name2 = 'm2'
    self.builder_name1 = 'b1'
    self.builder_name2 = 'b2'
    self.build_number1 = 1
    self.build_number2 = 2
    self.step_name1 = 's1'
    self.step_name2 = 's2'
    self.test_name1 = 't1'
    self.test_name2 = 't2'
    self.master_flake_analysis1 = self._CreateAndSaveMasterFlakeAnalysis(
        self.master_name1, self.builder_name1, self.build_number1,
        self.step_name1, self.test_name1)
    self.master_flake_analysis2 = self._CreateAndSaveMasterFlakeAnalysis(
        self.master_name2, self.builder_name2, self.build_number2,
        self.step_name2, self.test_name2)
    self.master_flake_analysis3 = self._CreateAndSaveMasterFlakeAnalysis(
        self.master_name2, self.builder_name2, self.build_number2,
        self.step_name2, self.test_name1)

  def testFilterMasterName(self):
    master_flake_analysis_query = MasterFlakeAnalysis.query()
    result = FilterMasterFlakeAnalysis(
        master_flake_analysis_query, self.master_name1, None, None, None, None)
    self.assertEqual(len(result), 1)
    self.assertTrue(result == [self.master_flake_analysis1])

  def testFilterBuilderName(self):
    master_flake_analysis_query = MasterFlakeAnalysis.query()
    result = FilterMasterFlakeAnalysis(master_flake_analysis_query, None,
                    self.builder_name1, None, None, None)
    self.assertEqual(len(result), 1)
    self.assertTrue(result == [self.master_flake_analysis1])

  def testFilterBuildNumber(self):
    master_flake_analysis_query = MasterFlakeAnalysis.query()
    result = FilterMasterFlakeAnalysis(master_flake_analysis_query, None, None,
                    self.build_number1, None, None)
    self.assertEqual(len(result), 1)
    self.assertTrue(result == [self.master_flake_analysis1])

  def testFilterStepName(self):
    master_flake_analysis_query = MasterFlakeAnalysis.query()
    result = FilterMasterFlakeAnalysis(
        master_flake_analysis_query, None, None, None, self.step_name1, None)
    self.assertEqual(len(result), 1)
    self.assertTrue(result == [self.master_flake_analysis1])

  def testFilterTestName(self):
    master_flake_analysis_query = MasterFlakeAnalysis.query()
    result = FilterMasterFlakeAnalysis(
        master_flake_analysis_query, None, None, None, None, self.test_name2)
    self.assertEqual(len(result), 1)
    self.assertTrue(result == [self.master_flake_analysis2])

  def testFilterMultipleMasterName(self):
    master_flake_analysis_query = MasterFlakeAnalysis.query()
    result = FilterMasterFlakeAnalysis(
        master_flake_analysis_query, self.master_name2, None, None, None, None)
    self.assertEqual(len(result), 2)
    self.assertTrue(result == [self.master_flake_analysis3,
                               self.master_flake_analysis2])

  def testFilterMultipleBuilderName(self):
    master_flake_analysis_query = MasterFlakeAnalysis.query()
    result = FilterMasterFlakeAnalysis(
        master_flake_analysis_query, None, self.builder_name2, None, None, None)
    self.assertEqual(len(result), 2)
    self.assertTrue(result == [self.master_flake_analysis3,
                               self.master_flake_analysis2])

  def testFilterMultipleBuildNumber(self):
    master_flake_analysis_query = MasterFlakeAnalysis.query()
    result = FilterMasterFlakeAnalysis(
        master_flake_analysis_query, None, None, self.build_number2, None, None)
    self.assertEqual(len(result), 2)
    self.assertTrue(result == [self.master_flake_analysis3,
                               self.master_flake_analysis2])

  def testFilterMultipleStepName(self):
    master_flake_analysis_query = MasterFlakeAnalysis.query()
    result = FilterMasterFlakeAnalysis(
        master_flake_analysis_query, None, None, None, self.step_name2, None)
    self.assertEqual(len(result), 2)
    self.assertTrue(result == [self.master_flake_analysis3,
                               self.master_flake_analysis2])

  def testFilterMultipleTestName(self):
    master_flake_analysis_query = MasterFlakeAnalysis.query()
    result = FilterMasterFlakeAnalysis(
        master_flake_analysis_query, None, None, None, None, self.test_name1)
    self.assertEqual(len(result), 2)
    self.assertTrue(result == [self.master_flake_analysis1,
                               self.master_flake_analysis3])

  def testNormalFlow(self):
    self.mock_current_user(user_email='test@chromium.org', is_admin=True)
    response = self.test_app.get('/waterfall/filter-flake')
    self.assertEquals(200, response.status_int)

  def testNormalFlowWithFilter(self):
    self.mock_current_user(user_email='test@chromium.org', is_admin=True)
    response = self.test_app.get(
        '/waterfall/filter-flake',
        params={'build_number': self.build_number1,
                'format': 'json'}
    )
    expected_result = {
        'master_flake_analyses': [
            {'step_name': self.step_name1,
             'master_name': self.master_name1,
             'build_number': self.build_number1,
             'builder_name': self.builder_name1,
             'test_name': self.test_name1
            }
        ]
    }
    self.assertEquals(response.json_body, expected_result)
    self.assertEquals(200, response.status_int)
