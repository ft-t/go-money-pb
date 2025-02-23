// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: gomoneypb/v1/transaction.proto

package gomoneypbv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionType int32

const (
	TransactionType_TRANSACTION_TYPE_UNSPECIFIED               TransactionType = 0
	TransactionType_TRANSACTION_TYPE_TRANSFER_BETWEEN_ACCOUNTS TransactionType = 1
	TransactionType_TRANSACTION_TYPE_DEPOSIT                   TransactionType = 2
	TransactionType_TRANSACTION_TYPE_WITHDRAWAL                TransactionType = 3
	TransactionType_TRANSACTION_TYPE_VOID                      TransactionType = 4
	TransactionType_TRANSACTION_TYPE_RECONCILIATION            TransactionType = 5
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "TRANSACTION_TYPE_UNSPECIFIED",
		1: "TRANSACTION_TYPE_TRANSFER_BETWEEN_ACCOUNTS",
		2: "TRANSACTION_TYPE_DEPOSIT",
		3: "TRANSACTION_TYPE_WITHDRAWAL",
		4: "TRANSACTION_TYPE_VOID",
		5: "TRANSACTION_TYPE_RECONCILIATION",
	}
	TransactionType_value = map[string]int32{
		"TRANSACTION_TYPE_UNSPECIFIED":               0,
		"TRANSACTION_TYPE_TRANSFER_BETWEEN_ACCOUNTS": 1,
		"TRANSACTION_TYPE_DEPOSIT":                   2,
		"TRANSACTION_TYPE_WITHDRAWAL":                3,
		"TRANSACTION_TYPE_VOID":                      4,
		"TRANSACTION_TYPE_RECONCILIATION":            5,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_gomoneypb_v1_transaction_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_gomoneypb_v1_transaction_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_gomoneypb_v1_transaction_proto_rawDescGZIP(), []int{0}
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SourceAmount         *string                `protobuf:"bytes,2,opt,name=source_amount,json=sourceAmount,proto3,oneof" json:"source_amount,omitempty"`
	SourceCurrency       *string                `protobuf:"bytes,3,opt,name=source_currency,json=sourceCurrency,proto3,oneof" json:"source_currency,omitempty"`
	DestinationAmount    *string                `protobuf:"bytes,4,opt,name=destination_amount,json=destinationAmount,proto3,oneof" json:"destination_amount,omitempty"`
	DestinationCurrency  *string                `protobuf:"bytes,5,opt,name=destination_currency,json=destinationCurrency,proto3,oneof" json:"destination_currency,omitempty"`
	SourceAccountId      *int32                 `protobuf:"varint,6,opt,name=source_account_id,json=sourceAccountId,proto3,oneof" json:"source_account_id,omitempty"`
	DestinationAccountId *int32                 `protobuf:"varint,7,opt,name=destination_account_id,json=destinationAccountId,proto3,oneof" json:"destination_account_id,omitempty"`
	LabelIds             []int32                `protobuf:"varint,8,rep,packed,name=label_ids,json=labelIds,proto3" json:"label_ids,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	TransactionDate      *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=transaction_date,json=transactionDate,proto3" json:"transaction_date,omitempty"`
	Type                 TransactionType        `protobuf:"varint,12,opt,name=type,proto3,enum=gomoneypb.v1.TransactionType" json:"type,omitempty"`
	Notes                string                 `protobuf:"bytes,13,opt,name=notes,proto3" json:"notes,omitempty"`
	Extra                map[string]string      `protobuf:"bytes,14,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomoneypb_v1_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_gomoneypb_v1_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_gomoneypb_v1_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Transaction) GetSourceAmount() string {
	if x != nil && x.SourceAmount != nil {
		return *x.SourceAmount
	}
	return ""
}

func (x *Transaction) GetSourceCurrency() string {
	if x != nil && x.SourceCurrency != nil {
		return *x.SourceCurrency
	}
	return ""
}

func (x *Transaction) GetDestinationAmount() string {
	if x != nil && x.DestinationAmount != nil {
		return *x.DestinationAmount
	}
	return ""
}

func (x *Transaction) GetDestinationCurrency() string {
	if x != nil && x.DestinationCurrency != nil {
		return *x.DestinationCurrency
	}
	return ""
}

func (x *Transaction) GetSourceAccountId() int32 {
	if x != nil && x.SourceAccountId != nil {
		return *x.SourceAccountId
	}
	return 0
}

func (x *Transaction) GetDestinationAccountId() int32 {
	if x != nil && x.DestinationAccountId != nil {
		return *x.DestinationAccountId
	}
	return 0
}

func (x *Transaction) GetLabelIds() []int32 {
	if x != nil {
		return x.LabelIds
	}
	return nil
}

func (x *Transaction) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Transaction) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Transaction) GetTransactionDate() *timestamppb.Timestamp {
	if x != nil {
		return x.TransactionDate
	}
	return nil
}

func (x *Transaction) GetType() TransactionType {
	if x != nil {
		return x.Type
	}
	return TransactionType_TRANSACTION_TYPE_UNSPECIFIED
}

func (x *Transaction) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *Transaction) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

var File_gomoneypb_v1_transaction_proto protoreflect.FileDescriptor

var file_gomoneypb_v1_transaction_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xed, 0x06, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x28, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x12, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x11, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x14, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x13, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04,
	0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x16, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x05, 0x52, 0x14, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x08, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x45, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x1a, 0x38,
	0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x15,
	0x0a, 0x13, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x14,
	0x0a, 0x12, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x2a,
	0xe2, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2e, 0x0a, 0x2a, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46,
	0x45, 0x52, 0x5f, 0x42, 0x45, 0x54, 0x57, 0x45, 0x45, 0x4e, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x53, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x4f, 0x53, 0x49,
	0x54, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57,
	0x41, 0x4c, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x4f, 0x49, 0x44, 0x10, 0x04, 0x12,
	0x23, 0x0a, 0x1f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x43, 0x49, 0x4c, 0x49, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x05, 0x42, 0xaf, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x74, 0x2d, 0x74, 0x2f, 0x67,
	0x6f, 0x2d, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2d, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x6f, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x70, 0x62, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x0c,
	0x47, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x47,
	0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x47, 0x6f,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x47, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x70, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gomoneypb_v1_transaction_proto_rawDescOnce sync.Once
	file_gomoneypb_v1_transaction_proto_rawDescData = file_gomoneypb_v1_transaction_proto_rawDesc
)

func file_gomoneypb_v1_transaction_proto_rawDescGZIP() []byte {
	file_gomoneypb_v1_transaction_proto_rawDescOnce.Do(func() {
		file_gomoneypb_v1_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_gomoneypb_v1_transaction_proto_rawDescData)
	})
	return file_gomoneypb_v1_transaction_proto_rawDescData
}

var file_gomoneypb_v1_transaction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gomoneypb_v1_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gomoneypb_v1_transaction_proto_goTypes = []any{
	(TransactionType)(0),          // 0: gomoneypb.v1.TransactionType
	(*Transaction)(nil),           // 1: gomoneypb.v1.Transaction
	nil,                           // 2: gomoneypb.v1.Transaction.ExtraEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_gomoneypb_v1_transaction_proto_depIdxs = []int32{
	3, // 0: gomoneypb.v1.Transaction.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: gomoneypb.v1.Transaction.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: gomoneypb.v1.Transaction.transaction_date:type_name -> google.protobuf.Timestamp
	0, // 3: gomoneypb.v1.Transaction.type:type_name -> gomoneypb.v1.TransactionType
	2, // 4: gomoneypb.v1.Transaction.extra:type_name -> gomoneypb.v1.Transaction.ExtraEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_gomoneypb_v1_transaction_proto_init() }
func file_gomoneypb_v1_transaction_proto_init() {
	if File_gomoneypb_v1_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gomoneypb_v1_transaction_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Transaction); i {
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
	file_gomoneypb_v1_transaction_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gomoneypb_v1_transaction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gomoneypb_v1_transaction_proto_goTypes,
		DependencyIndexes: file_gomoneypb_v1_transaction_proto_depIdxs,
		EnumInfos:         file_gomoneypb_v1_transaction_proto_enumTypes,
		MessageInfos:      file_gomoneypb_v1_transaction_proto_msgTypes,
	}.Build()
	File_gomoneypb_v1_transaction_proto = out.File
	file_gomoneypb_v1_transaction_proto_rawDesc = nil
	file_gomoneypb_v1_transaction_proto_goTypes = nil
	file_gomoneypb_v1_transaction_proto_depIdxs = nil
}
