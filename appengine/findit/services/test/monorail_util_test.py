# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
import datetime
import mock

from libs import time_util
from monorail_api import Issue
from monorail_api import IssueTrackerAPI
from services import issue_generator
from services import monorail_util
from waterfall.test import wf_testcase


class TestIssueGenerator(issue_generator.FlakyTestIssueGenerator):
  """A FlakyTestIssueGenerator used for testing."""

  def __init__(self,
               step_name='step',
               test_name='suite.test',
               test_label_name='*/suite.test/*'):
    super(TestIssueGenerator, self).__init__()
    self.step_name = step_name
    self.test_name = test_name
    self.test_label_name = test_label_name
    self._previous_tracking_bug_id = None

  def GetStepName(self):
    return self.step_name

  def GetTestName(self):
    return self.test_name

  def GetTestLabelName(self):
    return self.test_label_name

  def GetDescription(self):
    previous_tracking_bug_id = self.GetPreviousTrackingBugId()
    if previous_tracking_bug_id:
      return ('description with previous tracking bug id: %s.' %
              previous_tracking_bug_id)

    return 'description without previous tracking bug id.'

  def GetComment(self):
    previous_tracking_bug_id = self.GetPreviousTrackingBugId()
    if previous_tracking_bug_id:
      return ('comment with previous tracking bug id: %s.' %
              previous_tracking_bug_id)

    return 'comment without previous tracking bug id.'

  def ShouldRestoreChromiumSheriffLabel(self):
    # Sets to False as default value, if need to test this control flow, please
    # mock this method.
    return False

  def GetLabels(self):
    return ['label1', 'Sheriff-Chromium']


class MonorailUtilTest(wf_testcase.WaterfallTestCase):

  @mock.patch.object(
      time_util, 'GetUTCNow', return_value=datetime.datetime(2017, 1, 3))
  @mock.patch('services.monorail_util.IssueTrackerAPI')
  def testOpenBugAlreadyExistsForId(self, mock_api, _):
    mock_api.return_value.getIssue.return_value = None
    self.assertFalse(monorail_util.OpenBugAlreadyExistsForId(None))
    self.assertFalse(mock_api.return_value.getIssue.called)
    mock_api.reset_mock()

    mock_api.return_value.getIssue.return_value = None
    self.assertFalse(monorail_util.OpenBugAlreadyExistsForId(1234))
    self.assertTrue(mock_api.return_value.getIssue.called)
    args, _ = mock_api.return_value.getIssue.call_args
    self.assertEqual((1234,), args)
    mock_api.reset_mock()

    mock_issue = mock.MagicMock()
    mock_issue.open = True
    mock_issue.updated = datetime.datetime(2017, 1, 1)
    mock_issue.merged_into = None
    mock_api.return_value.getIssue.return_value = mock_issue
    self.assertTrue(monorail_util.OpenBugAlreadyExistsForId(1234))
    self.assertTrue(mock_api.return_value.getIssue.called)
    args, _ = mock_api.return_value.getIssue.call_args
    self.assertEqual((1234,), args)
    mock_api.reset_mock()

    mock_issue = mock.MagicMock()
    mock_issue.open = False
    mock_issue.updated = datetime.datetime(2017, 1, 2)
    mock_issue.merged_into = None
    mock_api.return_value.getIssue.return_value = mock_issue
    self.assertFalse(monorail_util.OpenBugAlreadyExistsForId(1234))
    self.assertTrue(mock_api.return_value.getIssue.called)
    args, _ = mock_api.return_value.getIssue.call_args
    self.assertEqual((1234,), args)
    mock_api.reset_mock()

  @mock.patch('services.monorail_util.IssueTrackerAPI')
  def testCreateBug(self, mock_api):
    summary = 'test summary'
    description = 'test description'
    project_id = 'proj'
    issue = Issue({
        'status': 'Available',
        'summary': summary,
        'description': description,
        'projectId': 'chromium',
        'state': 'open',
    })

    monorail_util.CreateBug(issue, project_id=project_id)
    mock_api.assert_has_calls(mock.call(project_id, use_staging=False))
    mock_api.return_value.create.assert_has_calls(mock.call(issue))

  @mock.patch('services.monorail_util.IssueTrackerAPI')
  def testUpdateBug(self, mock_api):
    summary = 'test summary'
    description = 'test description'
    project_id = 'proj'
    comment = 'test comment'
    issue = Issue({
        'status': 'Available',
        'summary': summary,
        'description': description,
        'projectId': 'chromium',
        'state': 'open',
    })

    monorail_util.UpdateBug(issue, comment, project_id=project_id)
    mock_api.assert_has_calls(mock.call(project_id, use_staging=False))
    mock_api.return_value.update.assert_has_calls(
        mock.call(issue, comment, send_email=True))

  @mock.patch('services.monorail_util.IssueTrackerAPI')
  def testGetMergedDestinationIssueWithoutMergeInto(self, mock_api):
    issue = Issue({'id': 12345})
    mock_api.return_value.getIssue.return_value = issue
    self.assertEqual(issue, monorail_util.GetMergedDestinationIssueForId(12345))

  @mock.patch('services.monorail_util.IssueTrackerAPI')
  def testGetMergedDestinationIssueWithMergeInto(self, mock_api):
    issue = Issue({'id': 12345, 'mergedInto': {'issueId': 56789}})

    another_issue = Issue({'id': 56789})

    def _return_issue(issue_id):
      if issue_id == 12345:
        return issue

      if issue_id == 56789:
        return another_issue

      return None

    mock_api.return_value.getIssue.side_effect = _return_issue
    self.assertEqual(another_issue,
                     monorail_util.GetMergedDestinationIssueForId(12345))

  @mock.patch('services.monorail_util.IssueTrackerAPI')
  def testGetMergedDestinationIssueWithMergeInCircle(self, mock_api):
    issue = Issue({'id': 12345, 'mergedInto': {'issueId': 56789}})

    another_issue = Issue({'id': 56789, 'mergedInto': {'issueId': 12345}})

    def _return_issue(issue_id):
      if issue_id == 12345:
        return issue

      if issue_id == 56789:
        return another_issue

      return None

    mock_api.return_value.getIssue.side_effect = _return_issue
    self.assertEqual(issue, monorail_util.GetMergedDestinationIssueForId(12345))

  # This test tests that creating issue via issue generator without previous
  # tracking bug id works properly.
  @mock.patch.object(monorail_util, 'UpdateBug')
  @mock.patch.object(monorail_util, 'CreateBug', return_value=12345)
  def testCreateIssueWithIssueGenerator(self, mock_create_bug_fn,
                                        mock_update_bug_fn):
    test_issue_generator = TestIssueGenerator()
    issue_id = monorail_util.CreateIssueWithIssueGenerator(
        issue_generator=test_issue_generator)

    self.assertTrue(mock_create_bug_fn.called)
    self.assertFalse(mock_update_bug_fn.called)
    self.assertEqual(12345, issue_id)
    issue = mock_create_bug_fn.call_args_list[0][0][0]
    self.assertEqual('Untriaged', issue.status)
    self.assertEqual('*/suite.test/* is flaky', issue.summary)
    self.assertEqual('description without previous tracking bug id.',
                     issue.description)
    self.assertEqual(['label1', 'Sheriff-Chromium', 'Pri-1'], issue.labels)
    self.assertEqual(1, len(issue.field_values))
    self.assertEqual('Flaky-Test', issue.field_values[0].to_dict()['fieldName'])
    self.assertEqual('suite.test',
                     issue.field_values[0].to_dict()['fieldValue'])

  # This test tests that creating issue via issue generator with previous
  # tracking bug id works properly.
  @mock.patch.object(monorail_util, 'UpdateBug')
  @mock.patch.object(monorail_util, 'CreateBug', return_value=12345)
  def testCreateIssueWithIssueGeneratorWithPreviousTrackingBugId(
      self, mock_create_bug_fn, mock_update_bug_fn):
    test_issue_generator = TestIssueGenerator()
    test_issue_generator.SetPreviousTrackingBugId(56789)
    issue_id = monorail_util.CreateIssueWithIssueGenerator(
        issue_generator=test_issue_generator)

    self.assertTrue(mock_create_bug_fn.called)
    self.assertFalse(mock_update_bug_fn.called)
    self.assertEqual(12345, issue_id)
    issue = mock_create_bug_fn.call_args_list[0][0][0]
    self.assertEqual('Untriaged', issue.status)
    self.assertEqual('*/suite.test/* is flaky', issue.summary)
    self.assertEqual('description with previous tracking bug id: 56789.',
                     issue.description)
    self.assertEqual(['label1', 'Sheriff-Chromium', 'Pri-1'], issue.labels)
    self.assertEqual(1, len(issue.field_values))
    self.assertEqual('Flaky-Test', issue.field_values[0].to_dict()['fieldName'])
    self.assertEqual('suite.test',
                     issue.field_values[0].to_dict()['fieldValue'])

  # This test tests that updating issue via issue generator without previous
  # tracking bug id works properly.
  @mock.patch.object(monorail_util, 'GetMergedDestinationIssueForId')
  @mock.patch.object(monorail_util, 'UpdateBug')
  @mock.patch.object(monorail_util, 'CreateBug')
  def testUpdateIssueWithIssueGenerator(
      self, mock_create_bug_fn, mock_update_bug_fn, mock_get_merged_issue):
    issue_id = 12345
    issue = Issue({
        'status': 'Available',
        'summary': 'summary',
        'description': 'description',
        'projectId': 'chromium',
        'labels': [],
        'fieldValues': [],
        'state': 'open',
    })
    mock_get_merged_issue.return_value = issue

    test_issue_generator = TestIssueGenerator()
    monorail_util.UpdateIssueWithIssueGenerator(
        issue_id=issue_id, issue_generator=test_issue_generator)

    self.assertFalse(mock_create_bug_fn.called)
    mock_update_bug_fn.assert_called_once_with(
        issue, 'comment without previous tracking bug id.', 'chromium')
    issue = mock_update_bug_fn.call_args_list[0][0][0]
    self.assertEqual(['label1'], issue.labels)
    self.assertEqual(1, len(issue.field_values))
    self.assertEqual('Flaky-Test', issue.field_values[0].to_dict()['fieldName'])
    self.assertEqual('suite.test',
                     issue.field_values[0].to_dict()['fieldValue'])

  # This test tests that updating issue via issue generator works properly if
  # the switch to turn on to restore Chromium Sheriffs label when update bugs.
  @mock.patch.object(
      TestIssueGenerator,
      'ShouldRestoreChromiumSheriffLabel',
      return_value=True)
  @mock.patch.object(monorail_util, 'GetMergedDestinationIssueForId')
  @mock.patch.object(monorail_util, 'UpdateBug')
  @mock.patch.object(monorail_util, 'CreateBug')
  def testUpdateIssueWithIssueGeneratorAndRestoreSheriffLabel(
      self, mock_create_bug_fn, mock_update_bug_fn, mock_get_merged_issue, _):
    issue_id = 12345
    issue = Issue({
        'status': 'Available',
        'summary': 'summary',
        'description': 'description',
        'projectId': 'chromium',
        'labels': [],
        'fieldValues': [],
        'state': 'open',
    })
    mock_get_merged_issue.return_value = issue

    test_issue_generator = TestIssueGenerator()
    monorail_util.UpdateIssueWithIssueGenerator(
        issue_id=issue_id, issue_generator=test_issue_generator)

    self.assertFalse(mock_create_bug_fn.called)
    self.assertTrue(mock_update_bug_fn.called)
    self.assertEqual(['label1', 'Sheriff-Chromium'], issue.labels)

  # This test tests that updating issue via issue generator with previous
  # tracking id works properly.
  @mock.patch.object(monorail_util, 'GetMergedDestinationIssueForId')
  @mock.patch.object(monorail_util, 'UpdateBug')
  @mock.patch.object(monorail_util, 'CreateBug')
  def testUpdateIssueWithIssueGeneratorWithPreviousTrackingId(
      self, mock_create_bug_fn, mock_update_bug_fn, mock_get_merged_issue):
    issue_id = 12345
    issue = Issue({
        'status': 'Available',
        'summary': 'summary',
        'description': 'description',
        'projectId': 'chromium',
        'labels': [],
        'fieldValues': [],
        'state': 'open',
    })
    mock_get_merged_issue.return_value = issue

    test_issue_generator = TestIssueGenerator()
    test_issue_generator.SetPreviousTrackingBugId(56789)
    monorail_util.UpdateIssueWithIssueGenerator(
        issue_id=issue_id, issue_generator=test_issue_generator)

    self.assertFalse(mock_create_bug_fn.called)
    mock_update_bug_fn.assert_called_once_with(
        issue, 'comment with previous tracking bug id: 56789.', 'chromium')

  @mock.patch('services.monorail_util.IssueTrackerAPI')
  def testGetMonorailIssueForIssueId(self, mocked_issue_tracker_api):
    issue = Issue({
        'status': 'Available',
        'summary': 'summary',
        'description': 'description',
        'projectId': 'chromium',
        'labels': [],
        'fieldValues': [],
        'state': 'open',
        'id': '12345'
    })
    mocked_issue_tracker_api.return_value.getIssue.return_value = issue
    self.assertEqual(
        issue, monorail_util.GetMonorailIssueForIssueId(12345, 'chromium'))
