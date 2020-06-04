// Copyright 2019 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.6.1
// source: infra/appengine/cros/lab_inventory/app/config/config.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LuciAuthGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *LuciAuthGroup) Reset() {
	*x = LuciAuthGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LuciAuthGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LuciAuthGroup) ProtoMessage() {}

func (x *LuciAuthGroup) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LuciAuthGroup.ProtoReflect.Descriptor instead.
func (*LuciAuthGroup) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescGZIP(), []int{0}
}

func (x *LuciAuthGroup) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Config is the configuration data served by luci-config for this app.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AdminService contains information about the skylab admin instances.
	AdminService *AdminService `protobuf:"bytes,2,opt,name=admin_service,json=adminService,proto3" json:"admin_service,omitempty"`
	// The access groups of the inventory.
	Readers                   *LuciAuthGroup `protobuf:"bytes,3,opt,name=readers,proto3" json:"readers,omitempty"`
	StatusWriters             *LuciAuthGroup `protobuf:"bytes,4,opt,name=status_writers,json=statusWriters,proto3" json:"status_writers,omitempty"`
	SetupWriters              *LuciAuthGroup `protobuf:"bytes,5,opt,name=setup_writers,json=setupWriters,proto3" json:"setup_writers,omitempty"`
	PrivilegedWriters         *LuciAuthGroup `protobuf:"bytes,6,opt,name=privileged_writers,json=privilegedWriters,proto3" json:"privileged_writers,omitempty"`
	HwidSecret                string         `protobuf:"bytes,7,opt,name=hwid_secret,json=hwidSecret,proto3" json:"hwid_secret,omitempty"`
	DeviceConfigSource        *Gitiles       `protobuf:"bytes,8,opt,name=device_config_source,json=deviceConfigSource,proto3" json:"device_config_source,omitempty"`
	ManufacturingConfigSource *Gitiles       `protobuf:"bytes,9,opt,name=manufacturing_config_source,json=manufacturingConfigSource,proto3" json:"manufacturing_config_source,omitempty"`
	// The git repo information of inventory v1.
	// TODO(guocb) remove this after migration.
	Inventory *InventoryV1Repo `protobuf:"bytes,12,opt,name=inventory,proto3" json:"inventory,omitempty"`
	// Environment managed by this instance of app, e.g. ENVIRONMENT_STAGING,
	// ENVIRONMENT_PROD, etc.
	Environment string `protobuf:"bytes,10,opt,name=environment,proto3" json:"environment,omitempty"`
	// The hostname of drone-queen service to push inventory to.
	QueenService string `protobuf:"bytes,11,opt,name=queen_service,json=queenService,proto3" json:"queen_service,omitempty"`
	// Report the DUT utilization or not.
	EnableInventoryReporting bool `protobuf:"varint,13,opt,name=enable_inventory_reporting,json=enableInventoryReporting,proto3" json:"enable_inventory_reporting,omitempty"`
	// HaRT PubSub Configs
	Hart          *HaRT          `protobuf:"bytes,14,opt,name=hart,proto3" json:"hart,omitempty"`
	PubsubPushers *LuciAuthGroup `protobuf:"bytes,15,opt,name=pubsub_pushers,json=pubsubPushers,proto3" json:"pubsub_pushers,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetAdminService() *AdminService {
	if x != nil {
		return x.AdminService
	}
	return nil
}

func (x *Config) GetReaders() *LuciAuthGroup {
	if x != nil {
		return x.Readers
	}
	return nil
}

func (x *Config) GetStatusWriters() *LuciAuthGroup {
	if x != nil {
		return x.StatusWriters
	}
	return nil
}

func (x *Config) GetSetupWriters() *LuciAuthGroup {
	if x != nil {
		return x.SetupWriters
	}
	return nil
}

func (x *Config) GetPrivilegedWriters() *LuciAuthGroup {
	if x != nil {
		return x.PrivilegedWriters
	}
	return nil
}

func (x *Config) GetHwidSecret() string {
	if x != nil {
		return x.HwidSecret
	}
	return ""
}

func (x *Config) GetDeviceConfigSource() *Gitiles {
	if x != nil {
		return x.DeviceConfigSource
	}
	return nil
}

func (x *Config) GetManufacturingConfigSource() *Gitiles {
	if x != nil {
		return x.ManufacturingConfigSource
	}
	return nil
}

func (x *Config) GetInventory() *InventoryV1Repo {
	if x != nil {
		return x.Inventory
	}
	return nil
}

func (x *Config) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *Config) GetQueenService() string {
	if x != nil {
		return x.QueenService
	}
	return ""
}

func (x *Config) GetEnableInventoryReporting() bool {
	if x != nil {
		return x.EnableInventoryReporting
	}
	return false
}

func (x *Config) GetHart() *HaRT {
	if x != nil {
		return x.Hart
	}
	return nil
}

func (x *Config) GetPubsubPushers() *LuciAuthGroup {
	if x != nil {
		return x.PubsubPushers
	}
	return nil
}

type AdminService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The skylab admin GAE server hosting the admin services.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *AdminService) Reset() {
	*x = AdminService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminService) ProtoMessage() {}

func (x *AdminService) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminService.ProtoReflect.Descriptor instead.
func (*AdminService) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescGZIP(), []int{2}
}

func (x *AdminService) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type Gitiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The gitiles host name, e.g. 'chrome-internal.googlesource.com'.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The project (repo) name, e.g. 'chromeos/infra/config'.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// The commit hash/branch to be checked out, e.g. 'refs/heads/master'.
	Committish string `protobuf:"bytes,3,opt,name=committish,proto3" json:"committish,omitempty"`
	// The path of the file to be downloaded, e.g. 'path/to/file.cfg'.
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Gitiles) Reset() {
	*x = Gitiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gitiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gitiles) ProtoMessage() {}

func (x *Gitiles) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gitiles.ProtoReflect.Descriptor instead.
func (*Gitiles) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescGZIP(), []int{3}
}

func (x *Gitiles) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Gitiles) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Gitiles) GetCommittish() string {
	if x != nil {
		return x.Committish
	}
	return ""
}

func (x *Gitiles) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type InventoryV1Repo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host                   string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Project                string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	Branch                 string `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	LabDataPath            string `protobuf:"bytes,4,opt,name=lab_data_path,json=labDataPath,proto3" json:"lab_data_path,omitempty"`
	InfrastructureDataPath string `protobuf:"bytes,5,opt,name=infrastructure_data_path,json=infrastructureDataPath,proto3" json:"infrastructure_data_path,omitempty"`
	Multifile              bool   `protobuf:"varint,6,opt,name=multifile,proto3" json:"multifile,omitempty"`
}

func (x *InventoryV1Repo) Reset() {
	*x = InventoryV1Repo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryV1Repo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryV1Repo) ProtoMessage() {}

func (x *InventoryV1Repo) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryV1Repo.ProtoReflect.Descriptor instead.
func (*InventoryV1Repo) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescGZIP(), []int{4}
}

func (x *InventoryV1Repo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *InventoryV1Repo) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *InventoryV1Repo) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *InventoryV1Repo) GetLabDataPath() string {
	if x != nil {
		return x.LabDataPath
	}
	return ""
}

func (x *InventoryV1Repo) GetInfrastructureDataPath() string {
	if x != nil {
		return x.InfrastructureDataPath
	}
	return ""
}

func (x *InventoryV1Repo) GetMultifile() bool {
	if x != nil {
		return x.Multifile
	}
	return false
}

type HaRT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project      string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Topic        string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Subscription string `protobuf:"bytes,3,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *HaRT) Reset() {
	*x = HaRT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HaRT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HaRT) ProtoMessage() {}

func (x *HaRT) ProtoReflect() protoreflect.Message {
	mi := &file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HaRT.ProtoReflect.Descriptor instead.
func (*HaRT) Descriptor() ([]byte, []int) {
	return file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescGZIP(), []int{5}
}

func (x *HaRT) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *HaRT) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *HaRT) GetSubscription() string {
	if x != nil {
		return x.Subscription
	}
	return ""
}

var File_infra_appengine_cros_lab_inventory_app_config_config_proto protoreflect.FileDescriptor

var file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6c, 0x61,
	0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x25, 0x0a, 0x0d, 0x4c, 0x75, 0x63, 0x69, 0x41, 0x75, 0x74, 0x68, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x97, 0x07, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x61,
	0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x0c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x07, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x75, 0x63, 0x69, 0x41, 0x75, 0x74, 0x68, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x07, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x4a, 0x0a, 0x0e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x75, 0x63, 0x69,
	0x41, 0x75, 0x74, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x73, 0x12, 0x48, 0x0a, 0x0d, 0x73, 0x65, 0x74, 0x75,
	0x70, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x75, 0x63, 0x69, 0x41, 0x75, 0x74, 0x68, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x0c, 0x73, 0x65, 0x74, 0x75, 0x70, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x52, 0x0a, 0x12, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64,
	0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4c, 0x75, 0x63, 0x69, 0x41, 0x75, 0x74, 0x68, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x11, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x77, 0x69, 0x64, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x77, 0x69,
	0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x4f, 0x0a, 0x14, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x69, 0x74,
	0x69, 0x6c, 0x65, 0x73, 0x52, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x1b, 0x6d, 0x61, 0x6e, 0x75,
	0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x19, 0x6d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x61, 0x62,
	0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x56, 0x31, 0x52, 0x65, 0x70,
	0x6f, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x71, 0x75, 0x65, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x71, 0x75, 0x65, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x1a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x2e, 0x0a, 0x04, 0x68, 0x61, 0x72, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6c, 0x61, 0x62, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x48, 0x61, 0x52, 0x54, 0x52, 0x04, 0x68, 0x61, 0x72,
	0x74, 0x12, 0x4a, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x75, 0x73, 0x68,
	0x65, 0x72, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6c, 0x61, 0x62, 0x5f,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x4c, 0x75, 0x63, 0x69, 0x41, 0x75, 0x74, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0d,
	0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x50, 0x75, 0x73, 0x68, 0x65, 0x72, 0x73, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x22, 0x22, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x07, 0x47, 0x69, 0x74, 0x69, 0x6c,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x69, 0x73, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x69, 0x73, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x22, 0xd3, 0x01, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x56, 0x31, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x22,
	0x0a, 0x0d, 0x6c, 0x61, 0x62, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x62, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x38, 0x0a, 0x18, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x5a, 0x0a, 0x04, 0x48, 0x61,
	0x52, 0x54, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x5a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescOnce sync.Once
	file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescData = file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDesc
)

func file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescGZIP() []byte {
	file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescOnce.Do(func() {
		file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescData)
	})
	return file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDescData
}

var file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_infra_appengine_cros_lab_inventory_app_config_config_proto_goTypes = []interface{}{
	(*LuciAuthGroup)(nil),   // 0: lab_inventory.config.LuciAuthGroup
	(*Config)(nil),          // 1: lab_inventory.config.Config
	(*AdminService)(nil),    // 2: lab_inventory.config.AdminService
	(*Gitiles)(nil),         // 3: lab_inventory.config.Gitiles
	(*InventoryV1Repo)(nil), // 4: lab_inventory.config.InventoryV1Repo
	(*HaRT)(nil),            // 5: lab_inventory.config.HaRT
}
var file_infra_appengine_cros_lab_inventory_app_config_config_proto_depIdxs = []int32{
	2,  // 0: lab_inventory.config.Config.admin_service:type_name -> lab_inventory.config.AdminService
	0,  // 1: lab_inventory.config.Config.readers:type_name -> lab_inventory.config.LuciAuthGroup
	0,  // 2: lab_inventory.config.Config.status_writers:type_name -> lab_inventory.config.LuciAuthGroup
	0,  // 3: lab_inventory.config.Config.setup_writers:type_name -> lab_inventory.config.LuciAuthGroup
	0,  // 4: lab_inventory.config.Config.privileged_writers:type_name -> lab_inventory.config.LuciAuthGroup
	3,  // 5: lab_inventory.config.Config.device_config_source:type_name -> lab_inventory.config.Gitiles
	3,  // 6: lab_inventory.config.Config.manufacturing_config_source:type_name -> lab_inventory.config.Gitiles
	4,  // 7: lab_inventory.config.Config.inventory:type_name -> lab_inventory.config.InventoryV1Repo
	5,  // 8: lab_inventory.config.Config.hart:type_name -> lab_inventory.config.HaRT
	0,  // 9: lab_inventory.config.Config.pubsub_pushers:type_name -> lab_inventory.config.LuciAuthGroup
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_infra_appengine_cros_lab_inventory_app_config_config_proto_init() }
func file_infra_appengine_cros_lab_inventory_app_config_config_proto_init() {
	if File_infra_appengine_cros_lab_inventory_app_config_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LuciAuthGroup); i {
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
		file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminService); i {
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
		file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gitiles); i {
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
		file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryV1Repo); i {
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
		file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HaRT); i {
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
			RawDescriptor: file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infra_appengine_cros_lab_inventory_app_config_config_proto_goTypes,
		DependencyIndexes: file_infra_appengine_cros_lab_inventory_app_config_config_proto_depIdxs,
		MessageInfos:      file_infra_appengine_cros_lab_inventory_app_config_config_proto_msgTypes,
	}.Build()
	File_infra_appengine_cros_lab_inventory_app_config_config_proto = out.File
	file_infra_appengine_cros_lab_inventory_app_config_config_proto_rawDesc = nil
	file_infra_appengine_cros_lab_inventory_app_config_config_proto_goTypes = nil
	file_infra_appengine_cros_lab_inventory_app_config_config_proto_depIdxs = nil
}
