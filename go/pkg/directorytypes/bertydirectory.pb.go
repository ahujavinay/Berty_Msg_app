// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: directorytypes/bertydirectory.proto

package directorytypes

import (
	_ "berty.tech/weshnet/pkg/protocoltypes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Register struct {
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02cf209c8aed6e0, []int{0}
}
func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (m *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(m, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

type Register_Request struct {
	VerifiedCredential      []byte `protobuf:"bytes,1,opt,name=verified_credential,json=verifiedCredential,proto3" json:"verified_credential,omitempty"`
	ExpirationDate          int64  `protobuf:"varint,2,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	LockedUntilDate         int64  `protobuf:"varint,3,opt,name=locked_until_date,json=lockedUntilDate,proto3" json:"locked_until_date,omitempty"`
	AccountURI              string `protobuf:"bytes,4,opt,name=account_uri,json=accountUri,proto3" json:"account_uri,omitempty"`
	OverwriteExistingRecord bool   `protobuf:"varint,5,opt,name=overwrite_existing_record,json=overwriteExistingRecord,proto3" json:"overwrite_existing_record,omitempty"`
}

func (m *Register_Request) Reset()         { *m = Register_Request{} }
func (m *Register_Request) String() string { return proto.CompactTextString(m) }
func (*Register_Request) ProtoMessage()    {}
func (*Register_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02cf209c8aed6e0, []int{0, 0}
}
func (m *Register_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register_Request.Unmarshal(m, b)
}
func (m *Register_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register_Request.Marshal(b, m, deterministic)
}
func (m *Register_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register_Request.Merge(m, src)
}
func (m *Register_Request) XXX_Size() int {
	return xxx_messageInfo_Register_Request.Size(m)
}
func (m *Register_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Register_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Register_Request proto.InternalMessageInfo

func (m *Register_Request) GetVerifiedCredential() []byte {
	if m != nil {
		return m.VerifiedCredential
	}
	return nil
}

func (m *Register_Request) GetExpirationDate() int64 {
	if m != nil {
		return m.ExpirationDate
	}
	return 0
}

func (m *Register_Request) GetLockedUntilDate() int64 {
	if m != nil {
		return m.LockedUntilDate
	}
	return 0
}

func (m *Register_Request) GetAccountURI() string {
	if m != nil {
		return m.AccountURI
	}
	return ""
}

func (m *Register_Request) GetOverwriteExistingRecord() bool {
	if m != nil {
		return m.OverwriteExistingRecord
	}
	return false
}

type Register_Reply struct {
	DirectoryRecordToken string `protobuf:"bytes,1,opt,name=directory_record_token,json=directoryRecordToken,proto3" json:"directory_record_token,omitempty"`
	DirectoryIdentifier  string `protobuf:"bytes,2,opt,name=directory_identifier,json=directoryIdentifier,proto3" json:"directory_identifier,omitempty"`
	ExpirationDate       int64  `protobuf:"varint,3,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	UnregisterToken      string `protobuf:"bytes,4,opt,name=unregister_token,json=unregisterToken,proto3" json:"unregister_token,omitempty"`
}

func (m *Register_Reply) Reset()         { *m = Register_Reply{} }
func (m *Register_Reply) String() string { return proto.CompactTextString(m) }
func (*Register_Reply) ProtoMessage()    {}
func (*Register_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02cf209c8aed6e0, []int{0, 1}
}
func (m *Register_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register_Reply.Unmarshal(m, b)
}
func (m *Register_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register_Reply.Marshal(b, m, deterministic)
}
func (m *Register_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register_Reply.Merge(m, src)
}
func (m *Register_Reply) XXX_Size() int {
	return xxx_messageInfo_Register_Reply.Size(m)
}
func (m *Register_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Register_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Register_Reply proto.InternalMessageInfo

func (m *Register_Reply) GetDirectoryRecordToken() string {
	if m != nil {
		return m.DirectoryRecordToken
	}
	return ""
}

func (m *Register_Reply) GetDirectoryIdentifier() string {
	if m != nil {
		return m.DirectoryIdentifier
	}
	return ""
}

func (m *Register_Reply) GetExpirationDate() int64 {
	if m != nil {
		return m.ExpirationDate
	}
	return 0
}

func (m *Register_Reply) GetUnregisterToken() string {
	if m != nil {
		return m.UnregisterToken
	}
	return ""
}

type Query struct {
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02cf209c8aed6e0, []int{1}
}
func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

type Query_Request struct {
	DirectoryIdentifiers []string `protobuf:"bytes,1,rep,name=directory_identifiers,json=directoryIdentifiers,proto3" json:"directory_identifiers,omitempty"`
}

func (m *Query_Request) Reset()         { *m = Query_Request{} }
func (m *Query_Request) String() string { return proto.CompactTextString(m) }
func (*Query_Request) ProtoMessage()    {}
func (*Query_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02cf209c8aed6e0, []int{1, 0}
}
func (m *Query_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query_Request.Unmarshal(m, b)
}
func (m *Query_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query_Request.Marshal(b, m, deterministic)
}
func (m *Query_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query_Request.Merge(m, src)
}
func (m *Query_Request) XXX_Size() int {
	return xxx_messageInfo_Query_Request.Size(m)
}
func (m *Query_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Query_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Query_Request proto.InternalMessageInfo

func (m *Query_Request) GetDirectoryIdentifiers() []string {
	if m != nil {
		return m.DirectoryIdentifiers
	}
	return nil
}

type Query_Reply struct {
	DirectoryIdentifier string `protobuf:"bytes,1,opt,name=directory_identifier,json=directoryIdentifier,proto3" json:"directory_identifier,omitempty"`
	ExpiresAt           int64  `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	AccountURI          string `protobuf:"bytes,3,opt,name=account_uri,json=accountUri,proto3" json:"account_uri,omitempty"`
	VerifiedCredential  []byte `protobuf:"bytes,4,opt,name=verified_credential,json=verifiedCredential,proto3" json:"verified_credential,omitempty"`
}

func (m *Query_Reply) Reset()         { *m = Query_Reply{} }
func (m *Query_Reply) String() string { return proto.CompactTextString(m) }
func (*Query_Reply) ProtoMessage()    {}
func (*Query_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02cf209c8aed6e0, []int{1, 1}
}
func (m *Query_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query_Reply.Unmarshal(m, b)
}
func (m *Query_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query_Reply.Marshal(b, m, deterministic)
}
func (m *Query_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query_Reply.Merge(m, src)
}
func (m *Query_Reply) XXX_Size() int {
	return xxx_messageInfo_Query_Reply.Size(m)
}
func (m *Query_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Query_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Query_Reply proto.InternalMessageInfo

func (m *Query_Reply) GetDirectoryIdentifier() string {
	if m != nil {
		return m.DirectoryIdentifier
	}
	return ""
}

func (m *Query_Reply) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *Query_Reply) GetAccountURI() string {
	if m != nil {
		return m.AccountURI
	}
	return ""
}

func (m *Query_Reply) GetVerifiedCredential() []byte {
	if m != nil {
		return m.VerifiedCredential
	}
	return nil
}

type Unregister struct {
}

func (m *Unregister) Reset()         { *m = Unregister{} }
func (m *Unregister) String() string { return proto.CompactTextString(m) }
func (*Unregister) ProtoMessage()    {}
func (*Unregister) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02cf209c8aed6e0, []int{2}
}
func (m *Unregister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unregister.Unmarshal(m, b)
}
func (m *Unregister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unregister.Marshal(b, m, deterministic)
}
func (m *Unregister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unregister.Merge(m, src)
}
func (m *Unregister) XXX_Size() int {
	return xxx_messageInfo_Unregister.Size(m)
}
func (m *Unregister) XXX_DiscardUnknown() {
	xxx_messageInfo_Unregister.DiscardUnknown(m)
}

var xxx_messageInfo_Unregister proto.InternalMessageInfo

type Unregister_Request struct {
	DirectoryIdentifier  string `protobuf:"bytes,1,opt,name=directory_identifier,json=directoryIdentifier,proto3" json:"directory_identifier,omitempty"`
	DirectoryRecordToken string `protobuf:"bytes,2,opt,name=directory_record_token,json=directoryRecordToken,proto3" json:"directory_record_token,omitempty"`
	UnregisterToken      string `protobuf:"bytes,3,opt,name=unregister_token,json=unregisterToken,proto3" json:"unregister_token,omitempty"`
}

func (m *Unregister_Request) Reset()         { *m = Unregister_Request{} }
func (m *Unregister_Request) String() string { return proto.CompactTextString(m) }
func (*Unregister_Request) ProtoMessage()    {}
func (*Unregister_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02cf209c8aed6e0, []int{2, 0}
}
func (m *Unregister_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unregister_Request.Unmarshal(m, b)
}
func (m *Unregister_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unregister_Request.Marshal(b, m, deterministic)
}
func (m *Unregister_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unregister_Request.Merge(m, src)
}
func (m *Unregister_Request) XXX_Size() int {
	return xxx_messageInfo_Unregister_Request.Size(m)
}
func (m *Unregister_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Unregister_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Unregister_Request proto.InternalMessageInfo

func (m *Unregister_Request) GetDirectoryIdentifier() string {
	if m != nil {
		return m.DirectoryIdentifier
	}
	return ""
}

func (m *Unregister_Request) GetDirectoryRecordToken() string {
	if m != nil {
		return m.DirectoryRecordToken
	}
	return ""
}

func (m *Unregister_Request) GetUnregisterToken() string {
	if m != nil {
		return m.UnregisterToken
	}
	return ""
}

type Unregister_Reply struct {
}

func (m *Unregister_Reply) Reset()         { *m = Unregister_Reply{} }
func (m *Unregister_Reply) String() string { return proto.CompactTextString(m) }
func (*Unregister_Reply) ProtoMessage()    {}
func (*Unregister_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02cf209c8aed6e0, []int{2, 1}
}
func (m *Unregister_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unregister_Reply.Unmarshal(m, b)
}
func (m *Unregister_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unregister_Reply.Marshal(b, m, deterministic)
}
func (m *Unregister_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unregister_Reply.Merge(m, src)
}
func (m *Unregister_Reply) XXX_Size() int {
	return xxx_messageInfo_Unregister_Reply.Size(m)
}
func (m *Unregister_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Unregister_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Unregister_Reply proto.InternalMessageInfo

type Record struct {
	DirectoryIdentifier  string `protobuf:"bytes,1,opt,name=directory_identifier,json=directoryIdentifier,proto3" json:"directory_identifier,omitempty" gorm:"index;primaryKey;autoIncrement:false"`
	DirectoryRecordToken string `protobuf:"bytes,2,opt,name=directory_record_token,json=directoryRecordToken,proto3" json:"directory_record_token,omitempty"`
	ExpiresAt            int64  `protobuf:"varint,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	LockedUntil          int64  `protobuf:"varint,4,opt,name=locked_until,json=lockedUntil,proto3" json:"locked_until,omitempty"`
	UnregisterToken      string `protobuf:"bytes,5,opt,name=unregister_token,json=unregisterToken,proto3" json:"unregister_token,omitempty"`
	AccountURI           string `protobuf:"bytes,6,opt,name=account_uri,json=accountUri,proto3" json:"account_uri,omitempty"`
	VerifiedCredential   []byte `protobuf:"bytes,7,opt,name=verified_credential,json=verifiedCredential,proto3" json:"verified_credential,omitempty"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02cf209c8aed6e0, []int{3}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetDirectoryIdentifier() string {
	if m != nil {
		return m.DirectoryIdentifier
	}
	return ""
}

func (m *Record) GetDirectoryRecordToken() string {
	if m != nil {
		return m.DirectoryRecordToken
	}
	return ""
}

func (m *Record) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *Record) GetLockedUntil() int64 {
	if m != nil {
		return m.LockedUntil
	}
	return 0
}

func (m *Record) GetUnregisterToken() string {
	if m != nil {
		return m.UnregisterToken
	}
	return ""
}

func (m *Record) GetAccountURI() string {
	if m != nil {
		return m.AccountURI
	}
	return ""
}

func (m *Record) GetVerifiedCredential() []byte {
	if m != nil {
		return m.VerifiedCredential
	}
	return nil
}

func init() {
	proto.RegisterType((*Register)(nil), "berty.directory.v1.Register")
	proto.RegisterType((*Register_Request)(nil), "berty.directory.v1.Register.Request")
	proto.RegisterType((*Register_Reply)(nil), "berty.directory.v1.Register.Reply")
	proto.RegisterType((*Query)(nil), "berty.directory.v1.Query")
	proto.RegisterType((*Query_Request)(nil), "berty.directory.v1.Query.Request")
	proto.RegisterType((*Query_Reply)(nil), "berty.directory.v1.Query.Reply")
	proto.RegisterType((*Unregister)(nil), "berty.directory.v1.Unregister")
	proto.RegisterType((*Unregister_Request)(nil), "berty.directory.v1.Unregister.Request")
	proto.RegisterType((*Unregister_Reply)(nil), "berty.directory.v1.Unregister.Reply")
	proto.RegisterType((*Record)(nil), "berty.directory.v1.Record")
}

func init() {
	proto.RegisterFile("directorytypes/bertydirectory.proto", fileDescriptor_d02cf209c8aed6e0)
}

var fileDescriptor_d02cf209c8aed6e0 = []byte{
	// 659 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x41, 0x4f, 0xd4, 0x40,
	0x18, 0xa5, 0x5b, 0x16, 0xd8, 0x81, 0x00, 0x0e, 0xa8, 0x6b, 0x13, 0x61, 0xa9, 0x04, 0x57, 0x4d,
	0xb6, 0x22, 0x9c, 0x20, 0x31, 0x01, 0xf1, 0x40, 0x8c, 0x07, 0x47, 0xf6, 0xc2, 0xa5, 0x29, 0xed,
	0x47, 0x9d, 0x50, 0x3a, 0x75, 0x3a, 0x5d, 0xe9, 0x0f, 0xf0, 0x6e, 0xe2, 0xd9, 0x9f, 0xe1, 0xc9,
	0xab, 0x3f, 0x80, 0xa3, 0x27, 0x0e, 0xf2, 0x0b, 0xf4, 0xec, 0xc1, 0xec, 0xb4, 0xdb, 0xee, 0xae,
	0xed, 0x6e, 0x42, 0x3c, 0xed, 0xe6, 0xbd, 0x37, 0x93, 0xf9, 0xde, 0x7b, 0xd3, 0x41, 0x0f, 0x1c,
	0xca, 0xc1, 0x16, 0x8c, 0xc7, 0x22, 0x0e, 0x20, 0x34, 0x4e, 0x80, 0x8b, 0x38, 0xc3, 0x5a, 0x01,
	0x67, 0x82, 0x61, 0x2c, 0xd1, 0x56, 0x0e, 0x77, 0x36, 0xb5, 0x65, 0x97, 0xb9, 0x4c, 0xd2, 0x46,
	0xf7, 0x5f, 0xa2, 0xd4, 0x96, 0xe4, 0x8f, 0xcd, 0x3c, 0xb9, 0x5b, 0x02, 0xea, 0x97, 0x2a, 0x9a,
	0x21, 0xe0, 0xd2, 0x50, 0x00, 0xd7, 0xfe, 0x28, 0x68, 0x9a, 0xc0, 0xfb, 0x08, 0x42, 0x81, 0x0d,
	0xb4, 0xd4, 0x01, 0x4e, 0x4f, 0x29, 0x38, 0xa6, 0xcd, 0xc1, 0x01, 0x5f, 0x50, 0xcb, 0xab, 0x2b,
	0x0d, 0xa5, 0x39, 0x47, 0x70, 0x8f, 0x7a, 0x91, 0x31, 0xf8, 0x21, 0x5a, 0x80, 0x8b, 0x80, 0x72,
	0x4b, 0x50, 0xe6, 0x9b, 0x8e, 0x25, 0xa0, 0x5e, 0x69, 0x28, 0x4d, 0x95, 0xcc, 0xe7, 0xf0, 0x81,
	0x25, 0x00, 0x3f, 0x46, 0xb7, 0x3c, 0x66, 0x9f, 0x81, 0x63, 0x46, 0xbe, 0xa0, 0x5e, 0x22, 0x55,
	0xa5, 0x74, 0x21, 0x21, 0xda, 0x5d, 0x5c, 0x6a, 0x0d, 0x34, 0x6b, 0xd9, 0x36, 0x8b, 0x7c, 0x61,
	0x46, 0x9c, 0xd6, 0x27, 0x1b, 0x4a, 0xb3, 0xb6, 0x3f, 0xff, 0xf3, 0x6a, 0x15, 0xed, 0x25, 0x70,
	0x9b, 0x1c, 0x12, 0x94, 0x4a, 0xda, 0x9c, 0xe2, 0x1d, 0x74, 0x8f, 0x75, 0x80, 0x7f, 0xe0, 0x54,
	0x80, 0x09, 0x17, 0x34, 0x14, 0xd4, 0x77, 0x4d, 0x0e, 0x36, 0xe3, 0x4e, 0xbd, 0xda, 0x50, 0x9a,
	0x33, 0xe4, 0x6e, 0x26, 0x78, 0x99, 0xf2, 0x44, 0xd2, 0xda, 0x77, 0x05, 0x55, 0x09, 0x04, 0x5e,
	0x8c, 0xb7, 0xd1, 0x9d, 0xcc, 0xd0, 0x74, 0xb1, 0x29, 0xd8, 0x19, 0xf8, 0x72, 0xfe, 0x1a, 0x59,
	0xce, 0xd8, 0x64, 0xe9, 0x51, 0x97, 0xc3, 0x9b, 0x28, 0xc7, 0x4d, 0x2a, 0x7d, 0x39, 0xa5, 0xc0,
	0xa5, 0x0d, 0x35, 0xb2, 0x94, 0x71, 0x87, 0x19, 0x55, 0x64, 0x9a, 0x5a, 0x68, 0xda, 0x23, 0xb4,
	0x18, 0xf9, 0x3c, 0x0d, 0x2a, 0x3d, 0x8b, 0x74, 0x83, 0x2c, 0xe4, 0xb8, 0x3c, 0x86, 0xfe, 0xb1,
	0x82, 0xaa, 0x6f, 0x22, 0xe0, 0xb1, 0xf6, 0x3c, 0x8f, 0x73, 0x0b, 0xdd, 0x2e, 0x3a, 0x5b, 0x58,
	0x57, 0x1a, 0xea, 0xc0, 0x40, 0xf9, 0xe1, 0x42, 0xed, 0x5b, 0x66, 0x48, 0xd9, 0x68, 0x4a, 0xf9,
	0x68, 0xf7, 0x11, 0x92, 0x33, 0x40, 0x68, 0x5a, 0x22, 0xad, 0x42, 0x2d, 0x45, 0xf6, 0xc4, 0x70,
	0xb2, 0xea, 0xd8, 0x64, 0x4b, 0x0a, 0x39, 0x59, 0x56, 0x48, 0xfd, 0xab, 0x82, 0x50, 0x3b, 0xf3,
	0x46, 0xfb, 0xd2, 0x57, 0xee, 0x1b, 0x8c, 0x53, 0x5e, 0x89, 0xca, 0x88, 0x4a, 0x14, 0xc5, 0xa6,
	0x16, 0xc6, 0xa6, 0x4d, 0xa7, 0x5e, 0xeb, 0xbf, 0x2a, 0x68, 0x2a, 0xd9, 0x03, 0x9f, 0x8c, 0x3a,
	0xe7, 0xbe, 0xf1, 0xfb, 0x6a, 0xf5, 0x89, 0xcb, 0xf8, 0xf9, 0x8e, 0x4e, 0x7d, 0x07, 0x2e, 0x76,
	0x03, 0x4e, 0xcf, 0x2d, 0x1e, 0xbf, 0x82, 0x78, 0xd7, 0x8a, 0x04, 0x3b, 0xf4, 0x6d, 0x0e, 0xe7,
	0xe0, 0x8b, 0x9d, 0x53, 0xcb, 0x0b, 0x41, 0xff, 0x9f, 0x83, 0x0d, 0xa6, 0xab, 0x0e, 0xa7, 0xbb,
	0x86, 0xe6, 0xfa, 0xef, 0xb8, 0x4c, 0x49, 0x25, 0xb3, 0x7d, 0xd7, 0xbb, 0xd0, 0x9a, 0x6a, 0xa1,
	0x35, 0xc3, 0x5d, 0x99, 0xba, 0x69, 0x57, 0xa6, 0xcb, 0xba, 0xf2, 0xec, 0x73, 0x05, 0x2d, 0x1e,
	0xf4, 0xe6, 0x7c, 0x0b, 0xbc, 0x43, 0x6d, 0xc0, 0x47, 0xf9, 0xa7, 0x11, 0xaf, 0xb7, 0xfe, 0xfd,
	0xce, 0xb6, 0x7a, 0x6c, 0x2b, 0xed, 0x95, 0xa6, 0x8f, 0x51, 0x75, 0xaf, 0xd2, 0xeb, 0xf4, 0x76,
	0xe2, 0xb5, 0x22, 0xb1, 0xa4, 0xb2, 0xfd, 0x56, 0x47, 0x49, 0x02, 0x2f, 0x7e, 0xaa, 0xe0, 0xe3,
	0xfe, 0x92, 0xe3, 0x8d, 0xa2, 0x05, 0x39, 0x9f, 0x6d, 0xbc, 0x3e, 0x56, 0x17, 0x78, 0xf1, 0xfe,
	0xf6, 0xa7, 0xeb, 0x95, 0x89, 0xcb, 0xeb, 0x95, 0x89, 0x1f, 0xd7, 0x2b, 0x13, 0xc7, 0x1b, 0xc9,
	0x12, 0x01, 0xf6, 0xbb, 0xe4, 0x29, 0x32, 0x5c, 0x66, 0x04, 0x67, 0xae, 0x31, 0xf8, 0x4c, 0x9d,
	0x4c, 0xc9, 0x97, 0x65, 0xeb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x0d, 0x1f, 0xf8, 0xbf,
	0x06, 0x00, 0x00,
}
