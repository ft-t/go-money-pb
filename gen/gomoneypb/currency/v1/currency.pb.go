// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: gomoneypb/currency/v1/currency.proto

package currencyv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromCurrency string `protobuf:"bytes,1,opt,name=from_currency,json=fromCurrency,proto3" json:"from_currency,omitempty"`
	ToCurrency   string `protobuf:"bytes,2,opt,name=to_currency,json=toCurrency,proto3" json:"to_currency,omitempty"`
	Amount       string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ExchangeRequest) Reset() {
	*x = ExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomoneypb_currency_v1_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeRequest) ProtoMessage() {}

func (x *ExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gomoneypb_currency_v1_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeRequest.ProtoReflect.Descriptor instead.
func (*ExchangeRequest) Descriptor() ([]byte, []int) {
	return file_gomoneypb_currency_v1_currency_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeRequest) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *ExchangeRequest) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

func (x *ExchangeRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type ExchangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ExchangeResponse) Reset() {
	*x = ExchangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gomoneypb_currency_v1_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeResponse) ProtoMessage() {}

func (x *ExchangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gomoneypb_currency_v1_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeResponse.ProtoReflect.Descriptor instead.
func (*ExchangeResponse) Descriptor() ([]byte, []int) {
	return file_gomoneypb_currency_v1_currency_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeResponse) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_gomoneypb_currency_v1_currency_proto protoreflect.FileDescriptor

var file_gomoneypb_currency_v1_currency_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70,
	0x62, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x6f, 0x0a,
	0x0f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2a,
	0x0a, 0x10, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x70, 0x0a, 0x0f, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a,
	0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x26, 0x2e, 0x67, 0x6f, 0x6d, 0x6f,
	0x6e, 0x65, 0x79, 0x70, 0x62, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xe2, 0x01, 0x0a,
	0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x74, 0x2d, 0x74, 0x2f, 0x67, 0x6f, 0x2d,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x2d, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x47, 0x43, 0x58, 0xaa, 0x02, 0x15, 0x47, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x2e,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x47, 0x6f,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x5c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x47, 0x6f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x70, 0x62, 0x5c,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x47, 0x6f, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x70, 0x62, 0x3a, 0x3a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gomoneypb_currency_v1_currency_proto_rawDescOnce sync.Once
	file_gomoneypb_currency_v1_currency_proto_rawDescData = file_gomoneypb_currency_v1_currency_proto_rawDesc
)

func file_gomoneypb_currency_v1_currency_proto_rawDescGZIP() []byte {
	file_gomoneypb_currency_v1_currency_proto_rawDescOnce.Do(func() {
		file_gomoneypb_currency_v1_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_gomoneypb_currency_v1_currency_proto_rawDescData)
	})
	return file_gomoneypb_currency_v1_currency_proto_rawDescData
}

var file_gomoneypb_currency_v1_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gomoneypb_currency_v1_currency_proto_goTypes = []any{
	(*ExchangeRequest)(nil),  // 0: gomoneypb.currency.v1.ExchangeRequest
	(*ExchangeResponse)(nil), // 1: gomoneypb.currency.v1.ExchangeResponse
}
var file_gomoneypb_currency_v1_currency_proto_depIdxs = []int32{
	0, // 0: gomoneypb.currency.v1.CurrencyService.Exchange:input_type -> gomoneypb.currency.v1.ExchangeRequest
	1, // 1: gomoneypb.currency.v1.CurrencyService.Exchange:output_type -> gomoneypb.currency.v1.ExchangeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gomoneypb_currency_v1_currency_proto_init() }
func file_gomoneypb_currency_v1_currency_proto_init() {
	if File_gomoneypb_currency_v1_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gomoneypb_currency_v1_currency_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExchangeRequest); i {
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
		file_gomoneypb_currency_v1_currency_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ExchangeResponse); i {
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
			RawDescriptor: file_gomoneypb_currency_v1_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gomoneypb_currency_v1_currency_proto_goTypes,
		DependencyIndexes: file_gomoneypb_currency_v1_currency_proto_depIdxs,
		MessageInfos:      file_gomoneypb_currency_v1_currency_proto_msgTypes,
	}.Build()
	File_gomoneypb_currency_v1_currency_proto = out.File
	file_gomoneypb_currency_v1_currency_proto_rawDesc = nil
	file_gomoneypb_currency_v1_currency_proto_goTypes = nil
	file_gomoneypb_currency_v1_currency_proto_depIdxs = nil
}
