# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/api_proto/projects.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from api.api_proto import common_pb2 as api_dot_api__proto_dot_common__pb2
from api.api_proto import project_objects_pb2 as api_dot_api__proto_dot_project__objects__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='api/api_proto/projects.proto',
  package='monorail',
  syntax='proto3',
  serialized_pb=_b('\n\x1c\x61pi/api_proto/projects.proto\x12\x08monorail\x1a\x1a\x61pi/api_proto/common.proto\x1a#api/api_proto/project_objects.proto\"<\n\x13ListProjectsRequest\x12\x11\n\tpage_size\x18\x01 \x01(\x05\x12\x12\n\npage_token\x18\x02 \x01(\t\"T\n\x14ListProjectsResponse\x12#\n\x08projects\x18\x01 \x03(\x0b\x32\x11.monorail.Project\x12\x17\n\x0fnext_page_token\x18\x02 \x01(\t\"O\n\x10GetConfigRequest\x12%\n\x05trace\x18\x01 \x01(\x0b\x32\x16.monorail.RequestTrace\x12\x14\n\x0cproject_name\x18\x02 \x01(\t\"Z\n\x1bGetCustomPermissionsRequest\x12%\n\x05trace\x18\x01 \x01(\x0b\x32\x16.monorail.RequestTrace\x12\x14\n\x0cproject_name\x18\x02 \x01(\t\"3\n\x1cGetCustomPermissionsResponse\x12\x13\n\x0bpermissions\x18\x01 \x03(\t\"W\n\x18GetVisibleMembersRequest\x12%\n\x05trace\x18\x01 \x01(\x0b\x32\x16.monorail.RequestTrace\x12\x14\n\x0cproject_name\x18\x02 \x01(\t\"h\n\x19GetVisibleMembersResponse\x12$\n\tuser_refs\x18\x01 \x03(\x0b\x32\x11.monorail.UserRef\x12%\n\ngroup_refs\x18\x02 \x03(\x0b\x32\x11.monorail.UserRef2\xe1\x02\n\x08Projects\x12O\n\x0cListProjects\x12\x1d.monorail.ListProjectsRequest\x1a\x1e.monorail.ListProjectsResponse\"\x00\x12;\n\tGetConfig\x12\x1a.monorail.GetConfigRequest\x1a\x10.monorail.Config\"\x00\x12g\n\x14GetCustomPermissions\x12%.monorail.GetCustomPermissionsRequest\x1a&.monorail.GetCustomPermissionsResponse\"\x00\x12^\n\x11GetVisibleMembers\x12\".monorail.GetVisibleMembersRequest\x1a#.monorail.GetVisibleMembersResponse\"\x00\x62\x06proto3')
  ,
  dependencies=[api_dot_api__proto_dot_common__pb2.DESCRIPTOR,api_dot_api__proto_dot_project__objects__pb2.DESCRIPTOR,])
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_LISTPROJECTSREQUEST = _descriptor.Descriptor(
  name='ListProjectsRequest',
  full_name='monorail.ListProjectsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='page_size', full_name='monorail.ListProjectsRequest.page_size', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='page_token', full_name='monorail.ListProjectsRequest.page_token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=107,
  serialized_end=167,
)


_LISTPROJECTSRESPONSE = _descriptor.Descriptor(
  name='ListProjectsResponse',
  full_name='monorail.ListProjectsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='projects', full_name='monorail.ListProjectsResponse.projects', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='next_page_token', full_name='monorail.ListProjectsResponse.next_page_token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=169,
  serialized_end=253,
)


_GETCONFIGREQUEST = _descriptor.Descriptor(
  name='GetConfigRequest',
  full_name='monorail.GetConfigRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='trace', full_name='monorail.GetConfigRequest.trace', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='project_name', full_name='monorail.GetConfigRequest.project_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=255,
  serialized_end=334,
)


_GETCUSTOMPERMISSIONSREQUEST = _descriptor.Descriptor(
  name='GetCustomPermissionsRequest',
  full_name='monorail.GetCustomPermissionsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='trace', full_name='monorail.GetCustomPermissionsRequest.trace', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='project_name', full_name='monorail.GetCustomPermissionsRequest.project_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=336,
  serialized_end=426,
)


_GETCUSTOMPERMISSIONSRESPONSE = _descriptor.Descriptor(
  name='GetCustomPermissionsResponse',
  full_name='monorail.GetCustomPermissionsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='permissions', full_name='monorail.GetCustomPermissionsResponse.permissions', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=428,
  serialized_end=479,
)


_GETVISIBLEMEMBERSREQUEST = _descriptor.Descriptor(
  name='GetVisibleMembersRequest',
  full_name='monorail.GetVisibleMembersRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='trace', full_name='monorail.GetVisibleMembersRequest.trace', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='project_name', full_name='monorail.GetVisibleMembersRequest.project_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=481,
  serialized_end=568,
)


_GETVISIBLEMEMBERSRESPONSE = _descriptor.Descriptor(
  name='GetVisibleMembersResponse',
  full_name='monorail.GetVisibleMembersResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_refs', full_name='monorail.GetVisibleMembersResponse.user_refs', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='group_refs', full_name='monorail.GetVisibleMembersResponse.group_refs', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=570,
  serialized_end=674,
)

_LISTPROJECTSRESPONSE.fields_by_name['projects'].message_type = api_dot_api__proto_dot_project__objects__pb2._PROJECT
_GETCONFIGREQUEST.fields_by_name['trace'].message_type = api_dot_api__proto_dot_common__pb2._REQUESTTRACE
_GETCUSTOMPERMISSIONSREQUEST.fields_by_name['trace'].message_type = api_dot_api__proto_dot_common__pb2._REQUESTTRACE
_GETVISIBLEMEMBERSREQUEST.fields_by_name['trace'].message_type = api_dot_api__proto_dot_common__pb2._REQUESTTRACE
_GETVISIBLEMEMBERSRESPONSE.fields_by_name['user_refs'].message_type = api_dot_api__proto_dot_common__pb2._USERREF
_GETVISIBLEMEMBERSRESPONSE.fields_by_name['group_refs'].message_type = api_dot_api__proto_dot_common__pb2._USERREF
DESCRIPTOR.message_types_by_name['ListProjectsRequest'] = _LISTPROJECTSREQUEST
DESCRIPTOR.message_types_by_name['ListProjectsResponse'] = _LISTPROJECTSRESPONSE
DESCRIPTOR.message_types_by_name['GetConfigRequest'] = _GETCONFIGREQUEST
DESCRIPTOR.message_types_by_name['GetCustomPermissionsRequest'] = _GETCUSTOMPERMISSIONSREQUEST
DESCRIPTOR.message_types_by_name['GetCustomPermissionsResponse'] = _GETCUSTOMPERMISSIONSRESPONSE
DESCRIPTOR.message_types_by_name['GetVisibleMembersRequest'] = _GETVISIBLEMEMBERSREQUEST
DESCRIPTOR.message_types_by_name['GetVisibleMembersResponse'] = _GETVISIBLEMEMBERSRESPONSE

ListProjectsRequest = _reflection.GeneratedProtocolMessageType('ListProjectsRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTPROJECTSREQUEST,
  __module__ = 'api.api_proto.projects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.ListProjectsRequest)
  ))
_sym_db.RegisterMessage(ListProjectsRequest)

ListProjectsResponse = _reflection.GeneratedProtocolMessageType('ListProjectsResponse', (_message.Message,), dict(
  DESCRIPTOR = _LISTPROJECTSRESPONSE,
  __module__ = 'api.api_proto.projects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.ListProjectsResponse)
  ))
_sym_db.RegisterMessage(ListProjectsResponse)

GetConfigRequest = _reflection.GeneratedProtocolMessageType('GetConfigRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETCONFIGREQUEST,
  __module__ = 'api.api_proto.projects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.GetConfigRequest)
  ))
_sym_db.RegisterMessage(GetConfigRequest)

GetCustomPermissionsRequest = _reflection.GeneratedProtocolMessageType('GetCustomPermissionsRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETCUSTOMPERMISSIONSREQUEST,
  __module__ = 'api.api_proto.projects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.GetCustomPermissionsRequest)
  ))
_sym_db.RegisterMessage(GetCustomPermissionsRequest)

GetCustomPermissionsResponse = _reflection.GeneratedProtocolMessageType('GetCustomPermissionsResponse', (_message.Message,), dict(
  DESCRIPTOR = _GETCUSTOMPERMISSIONSRESPONSE,
  __module__ = 'api.api_proto.projects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.GetCustomPermissionsResponse)
  ))
_sym_db.RegisterMessage(GetCustomPermissionsResponse)

GetVisibleMembersRequest = _reflection.GeneratedProtocolMessageType('GetVisibleMembersRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETVISIBLEMEMBERSREQUEST,
  __module__ = 'api.api_proto.projects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.GetVisibleMembersRequest)
  ))
_sym_db.RegisterMessage(GetVisibleMembersRequest)

GetVisibleMembersResponse = _reflection.GeneratedProtocolMessageType('GetVisibleMembersResponse', (_message.Message,), dict(
  DESCRIPTOR = _GETVISIBLEMEMBERSRESPONSE,
  __module__ = 'api.api_proto.projects_pb2'
  # @@protoc_insertion_point(class_scope:monorail.GetVisibleMembersResponse)
  ))
_sym_db.RegisterMessage(GetVisibleMembersResponse)


# @@protoc_insertion_point(module_scope)
