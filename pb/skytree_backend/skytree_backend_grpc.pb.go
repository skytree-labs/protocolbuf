// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package skytree_backend

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SkyteeBackendClient is the client API for SkyteeBackend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SkyteeBackendClient interface {
	GetCrossChainTxStatus(ctx context.Context, in *GetCrossChainTxStatusRequest, opts ...grpc.CallOption) (*GetCrossChainTxStatusResponse, error)
	GetCrossChainTxDetail(ctx context.Context, in *GetCrossChainTxDetailRequest, opts ...grpc.CallOption) (*GetCrossChainTxDetailResponse, error)
	DepositPoint(ctx context.Context, in *DepositPointRequest, opts ...grpc.CallOption) (*DepositPointResponse, error)
}

type skyteeBackendClient struct {
	cc grpc.ClientConnInterface
}

func NewSkyteeBackendClient(cc grpc.ClientConnInterface) SkyteeBackendClient {
	return &skyteeBackendClient{cc}
}

func (c *skyteeBackendClient) GetCrossChainTxStatus(ctx context.Context, in *GetCrossChainTxStatusRequest, opts ...grpc.CallOption) (*GetCrossChainTxStatusResponse, error) {
	out := new(GetCrossChainTxStatusResponse)
	err := c.cc.Invoke(ctx, "/skytree_backend.api.SkyteeBackend/GetCrossChainTxStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skyteeBackendClient) GetCrossChainTxDetail(ctx context.Context, in *GetCrossChainTxDetailRequest, opts ...grpc.CallOption) (*GetCrossChainTxDetailResponse, error) {
	out := new(GetCrossChainTxDetailResponse)
	err := c.cc.Invoke(ctx, "/skytree_backend.api.SkyteeBackend/GetCrossChainTxDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skyteeBackendClient) DepositPoint(ctx context.Context, in *DepositPointRequest, opts ...grpc.CallOption) (*DepositPointResponse, error) {
	out := new(DepositPointResponse)
	err := c.cc.Invoke(ctx, "/skytree_backend.api.SkyteeBackend/DepositPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SkyteeBackendServer is the server API for SkyteeBackend service.
// All implementations must embed UnimplementedSkyteeBackendServer
// for forward compatibility
type SkyteeBackendServer interface {
	GetCrossChainTxStatus(context.Context, *GetCrossChainTxStatusRequest) (*GetCrossChainTxStatusResponse, error)
	GetCrossChainTxDetail(context.Context, *GetCrossChainTxDetailRequest) (*GetCrossChainTxDetailResponse, error)
	DepositPoint(context.Context, *DepositPointRequest) (*DepositPointResponse, error)
	mustEmbedUnimplementedSkyteeBackendServer()
}

// UnimplementedSkyteeBackendServer must be embedded to have forward compatible implementations.
type UnimplementedSkyteeBackendServer struct {
}

func (UnimplementedSkyteeBackendServer) GetCrossChainTxStatus(context.Context, *GetCrossChainTxStatusRequest) (*GetCrossChainTxStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCrossChainTxStatus not implemented")
}
func (UnimplementedSkyteeBackendServer) GetCrossChainTxDetail(context.Context, *GetCrossChainTxDetailRequest) (*GetCrossChainTxDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCrossChainTxDetail not implemented")
}
func (UnimplementedSkyteeBackendServer) DepositPoint(context.Context, *DepositPointRequest) (*DepositPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositPoint not implemented")
}
func (UnimplementedSkyteeBackendServer) mustEmbedUnimplementedSkyteeBackendServer() {}

// UnsafeSkyteeBackendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SkyteeBackendServer will
// result in compilation errors.
type UnsafeSkyteeBackendServer interface {
	mustEmbedUnimplementedSkyteeBackendServer()
}

func RegisterSkyteeBackendServer(s grpc.ServiceRegistrar, srv SkyteeBackendServer) {
	s.RegisterService(&SkyteeBackend_ServiceDesc, srv)
}

func _SkyteeBackend_GetCrossChainTxStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCrossChainTxStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkyteeBackendServer).GetCrossChainTxStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skytree_backend.api.SkyteeBackend/GetCrossChainTxStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkyteeBackendServer).GetCrossChainTxStatus(ctx, req.(*GetCrossChainTxStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkyteeBackend_GetCrossChainTxDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCrossChainTxDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkyteeBackendServer).GetCrossChainTxDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skytree_backend.api.SkyteeBackend/GetCrossChainTxDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkyteeBackendServer).GetCrossChainTxDetail(ctx, req.(*GetCrossChainTxDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkyteeBackend_DepositPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkyteeBackendServer).DepositPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skytree_backend.api.SkyteeBackend/DepositPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkyteeBackendServer).DepositPoint(ctx, req.(*DepositPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SkyteeBackend_ServiceDesc is the grpc.ServiceDesc for SkyteeBackend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SkyteeBackend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "skytree_backend.api.SkyteeBackend",
	HandlerType: (*SkyteeBackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCrossChainTxStatus",
			Handler:    _SkyteeBackend_GetCrossChainTxStatus_Handler,
		},
		{
			MethodName: "GetCrossChainTxDetail",
			Handler:    _SkyteeBackend_GetCrossChainTxDetail_Handler,
		},
		{
			MethodName: "DepositPoint",
			Handler:    _SkyteeBackend_DepositPoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "skytree_backend.proto",
}
