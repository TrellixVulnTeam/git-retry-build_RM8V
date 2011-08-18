# Copyright (c) 2011 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Setups a local GAE instance to test against a live server for integration
tests.

It makes sure Google AppEngine SDK is found and starts the server on a free
inbound TCP port.
"""

import logging
import os
import re
import signal
import socket
import subprocess
import tempfile
import time
import sys
import urllib

GAE_SDK = None


def _load_modules():
  """Loads all the necessary modules.

  Update sys.path to be able to import chromium-status and GAE SDK.
  """
  global GAE_SDK
  if GAE_SDK:
    return
  root_dir = os.path.dirname(os.path.abspath(__file__))
  # First, verify the Google AppEngine SDK is available.
  while len(root_dir) > 3:
    if os.path.isfile(os.path.join(root_dir, 'google_appengine', 'VERSION')):
      break
    root_dir = os.path.dirname(root_dir)
  if len(root_dir) <= 3:
    raise Failure('Install google_appengine sdk in %s' % GAE_SDK)
  GAE_SDK = os.path.realpath(os.path.join(root_dir, 'google_appengine'))
  # Need yaml later.
  gae_sdk_lib = os.path.realpath(os.path.join(GAE_SDK, 'lib'))
  sys.path.insert(0, os.path.realpath(os.path.join(gae_sdk_lib, 'yaml', 'lib')))


class Failure(Exception):
  pass


def test_port(port):
  s = socket.socket()
  try:
    return s.connect_ex(('127.0.0.1', port)) == 0
  finally:
    s.close()


def find_free_port():
  """Finds an available port starting at 8080."""
  port = 8080
  max_val = (2<<16)
  while test_port(port) and port < max_val:
    port += 1
  if port == max_val:
    raise Failure('Having issues finding an available port')
  return port


class LocalGae(object):
  """Wraps up starting a GAE local instance for integration tests."""

  def __init__(self, base_dir=None):
    """base_dir defaults to .. from the file's directory."""
    # Paths
    self.base_dir = base_dir
    if not self.base_dir:
      self.base_dir = os.path.dirname(os.path.abspath(__file__))
      self.base_dir = os.path.realpath(os.path.join(self.base_dir, '..'))
    self.test_server = None
    self.port = None
    self.app_id = None
    self.url = None
    self.tmp_db = None

  def install_prerequisites(self):
    # Load GAE SDK.
    _load_modules()

    # Now safe to import GAE SDK modules.
    # Unable to import 'yaml'
    # pylint: disable=F0401

    import yaml
    self.app_id = yaml.load(
        open(os.path.join(self.base_dir, 'app.yaml'), 'r'))['application']
    logging.debug('Instance app id: %s' % self.app_id)
    assert self.app_id

  def start_server(self, verbose=False):
    self.install_prerequisites()
    self.port = find_free_port()
    if verbose:
      stdout = None
      stderr = None
    else:
      stdout = subprocess.PIPE
      stderr = subprocess.PIPE
    # Generate a friendly environment.
    env = os.environ.copy()
    env['LANGUAGE'] = 'en'
    self.tmp_db = tempfile.NamedTemporaryFile()
    cmd = [
        os.path.join(GAE_SDK, 'dev_appserver.py'),
        self.base_dir,
        '--port', str(self.port),
        '--datastore_path', self.tmp_db.name,
        '-c',
        '--skip_sdk_update_check',
    ]
    if verbose:
      cmd.extend(['-a', '0.0.0.0'])
      cmd.append('--debug')
    self.test_server = subprocess.Popen(
        cmd, stdout=stdout, stderr=stderr, env=env)
    # Loop until port 127.0.0.1:port opens or the process dies.
    while not test_port(self.port):
      self.test_server.poll()
      if self.test_server.returncode is not None:
        raise Failure(
            'Test GAE instance failed early on port %s' %
            self.port)
      time.sleep(0.001)
    self.url = 'http://localhost:%d/' % self.port

  def stop_server(self):
    if self.test_server:
      # pylint: disable=E1101
      if hasattr(self.test_server, 'kill'):
        self.test_server.kill()
      else:
        os.kill(self.test_server.pid, signal.SIGKILL)
      self.test_server = None
      self.port = None
      self.url = None
      self.tmp_db.close()
      self.tmp_db = None

  def get(self, suburl):
    return urllib.urlopen(self.url + suburl).read()

  def post(self, suburl, data):
    return urllib.urlopen(self.url + suburl, urllib.urlencode(data)).read()

  def query(self, cmd):
    """Lame way to modify the db remotely on dev server.

    Using remote_api inside the unit test is a bit too invasive.
    """
    data = {
        'code': 'from google.appengine.ext import db\n' + cmd,
    }
    result = self.post('_ah/admin/interactive/execute', data)
    match = re.search(
        re.escape(r'<pre id="output">') + r'(.*?)' +
        re.escape('</pre>\n</body>\n</html>\n'),
        result,
        re.DOTALL)
    return match.group(1)

# vim: ts=2:sw=2:tw=80:et:
