// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/appengine/rotation-proxy/proto/rotation_proxy.proto

package rotation_proxy

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request to update a list of rotations.
type BatchUpdateRotationsRequest struct {
	// The rotations to update.
	Requests             []*UpdateRotationRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BatchUpdateRotationsRequest) Reset()         { *m = BatchUpdateRotationsRequest{} }
func (m *BatchUpdateRotationsRequest) String() string { return proto.CompactTextString(m) }
func (*BatchUpdateRotationsRequest) ProtoMessage()    {}
func (*BatchUpdateRotationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a28b40e80ae262, []int{0}
}

func (m *BatchUpdateRotationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchUpdateRotationsRequest.Unmarshal(m, b)
}
func (m *BatchUpdateRotationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchUpdateRotationsRequest.Marshal(b, m, deterministic)
}
func (m *BatchUpdateRotationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchUpdateRotationsRequest.Merge(m, src)
}
func (m *BatchUpdateRotationsRequest) XXX_Size() int {
	return xxx_messageInfo_BatchUpdateRotationsRequest.Size(m)
}
func (m *BatchUpdateRotationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchUpdateRotationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchUpdateRotationsRequest proto.InternalMessageInfo

func (m *BatchUpdateRotationsRequest) GetRequests() []*UpdateRotationRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

type UpdateRotationRequest struct {
	// The rotation to update.
	Rotation             *Rotation `protobuf:"bytes,1,opt,name=rotation,proto3" json:"rotation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateRotationRequest) Reset()         { *m = UpdateRotationRequest{} }
func (m *UpdateRotationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRotationRequest) ProtoMessage()    {}
func (*UpdateRotationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a28b40e80ae262, []int{1}
}

func (m *UpdateRotationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRotationRequest.Unmarshal(m, b)
}
func (m *UpdateRotationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRotationRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRotationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRotationRequest.Merge(m, src)
}
func (m *UpdateRotationRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRotationRequest.Size(m)
}
func (m *UpdateRotationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRotationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRotationRequest proto.InternalMessageInfo

func (m *UpdateRotationRequest) GetRotation() *Rotation {
	if m != nil {
		return m.Rotation
	}
	return nil
}

type BatchUpdateRotationsResponse struct {
	// Rotation updated.
	Rotations            []*Rotation `protobuf:"bytes,1,rep,name=rotations,proto3" json:"rotations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BatchUpdateRotationsResponse) Reset()         { *m = BatchUpdateRotationsResponse{} }
func (m *BatchUpdateRotationsResponse) String() string { return proto.CompactTextString(m) }
func (*BatchUpdateRotationsResponse) ProtoMessage()    {}
func (*BatchUpdateRotationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a28b40e80ae262, []int{2}
}

func (m *BatchUpdateRotationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchUpdateRotationsResponse.Unmarshal(m, b)
}
func (m *BatchUpdateRotationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchUpdateRotationsResponse.Marshal(b, m, deterministic)
}
func (m *BatchUpdateRotationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchUpdateRotationsResponse.Merge(m, src)
}
func (m *BatchUpdateRotationsResponse) XXX_Size() int {
	return xxx_messageInfo_BatchUpdateRotationsResponse.Size(m)
}
func (m *BatchUpdateRotationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchUpdateRotationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchUpdateRotationsResponse proto.InternalMessageInfo

func (m *BatchUpdateRotationsResponse) GetRotations() []*Rotation {
	if m != nil {
		return m.Rotations
	}
	return nil
}

type BatchGetRotationsRequest struct {
	// The names of the rotations to fetch.
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchGetRotationsRequest) Reset()         { *m = BatchGetRotationsRequest{} }
func (m *BatchGetRotationsRequest) String() string { return proto.CompactTextString(m) }
func (*BatchGetRotationsRequest) ProtoMessage()    {}
func (*BatchGetRotationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a28b40e80ae262, []int{3}
}

func (m *BatchGetRotationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchGetRotationsRequest.Unmarshal(m, b)
}
func (m *BatchGetRotationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchGetRotationsRequest.Marshal(b, m, deterministic)
}
func (m *BatchGetRotationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchGetRotationsRequest.Merge(m, src)
}
func (m *BatchGetRotationsRequest) XXX_Size() int {
	return xxx_messageInfo_BatchGetRotationsRequest.Size(m)
}
func (m *BatchGetRotationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchGetRotationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchGetRotationsRequest proto.InternalMessageInfo

func (m *BatchGetRotationsRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type BatchGetRotationsResponse struct {
	// Details about the rotations requested.
	Rotations            []*Rotation `protobuf:"bytes,1,rep,name=rotations,proto3" json:"rotations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BatchGetRotationsResponse) Reset()         { *m = BatchGetRotationsResponse{} }
func (m *BatchGetRotationsResponse) String() string { return proto.CompactTextString(m) }
func (*BatchGetRotationsResponse) ProtoMessage()    {}
func (*BatchGetRotationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a28b40e80ae262, []int{4}
}

func (m *BatchGetRotationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchGetRotationsResponse.Unmarshal(m, b)
}
func (m *BatchGetRotationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchGetRotationsResponse.Marshal(b, m, deterministic)
}
func (m *BatchGetRotationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchGetRotationsResponse.Merge(m, src)
}
func (m *BatchGetRotationsResponse) XXX_Size() int {
	return xxx_messageInfo_BatchGetRotationsResponse.Size(m)
}
func (m *BatchGetRotationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchGetRotationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchGetRotationsResponse proto.InternalMessageInfo

func (m *BatchGetRotationsResponse) GetRotations() []*Rotation {
	if m != nil {
		return m.Rotations
	}
	return nil
}

// Contains information about a rotation.
type Rotation struct {
	// The unique name of the rotation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Shifts in a rotation.
	// Shifts will be sorted with the current shift as the first element.
	Shifts               []*Shift `protobuf:"bytes,2,rep,name=shifts,proto3" json:"shifts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rotation) Reset()         { *m = Rotation{} }
func (m *Rotation) String() string { return proto.CompactTextString(m) }
func (*Rotation) ProtoMessage()    {}
func (*Rotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a28b40e80ae262, []int{5}
}

func (m *Rotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rotation.Unmarshal(m, b)
}
func (m *Rotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rotation.Marshal(b, m, deterministic)
}
func (m *Rotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rotation.Merge(m, src)
}
func (m *Rotation) XXX_Size() int {
	return xxx_messageInfo_Rotation.Size(m)
}
func (m *Rotation) XXX_DiscardUnknown() {
	xxx_messageInfo_Rotation.DiscardUnknown(m)
}

var xxx_messageInfo_Rotation proto.InternalMessageInfo

func (m *Rotation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rotation) GetShifts() []*Shift {
	if m != nil {
		return m.Shifts
	}
	return nil
}

type Shift struct {
	// The list of oncall users for this shift.
	Oncalls []*OncallPerson `protobuf:"bytes,1,rep,name=oncalls,proto3" json:"oncalls,omitempty"`
	// Unix timestamp of the start of the shift.
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Unix timestamp of the end of the shift.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Shift) Reset()         { *m = Shift{} }
func (m *Shift) String() string { return proto.CompactTextString(m) }
func (*Shift) ProtoMessage()    {}
func (*Shift) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a28b40e80ae262, []int{6}
}

func (m *Shift) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shift.Unmarshal(m, b)
}
func (m *Shift) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shift.Marshal(b, m, deterministic)
}
func (m *Shift) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shift.Merge(m, src)
}
func (m *Shift) XXX_Size() int {
	return xxx_messageInfo_Shift.Size(m)
}
func (m *Shift) XXX_DiscardUnknown() {
	xxx_messageInfo_Shift.DiscardUnknown(m)
}

var xxx_messageInfo_Shift proto.InternalMessageInfo

func (m *Shift) GetOncalls() []*OncallPerson {
	if m != nil {
		return m.Oncalls
	}
	return nil
}

func (m *Shift) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Shift) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type OncallPerson struct {
	// Email of oncall person.
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OncallPerson) Reset()         { *m = OncallPerson{} }
func (m *OncallPerson) String() string { return proto.CompactTextString(m) }
func (*OncallPerson) ProtoMessage()    {}
func (*OncallPerson) Descriptor() ([]byte, []int) {
	return fileDescriptor_83a28b40e80ae262, []int{7}
}

func (m *OncallPerson) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OncallPerson.Unmarshal(m, b)
}
func (m *OncallPerson) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OncallPerson.Marshal(b, m, deterministic)
}
func (m *OncallPerson) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OncallPerson.Merge(m, src)
}
func (m *OncallPerson) XXX_Size() int {
	return xxx_messageInfo_OncallPerson.Size(m)
}
func (m *OncallPerson) XXX_DiscardUnknown() {
	xxx_messageInfo_OncallPerson.DiscardUnknown(m)
}

var xxx_messageInfo_OncallPerson proto.InternalMessageInfo

func (m *OncallPerson) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*BatchUpdateRotationsRequest)(nil), "rotation_proxy.BatchUpdateRotationsRequest")
	proto.RegisterType((*UpdateRotationRequest)(nil), "rotation_proxy.UpdateRotationRequest")
	proto.RegisterType((*BatchUpdateRotationsResponse)(nil), "rotation_proxy.BatchUpdateRotationsResponse")
	proto.RegisterType((*BatchGetRotationsRequest)(nil), "rotation_proxy.BatchGetRotationsRequest")
	proto.RegisterType((*BatchGetRotationsResponse)(nil), "rotation_proxy.BatchGetRotationsResponse")
	proto.RegisterType((*Rotation)(nil), "rotation_proxy.Rotation")
	proto.RegisterType((*Shift)(nil), "rotation_proxy.Shift")
	proto.RegisterType((*OncallPerson)(nil), "rotation_proxy.OncallPerson")
}

func init() {
	proto.RegisterFile("infra/appengine/rotation-proxy/proto/rotation_proxy.proto", fileDescriptor_83a28b40e80ae262)
}

var fileDescriptor_83a28b40e80ae262 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0x5a, 0xb6, 0xb5, 0x77, 0x08, 0x09, 0x6b, 0x13, 0x59, 0x99, 0xc4, 0x14, 0x09, 0xa9,
	0x13, 0x2c, 0x91, 0x8a, 0x36, 0xa9, 0x8f, 0xec, 0x01, 0x1e, 0x99, 0x5c, 0xe0, 0xb5, 0x72, 0xdb,
	0x9b, 0xc6, 0x52, 0x62, 0x7b, 0xb6, 0x37, 0xc1, 0x67, 0xf1, 0x47, 0xfc, 0x03, 0x3f, 0x80, 0x62,
	0xc7, 0x65, 0xcb, 0x02, 0x54, 0xda, 0x5b, 0xee, 0xf1, 0x39, 0xe7, 0x5e, 0x5f, 0x9f, 0xc0, 0x94,
	0x8b, 0x5c, 0xb3, 0x8c, 0x29, 0x85, 0x62, 0xcd, 0x05, 0x66, 0x5a, 0x5a, 0x66, 0xb9, 0x14, 0x67,
	0x4a, 0xcb, 0x6f, 0xdf, 0x33, 0xa5, 0xa5, 0x95, 0x1b, 0x70, 0xee, 0xc0, 0xd4, 0x81, 0xe4, 0xd9,
	0x7d, 0x74, 0xf4, 0x6a, 0x2d, 0xe5, 0xba, 0xc4, 0x8c, 0x29, 0x9e, 0xe5, 0x1c, 0xcb, 0xd5, 0x7c,
	0x81, 0x05, 0xbb, 0xe5, 0x52, 0x7b, 0xc1, 0x86, 0xe0, 0xaa, 0xc5, 0x4d, 0x9e, 0x59, 0x5e, 0xa1,
	0xb1, 0xac, 0x52, 0x9e, 0x90, 0x20, 0xbc, 0xbc, 0x64, 0x76, 0x59, 0x7c, 0x51, 0x2b, 0x66, 0x91,
	0x36, 0xf6, 0x86, 0xe2, 0xf5, 0x0d, 0x1a, 0x4b, 0x3e, 0xc0, 0x40, 0xfb, 0x4f, 0x13, 0x47, 0x27,
	0xfd, 0xf1, 0xfe, 0xe4, 0x75, 0xda, 0x9a, 0xec, 0xbe, 0xb2, 0x11, 0x5e, 0xf6, 0x7f, 0xbe, 0xef,
	0xd1, 0x8d, 0x36, 0xa1, 0x70, 0xd8, 0xc9, 0x23, 0x53, 0x18, 0x04, 0xbf, 0x38, 0x3a, 0x89, 0xc6,
	0xfb, 0x93, 0xb8, 0xdd, 0x20, 0x48, 0x82, 0x67, 0x53, 0x26, 0x5f, 0xe1, 0xb8, 0x7b, 0x74, 0xa3,
	0xa4, 0x30, 0x48, 0x2e, 0x60, 0x18, 0xb8, 0x61, 0xf8, 0xbf, 0x7a, 0xd3, 0x3f, 0xd4, 0xe4, 0x1c,
	0x62, 0xe7, 0xfb, 0x11, 0xed, 0x83, 0x7d, 0x1c, 0xc1, 0x8e, 0x60, 0x15, 0x7a, 0xbf, 0xa1, 0x9f,
	0xc8, 0x23, 0xc9, 0x0c, 0x8e, 0x3a, 0x64, 0x8f, 0x9c, 0x85, 0xc2, 0x20, 0xc0, 0xe4, 0x05, 0x3c,
	0xa9, 0x3b, 0xb9, 0x35, 0x35, 0xad, 0x1d, 0x40, 0xce, 0x60, 0xd7, 0x14, 0x3c, 0xb7, 0x26, 0xee,
	0x39, 0xe7, 0xc3, 0xb6, 0xf3, 0xac, 0x3e, 0xa5, 0x0d, 0x29, 0xf9, 0x11, 0xc1, 0x8e, 0x43, 0xc8,
	0x05, 0xec, 0x49, 0xb1, 0x64, 0x65, 0x19, 0x66, 0x3a, 0x6e, 0x2b, 0x3f, 0xb9, 0xe3, 0x2b, 0xd4,
	0x46, 0x0a, 0x1a, 0xc8, 0x64, 0x0a, 0x60, 0x2c, 0xd3, 0x76, 0x5e, 0xa7, 0x29, 0xee, 0xb9, 0x67,
	0x1b, 0xa5, 0x3e, 0x6a, 0x69, 0x88, 0x5a, 0xfa, 0x39, 0x44, 0x8d, 0x0e, 0x1d, 0xbb, 0xae, 0xc9,
	0x39, 0x0c, 0x50, 0xac, 0xbc, 0xb0, 0xff, 0x5f, 0xe1, 0x1e, 0x8a, 0x55, 0x5d, 0x25, 0xa7, 0xf0,
	0xf4, 0xee, 0x28, 0xf5, 0x3b, 0x60, 0xc5, 0x78, 0x79, 0x77, 0x19, 0x1e, 0x99, 0xfc, 0x8a, 0xe0,
	0x20, 0xec, 0xec, 0xaa, 0xbe, 0xc4, 0x0c, 0xf5, 0x2d, 0x5f, 0x22, 0xb9, 0x86, 0x83, 0xae, 0xbc,
	0x90, 0x37, 0xed, 0x4b, 0xff, 0xe3, 0x87, 0x18, 0xbd, 0xdd, 0x8e, 0xdc, 0x3c, 0x7b, 0x01, 0xcf,
	0x1f, 0x64, 0x82, 0x8c, 0x3b, 0x2d, 0x3a, 0xd2, 0x36, 0x3a, 0xdd, 0x82, 0xe9, 0x3b, 0x2d, 0x76,
	0xdd, 0xf6, 0xde, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x24, 0x51, 0x83, 0x5d, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RotationProxyServiceClient is the client API for RotationProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RotationProxyServiceClient interface {
	BatchUpdateRotations(ctx context.Context, in *BatchUpdateRotationsRequest, opts ...grpc.CallOption) (*BatchUpdateRotationsResponse, error)
	BatchGetRotations(ctx context.Context, in *BatchGetRotationsRequest, opts ...grpc.CallOption) (*BatchGetRotationsResponse, error)
}
type rotationProxyServicePRPCClient struct {
	client *prpc.Client
}

func NewRotationProxyServicePRPCClient(client *prpc.Client) RotationProxyServiceClient {
	return &rotationProxyServicePRPCClient{client}
}

func (c *rotationProxyServicePRPCClient) BatchUpdateRotations(ctx context.Context, in *BatchUpdateRotationsRequest, opts ...grpc.CallOption) (*BatchUpdateRotationsResponse, error) {
	out := new(BatchUpdateRotationsResponse)
	err := c.client.Call(ctx, "rotation_proxy.RotationProxyService", "BatchUpdateRotations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rotationProxyServicePRPCClient) BatchGetRotations(ctx context.Context, in *BatchGetRotationsRequest, opts ...grpc.CallOption) (*BatchGetRotationsResponse, error) {
	out := new(BatchGetRotationsResponse)
	err := c.client.Call(ctx, "rotation_proxy.RotationProxyService", "BatchGetRotations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type rotationProxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRotationProxyServiceClient(cc grpc.ClientConnInterface) RotationProxyServiceClient {
	return &rotationProxyServiceClient{cc}
}

func (c *rotationProxyServiceClient) BatchUpdateRotations(ctx context.Context, in *BatchUpdateRotationsRequest, opts ...grpc.CallOption) (*BatchUpdateRotationsResponse, error) {
	out := new(BatchUpdateRotationsResponse)
	err := c.cc.Invoke(ctx, "/rotation_proxy.RotationProxyService/BatchUpdateRotations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rotationProxyServiceClient) BatchGetRotations(ctx context.Context, in *BatchGetRotationsRequest, opts ...grpc.CallOption) (*BatchGetRotationsResponse, error) {
	out := new(BatchGetRotationsResponse)
	err := c.cc.Invoke(ctx, "/rotation_proxy.RotationProxyService/BatchGetRotations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RotationProxyServiceServer is the server API for RotationProxyService service.
type RotationProxyServiceServer interface {
	BatchUpdateRotations(context.Context, *BatchUpdateRotationsRequest) (*BatchUpdateRotationsResponse, error)
	BatchGetRotations(context.Context, *BatchGetRotationsRequest) (*BatchGetRotationsResponse, error)
}

// UnimplementedRotationProxyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRotationProxyServiceServer struct {
}

func (*UnimplementedRotationProxyServiceServer) BatchUpdateRotations(ctx context.Context, req *BatchUpdateRotationsRequest) (*BatchUpdateRotationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchUpdateRotations not implemented")
}
func (*UnimplementedRotationProxyServiceServer) BatchGetRotations(ctx context.Context, req *BatchGetRotationsRequest) (*BatchGetRotationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetRotations not implemented")
}

func RegisterRotationProxyServiceServer(s prpc.Registrar, srv RotationProxyServiceServer) {
	s.RegisterService(&_RotationProxyService_serviceDesc, srv)
}

func _RotationProxyService_BatchUpdateRotations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchUpdateRotationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotationProxyServiceServer).BatchUpdateRotations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotation_proxy.RotationProxyService/BatchUpdateRotations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotationProxyServiceServer).BatchUpdateRotations(ctx, req.(*BatchUpdateRotationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RotationProxyService_BatchGetRotations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetRotationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RotationProxyServiceServer).BatchGetRotations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rotation_proxy.RotationProxyService/BatchGetRotations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RotationProxyServiceServer).BatchGetRotations(ctx, req.(*BatchGetRotationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RotationProxyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rotation_proxy.RotationProxyService",
	HandlerType: (*RotationProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchUpdateRotations",
			Handler:    _RotationProxyService_BatchUpdateRotations_Handler,
		},
		{
			MethodName: "BatchGetRotations",
			Handler:    _RotationProxyService_BatchGetRotations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/appengine/rotation-proxy/proto/rotation_proxy.proto",
}
