// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package fleet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Event struct {
	Id *EventID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Asset:
	//	*Event_RackId
	//	*Event_MachineId
	Asset                isEvent_Asset        `protobuf_oneof:"asset"`
	EventLabel           EventType            `protobuf:"varint,4,opt,name=event_label,json=eventLabel,proto3,enum=fleet.EventType" json:"event_label,omitempty"`
	OldValue             string               `protobuf:"bytes,5,opt,name=old_value,json=oldValue,proto3" json:"old_value,omitempty"`
	NewValue             string               `protobuf:"bytes,6,opt,name=new_value,json=newValue,proto3" json:"new_value,omitempty"`
	UpdatedTime          *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	User                 *Event_User          `protobuf:"bytes,8,opt,name=user,proto3" json:"user,omitempty"`
	Comment              string               `protobuf:"bytes,9,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() *EventID {
	if m != nil {
		return m.Id
	}
	return nil
}

type isEvent_Asset interface {
	isEvent_Asset()
}

type Event_RackId struct {
	RackId *RackID `protobuf:"bytes,2,opt,name=rack_id,json=rackId,proto3,oneof"`
}

type Event_MachineId struct {
	MachineId *MachineID `protobuf:"bytes,3,opt,name=machine_id,json=machineId,proto3,oneof"`
}

func (*Event_RackId) isEvent_Asset() {}

func (*Event_MachineId) isEvent_Asset() {}

func (m *Event) GetAsset() isEvent_Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *Event) GetRackId() *RackID {
	if x, ok := m.GetAsset().(*Event_RackId); ok {
		return x.RackId
	}
	return nil
}

func (m *Event) GetMachineId() *MachineID {
	if x, ok := m.GetAsset().(*Event_MachineId); ok {
		return x.MachineId
	}
	return nil
}

func (m *Event) GetEventLabel() EventType {
	if m != nil {
		return m.EventLabel
	}
	return EventType_EVENT_INVALID
}

func (m *Event) GetOldValue() string {
	if m != nil {
		return m.OldValue
	}
	return ""
}

func (m *Event) GetNewValue() string {
	if m != nil {
		return m.NewValue
	}
	return ""
}

func (m *Event) GetUpdatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedTime
	}
	return nil
}

func (m *Event) GetUser() *Event_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Event) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_RackId)(nil),
		(*Event_MachineId)(nil),
	}
}

type Event_User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_User) Reset()         { *m = Event_User{} }
func (m *Event_User) String() string { return proto.CompactTextString(m) }
func (*Event_User) ProtoMessage()    {}
func (*Event_User) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0, 0}
}

func (m *Event_User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_User.Unmarshal(m, b)
}
func (m *Event_User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_User.Marshal(b, m, deterministic)
}
func (m *Event_User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_User.Merge(m, src)
}
func (m *Event_User) XXX_Size() int {
	return xxx_messageInfo_Event_User.Size(m)
}
func (m *Event_User) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_User.DiscardUnknown(m)
}

var xxx_messageInfo_Event_User proto.InternalMessageInfo

func (m *Event_User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event_User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "fleet.Event")
	proto.RegisterType((*Event_User)(nil), "fleet.Event.User")
}

func init() { proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4d, 0xcb, 0xda, 0x40,
	0x14, 0x85, 0x1b, 0x4d, 0x8c, 0x19, 0xab, 0xd8, 0xa1, 0x8b, 0x21, 0x85, 0x56, 0x0a, 0x85, 0xac,
	0xc6, 0x7e, 0x2c, 0x4b, 0x37, 0x62, 0xa1, 0x42, 0xbb, 0x19, 0x6c, 0x17, 0xdd, 0x84, 0x31, 0x73,
	0xb5, 0xc1, 0xc9, 0x07, 0xc9, 0x44, 0xf1, 0x3f, 0xf5, 0x47, 0xbe, 0xcc, 0x9d, 0x44, 0xdf, 0xdd,
	0x9c, 0x7b, 0x9e, 0xc3, 0x1c, 0x0e, 0x99, 0xc1, 0x05, 0x4a, 0xc3, 0xeb, 0xa6, 0x32, 0x15, 0x0d,
	0x8e, 0x1a, 0xc0, 0xc4, 0xef, 0x4e, 0x55, 0x75, 0xd2, 0xb0, 0xc6, 0xe3, 0xa1, 0x3b, 0xae, 0x4d,
	0x5e, 0x40, 0x6b, 0x64, 0x51, 0x3b, 0x2e, 0x5e, 0x60, 0x28, 0xcd, 0x55, 0xaf, 0x97, 0x4e, 0x9b,
	0x5b, 0x0d, 0xc3, 0xa5, 0x90, 0xd9, 0xbf, 0xbc, 0x84, 0x07, 0x33, 0x6f, 0x64, 0x76, 0xbe, 0xcb,
	0xf7, 0xff, 0xc7, 0x24, 0xf8, 0x6e, 0x53, 0xf4, 0x2d, 0x19, 0xe5, 0x8a, 0x79, 0x2b, 0x2f, 0x99,
	0x7d, 0x5e, 0x70, 0x6c, 0xc0, 0xd1, 0xd9, 0x6d, 0xc5, 0x28, 0x57, 0x34, 0x21, 0x61, 0x1f, 0x65,
	0x23, 0x84, 0xe6, 0x3d, 0x24, 0x64, 0x76, 0xde, 0x6d, 0x7f, 0xbc, 0x10, 0x13, 0xeb, 0xef, 0x14,
	0xfd, 0x44, 0xc8, 0xe3, 0x5b, 0x36, 0x46, 0x78, 0xd9, 0xc3, 0xbf, 0x9c, 0x81, 0x7c, 0xd4, 0x53,
	0x18, 0x71, 0x03, 0xa4, 0x5a, 0x1e, 0x40, 0x33, 0x7f, 0xe5, 0x25, 0x8b, 0x7b, 0x06, 0x5b, 0xec,
	0x6f, 0x35, 0x08, 0x82, 0xd0, 0x4f, 0xcb, 0xd0, 0x37, 0x24, 0xaa, 0xb4, 0x4a, 0x2f, 0x52, 0x77,
	0xc0, 0x82, 0x95, 0x97, 0x44, 0x62, 0x5a, 0x69, 0xf5, 0xc7, 0x6a, 0x6b, 0x96, 0x70, 0xed, 0xcd,
	0x89, 0x33, 0x4b, 0xb8, 0x3a, 0xf3, 0x1b, 0x79, 0xd9, 0xd5, 0x4a, 0x1a, 0x50, 0xa9, 0x5d, 0x94,
	0x85, 0xd8, 0x30, 0xe6, 0x6e, 0x6e, 0x3e, 0xcc, 0xcd, 0xf7, 0xc3, 0xdc, 0x62, 0xd6, 0xf3, 0xf6,
	0x42, 0x3f, 0x10, 0xbf, 0x6b, 0xa1, 0x61, 0x53, 0x8c, 0xbd, 0x7a, 0x5e, 0x92, 0xff, 0x6e, 0xa1,
	0x11, 0x68, 0x53, 0x46, 0xc2, 0xac, 0x2a, 0x0a, 0x28, 0x0d, 0x8b, 0xb0, 0xc0, 0x20, 0xe3, 0x8f,
	0xc4, 0xb7, 0x1c, 0xa5, 0xc4, 0x2f, 0x65, 0x01, 0xb8, 0x79, 0x24, 0xf0, 0x4d, 0x5f, 0x93, 0x00,
	0x0a, 0x99, 0x6b, 0xdc, 0x38, 0x12, 0x4e, 0x6c, 0x42, 0x12, 0xc8, 0xb6, 0x05, 0xb3, 0x89, 0xfe,
	0x86, 0xfc, 0x2b, 0x7e, 0x78, 0x98, 0x60, 0xcf, 0x2f, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15,
	0x15, 0xba, 0x4e, 0x3a, 0x02, 0x00, 0x00,
}
