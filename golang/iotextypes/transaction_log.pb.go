// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/types/transaction_log.proto

package iotextypes

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

type TransactionLogType int32

const (
	TransactionLogType_IN_CONTRACT_TRANSFER       TransactionLogType = 0
	TransactionLogType_BUCKET_WITHDRAW_AMOUNT     TransactionLogType = 1
	TransactionLogType_BUCKET_CREATE_AMOUNT       TransactionLogType = 2
	TransactionLogType_BUCKET_DEPOSIT_AMOUNT      TransactionLogType = 3
	TransactionLogType_CANDIDATE_SELF_STAKE       TransactionLogType = 4
	TransactionLogType_CANDIDATE_REGISTRATION_FEE TransactionLogType = 5
	TransactionLogType_GAS_FEE                    TransactionLogType = 6
	TransactionLogType_NATIVE_TRANSFER            TransactionLogType = 7
	TransactionLogType_DEPOSIT_TO_REWARDING_FUND  TransactionLogType = 8
	TransactionLogType_CLAIM_FROM_REWARDING_FUND  TransactionLogType = 9
)

var TransactionLogType_name = map[int32]string{
	0: "IN_CONTRACT_TRANSFER",
	1: "BUCKET_WITHDRAW_AMOUNT",
	2: "BUCKET_CREATE_AMOUNT",
	3: "BUCKET_DEPOSIT_AMOUNT",
	4: "CANDIDATE_SELF_STAKE",
	5: "CANDIDATE_REGISTRATION_FEE",
	6: "GAS_FEE",
	7: "NATIVE_TRANSFER",
	8: "DEPOSIT_TO_REWARDING_FUND",
	9: "CLAIM_FROM_REWARDING_FUND",
}

var TransactionLogType_value = map[string]int32{
	"IN_CONTRACT_TRANSFER":       0,
	"BUCKET_WITHDRAW_AMOUNT":     1,
	"BUCKET_CREATE_AMOUNT":       2,
	"BUCKET_DEPOSIT_AMOUNT":      3,
	"CANDIDATE_SELF_STAKE":       4,
	"CANDIDATE_REGISTRATION_FEE": 5,
	"GAS_FEE":                    6,
	"NATIVE_TRANSFER":            7,
	"DEPOSIT_TO_REWARDING_FUND":  8,
	"CLAIM_FROM_REWARDING_FUND":  9,
}

func (x TransactionLogType) String() string {
	return proto.EnumName(TransactionLogType_name, int32(x))
}

func (TransactionLogType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce17c7385702f675, []int{0}
}

type TransactionLog struct {
	ActionHash           []byte                        `protobuf:"bytes,1,opt,name=actionHash,proto3" json:"actionHash,omitempty"`
	NumTransactions      uint64                        `protobuf:"varint,2,opt,name=numTransactions,proto3" json:"numTransactions,omitempty"`
	Transactions         []*TransactionLog_Transaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *TransactionLog) Reset()         { *m = TransactionLog{} }
func (m *TransactionLog) String() string { return proto.CompactTextString(m) }
func (*TransactionLog) ProtoMessage()    {}
func (*TransactionLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce17c7385702f675, []int{0}
}

func (m *TransactionLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionLog.Unmarshal(m, b)
}
func (m *TransactionLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionLog.Marshal(b, m, deterministic)
}
func (m *TransactionLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionLog.Merge(m, src)
}
func (m *TransactionLog) XXX_Size() int {
	return xxx_messageInfo_TransactionLog.Size(m)
}
func (m *TransactionLog) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionLog.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionLog proto.InternalMessageInfo

func (m *TransactionLog) GetActionHash() []byte {
	if m != nil {
		return m.ActionHash
	}
	return nil
}

func (m *TransactionLog) GetNumTransactions() uint64 {
	if m != nil {
		return m.NumTransactions
	}
	return 0
}

func (m *TransactionLog) GetTransactions() []*TransactionLog_Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type TransactionLog_Transaction struct {
	Topic                []byte   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Amount               string   `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Sender               string   `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            string   `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionLog_Transaction) Reset()         { *m = TransactionLog_Transaction{} }
func (m *TransactionLog_Transaction) String() string { return proto.CompactTextString(m) }
func (*TransactionLog_Transaction) ProtoMessage()    {}
func (*TransactionLog_Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce17c7385702f675, []int{0, 0}
}

func (m *TransactionLog_Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionLog_Transaction.Unmarshal(m, b)
}
func (m *TransactionLog_Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionLog_Transaction.Marshal(b, m, deterministic)
}
func (m *TransactionLog_Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionLog_Transaction.Merge(m, src)
}
func (m *TransactionLog_Transaction) XXX_Size() int {
	return xxx_messageInfo_TransactionLog_Transaction.Size(m)
}
func (m *TransactionLog_Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionLog_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionLog_Transaction proto.InternalMessageInfo

func (m *TransactionLog_Transaction) GetTopic() []byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *TransactionLog_Transaction) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *TransactionLog_Transaction) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *TransactionLog_Transaction) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

type BlockTransactionLog struct {
	NumTransactions      uint64            `protobuf:"varint,1,opt,name=numTransactions,proto3" json:"numTransactions,omitempty"`
	TransactionLog       []*TransactionLog `protobuf:"bytes,2,rep,name=transactionLog,proto3" json:"transactionLog,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BlockTransactionLog) Reset()         { *m = BlockTransactionLog{} }
func (m *BlockTransactionLog) String() string { return proto.CompactTextString(m) }
func (*BlockTransactionLog) ProtoMessage()    {}
func (*BlockTransactionLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce17c7385702f675, []int{1}
}

func (m *BlockTransactionLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockTransactionLog.Unmarshal(m, b)
}
func (m *BlockTransactionLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockTransactionLog.Marshal(b, m, deterministic)
}
func (m *BlockTransactionLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTransactionLog.Merge(m, src)
}
func (m *BlockTransactionLog) XXX_Size() int {
	return xxx_messageInfo_BlockTransactionLog.Size(m)
}
func (m *BlockTransactionLog) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTransactionLog.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTransactionLog proto.InternalMessageInfo

func (m *BlockTransactionLog) GetNumTransactions() uint64 {
	if m != nil {
		return m.NumTransactions
	}
	return 0
}

func (m *BlockTransactionLog) GetTransactionLog() []*TransactionLog {
	if m != nil {
		return m.TransactionLog
	}
	return nil
}

func init() {
	proto.RegisterEnum("iotextypes.TransactionLogType", TransactionLogType_name, TransactionLogType_value)
	proto.RegisterType((*TransactionLog)(nil), "iotextypes.TransactionLog")
	proto.RegisterType((*TransactionLog_Transaction)(nil), "iotextypes.TransactionLog.Transaction")
	proto.RegisterType((*BlockTransactionLog)(nil), "iotextypes.BlockTransactionLog")
}

func init() { proto.RegisterFile("proto/types/transaction_log.proto", fileDescriptor_ce17c7385702f675) }

var fileDescriptor_ce17c7385702f675 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x6f, 0x6b, 0xda, 0x50,
	0x18, 0xc5, 0x17, 0xb5, 0x76, 0x3e, 0x96, 0x56, 0x6e, 0xbb, 0x92, 0xca, 0x56, 0x9c, 0x2f, 0x86,
	0x0c, 0x16, 0x61, 0x63, 0x1f, 0xe0, 0x9a, 0x5c, 0x6d, 0x56, 0x4d, 0xca, 0xcd, 0x75, 0xc2, 0x60,
	0x5c, 0xd2, 0x2c, 0xa4, 0xd9, 0x34, 0x37, 0x4b, 0xae, 0xb0, 0xbe, 0xde, 0xa7, 0xe8, 0xb7, 0x1d,
	0xf9, 0x23, 0x89, 0xd2, 0xbd, 0x3c, 0xe7, 0x77, 0x0e, 0xf7, 0xe1, 0x84, 0xc0, 0xdb, 0x38, 0x11,
	0x52, 0x8c, 0xe5, 0x63, 0xec, 0xa7, 0x63, 0x99, 0xb8, 0x51, 0xea, 0x7a, 0x32, 0x14, 0x11, 0x5f,
	0x8b, 0x40, 0xcb, 0x19, 0x82, 0x50, 0x48, 0xff, 0x4f, 0x9e, 0x18, 0x3e, 0x35, 0xe0, 0x94, 0x55,
	0xa9, 0xb9, 0x08, 0xd0, 0x35, 0x40, 0x21, 0x6e, 0xdc, 0xf4, 0x41, 0x55, 0x06, 0xca, 0xe8, 0x84,
	0xd6, 0x1c, 0x34, 0x82, 0xb3, 0x68, 0xbb, 0xa9, 0x95, 0x52, 0xb5, 0x31, 0x50, 0x46, 0x2d, 0x7a,
	0x68, 0xa3, 0x2f, 0x70, 0x22, 0xeb, 0xb1, 0xe6, 0xa0, 0x39, 0xea, 0x7e, 0x7c, 0xa7, 0x55, 0xef,
	0x6b, 0xfb, 0x6f, 0xd7, 0x25, 0xdd, 0xeb, 0xf6, 0x7f, 0x43, 0xb7, 0x06, 0xd1, 0x05, 0x1c, 0x49,
	0x11, 0x87, 0x5e, 0x79, 0x5f, 0x21, 0xd0, 0x25, 0xb4, 0xdd, 0x8d, 0xd8, 0x46, 0x32, 0xbf, 0xa8,
	0x43, 0x4b, 0x95, 0xf9, 0xa9, 0x1f, 0xfd, 0xf0, 0x13, 0xb5, 0x59, 0xf8, 0x85, 0x42, 0xaf, 0xa1,
	0x93, 0xf8, 0x5e, 0x18, 0x87, 0x7e, 0x24, 0xd5, 0x56, 0x8e, 0x2a, 0x63, 0xf8, 0x57, 0x81, 0xf3,
	0xc9, 0x5a, 0x78, 0xbf, 0x0e, 0x06, 0x7a, 0x66, 0x00, 0xe5, 0xf9, 0x01, 0x26, 0x70, 0x2a, 0xf7,
	0xba, 0x6a, 0x23, 0x9f, 0xa0, 0xff, 0xff, 0x09, 0xe8, 0x41, 0xe3, 0xfd, 0x53, 0x03, 0xd0, 0x7e,
	0x84, 0x3d, 0xc6, 0x3e, 0x52, 0xe1, 0xc2, 0xb4, 0xb8, 0x6e, 0x5b, 0x8c, 0x62, 0x9d, 0x71, 0x46,
	0xb1, 0xe5, 0x4c, 0x09, 0xed, 0xbd, 0x40, 0x7d, 0xb8, 0x9c, 0x2c, 0xf5, 0x5b, 0xc2, 0xf8, 0xca,
	0x64, 0x37, 0x06, 0xc5, 0x2b, 0x8e, 0x17, 0xf6, 0xd2, 0x62, 0x3d, 0x25, 0x6b, 0x95, 0x4c, 0xa7,
	0x04, 0x33, 0xb2, 0x23, 0x0d, 0x74, 0x05, 0xaf, 0x4a, 0x62, 0x90, 0x3b, 0xdb, 0x31, 0xd9, 0x0e,
	0x35, 0xb3, 0x92, 0x8e, 0x2d, 0xc3, 0x34, 0xb2, 0x82, 0x43, 0xe6, 0x53, 0xee, 0x30, 0x7c, 0x4b,
	0x7a, 0x2d, 0x74, 0x0d, 0xfd, 0x8a, 0x50, 0x32, 0x33, 0x1d, 0x46, 0x31, 0x33, 0x6d, 0x8b, 0x4f,
	0x09, 0xe9, 0x1d, 0xa1, 0x2e, 0x1c, 0xcf, 0xb0, 0x93, 0x8b, 0x36, 0x3a, 0x87, 0x33, 0x0b, 0x33,
	0xf3, 0x2b, 0xa9, 0x8e, 0x3d, 0x46, 0x6f, 0xe0, 0x6a, 0xf7, 0x1e, 0xb3, 0x39, 0x25, 0x2b, 0x4c,
	0x0d, 0xd3, 0x9a, 0xf1, 0xe9, 0xd2, 0x32, 0x7a, 0x2f, 0x33, 0xac, 0xcf, 0xb1, 0xb9, 0xe0, 0x53,
	0x6a, 0x2f, 0x0e, 0x71, 0x67, 0xf2, 0x1d, 0x86, 0x9e, 0xd8, 0x68, 0x41, 0x28, 0x1f, 0xb6, 0xf7,
	0xc5, 0xa6, 0x71, 0x22, 0x7e, 0xfa, 0x9e, 0xd4, 0x82, 0x24, 0xf6, 0xb4, 0x7c, 0xe0, 0x3b, 0xe5,
	0xdb, 0xe7, 0x32, 0xe1, 0x89, 0xcd, 0xb8, 0x9e, 0x2a, 0xc4, 0x87, 0xe2, 0x97, 0x09, 0xc4, 0xda,
	0x8d, 0x82, 0x71, 0xf5, 0x65, 0xee, 0xdb, 0x39, 0xf8, 0xf4, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x34,
	0xbb, 0x61, 0x22, 0x54, 0x03, 0x00, 0x00,
}
