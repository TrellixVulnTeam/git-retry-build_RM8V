# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from google.appengine.api import taskqueue

from common import constants
from gae_libs.appengine_util import IsStaging
from gae_libs.handlers.base_handler import BaseHandler
from gae_libs.handlers.base_handler import Permission
from services.flake_detection import detect_cq_false_rejection_flakes
from services.flake_detection import flake_report_util


class DetectCQFalseRejectionFlakesCronJob(BaseHandler):
  PERMISSION_LEVEL = Permission.APP_SELF

  def HandleGet(self):
    # Cron jobs run independently of each other. Therefore, there is no
    # guarantee that they will run either sequentially or simultaneously.
    #
    # Running flake detection tasks concurrently doesn't bring much benefits, so
    # use task queue to enforce that at most one detection task can be executed
    # at any time to avoid any potential race condition.
    taskqueue.add(
        method='GET',
        queue_name=constants.FLAKE_DETECTION_QUEUE,
        target=constants.FLAKE_DETECTION_BACKEND,
        url='/flake/detection/task/detect-cq-false-rejection-flakes')
    return {'return_code': 200}


class DetectCQFalseRejectionFlakes(BaseHandler):
  PERMISSION_LEVEL = Permission.APP_SELF

  def HandleGet(self):
    # Only triggers flake detections on staging for experimental and debugging
    # purposes.
    if not IsStaging():
      return {'return_code': 200}

    detect_cq_false_rejection_flakes.QueryAndStoreFlakes()
    flake_tuples_to_report = flake_report_util.GetFlakesWithEnoughOccurrences()
    flake_report_util.ReportFlakesToMonorail(flake_tuples_to_report)

    # Reporting to Flake Analyzer needs to come after reporting to Monorail
    # because it makes Flake Analyzer's job easier if it can reuse the issues
    # filed by Flake Detection.
    flake_report_util.ReportFlakesToFlakeAnalyzer(flake_tuples_to_report)
    return {'return_code': 200}
