// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go-internal/testutil.proto

package testutil

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EchoTest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoTest) Reset()         { *m = EchoTest{} }
func (m *EchoTest) String() string { return proto.CompactTextString(m) }
func (*EchoTest) ProtoMessage()    {}
func (*EchoTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2d367dd7245dca, []int{0}
}
func (m *EchoTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoTest.Unmarshal(m, b)
}
func (m *EchoTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoTest.Marshal(b, m, deterministic)
}
func (m *EchoTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoTest.Merge(m, src)
}
func (m *EchoTest) XXX_Size() int {
	return xxx_messageInfo_EchoTest.Size(m)
}
func (m *EchoTest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoTest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoTest proto.InternalMessageInfo

type EchoTest_Request struct {
	Delay                uint64   `protobuf:"varint,1,opt,name=delay,proto3" json:"delay,omitempty"`
	Echo                 string   `protobuf:"bytes,2,opt,name=echo,proto3" json:"echo,omitempty"`
	TriggerError         bool     `protobuf:"varint,3,opt,name=trigger_error,json=triggerError,proto3" json:"trigger_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoTest_Request) Reset()         { *m = EchoTest_Request{} }
func (m *EchoTest_Request) String() string { return proto.CompactTextString(m) }
func (*EchoTest_Request) ProtoMessage()    {}
func (*EchoTest_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2d367dd7245dca, []int{0, 0}
}
func (m *EchoTest_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoTest_Request.Unmarshal(m, b)
}
func (m *EchoTest_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoTest_Request.Marshal(b, m, deterministic)
}
func (m *EchoTest_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoTest_Request.Merge(m, src)
}
func (m *EchoTest_Request) XXX_Size() int {
	return xxx_messageInfo_EchoTest_Request.Size(m)
}
func (m *EchoTest_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoTest_Request.DiscardUnknown(m)
}

var xxx_messageInfo_EchoTest_Request proto.InternalMessageInfo

func (m *EchoTest_Request) GetDelay() uint64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

func (m *EchoTest_Request) GetEcho() string {
	if m != nil {
		return m.Echo
	}
	return ""
}

func (m *EchoTest_Request) GetTriggerError() bool {
	if m != nil {
		return m.TriggerError
	}
	return false
}

type EchoTest_Reply struct {
	Echo                 string   `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoTest_Reply) Reset()         { *m = EchoTest_Reply{} }
func (m *EchoTest_Reply) String() string { return proto.CompactTextString(m) }
func (*EchoTest_Reply) ProtoMessage()    {}
func (*EchoTest_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2d367dd7245dca, []int{0, 1}
}
func (m *EchoTest_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoTest_Reply.Unmarshal(m, b)
}
func (m *EchoTest_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoTest_Reply.Marshal(b, m, deterministic)
}
func (m *EchoTest_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoTest_Reply.Merge(m, src)
}
func (m *EchoTest_Reply) XXX_Size() int {
	return xxx_messageInfo_EchoTest_Reply.Size(m)
}
func (m *EchoTest_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoTest_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_EchoTest_Reply proto.InternalMessageInfo

func (m *EchoTest_Reply) GetEcho() string {
	if m != nil {
		return m.Echo
	}
	return ""
}

type EchoStreamTest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoStreamTest) Reset()         { *m = EchoStreamTest{} }
func (m *EchoStreamTest) String() string { return proto.CompactTextString(m) }
func (*EchoStreamTest) ProtoMessage()    {}
func (*EchoStreamTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2d367dd7245dca, []int{1}
}
func (m *EchoStreamTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoStreamTest.Unmarshal(m, b)
}
func (m *EchoStreamTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoStreamTest.Marshal(b, m, deterministic)
}
func (m *EchoStreamTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoStreamTest.Merge(m, src)
}
func (m *EchoStreamTest) XXX_Size() int {
	return xxx_messageInfo_EchoStreamTest.Size(m)
}
func (m *EchoStreamTest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoStreamTest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoStreamTest proto.InternalMessageInfo

type EchoStreamTest_Request struct {
	Delay                uint64   `protobuf:"varint,1,opt,name=delay,proto3" json:"delay,omitempty"`
	Echo                 string   `protobuf:"bytes,2,opt,name=echo,proto3" json:"echo,omitempty"`
	TriggerError         bool     `protobuf:"varint,3,opt,name=trigger_error,json=triggerError,proto3" json:"trigger_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoStreamTest_Request) Reset()         { *m = EchoStreamTest_Request{} }
func (m *EchoStreamTest_Request) String() string { return proto.CompactTextString(m) }
func (*EchoStreamTest_Request) ProtoMessage()    {}
func (*EchoStreamTest_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2d367dd7245dca, []int{1, 0}
}
func (m *EchoStreamTest_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoStreamTest_Request.Unmarshal(m, b)
}
func (m *EchoStreamTest_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoStreamTest_Request.Marshal(b, m, deterministic)
}
func (m *EchoStreamTest_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoStreamTest_Request.Merge(m, src)
}
func (m *EchoStreamTest_Request) XXX_Size() int {
	return xxx_messageInfo_EchoStreamTest_Request.Size(m)
}
func (m *EchoStreamTest_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoStreamTest_Request.DiscardUnknown(m)
}

var xxx_messageInfo_EchoStreamTest_Request proto.InternalMessageInfo

func (m *EchoStreamTest_Request) GetDelay() uint64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

func (m *EchoStreamTest_Request) GetEcho() string {
	if m != nil {
		return m.Echo
	}
	return ""
}

func (m *EchoStreamTest_Request) GetTriggerError() bool {
	if m != nil {
		return m.TriggerError
	}
	return false
}

type EchoStreamTest_Reply struct {
	Echo                 string   `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoStreamTest_Reply) Reset()         { *m = EchoStreamTest_Reply{} }
func (m *EchoStreamTest_Reply) String() string { return proto.CompactTextString(m) }
func (*EchoStreamTest_Reply) ProtoMessage()    {}
func (*EchoStreamTest_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2d367dd7245dca, []int{1, 1}
}
func (m *EchoStreamTest_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoStreamTest_Reply.Unmarshal(m, b)
}
func (m *EchoStreamTest_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoStreamTest_Reply.Marshal(b, m, deterministic)
}
func (m *EchoStreamTest_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoStreamTest_Reply.Merge(m, src)
}
func (m *EchoStreamTest_Reply) XXX_Size() int {
	return xxx_messageInfo_EchoStreamTest_Reply.Size(m)
}
func (m *EchoStreamTest_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoStreamTest_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_EchoStreamTest_Reply proto.InternalMessageInfo

func (m *EchoStreamTest_Reply) GetEcho() string {
	if m != nil {
		return m.Echo
	}
	return ""
}

type EchoDuplexTest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoDuplexTest) Reset()         { *m = EchoDuplexTest{} }
func (m *EchoDuplexTest) String() string { return proto.CompactTextString(m) }
func (*EchoDuplexTest) ProtoMessage()    {}
func (*EchoDuplexTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2d367dd7245dca, []int{2}
}
func (m *EchoDuplexTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoDuplexTest.Unmarshal(m, b)
}
func (m *EchoDuplexTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoDuplexTest.Marshal(b, m, deterministic)
}
func (m *EchoDuplexTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoDuplexTest.Merge(m, src)
}
func (m *EchoDuplexTest) XXX_Size() int {
	return xxx_messageInfo_EchoDuplexTest.Size(m)
}
func (m *EchoDuplexTest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoDuplexTest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoDuplexTest proto.InternalMessageInfo

type EchoDuplexTest_Request struct {
	Echo                 string   `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
	TriggerError         bool     `protobuf:"varint,2,opt,name=trigger_error,json=triggerError,proto3" json:"trigger_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoDuplexTest_Request) Reset()         { *m = EchoDuplexTest_Request{} }
func (m *EchoDuplexTest_Request) String() string { return proto.CompactTextString(m) }
func (*EchoDuplexTest_Request) ProtoMessage()    {}
func (*EchoDuplexTest_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2d367dd7245dca, []int{2, 0}
}
func (m *EchoDuplexTest_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoDuplexTest_Request.Unmarshal(m, b)
}
func (m *EchoDuplexTest_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoDuplexTest_Request.Marshal(b, m, deterministic)
}
func (m *EchoDuplexTest_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoDuplexTest_Request.Merge(m, src)
}
func (m *EchoDuplexTest_Request) XXX_Size() int {
	return xxx_messageInfo_EchoDuplexTest_Request.Size(m)
}
func (m *EchoDuplexTest_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoDuplexTest_Request.DiscardUnknown(m)
}

var xxx_messageInfo_EchoDuplexTest_Request proto.InternalMessageInfo

func (m *EchoDuplexTest_Request) GetEcho() string {
	if m != nil {
		return m.Echo
	}
	return ""
}

func (m *EchoDuplexTest_Request) GetTriggerError() bool {
	if m != nil {
		return m.TriggerError
	}
	return false
}

type EchoDuplexTest_Reply struct {
	Echo                 string   `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoDuplexTest_Reply) Reset()         { *m = EchoDuplexTest_Reply{} }
func (m *EchoDuplexTest_Reply) String() string { return proto.CompactTextString(m) }
func (*EchoDuplexTest_Reply) ProtoMessage()    {}
func (*EchoDuplexTest_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2d367dd7245dca, []int{2, 1}
}
func (m *EchoDuplexTest_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoDuplexTest_Reply.Unmarshal(m, b)
}
func (m *EchoDuplexTest_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoDuplexTest_Reply.Marshal(b, m, deterministic)
}
func (m *EchoDuplexTest_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoDuplexTest_Reply.Merge(m, src)
}
func (m *EchoDuplexTest_Reply) XXX_Size() int {
	return xxx_messageInfo_EchoDuplexTest_Reply.Size(m)
}
func (m *EchoDuplexTest_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoDuplexTest_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_EchoDuplexTest_Reply proto.InternalMessageInfo

func (m *EchoDuplexTest_Reply) GetEcho() string {
	if m != nil {
		return m.Echo
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoTest)(nil), "testutil.EchoTest")
	proto.RegisterType((*EchoTest_Request)(nil), "testutil.EchoTest.Request")
	proto.RegisterType((*EchoTest_Reply)(nil), "testutil.EchoTest.Reply")
	proto.RegisterType((*EchoStreamTest)(nil), "testutil.EchoStreamTest")
	proto.RegisterType((*EchoStreamTest_Request)(nil), "testutil.EchoStreamTest.Request")
	proto.RegisterType((*EchoStreamTest_Reply)(nil), "testutil.EchoStreamTest.Reply")
	proto.RegisterType((*EchoDuplexTest)(nil), "testutil.EchoDuplexTest")
	proto.RegisterType((*EchoDuplexTest_Request)(nil), "testutil.EchoDuplexTest.Request")
	proto.RegisterType((*EchoDuplexTest_Reply)(nil), "testutil.EchoDuplexTest.Reply")
}

func init() { proto.RegisterFile("go-internal/testutil.proto", fileDescriptor_0a2d367dd7245dca) }

var fileDescriptor_0a2d367dd7245dca = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0x99, 0x3e, 0xed, 0x63, 0x5c, 0x5f, 0x0e, 0x8b, 0x87, 0x10, 0x41, 0x42, 0x45, 0xcc,
	0xc5, 0xa4, 0xe8, 0x17, 0x90, 0x62, 0xbf, 0x40, 0x5a, 0x44, 0xbc, 0x48, 0x1a, 0x87, 0x24, 0xb0,
	0x76, 0xd3, 0xc9, 0x44, 0xcc, 0xd1, 0x93, 0x9f, 0x5a, 0x90, 0xa4, 0x2f, 0x4b, 0xd3, 0x2a, 0x78,
	0xf2, 0xb6, 0xf3, 0xb2, 0xff, 0xff, 0x8f, 0x99, 0x11, 0x4e, 0xa2, 0xaf, 0xb2, 0x19, 0x23, 0xcd,
	0x22, 0x15, 0x30, 0x16, 0x5c, 0x72, 0xa6, 0xfc, 0x9c, 0x34, 0x6b, 0x69, 0xad, 0xe2, 0xfe, 0x3b,
	0x08, 0x6b, 0x14, 0xa7, 0x7a, 0x82, 0x05, 0x3b, 0x0f, 0x62, 0x2f, 0xc4, 0x79, 0x89, 0x05, 0xcb,
	0x13, 0xd1, 0x7b, 0x46, 0x15, 0x55, 0x36, 0xb8, 0xe0, 0x75, 0xc3, 0x45, 0x20, 0xa5, 0xe8, 0x62,
	0x9c, 0x6a, 0xbb, 0xe3, 0x82, 0xb7, 0x1f, 0x36, 0x6f, 0x79, 0x2e, 0x8e, 0x98, 0xb2, 0x24, 0x41,
	0x7a, 0x42, 0x22, 0x4d, 0xf6, 0x3f, 0x17, 0x3c, 0x2b, 0x3c, 0x5c, 0x26, 0x47, 0x75, 0xce, 0x39,
	0x15, 0xbd, 0x10, 0x73, 0x65, 0x14, 0xc0, 0x28, 0xf4, 0x3f, 0x40, 0x1c, 0xd7, 0x0c, 0x63, 0x26,
	0x8c, 0x5e, 0xfe, 0x92, 0x64, 0xbe, 0x00, 0xb9, 0x2b, 0x73, 0x85, 0x6f, 0x0d, 0xc8, 0xd0, 0x80,
	0xec, 0xf8, 0xb0, 0x6d, 0xd9, 0xf9, 0xa5, 0xe5, 0xf5, 0x27, 0x88, 0x83, 0xda, 0x69, 0x8c, 0xf4,
	0x9a, 0xc5, 0x28, 0x6f, 0xcd, 0x3e, 0xa4, 0xe3, 0xaf, 0xf7, 0xb6, 0xca, 0xf9, 0x4b, 0x1a, 0xc7,
	0xde, 0x59, 0xab, 0x5d, 0x26, 0xed, 0x69, 0x4a, 0x77, 0xb3, 0xd7, 0x54, 0xd6, 0x6a, 0x67, 0x3f,
	0x74, 0xe4, 0xaa, 0x1a, 0x80, 0xbc, 0x6f, 0x8f, 0xa6, 0xad, 0x6a, 0x2a, 0xdf, 0xa9, 0x6e, 0x74,
	0xe4, 0xaa, 0xf2, 0x60, 0x00, 0xc3, 0xcb, 0xc7, 0x8b, 0x29, 0x12, 0x57, 0x3e, 0x63, 0x9c, 0x06,
	0xcd, 0x33, 0x48, 0x74, 0xb0, 0x75, 0xb9, 0xd3, 0xff, 0xcd, 0xe9, 0xde, 0x7c, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x48, 0x9f, 0x5d, 0xfb, 0xd8, 0x02, 0x00, 0x00,
}