# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: support_3pp/spec.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='support_3pp/spec.proto',
  package='recipe_modules.infra.support_3pp',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x16support_3pp/spec.proto\x12 recipe_modules.infra.support_3pp\"\xda\x08\n\x04Spec\x12=\n\x06\x63reate\x18\x01 \x03(\x0b\x32-.recipe_modules.infra.support_3pp.Spec.Create\x12=\n\x06upload\x18\x02 \x01(\x0b\x32-.recipe_modules.infra.support_3pp.Spec.Upload\x1a\xa2\x07\n\x06\x43reate\x12\x13\n\x0bplatform_re\x18\x01 \x01(\t\x12\x14\n\x0c\x65xperimental\x18\x02 \x01(\x08\x12\x13\n\x0bunsupported\x18\x03 \x01(\x08\x12\x44\n\x06source\x18\x04 \x01(\x0b\x32\x34.recipe_modules.infra.support_3pp.Spec.Create.Source\x12\x42\n\x05\x62uild\x18\x05 \x01(\x0b\x32\x33.recipe_modules.infra.support_3pp.Spec.Create.Build\x12\x46\n\x07package\x18\x06 \x01(\x0b\x32\x35.recipe_modules.infra.support_3pp.Spec.Create.Package\x12\x44\n\x06verify\x18\x07 \x01(\x0b\x32\x34.recipe_modules.infra.support_3pp.Spec.Create.Verify\x1a\xba\x02\n\x06Source\x12:\n\x03git\x18\x01 \x01(\x0b\x32+.recipe_modules.infra.support_3pp.GitSourceH\x00\x12<\n\x04\x63ipd\x18\x02 \x01(\x0b\x32,.recipe_modules.infra.support_3pp.CipdSourceH\x00\x12@\n\x06script\x18\x03 \x01(\x0b\x32..recipe_modules.infra.support_3pp.ScriptSourceH\x00\x12\x0e\n\x06subdir\x18\x04 \x01(\t\x12\x16\n\x0eunpack_archive\x18\x05 \x01(\x08\x12\x18\n\x10no_archive_prune\x18\x06 \x01(\x08\x12\x11\n\tpatch_dir\x18\x07 \x03(\t\x12\x15\n\rpatch_version\x18\x08 \x01(\tB\x08\n\x06method\x1aJ\n\x05\x42uild\x12\x0f\n\x07install\x18\x01 \x03(\t\x12\x0c\n\x04tool\x18\x02 \x03(\t\x12\x0b\n\x03\x64\x65p\x18\x03 \x03(\t\x12\x15\n\rno_docker_env\x18\x04 \x01(\x08\x1a\x9e\x01\n\x07Package\x12W\n\x0cinstall_mode\x18\x01 \x01(\x0e\x32\x41.recipe_modules.infra.support_3pp.Spec.Create.Package.InstallMode\x12\x14\n\x0cversion_file\x18\x02 \x01(\t\"$\n\x0bInstallMode\x12\x08\n\x04\x63opy\x10\x00\x12\x0b\n\x07symlink\x10\x01\x1a\x16\n\x06Verify\x12\x0c\n\x04test\x18\x01 \x03(\t\x1a/\n\x06Upload\x12\x12\n\npkg_prefix\x18\x01 \x01(\t\x12\x11\n\tuniversal\x18\x02 \x01(\x08\"D\n\tGitSource\x12\x0c\n\x04repo\x18\x01 \x01(\t\x12\x13\n\x0btag_pattern\x18\x02 \x01(\t\x12\x14\n\x0cversion_join\x18\x03 \x01(\t\"Q\n\nCipdSource\x12\x0b\n\x03pkg\x18\x01 \x01(\t\x12\x17\n\x0f\x64\x65\x66\x61ult_version\x18\x02 \x01(\t\x12\x1d\n\x15original_download_url\x18\x03 \x01(\t\"\x1c\n\x0cScriptSource\x12\x0c\n\x04name\x18\x01 \x03(\tb\x06proto3')
)



_SPEC_CREATE_PACKAGE_INSTALLMODE = _descriptor.EnumDescriptor(
  name='InstallMode',
  full_name='recipe_modules.infra.support_3pp.Spec.Create.Package.InstallMode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='copy', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='symlink', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1066,
  serialized_end=1102,
)
_sym_db.RegisterEnumDescriptor(_SPEC_CREATE_PACKAGE_INSTALLMODE)


_SPEC_CREATE_SOURCE = _descriptor.Descriptor(
  name='Source',
  full_name='recipe_modules.infra.support_3pp.Spec.Create.Source',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='git', full_name='recipe_modules.infra.support_3pp.Spec.Create.Source.git', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cipd', full_name='recipe_modules.infra.support_3pp.Spec.Create.Source.cipd', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='script', full_name='recipe_modules.infra.support_3pp.Spec.Create.Source.script', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='subdir', full_name='recipe_modules.infra.support_3pp.Spec.Create.Source.subdir', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='unpack_archive', full_name='recipe_modules.infra.support_3pp.Spec.Create.Source.unpack_archive', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='no_archive_prune', full_name='recipe_modules.infra.support_3pp.Spec.Create.Source.no_archive_prune', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='patch_dir', full_name='recipe_modules.infra.support_3pp.Spec.Create.Source.patch_dir', index=6,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='patch_version', full_name='recipe_modules.infra.support_3pp.Spec.Create.Source.patch_version', index=7,
      number=8, type=9, cpp_type=9, label=1,
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
    _descriptor.OneofDescriptor(
      name='method', full_name='recipe_modules.infra.support_3pp.Spec.Create.Source.method',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=551,
  serialized_end=865,
)

_SPEC_CREATE_BUILD = _descriptor.Descriptor(
  name='Build',
  full_name='recipe_modules.infra.support_3pp.Spec.Create.Build',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='install', full_name='recipe_modules.infra.support_3pp.Spec.Create.Build.install', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tool', full_name='recipe_modules.infra.support_3pp.Spec.Create.Build.tool', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dep', full_name='recipe_modules.infra.support_3pp.Spec.Create.Build.dep', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='no_docker_env', full_name='recipe_modules.infra.support_3pp.Spec.Create.Build.no_docker_env', index=3,
      number=4, type=8, cpp_type=7, label=1,
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
  serialized_start=867,
  serialized_end=941,
)

_SPEC_CREATE_PACKAGE = _descriptor.Descriptor(
  name='Package',
  full_name='recipe_modules.infra.support_3pp.Spec.Create.Package',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='install_mode', full_name='recipe_modules.infra.support_3pp.Spec.Create.Package.install_mode', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='version_file', full_name='recipe_modules.infra.support_3pp.Spec.Create.Package.version_file', index=1,
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
    _SPEC_CREATE_PACKAGE_INSTALLMODE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=944,
  serialized_end=1102,
)

_SPEC_CREATE_VERIFY = _descriptor.Descriptor(
  name='Verify',
  full_name='recipe_modules.infra.support_3pp.Spec.Create.Verify',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='test', full_name='recipe_modules.infra.support_3pp.Spec.Create.Verify.test', index=0,
      number=1, type=9, cpp_type=9, label=3,
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
  serialized_start=1104,
  serialized_end=1126,
)

_SPEC_CREATE = _descriptor.Descriptor(
  name='Create',
  full_name='recipe_modules.infra.support_3pp.Spec.Create',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='platform_re', full_name='recipe_modules.infra.support_3pp.Spec.Create.platform_re', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='experimental', full_name='recipe_modules.infra.support_3pp.Spec.Create.experimental', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='unsupported', full_name='recipe_modules.infra.support_3pp.Spec.Create.unsupported', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source', full_name='recipe_modules.infra.support_3pp.Spec.Create.source', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='build', full_name='recipe_modules.infra.support_3pp.Spec.Create.build', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='package', full_name='recipe_modules.infra.support_3pp.Spec.Create.package', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='verify', full_name='recipe_modules.infra.support_3pp.Spec.Create.verify', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_SPEC_CREATE_SOURCE, _SPEC_CREATE_BUILD, _SPEC_CREATE_PACKAGE, _SPEC_CREATE_VERIFY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=196,
  serialized_end=1126,
)

_SPEC_UPLOAD = _descriptor.Descriptor(
  name='Upload',
  full_name='recipe_modules.infra.support_3pp.Spec.Upload',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pkg_prefix', full_name='recipe_modules.infra.support_3pp.Spec.Upload.pkg_prefix', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='universal', full_name='recipe_modules.infra.support_3pp.Spec.Upload.universal', index=1,
      number=2, type=8, cpp_type=7, label=1,
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
  serialized_start=1128,
  serialized_end=1175,
)

_SPEC = _descriptor.Descriptor(
  name='Spec',
  full_name='recipe_modules.infra.support_3pp.Spec',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='create', full_name='recipe_modules.infra.support_3pp.Spec.create', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='upload', full_name='recipe_modules.infra.support_3pp.Spec.upload', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_SPEC_CREATE, _SPEC_UPLOAD, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=61,
  serialized_end=1175,
)


_GITSOURCE = _descriptor.Descriptor(
  name='GitSource',
  full_name='recipe_modules.infra.support_3pp.GitSource',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='repo', full_name='recipe_modules.infra.support_3pp.GitSource.repo', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tag_pattern', full_name='recipe_modules.infra.support_3pp.GitSource.tag_pattern', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='version_join', full_name='recipe_modules.infra.support_3pp.GitSource.version_join', index=2,
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
  serialized_start=1177,
  serialized_end=1245,
)


_CIPDSOURCE = _descriptor.Descriptor(
  name='CipdSource',
  full_name='recipe_modules.infra.support_3pp.CipdSource',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pkg', full_name='recipe_modules.infra.support_3pp.CipdSource.pkg', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_version', full_name='recipe_modules.infra.support_3pp.CipdSource.default_version', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='original_download_url', full_name='recipe_modules.infra.support_3pp.CipdSource.original_download_url', index=2,
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
  serialized_start=1247,
  serialized_end=1328,
)


_SCRIPTSOURCE = _descriptor.Descriptor(
  name='ScriptSource',
  full_name='recipe_modules.infra.support_3pp.ScriptSource',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='recipe_modules.infra.support_3pp.ScriptSource.name', index=0,
      number=1, type=9, cpp_type=9, label=3,
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
  serialized_start=1330,
  serialized_end=1358,
)

_SPEC_CREATE_SOURCE.fields_by_name['git'].message_type = _GITSOURCE
_SPEC_CREATE_SOURCE.fields_by_name['cipd'].message_type = _CIPDSOURCE
_SPEC_CREATE_SOURCE.fields_by_name['script'].message_type = _SCRIPTSOURCE
_SPEC_CREATE_SOURCE.containing_type = _SPEC_CREATE
_SPEC_CREATE_SOURCE.oneofs_by_name['method'].fields.append(
  _SPEC_CREATE_SOURCE.fields_by_name['git'])
_SPEC_CREATE_SOURCE.fields_by_name['git'].containing_oneof = _SPEC_CREATE_SOURCE.oneofs_by_name['method']
_SPEC_CREATE_SOURCE.oneofs_by_name['method'].fields.append(
  _SPEC_CREATE_SOURCE.fields_by_name['cipd'])
_SPEC_CREATE_SOURCE.fields_by_name['cipd'].containing_oneof = _SPEC_CREATE_SOURCE.oneofs_by_name['method']
_SPEC_CREATE_SOURCE.oneofs_by_name['method'].fields.append(
  _SPEC_CREATE_SOURCE.fields_by_name['script'])
_SPEC_CREATE_SOURCE.fields_by_name['script'].containing_oneof = _SPEC_CREATE_SOURCE.oneofs_by_name['method']
_SPEC_CREATE_BUILD.containing_type = _SPEC_CREATE
_SPEC_CREATE_PACKAGE.fields_by_name['install_mode'].enum_type = _SPEC_CREATE_PACKAGE_INSTALLMODE
_SPEC_CREATE_PACKAGE.containing_type = _SPEC_CREATE
_SPEC_CREATE_PACKAGE_INSTALLMODE.containing_type = _SPEC_CREATE_PACKAGE
_SPEC_CREATE_VERIFY.containing_type = _SPEC_CREATE
_SPEC_CREATE.fields_by_name['source'].message_type = _SPEC_CREATE_SOURCE
_SPEC_CREATE.fields_by_name['build'].message_type = _SPEC_CREATE_BUILD
_SPEC_CREATE.fields_by_name['package'].message_type = _SPEC_CREATE_PACKAGE
_SPEC_CREATE.fields_by_name['verify'].message_type = _SPEC_CREATE_VERIFY
_SPEC_CREATE.containing_type = _SPEC
_SPEC_UPLOAD.containing_type = _SPEC
_SPEC.fields_by_name['create'].message_type = _SPEC_CREATE
_SPEC.fields_by_name['upload'].message_type = _SPEC_UPLOAD
DESCRIPTOR.message_types_by_name['Spec'] = _SPEC
DESCRIPTOR.message_types_by_name['GitSource'] = _GITSOURCE
DESCRIPTOR.message_types_by_name['CipdSource'] = _CIPDSOURCE
DESCRIPTOR.message_types_by_name['ScriptSource'] = _SCRIPTSOURCE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Spec = _reflection.GeneratedProtocolMessageType('Spec', (_message.Message,), dict(

  Create = _reflection.GeneratedProtocolMessageType('Create', (_message.Message,), dict(

    Source = _reflection.GeneratedProtocolMessageType('Source', (_message.Message,), dict(
      DESCRIPTOR = _SPEC_CREATE_SOURCE,
      __module__ = 'support_3pp.spec_pb2'
      # @@protoc_insertion_point(class_scope:recipe_modules.infra.support_3pp.Spec.Create.Source)
      ))
    ,

    Build = _reflection.GeneratedProtocolMessageType('Build', (_message.Message,), dict(
      DESCRIPTOR = _SPEC_CREATE_BUILD,
      __module__ = 'support_3pp.spec_pb2'
      # @@protoc_insertion_point(class_scope:recipe_modules.infra.support_3pp.Spec.Create.Build)
      ))
    ,

    Package = _reflection.GeneratedProtocolMessageType('Package', (_message.Message,), dict(
      DESCRIPTOR = _SPEC_CREATE_PACKAGE,
      __module__ = 'support_3pp.spec_pb2'
      # @@protoc_insertion_point(class_scope:recipe_modules.infra.support_3pp.Spec.Create.Package)
      ))
    ,

    Verify = _reflection.GeneratedProtocolMessageType('Verify', (_message.Message,), dict(
      DESCRIPTOR = _SPEC_CREATE_VERIFY,
      __module__ = 'support_3pp.spec_pb2'
      # @@protoc_insertion_point(class_scope:recipe_modules.infra.support_3pp.Spec.Create.Verify)
      ))
    ,
    DESCRIPTOR = _SPEC_CREATE,
    __module__ = 'support_3pp.spec_pb2'
    # @@protoc_insertion_point(class_scope:recipe_modules.infra.support_3pp.Spec.Create)
    ))
  ,

  Upload = _reflection.GeneratedProtocolMessageType('Upload', (_message.Message,), dict(
    DESCRIPTOR = _SPEC_UPLOAD,
    __module__ = 'support_3pp.spec_pb2'
    # @@protoc_insertion_point(class_scope:recipe_modules.infra.support_3pp.Spec.Upload)
    ))
  ,
  DESCRIPTOR = _SPEC,
  __module__ = 'support_3pp.spec_pb2'
  # @@protoc_insertion_point(class_scope:recipe_modules.infra.support_3pp.Spec)
  ))
_sym_db.RegisterMessage(Spec)
_sym_db.RegisterMessage(Spec.Create)
_sym_db.RegisterMessage(Spec.Create.Source)
_sym_db.RegisterMessage(Spec.Create.Build)
_sym_db.RegisterMessage(Spec.Create.Package)
_sym_db.RegisterMessage(Spec.Create.Verify)
_sym_db.RegisterMessage(Spec.Upload)

GitSource = _reflection.GeneratedProtocolMessageType('GitSource', (_message.Message,), dict(
  DESCRIPTOR = _GITSOURCE,
  __module__ = 'support_3pp.spec_pb2'
  # @@protoc_insertion_point(class_scope:recipe_modules.infra.support_3pp.GitSource)
  ))
_sym_db.RegisterMessage(GitSource)

CipdSource = _reflection.GeneratedProtocolMessageType('CipdSource', (_message.Message,), dict(
  DESCRIPTOR = _CIPDSOURCE,
  __module__ = 'support_3pp.spec_pb2'
  # @@protoc_insertion_point(class_scope:recipe_modules.infra.support_3pp.CipdSource)
  ))
_sym_db.RegisterMessage(CipdSource)

ScriptSource = _reflection.GeneratedProtocolMessageType('ScriptSource', (_message.Message,), dict(
  DESCRIPTOR = _SCRIPTSOURCE,
  __module__ = 'support_3pp.spec_pb2'
  # @@protoc_insertion_point(class_scope:recipe_modules.infra.support_3pp.ScriptSource)
  ))
_sym_db.RegisterMessage(ScriptSource)


# @@protoc_insertion_point(module_scope)
