# Copyright 2014 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import datetime

from testing_utils import testing

from components import auth
from components import auth_testing
from components import utils

from cipd import acl
from cipd import api
from cipd import impl


class PackageRepositoryApiTest(testing.EndpointsTestCase):
  """Tests for API layer ONLY."""

  api_service_cls = api.PackageRepositoryApi

  def setUp(self):
    super(PackageRepositoryApiTest, self).setUp()
    auth_testing.mock_get_current_identity(self)
    auth_testing.mock_is_admin(self)
    self.repo_service = MockedRepoService()
    self.mock(impl, 'get_repo_service', lambda: self.repo_service)

  def test_fetch_package_ok(self):
    _, registered = self.repo_service.register_package(
        package_name='good/name',
        caller=auth.Identity.from_bytes('user:abc@example.com'),
        now=datetime.datetime(2014, 1, 1))
    self.assertTrue(registered)

    resp = self.call_api('fetch_package', {'package_name': 'good/name'})
    self.assertEqual({
      'package': {
        'package_name': 'good/name',
        'registered_by': 'user:abc@example.com',
        'registered_ts': '1388534400000000',
      },
      'status': 'SUCCESS',
    }, resp.json_body)

  def test_fetch_package_no_access(self):
    _, registered = self.repo_service.register_package(
        package_name='good/name',
        caller=auth.Identity.from_bytes('user:abc@example.com'),
        now=datetime.datetime(2014, 1, 1))
    self.assertTrue(registered)

    # Should return PACKAGE_NOT_FOUND even though package exists.
    self.mock(api.acl, 'can_fetch_package', lambda *_: False)
    resp = self.call_api('fetch_package', {'package_name': 'good/name'})
    self.assertEqual({'status': 'PACKAGE_NOT_FOUND'}, resp.json_body)

  def test_fetch_package_no_such_package(self):
    resp = self.call_api('fetch_package', {'package_name': 'good/name'})
    self.assertEqual({'status': 'PACKAGE_NOT_FOUND'}, resp.json_body)

  def test_fetch_package_bad_name(self):
    resp = self.call_api('fetch_package', {'package_name': 'bad name'})
    self.assertEqual({
      'status': 'ERROR',
      'error_message': 'Invalid package name',
    }, resp.json_body)

  def test_register_package(self):
    self.mock(utils, 'utcnow', lambda: datetime.datetime(2014, 1, 1))

    # Not yet registered.
    resp = self.call_api('register_package', {'package_name': 'good/name'})
    self.assertEqual(200, resp.status_code)
    self.assertEqual({
      'status': 'REGISTERED',
      'package': {
        'package_name': 'good/name',
        'registered_by': 'user:mocked@example.com',
        'registered_ts': '1388534400000000',
      },
    }, resp.json_body)

    # Check that it is indeed there.
    pkg = self.repo_service.get_package('good/name')
    self.assertTrue(pkg)
    expected = {
      'registered_by': auth.Identity(kind='user', name='mocked@example.com'),
      'registered_ts': datetime.datetime(2014, 1, 1, 0, 0),
    }
    self.assertEqual(expected, pkg.to_dict())

    # Attempt to register it again.
    resp = self.call_api('register_package', {'package_name': 'good/name'})
    self.assertEqual(200, resp.status_code)
    self.assertEqual({
      'status': 'ALREADY_REGISTERED',
      'package': {
        'package_name': 'good/name',
        'registered_by': 'user:mocked@example.com',
        'registered_ts': '1388534400000000',
      },
    }, resp.json_body)

  def test_register_package_bad_name(self):
    resp = self.call_api('register_package', {'package_name': 'bad name'})
    self.assertEqual({
      'status': 'ERROR',
      'error_message': 'Invalid package name',
    }, resp.json_body)

  def test_register_package_no_access(self):
    self.mock(api.acl, 'can_register_package', lambda *_: False)
    with self.call_should_fail(403):
      self.call_api('register_package', {'package_name': 'good/name'})

  def test_fetch_instance_ok(self):
    _, registered = self.repo_service.register_instance(
        package_name='good/name',
        instance_id='a'*40,
        caller=auth.Identity.from_bytes('user:abc@example.com'),
        now=datetime.datetime(2014, 1, 1))
    self.assertTrue(registered)

    resp = self.call_api('fetch_instance', {
      'package_name': 'good/name',
      'instance_id': 'a'*40,
    })
    self.assertEqual({
      'fetch_url': 'http://signed-url/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
      'instance': {
        'instance_id': 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
        'package_name': 'good/name',
        'registered_by': 'user:abc@example.com',
        'registered_ts': '1388534400000000',
      },
      'status': 'SUCCESS',
    }, resp.json_body)

  def test_fetch_instance_no_access(self):
    _, registered = self.repo_service.register_instance(
        package_name='good/name',
        instance_id='a'*40,
        caller=auth.Identity.from_bytes('user:abc@example.com'),
        now=datetime.datetime(2014, 1, 1))
    self.assertTrue(registered)

    # Should return PACKAGE_NOT_FOUND even though package exists.
    self.mock(api.acl, 'can_fetch_instance', lambda *_: False)
    resp = self.call_api('fetch_instance', {
      'package_name': 'good/name',
      'instance_id': 'a'*40,
    })
    self.assertEqual({'status': 'PACKAGE_NOT_FOUND'}, resp.json_body)

  def test_fetch_instance_no_such_package(self):
    resp = self.call_api('fetch_instance', {
      'package_name': 'good/name',
      'instance_id': 'a'*40,
    })
    self.assertEqual({'status': 'PACKAGE_NOT_FOUND'}, resp.json_body)

  def test_fetch_instance_no_such_instance(self):
    _, registered = self.repo_service.register_instance(
        package_name='good/name',
        instance_id='a'*40,
        caller=auth.Identity.from_bytes('user:abc@example.com'),
        now=datetime.datetime(2014, 1, 1))
    self.assertTrue(registered)

    resp = self.call_api('fetch_instance', {
      'package_name': 'good/name',
      'instance_id': 'b'*40,
    })
    self.assertEqual({'status': 'INSTANCE_NOT_FOUND'}, resp.json_body)

  def test_fetch_instance_bad_name(self):
    resp = self.call_api('fetch_instance', {
      'package_name': 'bad name',
      'instance_id': 'a'*40,
    })
    self.assertEqual({
      'status': 'ERROR',
      'error_message': 'Invalid package name',
    }, resp.json_body)

  def test_fetch_instance_bad_instance_id(self):
    resp = self.call_api('fetch_instance', {
      'package_name': 'good/name',
      'instance_id': 'bad instance id',
    })
    self.assertEqual({
      'status': 'ERROR',
      'error_message': 'Invalid package instance ID',
    }, resp.json_body)

  def test_fetch_instance_no_service(self):
    self.repo_service = None
    with self.call_should_fail(500):
      self.call_api('fetch_instance', {
        'package_name': 'good/name',
        'instance_id': 'a'*40,
      })

  def test_register_new_instance_flow(self):
    self.mock(utils, 'utcnow', lambda: datetime.datetime(2014, 1, 1))
    request = {
      'package_name': 'good/name',
      'instance_id': 'a'*40,
    }

    # Package is not uploaded yet. Should ask to upload.
    resp = self.call_api('register_instance', request)
    self.assertEqual(200, resp.status_code)
    self.assertEqual({
      'status': 'UPLOAD_FIRST',
      'upload_session_id': 'upload_session_id',
      'upload_url': 'http://upload_url',
    }, resp.json_body)

    # Pretend it is upload now.
    self.repo_service.uploaded.add('a'*40)

    # Should register the package.
    resp = self.call_api('register_instance', request)
    self.assertEqual(200, resp.status_code)
    self.assertEqual({
      'status': 'REGISTERED',
      'instance': {
        'instance_id': 'a'*40,
        'package_name': 'good/name',
        'registered_by': 'user:mocked@example.com',
        'registered_ts': '1388534400000000',
      },
    }, resp.json_body)

    # Check that it is indeed there.
    pkg = self.repo_service.get_instance('good/name', 'a'*40)
    self.assertTrue(pkg)
    expected = {
      'registered_by': auth.Identity(kind='user', name='mocked@example.com'),
      'registered_ts': datetime.datetime(2014, 1, 1, 0, 0),
    }
    self.assertEqual(expected, pkg.to_dict())

    # Attempt to register it again.
    resp = self.call_api('register_instance', request)
    self.assertEqual(200, resp.status_code)
    self.assertEqual({
      'status': 'ALREADY_REGISTERED',
      'instance': {
        'instance_id': 'a'*40,
        'package_name': 'good/name',
        'registered_by': 'user:mocked@example.com',
        'registered_ts': '1388534400000000',
      },
    }, resp.json_body)

  def test_register_instance_bad_name(self):
    resp = self.call_api('register_instance', {
      'package_name': 'bad name',
      'instance_id': 'a'*40,
    })
    self.assertEqual({
      'status': 'ERROR',
      'error_message': 'Invalid package name',
    }, resp.json_body)

  def test_register_instance_bad_instance_id(self):
    resp = self.call_api('register_instance', {
      'package_name': 'good/name',
      'instance_id': 'bad instance id',
    })
    self.assertEqual({
      'status': 'ERROR',
      'error_message': 'Invalid package instance ID',
    }, resp.json_body)

  def test_register_instance_no_access(self):
    self.mock(api.acl, 'can_register_instance', lambda *_: False)
    with self.call_should_fail(403):
      self.call_api('register_instance', {
        'package_name': 'good/name',
        'instance_id': 'a'*40,
      })

  def test_register_instance_no_access_to_register_package(self):
    self.mock(api.acl, 'can_register_package', lambda *_: False)
    with self.call_should_fail(403):
      self.call_api('register_instance', {
        'package_name': 'good/name',
        'instance_id': 'a'*40,
      })

  def test_register_instance_no_service(self):
    self.repo_service = None
    with self.call_should_fail(500):
      self.call_api('register_instance', {
        'package_name': 'good/name',
        'instance_id': 'a'*40,
      })

  def test_fetch_acl_ok(self):
    acl.modify_roles(
        changes=[
          acl.RoleChange(
              package_path='a',
              revoke=False,
              role='OWNER',
              user=auth.Identity.from_bytes('user:xyz@example.com'),
              group=None),
          acl.RoleChange(
              package_path='a/b/c',
              revoke=False,
              role='READER',
              user=None,
              group='reader-group'),
        ],
        caller=auth.Identity.from_bytes('user:abc@example.com'),
        now=datetime.datetime(2014, 1, 1))

    resp = self.call_api('fetch_acl', {'package_path': 'a/b/c/d'})
    self.assertEqual(200, resp.status_code)
    self.assertEqual({
      'status': 'SUCCESS',
      'acls': {
        'acls': [
          {
            'modified_by': 'user:abc@example.com',
            'modified_ts': '1388534400000000',
            'package_path': 'a',
            'principals': ['user:xyz@example.com'],
            'role': 'OWNER',
          },
          {
            'modified_by': 'user:abc@example.com',
            'modified_ts': '1388534400000000',
            'package_path': 'a/b/c',
            'principals': ['group:reader-group'],
            'role': 'READER',
          },
        ],
      },
    }, resp.json_body)

  def test_fetch_acl_missing(self):
    resp = self.call_api('fetch_acl', {'package_path': 'a/b/c/d'})
    self.assertEqual(200, resp.status_code)
    self.assertEqual({'status': 'SUCCESS', 'acls': {}}, resp.json_body)

  def test_fetch_acl_bad_package_name(self):
    resp = self.call_api('fetch_acl', {'package_path': 'bad name'})
    self.assertEqual(200, resp.status_code)
    self.assertEqual({
      'status': 'ERROR',
      'error_message': 'Invalid package path',
    }, resp.json_body)

  def test_fetch_acl_no_access(self):
    self.mock(api.acl, 'can_fetch_acl', lambda *_: False)
    resp = self.call_api('fetch_acl', {'package_path': 'a/b/c'})
    self.assertEqual(200, resp.status_code)
    self.assertEqual({'acls': {}, 'status': 'SUCCESS'}, resp.json_body)

  def test_modify_acl_ok(self):
    self.mock(utils, 'utcnow', lambda: datetime.datetime(2014, 1, 1))
    resp = self.call_api('modify_acl', {
      'package_path': 'a/b',
      'changes': [
        {
          'action': 'GRANT',
          'role': 'OWNER',
          'principal': 'user:abc@example.com',
        },
        {
          'action': 'GRANT',
          'role': 'READER',
          'principal': 'group:readers-group',
        },
        {
          'action': 'REVOKE',
          'role': 'WRITER',
          'principal': 'anonymous:anonymous',
        },
      ],
    })
    self.assertEqual(200, resp.status_code)
    self.assertEqual({'status': 'SUCCESS'}, resp.json_body)

    owner = acl.get_package_acls('a/b/c', 'OWNER')
    self.assertEqual(1, len(owner))
    self.assertEqual({
      'groups': [],
      'modified_by': auth.Identity(kind='user', name='mocked@example.com'),
      'modified_ts': datetime.datetime(2014, 1, 1, 0, 0),
      'rev': 1,
      'users': [auth.Identity(kind='user', name='abc@example.com')],
    }, owner[0].to_dict())

    reader = acl.get_package_acls('a/b/c', 'READER')
    self.assertEqual(1, len(reader))
    self.assertEqual({
      'groups': ['readers-group'],
      'modified_by': auth.Identity(kind='user', name='mocked@example.com'),
      'modified_ts': datetime.datetime(2014, 1, 1, 0, 0),
      'rev': 1,
      'users': [],
    }, reader[0].to_dict())

  def test_modify_acl_bad_role(self):
    resp = self.call_api('modify_acl', {
      'package_path': 'a/b',
      'changes': [
        {
          'action': 'GRANT',
          'role': 'UNKNOWN_ROLE',
          'principal': 'user:abc@example.com',
        },
      ],
    })
    self.assertEqual(200, resp.status_code)
    self.assertEqual({
      'status': 'ERROR',
      'error_message': 'Invalid role change request: Invalid role UNKNOWN_ROLE',
    }, resp.json_body)

  def test_modify_acl_bad_group_name(self):
    resp = self.call_api('modify_acl', {
      'package_path': 'a/b',
      'changes': [
        {
          'action': 'GRANT',
          'role': 'OWNER',
          'principal': 'group:bad/group/name',
        },
      ],
    })
    self.assertEqual(200, resp.status_code)
    self.assertEqual({
      'status': 'ERROR',
      'error_message': (
          'Invalid role change request: Invalid group name: "bad/group/name"'),
    }, resp.json_body)

  def test_modify_acl_bad_package_name(self):
    resp = self.call_api('modify_acl', {
      'package_path': 'bad name',
      'changes': [],
    })
    self.assertEqual(200, resp.status_code)
    self.assertEqual({
      'status': 'ERROR',
      'error_message': 'Invalid package path',
    }, resp.json_body)

  def test_modify_acl_no_access(self):
    self.mock(api.acl, 'can_modify_acl', lambda *_: False)
    with self.call_should_fail(403):
      self.call_api('modify_acl', {
        'package_path': 'a/b/c',
        'changes': [],
      })


class MockedRepoService(impl.RepoService):
  """Almost like a real one, except CAS part is stubbed."""

  def __init__(self):
    super(MockedRepoService, self).__init__(None)
    self.uploaded = set()

  def is_fetch_configured(self):
    return True

  def generate_fetch_url(self, instance):
    return 'http://signed-url/%s' % instance.instance_id

  def is_instance_file_uploaded(self, package_name, instance_id):
    return instance_id in self.uploaded

  def create_upload_session(self, package_name, instance_id, caller):
    return 'http://upload_url', 'upload_session_id'
