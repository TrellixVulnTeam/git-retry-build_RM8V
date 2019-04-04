# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: build.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
import common_pb2 as common__pb2
import step_pb2 as step__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='build.proto',
  package='buildbucket.v2',
  syntax='proto3',
  serialized_options=_b('Z4go.chromium.org/luci/buildbucket/proto;buildbucketpb'),
  serialized_pb=_b('\n\x0b\x62uild.proto\x12\x0e\x62uildbucket.v2\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x0c\x63ommon.proto\x1a\nstep.proto\"\xa6\x08\n\x05\x42uild\x12\n\n\x02id\x18\x01 \x01(\x03\x12*\n\x07\x62uilder\x18\x02 \x01(\x0b\x32\x19.buildbucket.v2.BuilderID\x12\x0e\n\x06number\x18\x03 \x01(\x05\x12\x12\n\ncreated_by\x18\x04 \x01(\t\x12\x13\n\x0b\x63\x61nceled_by\x18\x17 \x01(\t\x12/\n\x0b\x63reate_time\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nstart_time\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08\x65nd_time\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0bupdate_time\x18\t \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12&\n\x06status\x18\x0c \x01(\x0e\x32\x16.buildbucket.v2.Status\x12\x18\n\x10summary_markdown\x18\x14 \x01(\t\x12)\n\x08\x63ritical\x18\x15 \x01(\x0e\x32\x17.buildbucket.v2.Trinary\x12\x35\n\x0estatus_details\x18\x16 \x01(\x0b\x32\x1d.buildbucket.v2.StatusDetails\x12*\n\x05input\x18\x0f \x01(\x0b\x32\x1b.buildbucket.v2.Build.Input\x12,\n\x06output\x18\x10 \x01(\x0b\x32\x1c.buildbucket.v2.Build.Output\x12#\n\x05steps\x18\x11 \x03(\x0b\x32\x14.buildbucket.v2.Step\x12)\n\x05infra\x18\x12 \x01(\x0b\x32\x1a.buildbucket.v2.BuildInfra\x12(\n\x04tags\x18\x13 \x03(\x0b\x32\x1a.buildbucket.v2.StringPair\x12.\n\nexecutable\x18\x18 \x01(\x0b\x32\x1a.buildbucket.v2.Executable\x1a\xb7\x01\n\x05Input\x12+\n\nproperties\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x35\n\x0egitiles_commit\x18\x02 \x01(\x0b\x32\x1d.buildbucket.v2.GitilesCommit\x12\x34\n\x0egerrit_changes\x18\x03 \x03(\x0b\x32\x1c.buildbucket.v2.GerritChange\x12\x14\n\x0c\x65xperimental\x18\x05 \x01(\x08\x1ax\n\x06Output\x12+\n\nproperties\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x35\n\x0egitiles_commit\x18\x03 \x01(\x0b\x32\x1d.buildbucket.v2.GitilesCommitJ\x04\x08\x02\x10\x03J\x04\x08\x04\x10\x05J\x04\x08\x05\x10\x06J\x04\x08\r\x10\x0eJ\x04\x08\x0e\x10\x0f\"\xdc\x05\n\nBuildInfra\x12;\n\x0b\x62uildbucket\x18\x01 \x01(\x0b\x32&.buildbucket.v2.BuildInfra.Buildbucket\x12\x35\n\x08swarming\x18\x02 \x01(\x0b\x32#.buildbucket.v2.BuildInfra.Swarming\x12\x31\n\x06logdog\x18\x03 \x01(\x0b\x32!.buildbucket.v2.BuildInfra.LogDog\x12\x31\n\x06recipe\x18\x04 \x01(\x0b\x32!.buildbucket.v2.BuildInfra.Recipe\x1a\xb7\x01\n\x0b\x42uildbucket\x12\x1f\n\x17service_config_revision\x18\x02 \x01(\t\x12\x0e\n\x06\x63\x61nary\x18\x04 \x01(\x08\x12\x35\n\x14requested_properties\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12@\n\x14requested_dimensions\x18\x06 \x03(\x0b\x32\".buildbucket.v2.RequestedDimension\x1a\xce\x01\n\x08Swarming\x12\x10\n\x08hostname\x18\x01 \x01(\t\x12\x0f\n\x07task_id\x18\x02 \x01(\t\x12\x1c\n\x14task_service_account\x18\x03 \x01(\t\x12\x10\n\x08priority\x18\x04 \x01(\x05\x12;\n\x0ftask_dimensions\x18\x05 \x03(\x0b\x32\".buildbucket.v2.RequestedDimension\x12\x32\n\x0e\x62ot_dimensions\x18\x06 \x03(\x0b\x32\x1a.buildbucket.v2.StringPair\x1a;\n\x06LogDog\x12\x10\n\x08hostname\x18\x01 \x01(\t\x12\x0f\n\x07project\x18\x02 \x01(\t\x12\x0e\n\x06prefix\x18\x03 \x01(\t\x1a,\n\x06Recipe\x12\x14\n\x0c\x63ipd_package\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"=\n\tBuilderID\x12\x0f\n\x07project\x18\x01 \x01(\t\x12\x0e\n\x06\x62ucket\x18\x02 \x01(\t\x12\x0f\n\x07\x62uilder\x18\x03 \x01(\tB6Z4go.chromium.org/luci/buildbucket/proto;buildbucketpbb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,google_dot_protobuf_dot_struct__pb2.DESCRIPTOR,common__pb2.DESCRIPTOR,step__pb2.DESCRIPTOR,])




_BUILD_INPUT = _descriptor.Descriptor(
  name='Input',
  full_name='buildbucket.v2.Build.Input',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='properties', full_name='buildbucket.v2.Build.Input.properties', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='gitiles_commit', full_name='buildbucket.v2.Build.Input.gitiles_commit', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='gerrit_changes', full_name='buildbucket.v2.Build.Input.gerrit_changes', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='experimental', full_name='buildbucket.v2.Build.Input.experimental', index=3,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=860,
  serialized_end=1043,
)

_BUILD_OUTPUT = _descriptor.Descriptor(
  name='Output',
  full_name='buildbucket.v2.Build.Output',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='properties', full_name='buildbucket.v2.Build.Output.properties', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='gitiles_commit', full_name='buildbucket.v2.Build.Output.gitiles_commit', index=1,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1045,
  serialized_end=1165,
)

_BUILD = _descriptor.Descriptor(
  name='Build',
  full_name='buildbucket.v2.Build',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='buildbucket.v2.Build.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='builder', full_name='buildbucket.v2.Build.builder', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='number', full_name='buildbucket.v2.Build.number', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_by', full_name='buildbucket.v2.Build.created_by', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='canceled_by', full_name='buildbucket.v2.Build.canceled_by', index=4,
      number=23, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='create_time', full_name='buildbucket.v2.Build.create_time', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='buildbucket.v2.Build.start_time', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='buildbucket.v2.Build.end_time', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_time', full_name='buildbucket.v2.Build.update_time', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status', full_name='buildbucket.v2.Build.status', index=9,
      number=12, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='summary_markdown', full_name='buildbucket.v2.Build.summary_markdown', index=10,
      number=20, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='critical', full_name='buildbucket.v2.Build.critical', index=11,
      number=21, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status_details', full_name='buildbucket.v2.Build.status_details', index=12,
      number=22, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='input', full_name='buildbucket.v2.Build.input', index=13,
      number=15, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='output', full_name='buildbucket.v2.Build.output', index=14,
      number=16, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='steps', full_name='buildbucket.v2.Build.steps', index=15,
      number=17, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='infra', full_name='buildbucket.v2.Build.infra', index=16,
      number=18, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tags', full_name='buildbucket.v2.Build.tags', index=17,
      number=19, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='executable', full_name='buildbucket.v2.Build.executable', index=18,
      number=24, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_BUILD_INPUT, _BUILD_OUTPUT, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=121,
  serialized_end=1183,
)


_BUILDINFRA_BUILDBUCKET = _descriptor.Descriptor(
  name='Buildbucket',
  full_name='buildbucket.v2.BuildInfra.Buildbucket',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='service_config_revision', full_name='buildbucket.v2.BuildInfra.Buildbucket.service_config_revision', index=0,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='canary', full_name='buildbucket.v2.BuildInfra.Buildbucket.canary', index=1,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='requested_properties', full_name='buildbucket.v2.BuildInfra.Buildbucket.requested_properties', index=2,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='requested_dimensions', full_name='buildbucket.v2.BuildInfra.Buildbucket.requested_dimensions', index=3,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1419,
  serialized_end=1602,
)

_BUILDINFRA_SWARMING = _descriptor.Descriptor(
  name='Swarming',
  full_name='buildbucket.v2.BuildInfra.Swarming',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hostname', full_name='buildbucket.v2.BuildInfra.Swarming.hostname', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='task_id', full_name='buildbucket.v2.BuildInfra.Swarming.task_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='task_service_account', full_name='buildbucket.v2.BuildInfra.Swarming.task_service_account', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='priority', full_name='buildbucket.v2.BuildInfra.Swarming.priority', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='task_dimensions', full_name='buildbucket.v2.BuildInfra.Swarming.task_dimensions', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bot_dimensions', full_name='buildbucket.v2.BuildInfra.Swarming.bot_dimensions', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1605,
  serialized_end=1811,
)

_BUILDINFRA_LOGDOG = _descriptor.Descriptor(
  name='LogDog',
  full_name='buildbucket.v2.BuildInfra.LogDog',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hostname', full_name='buildbucket.v2.BuildInfra.LogDog.hostname', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project', full_name='buildbucket.v2.BuildInfra.LogDog.project', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='prefix', full_name='buildbucket.v2.BuildInfra.LogDog.prefix', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1813,
  serialized_end=1872,
)

_BUILDINFRA_RECIPE = _descriptor.Descriptor(
  name='Recipe',
  full_name='buildbucket.v2.BuildInfra.Recipe',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cipd_package', full_name='buildbucket.v2.BuildInfra.Recipe.cipd_package', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='buildbucket.v2.BuildInfra.Recipe.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1874,
  serialized_end=1918,
)

_BUILDINFRA = _descriptor.Descriptor(
  name='BuildInfra',
  full_name='buildbucket.v2.BuildInfra',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='buildbucket', full_name='buildbucket.v2.BuildInfra.buildbucket', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='swarming', full_name='buildbucket.v2.BuildInfra.swarming', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='logdog', full_name='buildbucket.v2.BuildInfra.logdog', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='recipe', full_name='buildbucket.v2.BuildInfra.recipe', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_BUILDINFRA_BUILDBUCKET, _BUILDINFRA_SWARMING, _BUILDINFRA_LOGDOG, _BUILDINFRA_RECIPE, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1186,
  serialized_end=1918,
)


_BUILDERID = _descriptor.Descriptor(
  name='BuilderID',
  full_name='buildbucket.v2.BuilderID',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project', full_name='buildbucket.v2.BuilderID.project', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bucket', full_name='buildbucket.v2.BuilderID.bucket', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='builder', full_name='buildbucket.v2.BuilderID.builder', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1920,
  serialized_end=1981,
)

_BUILD_INPUT.fields_by_name['properties'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_BUILD_INPUT.fields_by_name['gitiles_commit'].message_type = common__pb2._GITILESCOMMIT
_BUILD_INPUT.fields_by_name['gerrit_changes'].message_type = common__pb2._GERRITCHANGE
_BUILD_INPUT.containing_type = _BUILD
_BUILD_OUTPUT.fields_by_name['properties'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_BUILD_OUTPUT.fields_by_name['gitiles_commit'].message_type = common__pb2._GITILESCOMMIT
_BUILD_OUTPUT.containing_type = _BUILD
_BUILD.fields_by_name['builder'].message_type = _BUILDERID
_BUILD.fields_by_name['create_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_BUILD.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_BUILD.fields_by_name['end_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_BUILD.fields_by_name['update_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_BUILD.fields_by_name['status'].enum_type = common__pb2._STATUS
_BUILD.fields_by_name['critical'].enum_type = common__pb2._TRINARY
_BUILD.fields_by_name['status_details'].message_type = common__pb2._STATUSDETAILS
_BUILD.fields_by_name['input'].message_type = _BUILD_INPUT
_BUILD.fields_by_name['output'].message_type = _BUILD_OUTPUT
_BUILD.fields_by_name['steps'].message_type = step__pb2._STEP
_BUILD.fields_by_name['infra'].message_type = _BUILDINFRA
_BUILD.fields_by_name['tags'].message_type = common__pb2._STRINGPAIR
_BUILD.fields_by_name['executable'].message_type = common__pb2._EXECUTABLE
_BUILDINFRA_BUILDBUCKET.fields_by_name['requested_properties'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_BUILDINFRA_BUILDBUCKET.fields_by_name['requested_dimensions'].message_type = common__pb2._REQUESTEDDIMENSION
_BUILDINFRA_BUILDBUCKET.containing_type = _BUILDINFRA
_BUILDINFRA_SWARMING.fields_by_name['task_dimensions'].message_type = common__pb2._REQUESTEDDIMENSION
_BUILDINFRA_SWARMING.fields_by_name['bot_dimensions'].message_type = common__pb2._STRINGPAIR
_BUILDINFRA_SWARMING.containing_type = _BUILDINFRA
_BUILDINFRA_LOGDOG.containing_type = _BUILDINFRA
_BUILDINFRA_RECIPE.containing_type = _BUILDINFRA
_BUILDINFRA.fields_by_name['buildbucket'].message_type = _BUILDINFRA_BUILDBUCKET
_BUILDINFRA.fields_by_name['swarming'].message_type = _BUILDINFRA_SWARMING
_BUILDINFRA.fields_by_name['logdog'].message_type = _BUILDINFRA_LOGDOG
_BUILDINFRA.fields_by_name['recipe'].message_type = _BUILDINFRA_RECIPE
DESCRIPTOR.message_types_by_name['Build'] = _BUILD
DESCRIPTOR.message_types_by_name['BuildInfra'] = _BUILDINFRA
DESCRIPTOR.message_types_by_name['BuilderID'] = _BUILDERID
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Build = _reflection.GeneratedProtocolMessageType('Build', (_message.Message,), dict(

  Input = _reflection.GeneratedProtocolMessageType('Input', (_message.Message,), dict(
    DESCRIPTOR = _BUILD_INPUT,
    __module__ = 'build_pb2'
    # @@protoc_insertion_point(class_scope:buildbucket.v2.Build.Input)
    ))
  ,

  Output = _reflection.GeneratedProtocolMessageType('Output', (_message.Message,), dict(
    DESCRIPTOR = _BUILD_OUTPUT,
    __module__ = 'build_pb2'
    # @@protoc_insertion_point(class_scope:buildbucket.v2.Build.Output)
    ))
  ,
  DESCRIPTOR = _BUILD,
  __module__ = 'build_pb2'
  # @@protoc_insertion_point(class_scope:buildbucket.v2.Build)
  ))
_sym_db.RegisterMessage(Build)
_sym_db.RegisterMessage(Build.Input)
_sym_db.RegisterMessage(Build.Output)

BuildInfra = _reflection.GeneratedProtocolMessageType('BuildInfra', (_message.Message,), dict(

  Buildbucket = _reflection.GeneratedProtocolMessageType('Buildbucket', (_message.Message,), dict(
    DESCRIPTOR = _BUILDINFRA_BUILDBUCKET,
    __module__ = 'build_pb2'
    # @@protoc_insertion_point(class_scope:buildbucket.v2.BuildInfra.Buildbucket)
    ))
  ,

  Swarming = _reflection.GeneratedProtocolMessageType('Swarming', (_message.Message,), dict(
    DESCRIPTOR = _BUILDINFRA_SWARMING,
    __module__ = 'build_pb2'
    # @@protoc_insertion_point(class_scope:buildbucket.v2.BuildInfra.Swarming)
    ))
  ,

  LogDog = _reflection.GeneratedProtocolMessageType('LogDog', (_message.Message,), dict(
    DESCRIPTOR = _BUILDINFRA_LOGDOG,
    __module__ = 'build_pb2'
    # @@protoc_insertion_point(class_scope:buildbucket.v2.BuildInfra.LogDog)
    ))
  ,

  Recipe = _reflection.GeneratedProtocolMessageType('Recipe', (_message.Message,), dict(
    DESCRIPTOR = _BUILDINFRA_RECIPE,
    __module__ = 'build_pb2'
    # @@protoc_insertion_point(class_scope:buildbucket.v2.BuildInfra.Recipe)
    ))
  ,
  DESCRIPTOR = _BUILDINFRA,
  __module__ = 'build_pb2'
  # @@protoc_insertion_point(class_scope:buildbucket.v2.BuildInfra)
  ))
_sym_db.RegisterMessage(BuildInfra)
_sym_db.RegisterMessage(BuildInfra.Buildbucket)
_sym_db.RegisterMessage(BuildInfra.Swarming)
_sym_db.RegisterMessage(BuildInfra.LogDog)
_sym_db.RegisterMessage(BuildInfra.Recipe)

BuilderID = _reflection.GeneratedProtocolMessageType('BuilderID', (_message.Message,), dict(
  DESCRIPTOR = _BUILDERID,
  __module__ = 'build_pb2'
  # @@protoc_insertion_point(class_scope:buildbucket.v2.BuilderID)
  ))
_sym_db.RegisterMessage(BuilderID)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
