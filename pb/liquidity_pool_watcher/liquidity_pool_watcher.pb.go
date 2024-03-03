// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: liquidity_pool_watcher.proto

package liqudidity_pool_watcher

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

type PositionUpdatedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainId    uint64 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	PoolId     uint64 `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	LpAddress  string `protobuf:"bytes,3,opt,name=lp_address,json=lpAddress,proto3" json:"lp_address,omitempty"`
	IsIncrease bool   `protobuf:"varint,4,opt,name=is_increase,json=isIncrease,proto3" json:"is_increase,omitempty"`
	Amount     string `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PositionUpdatedRequest) Reset() {
	*x = PositionUpdatedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liquidity_pool_watcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionUpdatedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionUpdatedRequest) ProtoMessage() {}

func (x *PositionUpdatedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_liquidity_pool_watcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionUpdatedRequest.ProtoReflect.Descriptor instead.
func (*PositionUpdatedRequest) Descriptor() ([]byte, []int) {
	return file_liquidity_pool_watcher_proto_rawDescGZIP(), []int{0}
}

func (x *PositionUpdatedRequest) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *PositionUpdatedRequest) GetPoolId() uint64 {
	if x != nil {
		return x.PoolId
	}
	return 0
}

func (x *PositionUpdatedRequest) GetLpAddress() string {
	if x != nil {
		return x.LpAddress
	}
	return ""
}

func (x *PositionUpdatedRequest) GetIsIncrease() bool {
	if x != nil {
		return x.IsIncrease
	}
	return false
}

func (x *PositionUpdatedRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type PositionUpdatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PositionUpdatedResponse) Reset() {
	*x = PositionUpdatedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liquidity_pool_watcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionUpdatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionUpdatedResponse) ProtoMessage() {}

func (x *PositionUpdatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_liquidity_pool_watcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PositionUpdatedResponse.ProtoReflect.Descriptor instead.
func (*PositionUpdatedResponse) Descriptor() ([]byte, []int) {
	return file_liquidity_pool_watcher_proto_rawDescGZIP(), []int{1}
}

func (x *PositionUpdatedResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_liquidity_pool_watcher_proto protoreflect.FileDescriptor

var file_liquidity_pool_watcher_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x6c, 0x69, 0x71, 0x75, 0x64, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x22, 0xa4, 0x01, 0x0a, 0x16,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x70,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6c, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x69, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x31, 0x0a, 0x17, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x96, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x7e,
	0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x33, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x64, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x70,
	0x6f, 0x6f, 0x6c, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x6c, 0x69, 0x71, 0x75, 0x64, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1c,
	0x5a, 0x1a, 0x2e, 0x2f, 0x3b, 0x6c, 0x69, 0x71, 0x75, 0x64, 0x69, 0x64, 0x69, 0x74, 0x79, 0x5f,
	0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_liquidity_pool_watcher_proto_rawDescOnce sync.Once
	file_liquidity_pool_watcher_proto_rawDescData = file_liquidity_pool_watcher_proto_rawDesc
)

func file_liquidity_pool_watcher_proto_rawDescGZIP() []byte {
	file_liquidity_pool_watcher_proto_rawDescOnce.Do(func() {
		file_liquidity_pool_watcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_liquidity_pool_watcher_proto_rawDescData)
	})
	return file_liquidity_pool_watcher_proto_rawDescData
}

var file_liquidity_pool_watcher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_liquidity_pool_watcher_proto_goTypes = []interface{}{
	(*PositionUpdatedRequest)(nil),  // 0: liqudidity_pool_watcher.api.PositionUpdatedRequest
	(*PositionUpdatedResponse)(nil), // 1: liqudidity_pool_watcher.api.PositionUpdatedResponse
}
var file_liquidity_pool_watcher_proto_depIdxs = []int32{
	0, // 0: liqudidity_pool_watcher.api.LiquidityPoolWatcher.PositionUpdated:input_type -> liqudidity_pool_watcher.api.PositionUpdatedRequest
	1, // 1: liqudidity_pool_watcher.api.LiquidityPoolWatcher.PositionUpdated:output_type -> liqudidity_pool_watcher.api.PositionUpdatedResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_liquidity_pool_watcher_proto_init() }
func file_liquidity_pool_watcher_proto_init() {
	if File_liquidity_pool_watcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_liquidity_pool_watcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionUpdatedRequest); i {
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
		file_liquidity_pool_watcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionUpdatedResponse); i {
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
			RawDescriptor: file_liquidity_pool_watcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_liquidity_pool_watcher_proto_goTypes,
		DependencyIndexes: file_liquidity_pool_watcher_proto_depIdxs,
		MessageInfos:      file_liquidity_pool_watcher_proto_msgTypes,
	}.Build()
	File_liquidity_pool_watcher_proto = out.File
	file_liquidity_pool_watcher_proto_rawDesc = nil
	file_liquidity_pool_watcher_proto_goTypes = nil
	file_liquidity_pool_watcher_proto_depIdxs = nil
}
