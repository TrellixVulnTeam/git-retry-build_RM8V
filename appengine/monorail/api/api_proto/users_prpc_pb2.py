# Generated by the pRPC protocol buffer compiler plugin.  DO NOT EDIT!
# source: api/api_proto/users.proto

import base64
import zlib

from google.protobuf import descriptor_pb2

# Includes description of the api/api_proto/users.proto and all of its transitive
# dependencies. Includes source code info.
FILE_DESCRIPTOR_SET = descriptor_pb2.FileDescriptorSet()
FILE_DESCRIPTOR_SET.ParseFromString(zlib.decompress(base64.b64decode(
    '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')))
_INDEX = {
    f.name: {
      'descriptor': f,
      'services': {s.name: s for s in f.service},
    }
    for f in FILE_DESCRIPTOR_SET.file
}


UsersServiceDescription = {
  'file_descriptor_set': FILE_DESCRIPTOR_SET,
  'file_descriptor': _INDEX[u'api/api_proto/users.proto']['descriptor'],
  'service_descriptor': _INDEX[u'api/api_proto/users.proto']['services'][u'Users'],
}
