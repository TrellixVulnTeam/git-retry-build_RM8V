// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.1
// source: infra/cros/karte/api/observation.proto

package kartepb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An Observation describes a measurement during an action.
//
// Examples:
//      battery_level:70 - battery level on device is 70%
//      rrd:present      - rrd metric is present on device
//      disk_usage:35    - 35% internal storage used
//
// Next Tag: 2
type Observation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the observation. Names are generated
	// automatically when a new observation is created.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Observation) Reset() {
	*x = Observation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_cros_karte_api_observation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Observation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Observation) ProtoMessage() {}

func (x *Observation) ProtoReflect() protoreflect.Message {
	mi := &file_infra_cros_karte_api_observation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Observation.ProtoReflect.Descriptor instead.
func (*Observation) Descriptor() ([]byte, []int) {
	return file_infra_cros_karte_api_observation_proto_rawDescGZIP(), []int{0}
}

func (x *Observation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_infra_cros_karte_api_observation_proto protoreflect.FileDescriptor

var file_infra_cros_karte_api_observation_proto_rawDesc = []byte{
	0x0a, 0x26, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x6b, 0x61, 0x72,
	0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65,
	0x6f, 0x73, 0x2e, 0x6b, 0x61, 0x72, 0x74, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0b, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x3e, 0xea, 0x41, 0x3b, 0x0a, 0x1d, 0x6b, 0x61, 0x72,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x42, 0x1e, 0x5a, 0x1c, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f,
	0x63, 0x72, 0x6f, 0x73, 0x2f, 0x6b, 0x61, 0x72, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x6b,
	0x61, 0x72, 0x74, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_cros_karte_api_observation_proto_rawDescOnce sync.Once
	file_infra_cros_karte_api_observation_proto_rawDescData = file_infra_cros_karte_api_observation_proto_rawDesc
)

func file_infra_cros_karte_api_observation_proto_rawDescGZIP() []byte {
	file_infra_cros_karte_api_observation_proto_rawDescOnce.Do(func() {
		file_infra_cros_karte_api_observation_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_cros_karte_api_observation_proto_rawDescData)
	})
	return file_infra_cros_karte_api_observation_proto_rawDescData
}

var file_infra_cros_karte_api_observation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infra_cros_karte_api_observation_proto_goTypes = []interface{}{
	(*Observation)(nil), // 0: chromeos.karte.Observation
}
var file_infra_cros_karte_api_observation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_infra_cros_karte_api_observation_proto_init() }
func file_infra_cros_karte_api_observation_proto_init() {
	if File_infra_cros_karte_api_observation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_cros_karte_api_observation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Observation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infra_cros_karte_api_observation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_cros_karte_api_observation_proto_goTypes,
		DependencyIndexes: file_infra_cros_karte_api_observation_proto_depIdxs,
		MessageInfos:      file_infra_cros_karte_api_observation_proto_msgTypes,
	}.Build()
	File_infra_cros_karte_api_observation_proto = out.File
	file_infra_cros_karte_api_observation_proto_rawDesc = nil
	file_infra_cros_karte_api_observation_proto_goTypes = nil
	file_infra_cros_karte_api_observation_proto_depIdxs = nil
}
