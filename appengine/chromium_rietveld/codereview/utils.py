# Copyright 2011 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Collection of helper functions."""

import re
import urlparse

from google.appengine.ext import db

from codereview.exceptions import FetchError


def make_url(base, filename, rev):
  """Helper to construct the URL to fetch.

  Args:
    base: The base property of the Issue to which the Patch belongs.
    filename: The filename property of the Patch instance.
    rev: Revision number, or None for head revision.

  Returns:
    A URL referring to the given revision of the file.
  """
  scheme, netloc, path, _, _, _ = urlparse.urlparse(base)
  if netloc.endswith(".googlecode.com"):
    # Handle Google code repositories
    if rev is None:
      raise FetchError("Can't access googlecode.com without a revision")
    if not path.startswith("/svn/"):
      raise FetchError( "Malformed googlecode.com URL (%s)" % base)
    path = path[5:]  # Strip "/svn/"
    url = "%s://%s/svn-history/r%d/%s/%s" % (scheme, netloc, rev,
                                             path, filename)
    return url
  elif netloc.endswith("sourceforge.net") and rev is not None:
    if path.strip().endswith("/"):
      path = path.strip()[:-1]
    else:
      path = path.strip()
    splitted_path = path.split("/")
    url = "%s://%s/%s/!svn/bc/%d/%s/%s" % (scheme, netloc,
                                           "/".join(splitted_path[1:3]), rev,
                                           "/".join(splitted_path[3:]),
                                           filename)
    return url
  # Default for viewvc-based URLs (svn.python.org)
  url = base
  if not url.endswith('/'):
    url += '/'
  url += filename
  if rev is not None:
    url += '?rev=%s' % rev
  return url


def to_dbtext(text):
  """Helper to turn a string into a db.Text instance.

  Args:
    text: a string.

  Returns:
    A db.Text instance.
  """
  if isinstance(text, unicode):
    # A TypeError is raised if text is unicode and an encoding is given.
    return db.Text(text)
  else:
    try:
      return db.Text(text, encoding='utf-8')
    except UnicodeDecodeError:
      return db.Text(text, encoding='latin-1')


def unify_linebreaks(text):
  """Helper to return a string with all line breaks converted to LF.

  Args:
    text: a string.

  Returns:
    A string with all line breaks converted to LF.
  """
  return text.replace('\r\n', '\n').replace('\r', '\n')


_CQ_STATUS_REGEX = re.compile(
    '(dry run: )?CQ is trying da patch. Follow status at\s+'
    '(https://.+/patch-status/(.+/)?(\d+)/(\d+))\s*', re.I)


def parse_cq_status_url_message(msg):
  """Returns url, issue, patchset parsed from CQ status message.

  If parsing failed, returns None, None, None.
  """
  # Example of message, Dry Run prefix is optional.
  # Dry run: CQ is trying da patch. Follow status at
  # https://chromium-cq-status.appspot.com/v2/patch-status/codereview.chromium.org/2131593002/1
  match = _CQ_STATUS_REGEX.match(msg)
  if not match:
    return None, None, None
  _, url, _, issue, patchset = match.groups()
  return url, int(issue), int(patchset)
