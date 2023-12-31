// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: berty/errcode.proto

package errcode

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type ErrCode int32

const (
	Undefined                      ErrCode = 0
	TODO                           ErrCode = 666
	ErrNotImplemented              ErrCode = 777
	ErrInternal                    ErrCode = 888
	ErrInvalidInput                ErrCode = 100
	ErrInvalidRange                ErrCode = 101
	ErrMissingInput                ErrCode = 102
	ErrSerialization               ErrCode = 103
	ErrDeserialization             ErrCode = 104
	ErrStreamRead                  ErrCode = 105
	ErrStreamWrite                 ErrCode = 106
	ErrStreamTransform             ErrCode = 110
	ErrStreamSendAndClose          ErrCode = 111
	ErrStreamHeaderWrite           ErrCode = 112
	ErrStreamHeaderRead            ErrCode = 115
	ErrStreamSink                  ErrCode = 113
	ErrStreamCloseAndRecv          ErrCode = 114
	ErrMissingMapKey               ErrCode = 107
	ErrDBWrite                     ErrCode = 108
	ErrDBRead                      ErrCode = 109
	ErrDBDestroy                   ErrCode = 120
	ErrDBMigrate                   ErrCode = 121
	ErrDBReplay                    ErrCode = 122
	ErrDBRestore                   ErrCode = 123
	ErrDBOpen                      ErrCode = 124
	ErrDBClose                     ErrCode = 125
	ErrCryptoRandomGeneration      ErrCode = 200
	ErrCryptoKeyGeneration         ErrCode = 201
	ErrCryptoNonceGeneration       ErrCode = 202
	ErrCryptoSignature             ErrCode = 203
	ErrCryptoSignatureVerification ErrCode = 204
	ErrCryptoDecrypt               ErrCode = 205
	ErrCryptoDecryptPayload        ErrCode = 206
	ErrCryptoEncrypt               ErrCode = 207
	ErrCryptoKeyConversion         ErrCode = 208
	ErrCryptoCipherInit            ErrCode = 209
	ErrCryptoKeyDerivation         ErrCode = 210
	ErrMap                         ErrCode = 300
	ErrForEach                     ErrCode = 301
	ErrKeystoreGet                 ErrCode = 400
	ErrKeystorePut                 ErrCode = 401
	ErrNotFound                    ErrCode = 404
	ErrIPFSAdd                     ErrCode = 1050
	ErrIPFSGet                     ErrCode = 1051
	ErrIPFSInit                    ErrCode = 1052
	ErrIPFSSetupConfig             ErrCode = 1053
	ErrIPFSSetupRepo               ErrCode = 1054
	ErrIPFSSetupHost               ErrCode = 1055
	// Event errors
	ErrEventListMetadata                   ErrCode = 1400
	ErrEventListMessage                    ErrCode = 1401
	ErrBridgeInterrupted                   ErrCode = 1600
	ErrBridgeNotRunning                    ErrCode = 1601
	ErrMessengerInvalidDeepLink            ErrCode = 2000
	ErrMessengerDeepLinkRequiresPassphrase ErrCode = 2001
	ErrMessengerDeepLinkInvalidPassphrase  ErrCode = 2002
	ErrMessengerStreamEvent                ErrCode = 2003
	ErrMessengerContactMetadataUnmarshal   ErrCode = 2004
	ErrDBEntryAlreadyExists                ErrCode = 2100
	ErrDBAddConversation                   ErrCode = 2101
	ErrDBAddContactRequestOutgoingSent     ErrCode = 2102
	ErrDBAddContactRequestOutgoingEnqueud  ErrCode = 2103
	ErrDBAddContactRequestIncomingReceived ErrCode = 2104
	ErrDBAddContactRequestIncomingAccepted ErrCode = 2105
	ErrDBAddGroupMemberDeviceAdded         ErrCode = 2106
	ErrDBMultipleRecords                   ErrCode = 2107
	ErrReplayProcessGroupMetadata          ErrCode = 2200
	ErrReplayProcessGroupMessage           ErrCode = 2201
	ErrAttachmentPrepare                   ErrCode = 2300
	ErrAttachmentRetrieve                  ErrCode = 2301
	ErrProtocolSend                        ErrCode = 2302
	ErrProtocolEventUnmarshal              ErrCode = 2303
	ErrProtocolGetGroupInfo                ErrCode = 2304
	// Test Error
	ErrTestEcho                                          ErrCode = 2401
	ErrTestEchoRecv                                      ErrCode = 2402
	ErrTestEchoSend                                      ErrCode = 2403
	ErrCLINoTermcaps                                     ErrCode = 3001
	ErrServicesDirectory                                 ErrCode = 4200
	ErrServicesDirectoryInvalidVerifiedCredentialSubject ErrCode = 4201
	ErrServicesDirectoryExistingRecordNotFound           ErrCode = 4202
	ErrServicesDirectoryRecordLockedAndCantBeReplaced    ErrCode = 4203
	ErrServicesDirectoryExplicitReplaceFlagRequired      ErrCode = 4204
	ErrServicesDirectoryInvalidVerifiedCredential        ErrCode = 4205
	ErrServicesDirectoryExpiredVerifiedCredential        ErrCode = 4206
	ErrServicesDirectoryInvalidVerifiedCredentialID      ErrCode = 4207
	ErrBertyAccount                                      ErrCode = 5000
	ErrBertyAccountNoIDSpecified                         ErrCode = 5001
	ErrBertyAccountAlreadyOpened                         ErrCode = 5002
	ErrBertyAccountInvalidIDFormat                       ErrCode = 5003
	ErrBertyAccountLoggerDecorator                       ErrCode = 5004
	ErrBertyAccountGRPCClient                            ErrCode = 5005
	ErrBertyAccountOpenAccount                           ErrCode = 5006
	ErrBertyAccountDataNotFound                          ErrCode = 5007
	ErrBertyAccountMetadataUpdate                        ErrCode = 5008
	ErrBertyAccountManagerOpen                           ErrCode = 5009
	ErrBertyAccountManagerClose                          ErrCode = 5010
	ErrBertyAccountInvalidCLIArgs                        ErrCode = 5011
	ErrBertyAccountFSError                               ErrCode = 5012
	ErrBertyAccountAlreadyExists                         ErrCode = 5013
	ErrBertyAccountNoBackupSpecified                     ErrCode = 5014
	ErrBertyAccountIDGenFailed                           ErrCode = 5015
	ErrBertyAccountCreationFailed                        ErrCode = 5016
	ErrBertyAccountUpdateFailed                          ErrCode = 5017
	ErrAppStorageNotSupported                            ErrCode = 5018
)

var ErrCode_name = map[int32]string{
	0:    "Undefined",
	666:  "TODO",
	777:  "ErrNotImplemented",
	888:  "ErrInternal",
	100:  "ErrInvalidInput",
	101:  "ErrInvalidRange",
	102:  "ErrMissingInput",
	103:  "ErrSerialization",
	104:  "ErrDeserialization",
	105:  "ErrStreamRead",
	106:  "ErrStreamWrite",
	110:  "ErrStreamTransform",
	111:  "ErrStreamSendAndClose",
	112:  "ErrStreamHeaderWrite",
	115:  "ErrStreamHeaderRead",
	113:  "ErrStreamSink",
	114:  "ErrStreamCloseAndRecv",
	107:  "ErrMissingMapKey",
	108:  "ErrDBWrite",
	109:  "ErrDBRead",
	120:  "ErrDBDestroy",
	121:  "ErrDBMigrate",
	122:  "ErrDBReplay",
	123:  "ErrDBRestore",
	124:  "ErrDBOpen",
	125:  "ErrDBClose",
	200:  "ErrCryptoRandomGeneration",
	201:  "ErrCryptoKeyGeneration",
	202:  "ErrCryptoNonceGeneration",
	203:  "ErrCryptoSignature",
	204:  "ErrCryptoSignatureVerification",
	205:  "ErrCryptoDecrypt",
	206:  "ErrCryptoDecryptPayload",
	207:  "ErrCryptoEncrypt",
	208:  "ErrCryptoKeyConversion",
	209:  "ErrCryptoCipherInit",
	210:  "ErrCryptoKeyDerivation",
	300:  "ErrMap",
	301:  "ErrForEach",
	400:  "ErrKeystoreGet",
	401:  "ErrKeystorePut",
	404:  "ErrNotFound",
	1050: "ErrIPFSAdd",
	1051: "ErrIPFSGet",
	1052: "ErrIPFSInit",
	1053: "ErrIPFSSetupConfig",
	1054: "ErrIPFSSetupRepo",
	1055: "ErrIPFSSetupHost",
	1400: "ErrEventListMetadata",
	1401: "ErrEventListMessage",
	1600: "ErrBridgeInterrupted",
	1601: "ErrBridgeNotRunning",
	2000: "ErrMessengerInvalidDeepLink",
	2001: "ErrMessengerDeepLinkRequiresPassphrase",
	2002: "ErrMessengerDeepLinkInvalidPassphrase",
	2003: "ErrMessengerStreamEvent",
	2004: "ErrMessengerContactMetadataUnmarshal",
	2100: "ErrDBEntryAlreadyExists",
	2101: "ErrDBAddConversation",
	2102: "ErrDBAddContactRequestOutgoingSent",
	2103: "ErrDBAddContactRequestOutgoingEnqueud",
	2104: "ErrDBAddContactRequestIncomingReceived",
	2105: "ErrDBAddContactRequestIncomingAccepted",
	2106: "ErrDBAddGroupMemberDeviceAdded",
	2107: "ErrDBMultipleRecords",
	2200: "ErrReplayProcessGroupMetadata",
	2201: "ErrReplayProcessGroupMessage",
	2300: "ErrAttachmentPrepare",
	2301: "ErrAttachmentRetrieve",
	2302: "ErrProtocolSend",
	2303: "ErrProtocolEventUnmarshal",
	2304: "ErrProtocolGetGroupInfo",
	2401: "ErrTestEcho",
	2402: "ErrTestEchoRecv",
	2403: "ErrTestEchoSend",
	3001: "ErrCLINoTermcaps",
	4200: "ErrServicesDirectory",
	4201: "ErrServicesDirectoryInvalidVerifiedCredentialSubject",
	4202: "ErrServicesDirectoryExistingRecordNotFound",
	4203: "ErrServicesDirectoryRecordLockedAndCantBeReplaced",
	4204: "ErrServicesDirectoryExplicitReplaceFlagRequired",
	4205: "ErrServicesDirectoryInvalidVerifiedCredential",
	4206: "ErrServicesDirectoryExpiredVerifiedCredential",
	4207: "ErrServicesDirectoryInvalidVerifiedCredentialID",
	5000: "ErrBertyAccount",
	5001: "ErrBertyAccountNoIDSpecified",
	5002: "ErrBertyAccountAlreadyOpened",
	5003: "ErrBertyAccountInvalidIDFormat",
	5004: "ErrBertyAccountLoggerDecorator",
	5005: "ErrBertyAccountGRPCClient",
	5006: "ErrBertyAccountOpenAccount",
	5007: "ErrBertyAccountDataNotFound",
	5008: "ErrBertyAccountMetadataUpdate",
	5009: "ErrBertyAccountManagerOpen",
	5010: "ErrBertyAccountManagerClose",
	5011: "ErrBertyAccountInvalidCLIArgs",
	5012: "ErrBertyAccountFSError",
	5013: "ErrBertyAccountAlreadyExists",
	5014: "ErrBertyAccountNoBackupSpecified",
	5015: "ErrBertyAccountIDGenFailed",
	5016: "ErrBertyAccountCreationFailed",
	5017: "ErrBertyAccountUpdateFailed",
	5018: "ErrAppStorageNotSupported",
}

var ErrCode_value = map[string]int32{
	"Undefined":                              0,
	"TODO":                                   666,
	"ErrNotImplemented":                      777,
	"ErrInternal":                            888,
	"ErrInvalidInput":                        100,
	"ErrInvalidRange":                        101,
	"ErrMissingInput":                        102,
	"ErrSerialization":                       103,
	"ErrDeserialization":                     104,
	"ErrStreamRead":                          105,
	"ErrStreamWrite":                         106,
	"ErrStreamTransform":                     110,
	"ErrStreamSendAndClose":                  111,
	"ErrStreamHeaderWrite":                   112,
	"ErrStreamHeaderRead":                    115,
	"ErrStreamSink":                          113,
	"ErrStreamCloseAndRecv":                  114,
	"ErrMissingMapKey":                       107,
	"ErrDBWrite":                             108,
	"ErrDBRead":                              109,
	"ErrDBDestroy":                           120,
	"ErrDBMigrate":                           121,
	"ErrDBReplay":                            122,
	"ErrDBRestore":                           123,
	"ErrDBOpen":                              124,
	"ErrDBClose":                             125,
	"ErrCryptoRandomGeneration":              200,
	"ErrCryptoKeyGeneration":                 201,
	"ErrCryptoNonceGeneration":               202,
	"ErrCryptoSignature":                     203,
	"ErrCryptoSignatureVerification":         204,
	"ErrCryptoDecrypt":                       205,
	"ErrCryptoDecryptPayload":                206,
	"ErrCryptoEncrypt":                       207,
	"ErrCryptoKeyConversion":                 208,
	"ErrCryptoCipherInit":                    209,
	"ErrCryptoKeyDerivation":                 210,
	"ErrMap":                                 300,
	"ErrForEach":                             301,
	"ErrKeystoreGet":                         400,
	"ErrKeystorePut":                         401,
	"ErrNotFound":                            404,
	"ErrIPFSAdd":                             1050,
	"ErrIPFSGet":                             1051,
	"ErrIPFSInit":                            1052,
	"ErrIPFSSetupConfig":                     1053,
	"ErrIPFSSetupRepo":                       1054,
	"ErrIPFSSetupHost":                       1055,
	"ErrEventListMetadata":                   1400,
	"ErrEventListMessage":                    1401,
	"ErrBridgeInterrupted":                   1600,
	"ErrBridgeNotRunning":                    1601,
	"ErrMessengerInvalidDeepLink":            2000,
	"ErrMessengerDeepLinkRequiresPassphrase": 2001,
	"ErrMessengerDeepLinkInvalidPassphrase":  2002,
	"ErrMessengerStreamEvent":                2003,
	"ErrMessengerContactMetadataUnmarshal":   2004,
	"ErrDBEntryAlreadyExists":                2100,
	"ErrDBAddConversation":                   2101,
	"ErrDBAddContactRequestOutgoingSent":     2102,
	"ErrDBAddContactRequestOutgoingEnqueud":  2103,
	"ErrDBAddContactRequestIncomingReceived": 2104,
	"ErrDBAddContactRequestIncomingAccepted": 2105,
	"ErrDBAddGroupMemberDeviceAdded":         2106,
	"ErrDBMultipleRecords":                   2107,
	"ErrReplayProcessGroupMetadata":          2200,
	"ErrReplayProcessGroupMessage":           2201,
	"ErrAttachmentPrepare":                   2300,
	"ErrAttachmentRetrieve":                  2301,
	"ErrProtocolSend":                        2302,
	"ErrProtocolEventUnmarshal":              2303,
	"ErrProtocolGetGroupInfo":                2304,
	"ErrTestEcho":                            2401,
	"ErrTestEchoRecv":                        2402,
	"ErrTestEchoSend":                        2403,
	"ErrCLINoTermcaps":                       3001,
	"ErrServicesDirectory":                   4200,
	"ErrServicesDirectoryInvalidVerifiedCredentialSubject": 4201,
	"ErrServicesDirectoryExistingRecordNotFound":           4202,
	"ErrServicesDirectoryRecordLockedAndCantBeReplaced":    4203,
	"ErrServicesDirectoryExplicitReplaceFlagRequired":      4204,
	"ErrServicesDirectoryInvalidVerifiedCredential":        4205,
	"ErrServicesDirectoryExpiredVerifiedCredential":        4206,
	"ErrServicesDirectoryInvalidVerifiedCredentialID":      4207,
	"ErrBertyAccount":                  5000,
	"ErrBertyAccountNoIDSpecified":     5001,
	"ErrBertyAccountAlreadyOpened":     5002,
	"ErrBertyAccountInvalidIDFormat":   5003,
	"ErrBertyAccountLoggerDecorator":   5004,
	"ErrBertyAccountGRPCClient":        5005,
	"ErrBertyAccountOpenAccount":       5006,
	"ErrBertyAccountDataNotFound":      5007,
	"ErrBertyAccountMetadataUpdate":    5008,
	"ErrBertyAccountManagerOpen":       5009,
	"ErrBertyAccountManagerClose":      5010,
	"ErrBertyAccountInvalidCLIArgs":    5011,
	"ErrBertyAccountFSError":           5012,
	"ErrBertyAccountAlreadyExists":     5013,
	"ErrBertyAccountNoBackupSpecified": 5014,
	"ErrBertyAccountIDGenFailed":       5015,
	"ErrBertyAccountCreationFailed":    5016,
	"ErrBertyAccountUpdateFailed":      5017,
	"ErrAppStorageNotSupported":        5018,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36ad8b3d0b181e72, []int{0}
}

type ErrDetails struct {
	Codes                []ErrCode `protobuf:"varint,1,rep,packed,name=codes,proto3,enum=berty.errcode.ErrCode" json:"codes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ErrDetails) Reset()         { *m = ErrDetails{} }
func (m *ErrDetails) String() string { return proto.CompactTextString(m) }
func (*ErrDetails) ProtoMessage()    {}
func (*ErrDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_36ad8b3d0b181e72, []int{0}
}
func (m *ErrDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ErrDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ErrDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ErrDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrDetails.Merge(m, src)
}
func (m *ErrDetails) XXX_Size() int {
	return m.Size()
}
func (m *ErrDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrDetails.DiscardUnknown(m)
}

var xxx_messageInfo_ErrDetails proto.InternalMessageInfo

func (m *ErrDetails) GetCodes() []ErrCode {
	if m != nil {
		return m.Codes
	}
	return nil
}

func init() {
	proto.RegisterEnum("berty.errcode.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterType((*ErrDetails)(nil), "berty.errcode.ErrDetails")
}

func init() { proto.RegisterFile("berty/errcode.proto", fileDescriptor_36ad8b3d0b181e72) }

var fileDescriptor_36ad8b3d0b181e72 = []byte{
	// 1510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0x4b, 0x73, 0x63, 0x47,
	0x15, 0x8e, 0x32, 0x64, 0x70, 0x3a, 0x33, 0xe3, 0xa3, 0x9e, 0xc9, 0x78, 0x66, 0x92, 0x38, 0x62,
	0x20, 0x08, 0x06, 0xb0, 0x2b, 0x21, 0x45, 0x15, 0xec, 0xf4, 0x74, 0x54, 0xe3, 0x87, 0x4a, 0x72,
	0xa0, 0x8a, 0x5d, 0xfb, 0xf6, 0xf1, 0x55, 0xc7, 0x57, 0xdd, 0x77, 0x4e, 0xf7, 0x55, 0x8d, 0x02,
	0x54, 0xb1, 0x24, 0x3c, 0x13, 0x08, 0x90, 0x0c, 0xaf, 0x3f, 0x00, 0x0b, 0xaa, 0x78, 0x05, 0x36,
	0xb0, 0x0b, 0x6f, 0x27, 0xf0, 0x03, 0x88, 0xd9, 0xf0, 0x66, 0x1b, 0xaa, 0x78, 0xd5, 0xed, 0xdb,
	0x92, 0x65, 0x5b, 0x4e, 0xc1, 0xca, 0xd6, 0xf9, 0xbe, 0x73, 0x4e, 0x9f, 0x67, 0xf7, 0x65, 0x17,
	0x77, 0x90, 0xdc, 0x78, 0x15, 0x89, 0x22, 0x23, 0x71, 0x25, 0x25, 0xe3, 0x0c, 0x3f, 0xef, 0x85,
	0x2b, 0x41, 0x78, 0xed, 0x52, 0x6c, 0x62, 0xe3, 0x91, 0xd5, 0xfc, 0xbf, 0x82, 0x74, 0xfd, 0x03,
	0x8c, 0xb5, 0x88, 0x9a, 0xe8, 0x84, 0x4a, 0x2c, 0x7f, 0x37, 0xbb, 0x27, 0xe7, 0xda, 0x2b, 0xa5,
	0xca, 0x99, 0x77, 0x5c, 0x78, 0xec, 0xf2, 0xca, 0x11, 0x13, 0x2b, 0x2d, 0xa2, 0x86, 0x91, 0xd8,
	0x2b, 0x48, 0x37, 0xbe, 0xbd, 0xc4, 0xde, 0x1c, 0x44, 0xfc, 0x3c, 0xbb, 0xf7, 0x49, 0x2d, 0x71,
	0x57, 0x69, 0x94, 0x70, 0x17, 0xbf, 0x97, 0xbd, 0x69, 0x7b, 0xab, 0xb9, 0x05, 0x77, 0xee, 0xe1,
	0x97, 0x59, 0xb9, 0x45, 0xb4, 0x69, 0x5c, 0x67, 0x98, 0x26, 0x38, 0x44, 0xed, 0x50, 0xc2, 0x33,
	0x67, 0x39, 0xb0, 0xfb, 0x5a, 0x44, 0x1d, 0xed, 0x90, 0xb4, 0x48, 0xe0, 0xf5, 0xb3, 0xfc, 0x22,
	0x5b, 0xf4, 0x92, 0x91, 0x48, 0x94, 0xec, 0xe8, 0x34, 0x73, 0x20, 0x8f, 0x0a, 0x7b, 0x42, 0xc7,
	0x08, 0x18, 0x84, 0x1b, 0xca, 0x5a, 0xa5, 0xe3, 0x82, 0xb9, 0xcb, 0x2f, 0x31, 0x68, 0x11, 0xf5,
	0x91, 0x94, 0x48, 0xd4, 0xd3, 0xc2, 0x29, 0xa3, 0x21, 0xe6, 0x97, 0x19, 0xf7, 0x01, 0xda, 0x23,
	0xf2, 0x01, 0x2f, 0xb3, 0xf3, 0x39, 0xdb, 0x11, 0x8a, 0x61, 0x0f, 0x85, 0x04, 0xc5, 0x39, 0xbb,
	0x30, 0x15, 0x7d, 0x88, 0x94, 0x43, 0x78, 0x2a, 0xa8, 0x17, 0xb2, 0x6d, 0x12, 0xda, 0xee, 0x1a,
	0x1a, 0x82, 0xe6, 0x57, 0xd9, 0xfd, 0x53, 0x79, 0x1f, 0xb5, 0xac, 0x69, 0xd9, 0x48, 0x8c, 0x45,
	0x30, 0xfc, 0x0a, 0xbb, 0x34, 0x85, 0x9e, 0x40, 0x21, 0x91, 0x0a, 0x63, 0x29, 0x5f, 0x62, 0x17,
	0x8f, 0x21, 0xde, 0xb3, 0x3d, 0x72, 0x98, 0xbe, 0xd2, 0x7b, 0x70, 0xeb, 0x88, 0x03, 0x6f, 0xb9,
	0xa6, 0x65, 0x0f, 0xa3, 0x11, 0x50, 0x08, 0x34, 0x44, 0xbf, 0x21, 0xd2, 0x9b, 0x38, 0x86, 0x3d,
	0x7e, 0xa1, 0xa8, 0x64, 0xbd, 0x70, 0x96, 0xe4, 0x15, 0xf1, 0xbf, 0xbd, 0x8b, 0x21, 0x07, 0x76,
	0xce, 0xff, 0x6c, 0xa2, 0x75, 0x64, 0xc6, 0x70, 0x7b, 0x2a, 0xd9, 0x50, 0x31, 0x09, 0x87, 0x30,
	0xe6, 0x8b, 0xbe, 0x24, 0xb9, 0x4a, 0x9a, 0x88, 0x31, 0x3c, 0x3d, 0xa5, 0xf4, 0xd0, 0x3a, 0x43,
	0x08, 0x1f, 0x99, 0x5a, 0xdd, 0x4a, 0x51, 0xc3, 0x47, 0xa7, 0x4e, 0x8b, 0xd8, 0x3f, 0xc6, 0x97,
	0xd9, 0xd5, 0xbc, 0x23, 0x68, 0x9c, 0x3a, 0xd3, 0x13, 0x5a, 0x9a, 0xe1, 0x1a, 0x6a, 0xa4, 0x22,
	0xe9, 0x2f, 0x97, 0xf8, 0x03, 0xec, 0xf2, 0x14, 0xbf, 0x89, 0xe3, 0x19, 0xf0, 0xa7, 0x25, 0xfe,
	0x10, 0xbb, 0x32, 0x05, 0x37, 0x8d, 0x8e, 0x70, 0x06, 0xfe, 0x59, 0x89, 0x2f, 0xf9, 0x52, 0x14,
	0x70, 0x5f, 0xc5, 0x5a, 0xb8, 0x8c, 0x10, 0x7e, 0x5e, 0xe2, 0x6f, 0x65, 0xcb, 0x27, 0x81, 0x0f,
	0x22, 0xa9, 0x5d, 0x15, 0x15, 0xda, 0xbf, 0x28, 0xf1, 0xfb, 0x7d, 0xd2, 0x0a, 0x52, 0x13, 0xa3,
	0xfc, 0x2f, 0xfc, 0xb2, 0xc4, 0x1f, 0x64, 0x4b, 0xc7, 0xc5, 0x5d, 0x31, 0x4e, 0x8c, 0x90, 0xf0,
	0xab, 0xa3, 0x4a, 0x2d, 0x5d, 0x28, 0xfd, 0xfa, 0x44, 0x14, 0x0d, 0xa3, 0x47, 0x48, 0x36, 0x77,
	0xb4, 0x5f, 0xe2, 0x57, 0x7c, 0x91, 0x0b, 0xb0, 0xa1, 0xd2, 0x01, 0x52, 0x47, 0x2b, 0x07, 0xaf,
	0x9c, 0x50, 0x6b, 0x22, 0xa9, 0x51, 0x71, 0xbe, 0x57, 0x4b, 0xfc, 0x3e, 0x76, 0x36, 0x2f, 0xaa,
	0x48, 0xe1, 0x9b, 0x77, 0xf3, 0x45, 0x9f, 0xd6, 0xb6, 0xa1, 0x96, 0x88, 0x06, 0xf0, 0xad, 0xbb,
	0xf9, 0x45, 0xdf, 0x9a, 0x37, 0x71, 0xec, 0xeb, 0xb0, 0x86, 0x0e, 0x9e, 0x3d, 0x73, 0x4c, 0xd8,
	0xcd, 0x1c, 0x3c, 0x77, 0x26, 0x8c, 0xd5, 0xa6, 0x71, 0x6d, 0x93, 0x69, 0x09, 0xcf, 0x9f, 0x09,
	0xc6, 0x3a, 0xdd, 0x76, 0xbf, 0x26, 0x25, 0xdc, 0x59, 0x98, 0x11, 0xe4, 0x86, 0xbe, 0xb2, 0x30,
	0x19, 0xc5, 0x6e, 0xbb, 0xef, 0x8f, 0xfa, 0xd5, 0x85, 0x90, 0xeb, 0x5c, 0xd2, 0x47, 0x97, 0xa5,
	0x0d, 0xa3, 0x77, 0x55, 0x0c, 0x5f, 0x5b, 0x08, 0x19, 0x99, 0x02, 0x3d, 0x4c, 0x0d, 0x7c, 0xfd,
	0x84, 0xf8, 0x09, 0x63, 0x1d, 0x7c, 0x63, 0x81, 0x5f, 0xf5, 0xa3, 0xd0, 0x1a, 0xa1, 0x76, 0xeb,
	0xca, 0xba, 0x0d, 0x74, 0x42, 0x0a, 0x27, 0xe0, 0x75, 0x16, 0xd2, 0x34, 0x03, 0x59, 0x2b, 0x62,
	0x84, 0x7f, 0xb0, 0xa0, 0x54, 0x27, 0x25, 0x63, 0xf4, 0xeb, 0x81, 0xb2, 0x34, 0xdf, 0x19, 0x3f,
	0x3e, 0x17, 0x94, 0x0a, 0x68, 0xd3, 0xb8, 0x5e, 0xa6, 0xb5, 0xd2, 0x31, 0xfc, 0xe4, 0x1c, 0xaf,
	0xb0, 0x07, 0xf2, 0xf4, 0xa1, 0xb5, 0xa8, 0x63, 0x9c, 0xec, 0x8b, 0x26, 0x62, 0xba, 0x9e, 0xcf,
	0xd3, 0xfe, 0x22, 0x7f, 0x17, 0x7b, 0xfb, 0x2c, 0x63, 0x02, 0xf5, 0xf0, 0x56, 0xa6, 0x08, 0x6d,
	0x57, 0x58, 0x9b, 0x0e, 0x48, 0x58, 0x84, 0x57, 0x16, 0xf9, 0x0d, 0xf6, 0xc8, 0x3c, 0x72, 0x30,
	0x3b, 0xc3, 0x7d, 0x75, 0x31, 0xb4, 0xd0, 0x94, 0x5b, 0x8c, 0xac, 0x0f, 0x0c, 0x7e, 0xb3, 0xc8,
	0xdf, 0xc9, 0xde, 0x36, 0x8b, 0x36, 0x8c, 0x76, 0x22, 0x9a, 0x66, 0xe2, 0x49, 0x3d, 0x14, 0x64,
	0x07, 0x22, 0x81, 0xdf, 0x4e, 0x0c, 0x35, 0xeb, 0x2d, 0xed, 0x68, 0x5c, 0x4b, 0x08, 0x85, 0x1c,
	0xb7, 0x6e, 0x2b, 0xeb, 0x2c, 0x7c, 0x07, 0x42, 0x5a, 0x9a, 0xf5, 0x9a, 0x94, 0xa1, 0xe1, 0x8a,
	0xde, 0xf9, 0x2e, 0xf0, 0x2a, 0xbb, 0x3e, 0x03, 0xe5, 0xf6, 0xf3, 0xa8, 0xd0, 0xba, 0xad, 0xcc,
	0xc5, 0x46, 0xe9, 0xb8, 0x9f, 0x1f, 0xe6, 0x7b, 0x10, 0xc2, 0x7a, 0x03, 0x62, 0x4b, 0xdf, 0xca,
	0x30, 0x93, 0xf0, 0x7d, 0x08, 0xf9, 0x9a, 0xc3, 0xed, 0xe8, 0xc8, 0x0c, 0x95, 0x8e, 0x7b, 0x18,
	0xa1, 0x1a, 0xa1, 0x84, 0x1f, 0xfc, 0x0f, 0xe4, 0x5a, 0x14, 0xa1, 0xaf, 0xe2, 0x4b, 0x10, 0xe6,
	0xd5, 0x93, 0xd7, 0xc8, 0x64, 0xe9, 0x06, 0x0e, 0x77, 0xf2, 0x1c, 0x8f, 0x54, 0x84, 0x35, 0x29,
	0x51, 0xc2, 0x0f, 0x0f, 0xc3, 0xdd, 0xc8, 0x12, 0xa7, 0xd2, 0x04, 0x7b, 0x18, 0x19, 0x92, 0x16,
	0x7e, 0x04, 0xfc, 0x3a, 0x7b, 0xa8, 0x45, 0x54, 0x2c, 0xa9, 0x2e, 0x99, 0x08, 0xad, 0x0d, 0x76,
	0x42, 0x7b, 0xbd, 0x50, 0xe6, 0x6f, 0x61, 0x0f, 0x9e, 0xc2, 0x29, 0xfa, 0xec, 0xc5, 0x72, 0xf0,
	0x50, 0x73, 0x4e, 0x44, 0x83, 0xfc, 0x5a, 0xea, 0x12, 0xa6, 0x82, 0x10, 0xfe, 0x59, 0xe6, 0xd7,
	0xfc, 0xf2, 0x3d, 0x84, 0x7a, 0xe8, 0x48, 0xe1, 0x08, 0xe1, 0x5f, 0x65, 0x7e, 0xc9, 0xdf, 0x3d,
	0xdd, 0xfc, 0xf6, 0x8c, 0x4c, 0x92, 0xef, 0x7e, 0xf8, 0x77, 0x39, 0x2c, 0xbe, 0x89, 0xd4, 0x57,
	0xff, 0xb0, 0xb6, 0xff, 0x29, 0x87, 0xda, 0x4e, 0xf0, 0x35, 0x74, 0xfe, 0x34, 0x1d, 0xbd, 0x6b,
	0xe0, 0xe3, 0x3c, 0x0c, 0xe0, 0x36, 0x5a, 0xd7, 0x8a, 0x06, 0x06, 0x5e, 0xe3, 0xc1, 0xcb, 0x44,
	0xe2, 0x17, 0xff, 0xc1, 0x71, 0xa9, 0xf7, 0xfd, 0x7b, 0x3e, 0xd9, 0x52, 0xeb, 0x9d, 0x4d, 0xb3,
	0x8d, 0x34, 0x8c, 0x44, 0x6a, 0xe1, 0xa5, 0xa5, 0x10, 0x5f, 0x1f, 0x29, 0xcf, 0xab, 0x6d, 0x2a,
	0xc2, 0xc8, 0x19, 0x1a, 0xc3, 0x1f, 0x2a, 0xfc, 0xfd, 0xec, 0xf1, 0x79, 0x50, 0x68, 0xef, 0x62,
	0x75, 0xa2, 0x6c, 0x10, 0x4a, 0xd4, 0x4e, 0x89, 0xa4, 0x9f, 0xed, 0x3c, 0x85, 0x91, 0x83, 0x3f,
	0x56, 0xf8, 0x2a, 0xbb, 0x31, 0x4f, 0xd5, 0xf7, 0x69, 0xd1, 0x14, 0x86, 0xe4, 0x74, 0xfd, 0xfc,
	0xa9, 0xc2, 0xdf, 0xc7, 0x1e, 0x9d, 0xa7, 0x50, 0x10, 0xd7, 0x4d, 0xb4, 0x87, 0xfe, 0xf2, 0x14,
	0xda, 0xd5, 0xd1, 0xd7, 0x2b, 0x42, 0x09, 0x7f, 0xae, 0xf0, 0xc7, 0xd9, 0xea, 0x7c, 0x47, 0x69,
	0xa2, 0x22, 0xe5, 0x02, 0xb5, 0x9d, 0x88, 0x38, 0x8c, 0xb0, 0x84, 0xbf, 0x54, 0xf8, 0x63, 0xec,
	0x3d, 0xff, 0x57, 0x64, 0xf0, 0xd7, 0x53, 0x75, 0x5a, 0xb7, 0xd3, 0xdc, 0xea, 0x1c, 0x9d, 0xbf,
	0x9d, 0x7a, 0xba, 0x53, 0xfd, 0x74, 0x9a, 0xf0, 0xf7, 0x4a, 0xa8, 0x5f, 0x3d, 0x7f, 0x54, 0xd5,
	0xa2, 0xc8, 0x64, 0xda, 0xc1, 0x27, 0xaa, 0xa1, 0x57, 0x67, 0xa5, 0x9b, 0xa6, 0xd3, 0xec, 0xa7,
	0x18, 0x79, 0x23, 0xf0, 0xcc, 0x3c, 0x4a, 0xd8, 0x0f, 0xf9, 0x3d, 0x8c, 0x12, 0x3e, 0x59, 0x0d,
	0x53, 0x35, 0x4b, 0x99, 0xbc, 0xa4, 0x9a, 0x6d, 0x43, 0x43, 0xe1, 0xe0, 0x53, 0xf3, 0x48, 0xeb,
	0x26, 0xf6, 0x0b, 0x2e, 0x32, 0x24, 0x9c, 0x21, 0xf8, 0x74, 0x35, 0xf4, 0xf2, 0x2c, 0x69, 0xad,
	0xd7, 0x6d, 0x34, 0x12, 0x95, 0x6f, 0x91, 0xcf, 0x54, 0xf9, 0xc3, 0xec, 0xda, 0x31, 0x3c, 0x3f,
	0xc5, 0x24, 0xa0, 0xcf, 0x56, 0xc3, 0x32, 0x9e, 0x25, 0x34, 0x85, 0x13, 0xd3, 0xa6, 0xf8, 0x5c,
	0x35, 0x8c, 0xf0, 0x2c, 0x63, 0xba, 0x11, 0x53, 0x99, 0x3f, 0x46, 0x9e, 0x9d, 0xe7, 0x66, 0x43,
	0x68, 0x11, 0x23, 0xf9, 0xb7, 0xc7, 0x73, 0xf3, 0xdc, 0x04, 0x42, 0xf1, 0x1a, 0xf9, 0xfc, 0x3c,
	0x37, 0x21, 0x27, 0x8d, 0xf5, 0x4e, 0x8d, 0x62, 0x0b, 0x5f, 0xa8, 0x86, 0x5b, 0x79, 0x96, 0xd3,
	0xee, 0xb7, 0x88, 0x0c, 0xc1, 0xf3, 0x6f, 0x90, 0xf7, 0xb0, 0x97, 0xbf, 0x58, 0xe5, 0x8f, 0xb0,
	0xca, 0x89, 0xea, 0xd5, 0x45, 0xb4, 0x97, 0xa5, 0x87, 0x15, 0xfc, 0xd2, 0xbc, 0x68, 0x3a, 0xcd,
	0x35, 0xd4, 0x6d, 0xa1, 0x12, 0x94, 0xf0, 0xe5, 0x79, 0x67, 0x6d, 0x10, 0xfa, 0x15, 0x1f, 0x38,
	0x2f, 0xcc, 0x8b, 0xb8, 0x48, 0x57, 0x60, 0xbc, 0x38, 0xa9, 0x5d, 0x2d, 0x4d, 0xfb, 0xce, 0x90,
	0xf0, 0xb7, 0x64, 0x3f, 0x4b, 0x53, 0x43, 0xf9, 0xee, 0xbd, 0x53, 0xad, 0x3f, 0xba, 0xff, 0xda,
	0xf2, 0x5d, 0x2f, 0x1f, 0x2c, 0x97, 0xf6, 0x0f, 0x96, 0x4b, 0xbf, 0x3b, 0x58, 0x2e, 0x7d, 0xf8,
	0xe1, 0xe2, 0x8d, 0xef, 0x30, 0x1a, 0xac, 0x16, 0x9f, 0x11, 0xb1, 0x59, 0x4d, 0xf7, 0xe2, 0xc9,
	0xd7, 0xc4, 0xce, 0x59, 0xff, 0xa5, 0xf0, 0xde, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xa3,
	0xb0, 0x74, 0x65, 0x0c, 0x00, 0x00,
}

func (m *ErrDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ErrDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ErrDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Codes) > 0 {
		dAtA2 := make([]byte, len(m.Codes)*10)
		var j1 int
		for _, num := range m.Codes {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintErrcode(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintErrcode(dAtA []byte, offset int, v uint64) int {
	offset -= sovErrcode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ErrDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Codes) > 0 {
		l = 0
		for _, e := range m.Codes {
			l += sovErrcode(uint64(e))
		}
		n += 1 + sovErrcode(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovErrcode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErrcode(x uint64) (n int) {
	return sovErrcode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ErrDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrcode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ErrDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ErrDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v ErrCode
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowErrcode
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ErrCode(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Codes = append(m.Codes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowErrcode
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthErrcode
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthErrcode
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Codes) == 0 {
					m.Codes = make([]ErrCode, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ErrCode
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowErrcode
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ErrCode(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Codes = append(m.Codes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Codes", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipErrcode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErrcode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipErrcode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrcode
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrcode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowErrcode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthErrcode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErrcode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErrcode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErrcode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrcode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErrcode = fmt.Errorf("proto: unexpected end of group")
)
