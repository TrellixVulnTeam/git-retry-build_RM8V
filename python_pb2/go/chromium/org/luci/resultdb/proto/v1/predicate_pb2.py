# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: go.chromium.org/luci/resultdb/proto/v1/predicate.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from go.chromium.org.luci.resultdb.proto.v1 import common_pb2 as go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='go.chromium.org/luci/resultdb/proto/v1/predicate.proto',
  package='luci.resultdb.v1',
  syntax='proto3',
  serialized_options=b'Z/go.chromium.org/luci/resultdb/proto/v1;resultpb',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n6go.chromium.org/luci/resultdb/proto/v1/predicate.proto\x12\x10luci.resultdb.v1\x1a\x33go.chromium.org/luci/resultdb/proto/v1/common.proto\"\xac\x02\n\x13TestResultPredicate\x12\x16\n\x0etest_id_regexp\x18\x01 \x01(\t\x12\x33\n\x07variant\x18\x02 \x01(\x0b\x32\".luci.resultdb.v1.VariantPredicate\x12\x44\n\nexpectancy\x18\x03 \x01(\x0e\x32\x30.luci.resultdb.v1.TestResultPredicate.Expectancy\x12\x1a\n\x12\x65xclude_exonerated\x18\x04 \x01(\x08\"f\n\nExpectancy\x12\x07\n\x03\x41LL\x10\x00\x12$\n VARIANTS_WITH_UNEXPECTED_RESULTS\x10\x01\x12)\n%VARIANTS_WITH_ONLY_UNEXPECTED_RESULTS\x10\x02\"g\n\x18TestExonerationPredicate\x12\x16\n\x0etest_id_regexp\x18\x01 \x01(\t\x12\x33\n\x07variant\x18\x02 \x01(\x0b\x32\".luci.resultdb.v1.VariantPredicate\"{\n\x10VariantPredicate\x12+\n\x06\x65quals\x18\x01 \x01(\x0b\x32\x19.luci.resultdb.v1.VariantH\x00\x12-\n\x08\x63ontains\x18\x02 \x01(\x0b\x32\x19.luci.resultdb.v1.VariantH\x00\x42\x0b\n\tpredicate\"\x80\x02\n\x11\x41rtifactPredicate\x12\x45\n\x0c\x66ollow_edges\x18\x01 \x01(\x0b\x32/.luci.resultdb.v1.ArtifactPredicate.EdgeTypeSet\x12\x44\n\x15test_result_predicate\x18\x02 \x01(\x0b\x32%.luci.resultdb.v1.TestResultPredicate\x12\x1b\n\x13\x63ontent_type_regexp\x18\x03 \x01(\t\x1a\x41\n\x0b\x45\x64geTypeSet\x12\x1c\n\x14included_invocations\x18\x01 \x01(\x08\x12\x14\n\x0ctest_results\x18\x02 \x01(\x08\x42\x31Z/go.chromium.org/luci/resultdb/proto/v1;resultpbb\x06proto3'
  ,
  dependencies=[go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_common__pb2.DESCRIPTOR,])



_TESTRESULTPREDICATE_EXPECTANCY = _descriptor.EnumDescriptor(
  name='Expectancy',
  full_name='luci.resultdb.v1.TestResultPredicate.Expectancy',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ALL', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='VARIANTS_WITH_UNEXPECTED_RESULTS', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='VARIANTS_WITH_ONLY_UNEXPECTED_RESULTS', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=328,
  serialized_end=430,
)
_sym_db.RegisterEnumDescriptor(_TESTRESULTPREDICATE_EXPECTANCY)


_TESTRESULTPREDICATE = _descriptor.Descriptor(
  name='TestResultPredicate',
  full_name='luci.resultdb.v1.TestResultPredicate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='test_id_regexp', full_name='luci.resultdb.v1.TestResultPredicate.test_id_regexp', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='variant', full_name='luci.resultdb.v1.TestResultPredicate.variant', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='expectancy', full_name='luci.resultdb.v1.TestResultPredicate.expectancy', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='exclude_exonerated', full_name='luci.resultdb.v1.TestResultPredicate.exclude_exonerated', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _TESTRESULTPREDICATE_EXPECTANCY,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=130,
  serialized_end=430,
)


_TESTEXONERATIONPREDICATE = _descriptor.Descriptor(
  name='TestExonerationPredicate',
  full_name='luci.resultdb.v1.TestExonerationPredicate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='test_id_regexp', full_name='luci.resultdb.v1.TestExonerationPredicate.test_id_regexp', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='variant', full_name='luci.resultdb.v1.TestExonerationPredicate.variant', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=432,
  serialized_end=535,
)


_VARIANTPREDICATE = _descriptor.Descriptor(
  name='VariantPredicate',
  full_name='luci.resultdb.v1.VariantPredicate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='equals', full_name='luci.resultdb.v1.VariantPredicate.equals', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='contains', full_name='luci.resultdb.v1.VariantPredicate.contains', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
    _descriptor.OneofDescriptor(
      name='predicate', full_name='luci.resultdb.v1.VariantPredicate.predicate',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=537,
  serialized_end=660,
)


_ARTIFACTPREDICATE_EDGETYPESET = _descriptor.Descriptor(
  name='EdgeTypeSet',
  full_name='luci.resultdb.v1.ArtifactPredicate.EdgeTypeSet',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='included_invocations', full_name='luci.resultdb.v1.ArtifactPredicate.EdgeTypeSet.included_invocations', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='test_results', full_name='luci.resultdb.v1.ArtifactPredicate.EdgeTypeSet.test_results', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=854,
  serialized_end=919,
)

_ARTIFACTPREDICATE = _descriptor.Descriptor(
  name='ArtifactPredicate',
  full_name='luci.resultdb.v1.ArtifactPredicate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='follow_edges', full_name='luci.resultdb.v1.ArtifactPredicate.follow_edges', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='test_result_predicate', full_name='luci.resultdb.v1.ArtifactPredicate.test_result_predicate', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='content_type_regexp', full_name='luci.resultdb.v1.ArtifactPredicate.content_type_regexp', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_ARTIFACTPREDICATE_EDGETYPESET, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=663,
  serialized_end=919,
)

_TESTRESULTPREDICATE.fields_by_name['variant'].message_type = _VARIANTPREDICATE
_TESTRESULTPREDICATE.fields_by_name['expectancy'].enum_type = _TESTRESULTPREDICATE_EXPECTANCY
_TESTRESULTPREDICATE_EXPECTANCY.containing_type = _TESTRESULTPREDICATE
_TESTEXONERATIONPREDICATE.fields_by_name['variant'].message_type = _VARIANTPREDICATE
_VARIANTPREDICATE.fields_by_name['equals'].message_type = go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_common__pb2._VARIANT
_VARIANTPREDICATE.fields_by_name['contains'].message_type = go_dot_chromium_dot_org_dot_luci_dot_resultdb_dot_proto_dot_v1_dot_common__pb2._VARIANT
_VARIANTPREDICATE.oneofs_by_name['predicate'].fields.append(
  _VARIANTPREDICATE.fields_by_name['equals'])
_VARIANTPREDICATE.fields_by_name['equals'].containing_oneof = _VARIANTPREDICATE.oneofs_by_name['predicate']
_VARIANTPREDICATE.oneofs_by_name['predicate'].fields.append(
  _VARIANTPREDICATE.fields_by_name['contains'])
_VARIANTPREDICATE.fields_by_name['contains'].containing_oneof = _VARIANTPREDICATE.oneofs_by_name['predicate']
_ARTIFACTPREDICATE_EDGETYPESET.containing_type = _ARTIFACTPREDICATE
_ARTIFACTPREDICATE.fields_by_name['follow_edges'].message_type = _ARTIFACTPREDICATE_EDGETYPESET
_ARTIFACTPREDICATE.fields_by_name['test_result_predicate'].message_type = _TESTRESULTPREDICATE
DESCRIPTOR.message_types_by_name['TestResultPredicate'] = _TESTRESULTPREDICATE
DESCRIPTOR.message_types_by_name['TestExonerationPredicate'] = _TESTEXONERATIONPREDICATE
DESCRIPTOR.message_types_by_name['VariantPredicate'] = _VARIANTPREDICATE
DESCRIPTOR.message_types_by_name['ArtifactPredicate'] = _ARTIFACTPREDICATE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TestResultPredicate = _reflection.GeneratedProtocolMessageType('TestResultPredicate', (_message.Message,), {
  'DESCRIPTOR' : _TESTRESULTPREDICATE,
  '__module__' : 'go.chromium.org.luci.resultdb.proto.v1.predicate_pb2'
  # @@protoc_insertion_point(class_scope:luci.resultdb.v1.TestResultPredicate)
  })
_sym_db.RegisterMessage(TestResultPredicate)

TestExonerationPredicate = _reflection.GeneratedProtocolMessageType('TestExonerationPredicate', (_message.Message,), {
  'DESCRIPTOR' : _TESTEXONERATIONPREDICATE,
  '__module__' : 'go.chromium.org.luci.resultdb.proto.v1.predicate_pb2'
  # @@protoc_insertion_point(class_scope:luci.resultdb.v1.TestExonerationPredicate)
  })
_sym_db.RegisterMessage(TestExonerationPredicate)

VariantPredicate = _reflection.GeneratedProtocolMessageType('VariantPredicate', (_message.Message,), {
  'DESCRIPTOR' : _VARIANTPREDICATE,
  '__module__' : 'go.chromium.org.luci.resultdb.proto.v1.predicate_pb2'
  # @@protoc_insertion_point(class_scope:luci.resultdb.v1.VariantPredicate)
  })
_sym_db.RegisterMessage(VariantPredicate)

ArtifactPredicate = _reflection.GeneratedProtocolMessageType('ArtifactPredicate', (_message.Message,), {

  'EdgeTypeSet' : _reflection.GeneratedProtocolMessageType('EdgeTypeSet', (_message.Message,), {
    'DESCRIPTOR' : _ARTIFACTPREDICATE_EDGETYPESET,
    '__module__' : 'go.chromium.org.luci.resultdb.proto.v1.predicate_pb2'
    # @@protoc_insertion_point(class_scope:luci.resultdb.v1.ArtifactPredicate.EdgeTypeSet)
    })
  ,
  'DESCRIPTOR' : _ARTIFACTPREDICATE,
  '__module__' : 'go.chromium.org.luci.resultdb.proto.v1.predicate_pb2'
  # @@protoc_insertion_point(class_scope:luci.resultdb.v1.ArtifactPredicate)
  })
_sym_db.RegisterMessage(ArtifactPredicate)
_sym_db.RegisterMessage(ArtifactPredicate.EdgeTypeSet)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
