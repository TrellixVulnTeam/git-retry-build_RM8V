// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event_id.proto

package fleet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type EventID struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventID) Reset()         { *m = EventID{} }
func (m *EventID) String() string { return proto.CompactTextString(m) }
func (*EventID) ProtoMessage()    {}
func (*EventID) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a3b58c5ac123ae0, []int{0}
}

func (m *EventID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventID.Unmarshal(m, b)
}
func (m *EventID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventID.Marshal(b, m, deterministic)
}
func (m *EventID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventID.Merge(m, src)
}
func (m *EventID) XXX_Size() int {
	return xxx_messageInfo_EventID.Size(m)
}
func (m *EventID) XXX_DiscardUnknown() {
	xxx_messageInfo_EventID.DiscardUnknown(m)
}

var xxx_messageInfo_EventID proto.InternalMessageInfo

func (m *EventID) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*EventID)(nil), "fleet.EventID")
}

func init() { proto.RegisterFile("event_id.proto", fileDescriptor_1a3b58c5ac123ae0) }

var fileDescriptor_1a3b58c5ac123ae0 = []byte{
	// 87 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2d, 0x4b, 0xcd,
	0x2b, 0x89, 0xcf, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0xcb, 0x49, 0x4d,
	0x2d, 0x51, 0x92, 0xe7, 0x62, 0x77, 0x05, 0x49, 0x78, 0xba, 0x08, 0x89, 0x70, 0xb1, 0x96, 0x25,
	0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x4e, 0x9c, 0x51, 0xec,
	0x7a, 0xd6, 0x60, 0xb5, 0x49, 0x6c, 0x60, 0x9d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4,
	0x53, 0xd7, 0x27, 0x4b, 0x00, 0x00, 0x00,
}
