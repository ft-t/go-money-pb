// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: gomoneypb/v1/account.proto

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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Currency        string                 `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	CurrencyBalance string                 `protobuf:"bytes,4,opt,name=currency_balance,json=currencyBalance,proto3" json:"currency_balance,omitempty"`
	Extra           map[string]string      `protobuf:"bytes,5,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt       *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
	Type            string                 `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomoneypb_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_gomoneypb_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_gomoneypb_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Account) GetCurrencyBalance() string {
	if x != nil {
		return x.CurrencyBalance
	}
	return ""
}

func (x *Account) GetExtra() map[string]string {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *Account) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Account) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Account) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_gomoneypb_v1_account_proto protoreflect.FileDescriptor

var file_gomoneypb_v1_account_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x6f,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x03, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x42, 0xab, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x74, 0x2d, 0x74, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x2d, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x70, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x47, 0x6f, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x70, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x47, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x70, 0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x47, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70,
	0x62, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0d, 0x47, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gomoneypb_v1_account_proto_rawDescOnce sync.Once
	file_gomoneypb_v1_account_proto_rawDescData = file_gomoneypb_v1_account_proto_rawDesc
)

func file_gomoneypb_v1_account_proto_rawDescGZIP() []byte {
	file_gomoneypb_v1_account_proto_rawDescOnce.Do(func() {
		file_gomoneypb_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_gomoneypb_v1_account_proto_rawDescData)
	})
	return file_gomoneypb_v1_account_proto_rawDescData
}

var file_gomoneypb_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gomoneypb_v1_account_proto_goTypes = []any{
	(*Account)(nil),               // 0: gomoneypb.v1.Account
	nil,                           // 1: gomoneypb.v1.Account.ExtraEntry
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_gomoneypb_v1_account_proto_depIdxs = []int32{
	1, // 0: gomoneypb.v1.Account.extra:type_name -> gomoneypb.v1.Account.ExtraEntry
	2, // 1: gomoneypb.v1.Account.updated_at:type_name -> google.protobuf.Timestamp
	2, // 2: gomoneypb.v1.Account.deleted_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gomoneypb_v1_account_proto_init() }
func file_gomoneypb_v1_account_proto_init() {
	if File_gomoneypb_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gomoneypb_v1_account_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Account); i {
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
	file_gomoneypb_v1_account_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gomoneypb_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gomoneypb_v1_account_proto_goTypes,
		DependencyIndexes: file_gomoneypb_v1_account_proto_depIdxs,
		MessageInfos:      file_gomoneypb_v1_account_proto_msgTypes,
	}.Build()
	File_gomoneypb_v1_account_proto = out.File
	file_gomoneypb_v1_account_proto_rawDesc = nil
	file_gomoneypb_v1_account_proto_goTypes = nil
	file_gomoneypb_v1_account_proto_depIdxs = nil
}
