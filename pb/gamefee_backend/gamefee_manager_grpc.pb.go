// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gamefee_manager

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

// GameFeeManagerClient is the client API for GameFeeManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameFeeManagerClient interface {
	ReturnGameFee(ctx context.Context, in *ReturnGameFeeRequest, opts ...grpc.CallOption) (*ReturnGameFeeResponse, error)
}

type gameFeeManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewGameFeeManagerClient(cc grpc.ClientConnInterface) GameFeeManagerClient {
	return &gameFeeManagerClient{cc}
}

func (c *gameFeeManagerClient) ReturnGameFee(ctx context.Context, in *ReturnGameFeeRequest, opts ...grpc.CallOption) (*ReturnGameFeeResponse, error) {
	out := new(ReturnGameFeeResponse)
	err := c.cc.Invoke(ctx, "/gamefee_manager.api.GameFeeManager/ReturnGameFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameFeeManagerServer is the server API for GameFeeManager service.
// All implementations must embed UnimplementedGameFeeManagerServer
// for forward compatibility
type GameFeeManagerServer interface {
	ReturnGameFee(context.Context, *ReturnGameFeeRequest) (*ReturnGameFeeResponse, error)
	mustEmbedUnimplementedGameFeeManagerServer()
}

// UnimplementedGameFeeManagerServer must be embedded to have forward compatible implementations.
type UnimplementedGameFeeManagerServer struct {
}

func (UnimplementedGameFeeManagerServer) ReturnGameFee(context.Context, *ReturnGameFeeRequest) (*ReturnGameFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnGameFee not implemented")
}
func (UnimplementedGameFeeManagerServer) mustEmbedUnimplementedGameFeeManagerServer() {}

// UnsafeGameFeeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameFeeManagerServer will
// result in compilation errors.
type UnsafeGameFeeManagerServer interface {
	mustEmbedUnimplementedGameFeeManagerServer()
}

func RegisterGameFeeManagerServer(s grpc.ServiceRegistrar, srv GameFeeManagerServer) {
	s.RegisterService(&GameFeeManager_ServiceDesc, srv)
}

func _GameFeeManager_ReturnGameFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnGameFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameFeeManagerServer).ReturnGameFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gamefee_manager.api.GameFeeManager/ReturnGameFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameFeeManagerServer).ReturnGameFee(ctx, req.(*ReturnGameFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GameFeeManager_ServiceDesc is the grpc.ServiceDesc for GameFeeManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameFeeManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gamefee_manager.api.GameFeeManager",
	HandlerType: (*GameFeeManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReturnGameFee",
			Handler:    _GameFeeManager_ReturnGameFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gamefee_manager.proto",
}
