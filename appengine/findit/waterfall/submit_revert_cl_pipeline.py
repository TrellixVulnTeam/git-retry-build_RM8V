# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from gae_libs.pipeline_wrapper import BasePipeline
from libs import analysis_status as status
from model.wf_suspected_cl import WfSuspectedCL
from services import revert


class SubmitRevertCLPipeline(BasePipeline):

  def __init__(self, repo_name, revision, _):
    super(SubmitRevertCLPipeline, self).__init__(repo_name, revision, _)
    self.repo_name = repo_name
    self.revision = revision

  def _LogUnexpectedAborting(self, was_aborted):
    if not was_aborted:  # pragma: no cover
      return

    culprit = WfSuspectedCL.Get(self.repo_name, self.revision)

    if culprit.submit_revert_pipeline_id == self.pipeline_id:
      if (culprit.revert_submission_status and
          culprit.revert_submission_status != status.COMPLETED):
        culprit.revert_submission_status = status.ERROR
      culprit.submit_revert_pipeline_id = None
      culprit.put()

  def finalized(self):  # pragma: no cover
    self._LogUnexpectedAborting(self.was_aborted)

  # Arguments number differs from overridden method - pylint: disable=W0221
  def run(self, repo_name, revision, revert_status):
    return revert.CommitRevert(repo_name, revision, revert_status,
                               self.pipeline_id)
