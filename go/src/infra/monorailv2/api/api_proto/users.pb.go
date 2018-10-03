// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api_proto/users.proto

package monorail

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// TODO(jojwang): monorail:1701, fill User with all info necessary for
// creating a user profile page.
type User struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_e517996dd141ad63, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

// Next available tag: 3
type ListReferencedUsersRequest struct {
	Trace                *RequestTrace `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
	Emails               []string      `protobuf:"bytes,2,rep,name=emails,proto3" json:"emails,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListReferencedUsersRequest) Reset()         { *m = ListReferencedUsersRequest{} }
func (m *ListReferencedUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListReferencedUsersRequest) ProtoMessage()    {}
func (*ListReferencedUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e517996dd141ad63, []int{1}
}

func (m *ListReferencedUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReferencedUsersRequest.Unmarshal(m, b)
}
func (m *ListReferencedUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReferencedUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListReferencedUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReferencedUsersRequest.Merge(m, src)
}
func (m *ListReferencedUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListReferencedUsersRequest.Size(m)
}
func (m *ListReferencedUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReferencedUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListReferencedUsersRequest proto.InternalMessageInfo

func (m *ListReferencedUsersRequest) GetTrace() *RequestTrace {
	if m != nil {
		return m.Trace
	}
	return nil
}

func (m *ListReferencedUsersRequest) GetEmails() []string {
	if m != nil {
		return m.Emails
	}
	return nil
}

type ListReferencedUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListReferencedUsersResponse) Reset()         { *m = ListReferencedUsersResponse{} }
func (m *ListReferencedUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListReferencedUsersResponse) ProtoMessage()    {}
func (*ListReferencedUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e517996dd141ad63, []int{2}
}

func (m *ListReferencedUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReferencedUsersResponse.Unmarshal(m, b)
}
func (m *ListReferencedUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReferencedUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListReferencedUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReferencedUsersResponse.Merge(m, src)
}
func (m *ListReferencedUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListReferencedUsersResponse.Size(m)
}
func (m *ListReferencedUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReferencedUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListReferencedUsersResponse proto.InternalMessageInfo

func (m *ListReferencedUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

// Next available tag: 3
type GetMembershipsRequest struct {
	Trace                *RequestTrace `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
	UserRef              *UserRef      `protobuf:"bytes,2,opt,name=user_ref,json=userRef,proto3" json:"user_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetMembershipsRequest) Reset()         { *m = GetMembershipsRequest{} }
func (m *GetMembershipsRequest) String() string { return proto.CompactTextString(m) }
func (*GetMembershipsRequest) ProtoMessage()    {}
func (*GetMembershipsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e517996dd141ad63, []int{3}
}

func (m *GetMembershipsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMembershipsRequest.Unmarshal(m, b)
}
func (m *GetMembershipsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMembershipsRequest.Marshal(b, m, deterministic)
}
func (m *GetMembershipsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMembershipsRequest.Merge(m, src)
}
func (m *GetMembershipsRequest) XXX_Size() int {
	return xxx_messageInfo_GetMembershipsRequest.Size(m)
}
func (m *GetMembershipsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMembershipsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMembershipsRequest proto.InternalMessageInfo

func (m *GetMembershipsRequest) GetTrace() *RequestTrace {
	if m != nil {
		return m.Trace
	}
	return nil
}

func (m *GetMembershipsRequest) GetUserRef() *UserRef {
	if m != nil {
		return m.UserRef
	}
	return nil
}

// Next available tag: 2
type GetMembershipsResponse struct {
	GroupRefs            []*UserRef `protobuf:"bytes,1,rep,name=group_refs,json=groupRefs,proto3" json:"group_refs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetMembershipsResponse) Reset()         { *m = GetMembershipsResponse{} }
func (m *GetMembershipsResponse) String() string { return proto.CompactTextString(m) }
func (*GetMembershipsResponse) ProtoMessage()    {}
func (*GetMembershipsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e517996dd141ad63, []int{4}
}

func (m *GetMembershipsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMembershipsResponse.Unmarshal(m, b)
}
func (m *GetMembershipsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMembershipsResponse.Marshal(b, m, deterministic)
}
func (m *GetMembershipsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMembershipsResponse.Merge(m, src)
}
func (m *GetMembershipsResponse) XXX_Size() int {
	return xxx_messageInfo_GetMembershipsResponse.Size(m)
}
func (m *GetMembershipsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMembershipsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMembershipsResponse proto.InternalMessageInfo

func (m *GetMembershipsResponse) GetGroupRefs() []*UserRef {
	if m != nil {
		return m.GroupRefs
	}
	return nil
}

// Next available tag: 4
type GetUserCommitsRequest struct {
	Trace                *RequestTrace `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
	Email                string        `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FromTimestamp        uint32        `protobuf:"fixed32,3,opt,name=from_timestamp,json=fromTimestamp,proto3" json:"from_timestamp,omitempty"`
	UntilTimestamp       uint32        `protobuf:"fixed32,4,opt,name=until_timestamp,json=untilTimestamp,proto3" json:"until_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetUserCommitsRequest) Reset()         { *m = GetUserCommitsRequest{} }
func (m *GetUserCommitsRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserCommitsRequest) ProtoMessage()    {}
func (*GetUserCommitsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e517996dd141ad63, []int{5}
}

func (m *GetUserCommitsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserCommitsRequest.Unmarshal(m, b)
}
func (m *GetUserCommitsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserCommitsRequest.Marshal(b, m, deterministic)
}
func (m *GetUserCommitsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserCommitsRequest.Merge(m, src)
}
func (m *GetUserCommitsRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserCommitsRequest.Size(m)
}
func (m *GetUserCommitsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserCommitsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserCommitsRequest proto.InternalMessageInfo

func (m *GetUserCommitsRequest) GetTrace() *RequestTrace {
	if m != nil {
		return m.Trace
	}
	return nil
}

func (m *GetUserCommitsRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GetUserCommitsRequest) GetFromTimestamp() uint32 {
	if m != nil {
		return m.FromTimestamp
	}
	return 0
}

func (m *GetUserCommitsRequest) GetUntilTimestamp() uint32 {
	if m != nil {
		return m.UntilTimestamp
	}
	return 0
}

type GetUserCommitsResponse struct {
	UserCommits          []*Commit `protobuf:"bytes,1,rep,name=user_commits,json=userCommits,proto3" json:"user_commits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetUserCommitsResponse) Reset()         { *m = GetUserCommitsResponse{} }
func (m *GetUserCommitsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserCommitsResponse) ProtoMessage()    {}
func (*GetUserCommitsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e517996dd141ad63, []int{6}
}

func (m *GetUserCommitsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserCommitsResponse.Unmarshal(m, b)
}
func (m *GetUserCommitsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserCommitsResponse.Marshal(b, m, deterministic)
}
func (m *GetUserCommitsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserCommitsResponse.Merge(m, src)
}
func (m *GetUserCommitsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserCommitsResponse.Size(m)
}
func (m *GetUserCommitsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserCommitsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserCommitsResponse proto.InternalMessageInfo

func (m *GetUserCommitsResponse) GetUserCommits() []*Commit {
	if m != nil {
		return m.UserCommits
	}
	return nil
}

// Next available tag: 3
type GetUserStarCountRequest struct {
	Trace                *RequestTrace `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
	UserRef              *UserRef      `protobuf:"bytes,2,opt,name=user_ref,json=userRef,proto3" json:"user_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetUserStarCountRequest) Reset()         { *m = GetUserStarCountRequest{} }
func (m *GetUserStarCountRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserStarCountRequest) ProtoMessage()    {}
func (*GetUserStarCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e517996dd141ad63, []int{7}
}

func (m *GetUserStarCountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserStarCountRequest.Unmarshal(m, b)
}
func (m *GetUserStarCountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserStarCountRequest.Marshal(b, m, deterministic)
}
func (m *GetUserStarCountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserStarCountRequest.Merge(m, src)
}
func (m *GetUserStarCountRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserStarCountRequest.Size(m)
}
func (m *GetUserStarCountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserStarCountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserStarCountRequest proto.InternalMessageInfo

func (m *GetUserStarCountRequest) GetTrace() *RequestTrace {
	if m != nil {
		return m.Trace
	}
	return nil
}

func (m *GetUserStarCountRequest) GetUserRef() *UserRef {
	if m != nil {
		return m.UserRef
	}
	return nil
}

// Next available tag: 2
type GetUserStarCountResponse struct {
	StarCount            uint32   `protobuf:"varint,1,opt,name=star_count,json=starCount,proto3" json:"star_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserStarCountResponse) Reset()         { *m = GetUserStarCountResponse{} }
func (m *GetUserStarCountResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserStarCountResponse) ProtoMessage()    {}
func (*GetUserStarCountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e517996dd141ad63, []int{8}
}

func (m *GetUserStarCountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserStarCountResponse.Unmarshal(m, b)
}
func (m *GetUserStarCountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserStarCountResponse.Marshal(b, m, deterministic)
}
func (m *GetUserStarCountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserStarCountResponse.Merge(m, src)
}
func (m *GetUserStarCountResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserStarCountResponse.Size(m)
}
func (m *GetUserStarCountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserStarCountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserStarCountResponse proto.InternalMessageInfo

func (m *GetUserStarCountResponse) GetStarCount() uint32 {
	if m != nil {
		return m.StarCount
	}
	return 0
}

// Next available tag: 4
type StarUserRequest struct {
	Trace                *RequestTrace `protobuf:"bytes,1,opt,name=trace,proto3" json:"trace,omitempty"`
	UserRef              *UserRef      `protobuf:"bytes,2,opt,name=user_ref,json=userRef,proto3" json:"user_ref,omitempty"`
	Starred              bool          `protobuf:"varint,3,opt,name=starred,proto3" json:"starred,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *StarUserRequest) Reset()         { *m = StarUserRequest{} }
func (m *StarUserRequest) String() string { return proto.CompactTextString(m) }
func (*StarUserRequest) ProtoMessage()    {}
func (*StarUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e517996dd141ad63, []int{9}
}

func (m *StarUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StarUserRequest.Unmarshal(m, b)
}
func (m *StarUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StarUserRequest.Marshal(b, m, deterministic)
}
func (m *StarUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StarUserRequest.Merge(m, src)
}
func (m *StarUserRequest) XXX_Size() int {
	return xxx_messageInfo_StarUserRequest.Size(m)
}
func (m *StarUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StarUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StarUserRequest proto.InternalMessageInfo

func (m *StarUserRequest) GetTrace() *RequestTrace {
	if m != nil {
		return m.Trace
	}
	return nil
}

func (m *StarUserRequest) GetUserRef() *UserRef {
	if m != nil {
		return m.UserRef
	}
	return nil
}

func (m *StarUserRequest) GetStarred() bool {
	if m != nil {
		return m.Starred
	}
	return false
}

// Next available tag: 2
type StarUserResponse struct {
	StarCount            uint32   `protobuf:"varint,1,opt,name=star_count,json=starCount,proto3" json:"star_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StarUserResponse) Reset()         { *m = StarUserResponse{} }
func (m *StarUserResponse) String() string { return proto.CompactTextString(m) }
func (*StarUserResponse) ProtoMessage()    {}
func (*StarUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e517996dd141ad63, []int{10}
}

func (m *StarUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StarUserResponse.Unmarshal(m, b)
}
func (m *StarUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StarUserResponse.Marshal(b, m, deterministic)
}
func (m *StarUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StarUserResponse.Merge(m, src)
}
func (m *StarUserResponse) XXX_Size() int {
	return xxx_messageInfo_StarUserResponse.Size(m)
}
func (m *StarUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StarUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StarUserResponse proto.InternalMessageInfo

func (m *StarUserResponse) GetStarCount() uint32 {
	if m != nil {
		return m.StarCount
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "monorail.User")
	proto.RegisterType((*ListReferencedUsersRequest)(nil), "monorail.ListReferencedUsersRequest")
	proto.RegisterType((*ListReferencedUsersResponse)(nil), "monorail.ListReferencedUsersResponse")
	proto.RegisterType((*GetMembershipsRequest)(nil), "monorail.GetMembershipsRequest")
	proto.RegisterType((*GetMembershipsResponse)(nil), "monorail.GetMembershipsResponse")
	proto.RegisterType((*GetUserCommitsRequest)(nil), "monorail.GetUserCommitsRequest")
	proto.RegisterType((*GetUserCommitsResponse)(nil), "monorail.GetUserCommitsResponse")
	proto.RegisterType((*GetUserStarCountRequest)(nil), "monorail.GetUserStarCountRequest")
	proto.RegisterType((*GetUserStarCountResponse)(nil), "monorail.GetUserStarCountResponse")
	proto.RegisterType((*StarUserRequest)(nil), "monorail.StarUserRequest")
	proto.RegisterType((*StarUserResponse)(nil), "monorail.StarUserResponse")
}

func init() { proto.RegisterFile("api/api_proto/users.proto", fileDescriptor_e517996dd141ad63) }

var fileDescriptor_e517996dd141ad63 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0x93, 0xe6, 0x75, 0x43, 0xd3, 0x30, 0x40, 0xea, 0x1a, 0x21, 0xcc, 0xa8, 0x15, 0x59,
	0x54, 0x29, 0xa4, 0x62, 0xc1, 0x3a, 0x0b, 0x04, 0xa2, 0x9b, 0xa1, 0x5d, 0xb1, 0x88, 0x9c, 0xe4,
	0x1a, 0x8c, 0x62, 0x8f, 0x99, 0x19, 0x7f, 0x03, 0x1f, 0xc3, 0xb7, 0xf1, 0x0f, 0x68, 0xc6, 0x76,
	0x9c, 0xd8, 0x2e, 0xa0, 0x4a, 0x74, 0x97, 0x99, 0x39, 0xf7, 0x9c, 0x73, 0x5f, 0x31, 0x9c, 0x78,
	0x71, 0x70, 0xe1, 0xc5, 0xc1, 0x22, 0x16, 0x5c, 0xf1, 0x8b, 0x44, 0xa2, 0x90, 0x53, 0xf3, 0x9b,
	0xf4, 0x42, 0x1e, 0x71, 0xe1, 0x05, 0x1b, 0xc7, 0xad, 0x82, 0x16, 0x7c, 0xf9, 0x0d, 0x57, 0x2a,
	0xc3, 0x3a, 0xce, 0x3e, 0x62, 0xc5, 0xc3, 0x90, 0x47, 0xe9, 0x1b, 0x7d, 0x03, 0x07, 0x37, 0x12,
	0x05, 0x79, 0x0c, 0x6d, 0x0c, 0xbd, 0x60, 0x63, 0x5b, 0xae, 0x35, 0xe9, 0xb3, 0xf4, 0x40, 0x8e,
	0xa1, 0x6b, 0xf8, 0x82, 0xb5, 0xdd, 0x74, 0xad, 0x49, 0x8b, 0x75, 0xf4, 0xf1, 0xfd, 0x9a, 0x2e,
	0xc1, 0xf9, 0x18, 0x48, 0xc5, 0xd0, 0x47, 0x81, 0xd1, 0x0a, 0xd7, 0x9a, 0x44, 0x32, 0xfc, 0x9e,
	0xa0, 0x54, 0xe4, 0x1c, 0xda, 0x4a, 0x78, 0x2b, 0x34, 0x64, 0x83, 0xd9, 0x78, 0x9a, 0x9b, 0x9d,
	0x66, 0x88, 0x6b, 0xfd, 0xca, 0x52, 0x10, 0x19, 0x43, 0xc7, 0xa8, 0x49, 0xbb, 0xe9, 0xb6, 0x26,
	0x7d, 0x96, 0x9d, 0xe8, 0x1c, 0x9e, 0xd6, 0x6a, 0xc8, 0x98, 0x47, 0x12, 0xc9, 0x29, 0xb4, 0x4d,
	0x41, 0x6c, 0xcb, 0x6d, 0x4d, 0x06, 0xb3, 0x61, 0x21, 0xa2, 0x71, 0x2c, 0x7d, 0xa4, 0x12, 0x9e,
	0xbc, 0x43, 0x75, 0x85, 0xe1, 0x12, 0x85, 0xfc, 0x1a, 0xc4, 0x77, 0xf4, 0x78, 0x0e, 0x3d, 0x53,
	0x08, 0x81, 0xbe, 0xa9, 0xc4, 0x60, 0xf6, 0xb0, 0xa4, 0x87, 0x3e, 0x33, 0xb5, 0x62, 0xe8, 0xd3,
	0x0f, 0x30, 0x2e, 0x8b, 0x66, 0xa6, 0x5f, 0x01, 0x7c, 0x11, 0x3c, 0x89, 0x35, 0x51, 0xee, 0xbc,
	0x86, 0xa9, 0x6f, 0x40, 0x0c, 0x7d, 0x49, 0x7f, 0x5a, 0x26, 0x03, 0xfd, 0x32, 0xe7, 0x61, 0x18,
	0xa8, 0x3b, 0x66, 0xb0, 0x6d, 0x70, 0x73, 0xb7, 0xc1, 0x67, 0x30, 0xf4, 0x05, 0x0f, 0x17, 0x2a,
	0x08, 0x51, 0x2a, 0x2f, 0x8c, 0xed, 0x96, 0x6b, 0x4d, 0xba, 0xec, 0x50, 0xdf, 0x5e, 0xe7, 0x97,
	0xe4, 0x25, 0x1c, 0x25, 0x91, 0x0a, 0x36, 0x3b, 0xb8, 0x03, 0x83, 0x1b, 0x9a, 0xeb, 0x2d, 0x90,
	0x5e, 0x99, 0xcc, 0xf7, 0xcc, 0x66, 0x99, 0x5f, 0xc2, 0x03, 0x53, 0xc1, 0x55, 0x7a, 0x9f, 0xe5,
	0x3e, 0x2a, 0x4c, 0xa7, 0x01, 0x6c, 0x90, 0x14, 0xc1, 0x34, 0x81, 0xe3, 0x8c, 0xee, 0x93, 0xf2,
	0xc4, 0x9c, 0x27, 0x91, 0xba, 0x8f, 0xfe, 0xbd, 0x05, 0xbb, 0x2a, 0x9b, 0xe5, 0xf1, 0x0c, 0x40,
	0x2a, 0x4f, 0xe7, 0x91, 0x44, 0xca, 0x88, 0x1f, 0xb2, 0xbe, 0xcc, 0x61, 0xf4, 0x87, 0x05, 0x47,
	0x3a, 0x28, 0xe5, 0xfc, 0xef, 0x56, 0x89, 0x0d, 0x5d, 0x2d, 0x2e, 0x70, 0x6d, 0x3a, 0xd7, 0x63,
	0xf9, 0x91, 0xbe, 0x86, 0x51, 0x61, 0xe4, 0x9f, 0xcc, 0xcf, 0x7e, 0xb5, 0xa0, 0x6d, 0x96, 0x8c,
	0x4c, 0xa1, 0x9b, 0x55, 0x80, 0x54, 0xd5, 0x9d, 0xd2, 0xae, 0xd1, 0x06, 0x59, 0xc3, 0xa3, 0x9a,
	0x5d, 0x25, 0xa7, 0x05, 0xf0, 0xf6, 0xbf, 0x0b, 0xe7, 0xec, 0x2f, 0xa8, 0xd4, 0x3c, 0x6d, 0x90,
	0x1b, 0x18, 0xee, 0xef, 0x15, 0x79, 0x5e, 0x84, 0xd6, 0xae, 0xb9, 0xe3, 0xde, 0x0e, 0x28, 0xd1,
	0xee, 0x0c, 0x6d, 0x89, 0xb6, 0xba, 0x7b, 0x25, 0xda, 0x9a, 0x79, 0xa7, 0x0d, 0xf2, 0x19, 0x46,
	0xe5, 0x29, 0x22, 0x2f, 0x2a, 0x71, 0xe5, 0xc1, 0x76, 0xe8, 0x9f, 0x20, 0x5b, 0xf2, 0x39, 0xf4,
	0xf2, 0xee, 0x92, 0x93, 0x22, 0xa2, 0x34, 0x7a, 0x8e, 0x53, 0xf7, 0x94, 0x93, 0x2c, 0x3b, 0xe6,
	0x1b, 0x70, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x0f, 0x48, 0x0e, 0x68, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	GetUser(ctx context.Context, in *UserRef, opts ...grpc.CallOption) (*User, error)
	ListReferencedUsers(ctx context.Context, in *ListReferencedUsersRequest, opts ...grpc.CallOption) (*ListReferencedUsersResponse, error)
	GetMemberships(ctx context.Context, in *GetMembershipsRequest, opts ...grpc.CallOption) (*GetMembershipsResponse, error)
	GetUserCommits(ctx context.Context, in *GetUserCommitsRequest, opts ...grpc.CallOption) (*GetUserCommitsResponse, error)
	GetUserStarCount(ctx context.Context, in *GetUserStarCountRequest, opts ...grpc.CallOption) (*GetUserStarCountResponse, error)
	StarUser(ctx context.Context, in *StarUserRequest, opts ...grpc.CallOption) (*StarUserResponse, error)
}
type usersPRPCClient struct {
	client *prpc.Client
}

func NewUsersPRPCClient(client *prpc.Client) UsersClient {
	return &usersPRPCClient{client}
}

func (c *usersPRPCClient) GetUser(ctx context.Context, in *UserRef, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.client.Call(ctx, "monorail.Users", "GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersPRPCClient) ListReferencedUsers(ctx context.Context, in *ListReferencedUsersRequest, opts ...grpc.CallOption) (*ListReferencedUsersResponse, error) {
	out := new(ListReferencedUsersResponse)
	err := c.client.Call(ctx, "monorail.Users", "ListReferencedUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersPRPCClient) GetMemberships(ctx context.Context, in *GetMembershipsRequest, opts ...grpc.CallOption) (*GetMembershipsResponse, error) {
	out := new(GetMembershipsResponse)
	err := c.client.Call(ctx, "monorail.Users", "GetMemberships", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersPRPCClient) GetUserCommits(ctx context.Context, in *GetUserCommitsRequest, opts ...grpc.CallOption) (*GetUserCommitsResponse, error) {
	out := new(GetUserCommitsResponse)
	err := c.client.Call(ctx, "monorail.Users", "GetUserCommits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersPRPCClient) GetUserStarCount(ctx context.Context, in *GetUserStarCountRequest, opts ...grpc.CallOption) (*GetUserStarCountResponse, error) {
	out := new(GetUserStarCountResponse)
	err := c.client.Call(ctx, "monorail.Users", "GetUserStarCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersPRPCClient) StarUser(ctx context.Context, in *StarUserRequest, opts ...grpc.CallOption) (*StarUserResponse, error) {
	out := new(StarUserResponse)
	err := c.client.Call(ctx, "monorail.Users", "StarUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) GetUser(ctx context.Context, in *UserRef, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/monorail.Users/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) ListReferencedUsers(ctx context.Context, in *ListReferencedUsersRequest, opts ...grpc.CallOption) (*ListReferencedUsersResponse, error) {
	out := new(ListReferencedUsersResponse)
	err := c.cc.Invoke(ctx, "/monorail.Users/ListReferencedUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetMemberships(ctx context.Context, in *GetMembershipsRequest, opts ...grpc.CallOption) (*GetMembershipsResponse, error) {
	out := new(GetMembershipsResponse)
	err := c.cc.Invoke(ctx, "/monorail.Users/GetMemberships", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserCommits(ctx context.Context, in *GetUserCommitsRequest, opts ...grpc.CallOption) (*GetUserCommitsResponse, error) {
	out := new(GetUserCommitsResponse)
	err := c.cc.Invoke(ctx, "/monorail.Users/GetUserCommits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetUserStarCount(ctx context.Context, in *GetUserStarCountRequest, opts ...grpc.CallOption) (*GetUserStarCountResponse, error) {
	out := new(GetUserStarCountResponse)
	err := c.cc.Invoke(ctx, "/monorail.Users/GetUserStarCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) StarUser(ctx context.Context, in *StarUserRequest, opts ...grpc.CallOption) (*StarUserResponse, error) {
	out := new(StarUserResponse)
	err := c.cc.Invoke(ctx, "/monorail.Users/StarUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	GetUser(context.Context, *UserRef) (*User, error)
	ListReferencedUsers(context.Context, *ListReferencedUsersRequest) (*ListReferencedUsersResponse, error)
	GetMemberships(context.Context, *GetMembershipsRequest) (*GetMembershipsResponse, error)
	GetUserCommits(context.Context, *GetUserCommitsRequest) (*GetUserCommitsResponse, error)
	GetUserStarCount(context.Context, *GetUserStarCountRequest) (*GetUserStarCountResponse, error)
	StarUser(context.Context, *StarUserRequest) (*StarUserResponse, error)
}

func RegisterUsersServer(s prpc.Registrar, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.Users/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUser(ctx, req.(*UserRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_ListReferencedUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReferencedUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).ListReferencedUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.Users/ListReferencedUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).ListReferencedUsers(ctx, req.(*ListReferencedUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetMemberships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMembershipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetMemberships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.Users/GetMemberships",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetMemberships(ctx, req.(*GetMembershipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserCommits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCommitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserCommits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.Users/GetUserCommits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserCommits(ctx, req.(*GetUserCommitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_GetUserStarCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStarCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).GetUserStarCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.Users/GetUserStarCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).GetUserStarCount(ctx, req.(*GetUserStarCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_StarUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StarUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).StarUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monorail.Users/StarUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).StarUser(ctx, req.(*StarUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monorail.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Users_GetUser_Handler,
		},
		{
			MethodName: "ListReferencedUsers",
			Handler:    _Users_ListReferencedUsers_Handler,
		},
		{
			MethodName: "GetMemberships",
			Handler:    _Users_GetMemberships_Handler,
		},
		{
			MethodName: "GetUserCommits",
			Handler:    _Users_GetUserCommits_Handler,
		},
		{
			MethodName: "GetUserStarCount",
			Handler:    _Users_GetUserStarCount_Handler,
		},
		{
			MethodName: "StarUser",
			Handler:    _Users_StarUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api_proto/users.proto",
}
