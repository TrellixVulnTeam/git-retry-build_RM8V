# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
"""Utility functions for code coverage."""

import json
import urllib2

from common.findit_http_client import FinditHttpClient
from gae_libs.caches import PickledMemCache
from libs.cache_decorator import Cached

# Mapping from metric names to detailed explanations, and one use case is to use
# as the tooltips.
_METRIC_NAME_DETAIL_MAPPING = {
    'line': (
        "Line coverage is the percentage of code lines which have been "
        "executed at least once. Only executable lines within function bodies "
        "are considered to be code lines."),
    'function': (
        "Function coverage is the percentage of functions which have been "
        "executed at least once. A function is considered to be executed if "
        "any of its instantiations are executed."),
    'region': (
        "Region coverage is the percentage of code regions which have been "
        "executed at least once. A code region may span multiple lines (e.g in "
        "a large function body with no control flow). However, it's also "
        "possible for a single line to contain multiple code regions (e.g in "
        "'return x || y &amp;&amp; z')."),
    'branch': (
        "Branch coverage is the percentage of branches from each decision "
        "point is executed at least once."),
    'instruction': (
        "Java instruction coverage is the percentage of the Java byte code "
        "instructions which have been executed at least once."),
}

# List of patchset kinds that are applicable for sharing coverage data between
# patchsets, and the list of possible kinds is defined in:
# https://gerrit-review.googlesource.com/Documentation/json.html
_NON_CONFLICT_CHANGE_KIND = [
    # Conflict-free merge between the new parent and the prior patch set.
    'TRIVIAL_REBASE',
    # Conflict-free change of first (left) parent of a merge commit.
    'MERGE_FIRST_PARENT_UPDATE',
    # No code changed; same tree and same parent tree.
    'NO_CODE_CHANGE',
    # No changes; same commit message, same tree and same parent tree.
    'NO_CHANGE'
]


def GetMetricsBasedOnCoverageTool(coverage_tool):
  """Gets a list of metrics for the given coverage tool.

  Args:
    coverage_tool(str): Name of the coverage tool, such as clang and jacoco.

  Returns:
    A list of dict of following format:
    {'name': clang, 'detail': blala}, where the name is the name of the metric
    and detail is an explanation of what the metric stands for.
  """
  assert coverage_tool in ('clang', 'jacoco'), (
      'Unrecognized coverage tool: %s' % coverage_tool)

  metrics = []
  if coverage_tool == 'clang':
    metrics = ['line', 'function', 'region']
  else:
    metrics = ['line', 'branch', 'instruction']

  return [{
      'name': m,
      'detail': _METRIC_NAME_DETAIL_MAPPING.get(m, '')
  } for m in metrics]


@Cached(PickledMemCache(), namespace='coverage_equivalent_patchsets')
def GetEquivalentPatchsets(host, project, change, patchset):
  """Gets equivalent patchsets that are applicable for sharing coverage data.

  The reason why this is not just the current patchset number is because there
  may have been a succession of "trivial" changes before the current patchset.

  Args:
    host (str): The host name.
    project (str): The project name.
    change (int): The change number.
    patchset (int): The patchset number.

  Returns:
    A list of equivalent patchset numbers in descending order.
  """
  assert isinstance(change, int), 'Change is expected to be an integer'
  assert isinstance(patchset, int), 'Patchset is expected to be an integer'

  project_quoted = urllib2.quote(project, safe='')

  # Uses the Get Change API to get and parse the details of this change.
  # https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-change.
  template_to_get_change = (
      'https://%s/changes/%s~%d?o=ALL_REVISIONS&o=SKIP_MERGEABLE')
  url = template_to_get_change % (host, project_quoted, change)
  status_code, content, _ = FinditHttpClient().Get(url)
  if status_code != 200:
    raise RuntimeError(
        'Failed to get change details with status code: %d' % status_code)

  # Remove XSSI magic prefix
  if content.startswith(')]}\''):
    content = content[4:]
  change_details = json.loads(content)

  revisions = change_details['revisions'].values()
  revisions.sort(key=lambda r: r['_number'], reverse=True)
  patchsets = []
  for i, r in enumerate(revisions):
    if i == 0 and change_details['status'] == 'MERGED':
      # Depending on the submit strategy, the last patchset of submitted CLs
      # might be autogenerated and whose kind is labeled as 'REWORK' even though
      # it's actually trivial rebase.
      #
      # This function assumes that the submit strategy is 'Rebase Always' (such
      # as Chromium project), and it may break for projects with other submit
      # strategies, for example: crbug.com/809182.
      #
      # TODO(crbug.com/809182): Make the equivalent patchsets logic generic
      # across all projects. Note that the bug specifically refers to
      # buildbucket, but the same reasonings apply here.
      continue

    current_patchset = r['_number']
    if current_patchset > patchset:
      continue

    patchsets.append(current_patchset)
    if r['kind'] not in _NON_CONFLICT_CHANGE_KIND:
      # If this revision was a non-trivial change, don't consider patchsets
      # prior to it.
      break

  return patchsets
