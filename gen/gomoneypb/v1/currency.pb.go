// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: gomoneypb/v1/currency.proto

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

type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Rate          string                 `protobuf:"bytes,2,opt,name=rate,proto3" json:"rate,omitempty"`
	IsActive      bool                   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	DecimalPlaces int32                  `protobuf:"varint,4,opt,name=decimal_places,json=decimalPlaces,proto3" json:"decimal_places,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomoneypb_v1_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_gomoneypb_v1_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Currency.ProtoReflect.Descriptor instead.
func (*Currency) Descriptor() ([]byte, []int) {
	return file_gomoneypb_v1_currency_proto_rawDescGZIP(), []int{0}
}

func (x *Currency) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Currency) GetRate() string {
	if x != nil {
		return x.Rate
	}
	return ""
}

func (x *Currency) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Currency) GetDecimalPlaces() int32 {
	if x != nil {
		return x.DecimalPlaces
	}
	return 0
}

func (x *Currency) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Currency) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_gomoneypb_v1_currency_proto protoreflect.FileDescriptor

var file_gomoneypb_v1_currency_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67,
	0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a,
	0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x73, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0xac, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67,
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
	file_gomoneypb_v1_currency_proto_rawDescOnce sync.Once
	file_gomoneypb_v1_currency_proto_rawDescData = file_gomoneypb_v1_currency_proto_rawDesc
)

func file_gomoneypb_v1_currency_proto_rawDescGZIP() []byte {
	file_gomoneypb_v1_currency_proto_rawDescOnce.Do(func() {
		file_gomoneypb_v1_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_gomoneypb_v1_currency_proto_rawDescData)
	})
	return file_gomoneypb_v1_currency_proto_rawDescData
}

var file_gomoneypb_v1_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gomoneypb_v1_currency_proto_goTypes = []any{
	(*Currency)(nil),              // 0: gomoneypb.v1.Currency
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_gomoneypb_v1_currency_proto_depIdxs = []int32{
	1, // 0: gomoneypb.v1.Currency.updated_at:type_name -> google.protobuf.Timestamp
	1, // 1: gomoneypb.v1.Currency.deleted_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gomoneypb_v1_currency_proto_init() }
func file_gomoneypb_v1_currency_proto_init() {
	if File_gomoneypb_v1_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gomoneypb_v1_currency_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Currency); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gomoneypb_v1_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gomoneypb_v1_currency_proto_goTypes,
		DependencyIndexes: file_gomoneypb_v1_currency_proto_depIdxs,
		MessageInfos:      file_gomoneypb_v1_currency_proto_msgTypes,
	}.Build()
	File_gomoneypb_v1_currency_proto = out.File
	file_gomoneypb_v1_currency_proto_rawDesc = nil
	file_gomoneypb_v1_currency_proto_goTypes = nil
	file_gomoneypb_v1_currency_proto_depIdxs = nil
}
