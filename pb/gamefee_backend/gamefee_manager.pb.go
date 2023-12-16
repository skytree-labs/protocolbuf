// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: gamefee_manager.proto

package gamefee_manager

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

type ReturnGameFeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId   int64  `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	To        string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	CrossTxId int64  `protobuf:"varint,3,opt,name=cross_tx_id,json=crossTxId,proto3" json:"cross_tx_id,omitempty"`
	TossId    int64  `protobuf:"varint,4,opt,name=toss_id,json=tossId,proto3" json:"toss_id,omitempty"`
}

func (x *ReturnGameFeeRequest) Reset() {
	*x = ReturnGameFeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamefee_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnGameFeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnGameFeeRequest) ProtoMessage() {}

func (x *ReturnGameFeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gamefee_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnGameFeeRequest.ProtoReflect.Descriptor instead.
func (*ReturnGameFeeRequest) Descriptor() ([]byte, []int) {
	return file_gamefee_manager_proto_rawDescGZIP(), []int{0}
}

func (x *ReturnGameFeeRequest) GetChainId() int64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *ReturnGameFeeRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *ReturnGameFeeRequest) GetCrossTxId() int64 {
	if x != nil {
		return x.CrossTxId
	}
	return 0
}

func (x *ReturnGameFeeRequest) GetTossId() int64 {
	if x != nil {
		return x.TossId
	}
	return 0
}

type ReturnGameFeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReturnGameFeeResponse) Reset() {
	*x = ReturnGameFeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamefee_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnGameFeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnGameFeeResponse) ProtoMessage() {}

func (x *ReturnGameFeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gamefee_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnGameFeeResponse.ProtoReflect.Descriptor instead.
func (*ReturnGameFeeResponse) Descriptor() ([]byte, []int) {
	return file_gamefee_manager_proto_rawDescGZIP(), []int{1}
}

var File_gamefee_manager_proto protoreflect.FileDescriptor

var file_gamefee_manager_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x61, 0x6d, 0x65, 0x66, 0x65, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x61, 0x6d, 0x65, 0x66, 0x65, 0x65,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x7a, 0x0a, 0x14,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x1e, 0x0a, 0x0b, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x5f, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x54, 0x78, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x6f, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x74, 0x6f, 0x73, 0x73, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x7a, 0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x46, 0x65, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x68, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x47, 0x61, 0x6d,
	0x65, 0x46, 0x65, 0x65, 0x12, 0x29, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x66, 0x65, 0x65, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x47, 0x61, 0x6d, 0x65, 0x46, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x66, 0x65, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x47, 0x61, 0x6d, 0x65,
	0x46, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a,
	0x12, 0x2e, 0x2f, 0x3b, 0x67, 0x61, 0x6d, 0x65, 0x66, 0x65, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gamefee_manager_proto_rawDescOnce sync.Once
	file_gamefee_manager_proto_rawDescData = file_gamefee_manager_proto_rawDesc
)

func file_gamefee_manager_proto_rawDescGZIP() []byte {
	file_gamefee_manager_proto_rawDescOnce.Do(func() {
		file_gamefee_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_gamefee_manager_proto_rawDescData)
	})
	return file_gamefee_manager_proto_rawDescData
}

var file_gamefee_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gamefee_manager_proto_goTypes = []interface{}{
	(*ReturnGameFeeRequest)(nil),  // 0: gamefee_manager.api.ReturnGameFeeRequest
	(*ReturnGameFeeResponse)(nil), // 1: gamefee_manager.api.ReturnGameFeeResponse
}
var file_gamefee_manager_proto_depIdxs = []int32{
	0, // 0: gamefee_manager.api.GameFeeManager.ReturnGameFee:input_type -> gamefee_manager.api.ReturnGameFeeRequest
	1, // 1: gamefee_manager.api.GameFeeManager.ReturnGameFee:output_type -> gamefee_manager.api.ReturnGameFeeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gamefee_manager_proto_init() }
func file_gamefee_manager_proto_init() {
	if File_gamefee_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gamefee_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnGameFeeRequest); i {
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
		file_gamefee_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnGameFeeResponse); i {
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
			RawDescriptor: file_gamefee_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gamefee_manager_proto_goTypes,
		DependencyIndexes: file_gamefee_manager_proto_depIdxs,
		MessageInfos:      file_gamefee_manager_proto_msgTypes,
	}.Build()
	File_gamefee_manager_proto = out.File
	file_gamefee_manager_proto_rawDesc = nil
	file_gamefee_manager_proto_goTypes = nil
	file_gamefee_manager_proto_depIdxs = nil
}