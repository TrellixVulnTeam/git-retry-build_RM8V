// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api_proto/features_objects.proto

package monorail

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Next available tag: 5
type Hotlist struct {
	OwnerRef             *UserRef `protobuf:"bytes,1,opt,name=owner_ref,json=ownerRef,proto3" json:"owner_ref,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Summary              string   `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hotlist) Reset()         { *m = Hotlist{} }
func (m *Hotlist) String() string { return proto.CompactTextString(m) }
func (*Hotlist) ProtoMessage()    {}
func (*Hotlist) Descriptor() ([]byte, []int) {
	return fileDescriptor_806b6b78af767289, []int{0}
}

func (m *Hotlist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hotlist.Unmarshal(m, b)
}
func (m *Hotlist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hotlist.Marshal(b, m, deterministic)
}
func (m *Hotlist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hotlist.Merge(m, src)
}
func (m *Hotlist) XXX_Size() int {
	return xxx_messageInfo_Hotlist.Size(m)
}
func (m *Hotlist) XXX_DiscardUnknown() {
	xxx_messageInfo_Hotlist.DiscardUnknown(m)
}

var xxx_messageInfo_Hotlist proto.InternalMessageInfo

func (m *Hotlist) GetOwnerRef() *UserRef {
	if m != nil {
		return m.OwnerRef
	}
	return nil
}

func (m *Hotlist) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Hotlist) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *Hotlist) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Next available tag: 6
type HotlistItem struct {
	Issue                *Issue   `protobuf:"bytes,1,opt,name=issue,proto3" json:"issue,omitempty"`
	Rank                 uint32   `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	AdderRef             *UserRef `protobuf:"bytes,3,opt,name=adder_ref,json=adderRef,proto3" json:"adder_ref,omitempty"`
	AddedTimestamp       uint32   `protobuf:"varint,4,opt,name=added_timestamp,json=addedTimestamp,proto3" json:"added_timestamp,omitempty"`
	Note                 string   `protobuf:"bytes,5,opt,name=note,proto3" json:"note,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotlistItem) Reset()         { *m = HotlistItem{} }
func (m *HotlistItem) String() string { return proto.CompactTextString(m) }
func (*HotlistItem) ProtoMessage()    {}
func (*HotlistItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_806b6b78af767289, []int{1}
}

func (m *HotlistItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotlistItem.Unmarshal(m, b)
}
func (m *HotlistItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotlistItem.Marshal(b, m, deterministic)
}
func (m *HotlistItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotlistItem.Merge(m, src)
}
func (m *HotlistItem) XXX_Size() int {
	return xxx_messageInfo_HotlistItem.Size(m)
}
func (m *HotlistItem) XXX_DiscardUnknown() {
	xxx_messageInfo_HotlistItem.DiscardUnknown(m)
}

var xxx_messageInfo_HotlistItem proto.InternalMessageInfo

func (m *HotlistItem) GetIssue() *Issue {
	if m != nil {
		return m.Issue
	}
	return nil
}

func (m *HotlistItem) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *HotlistItem) GetAdderRef() *UserRef {
	if m != nil {
		return m.AdderRef
	}
	return nil
}

func (m *HotlistItem) GetAddedTimestamp() uint32 {
	if m != nil {
		return m.AddedTimestamp
	}
	return 0
}

func (m *HotlistItem) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func init() {
	proto.RegisterType((*Hotlist)(nil), "monorail.Hotlist")
	proto.RegisterType((*HotlistItem)(nil), "monorail.HotlistItem")
}

func init() {
	proto.RegisterFile("api/api_proto/features_objects.proto", fileDescriptor_806b6b78af767289)
}

var fileDescriptor_806b6b78af767289 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4e, 0xeb, 0x30,
	0x10, 0x86, 0x95, 0xd7, 0xf6, 0xb5, 0x75, 0x54, 0x2a, 0xbc, 0xb2, 0xb2, 0x0a, 0x15, 0x88, 0xae,
	0x52, 0x09, 0x2e, 0x41, 0xb7, 0x16, 0xac, 0x23, 0x37, 0x99, 0x48, 0x86, 0xda, 0x13, 0xd9, 0x13,
	0x21, 0x8e, 0xc0, 0x81, 0xb8, 0x1f, 0xca, 0x34, 0x11, 0x64, 0xc1, 0x6e, 0xfc, 0xcf, 0x2f, 0xfb,
	0xfb, 0x2c, 0x6e, 0x4d, 0x6b, 0x0f, 0xa6, 0xb5, 0x65, 0x1b, 0x90, 0xf0, 0xd0, 0x80, 0xa1, 0x2e,
	0x40, 0x2c, 0xf1, 0xf4, 0x0a, 0x15, 0xc5, 0x82, 0x63, 0xb9, 0x72, 0xe8, 0x31, 0x18, 0x7b, 0xce,
	0xb2, 0x69, 0xbf, 0x42, 0xe7, 0xd0, 0x5f, 0x5a, 0xd9, 0xcd, 0x74, 0x67, 0x63, 0xec, 0x60, 0x7a,
	0xd1, 0xee, 0x33, 0x11, 0xcb, 0x27, 0xa4, 0xb3, 0x8d, 0x24, 0x0b, 0xb1, 0xc6, 0x77, 0x0f, 0xa1,
	0x0c, 0xd0, 0xa8, 0x24, 0x4f, 0xf6, 0xe9, 0xc3, 0x75, 0x31, 0x3e, 0x54, 0xbc, 0x44, 0x08, 0x1a,
	0x1a, 0xbd, 0xe2, 0x8e, 0x86, 0x46, 0x4a, 0x31, 0xf7, 0xc6, 0x81, 0xfa, 0x97, 0x27, 0xfb, 0xb5,
	0xe6, 0x59, 0x2a, 0xb1, 0x8c, 0x9d, 0x73, 0x26, 0x7c, 0xa8, 0x19, 0xc7, 0xe3, 0x51, 0xe6, 0x22,
	0xad, 0x21, 0x56, 0xc1, 0xb6, 0x64, 0xd1, 0xab, 0x39, 0x6f, 0x7f, 0x47, 0xbb, 0xaf, 0x44, 0xa4,
	0x03, 0xcb, 0x91, 0xc0, 0xc9, 0x3b, 0xb1, 0x60, 0xe4, 0x81, 0x65, 0xfb, 0xc3, 0x72, 0xec, 0x63,
	0x7d, 0xd9, 0xf6, 0x18, 0xc1, 0xf8, 0x37, 0xc6, 0xd8, 0x68, 0x9e, 0x7b, 0x15, 0x53, 0xd7, 0x83,
	0xca, 0xec, 0x4f, 0x15, 0xee, 0xf4, 0x2a, 0xf7, 0x62, 0xdb, 0xcf, 0x75, 0x49, 0xd6, 0x41, 0x24,
	0xe3, 0x5a, 0x06, 0xdc, 0xe8, 0x2b, 0x8e, 0x9f, 0xc7, 0x94, 0x9d, 0x91, 0x40, 0x2d, 0x06, 0x67,
	0x24, 0x38, 0xfd, 0xe7, 0xaf, 0x7c, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x45, 0x11, 0xac, 0xcd,
	0xbb, 0x01, 0x00, 0x00,
}
