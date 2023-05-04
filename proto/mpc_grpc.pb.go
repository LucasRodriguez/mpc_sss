// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0--rc2
// source: proto/mpc.proto

package proto

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

const (
	MPC_ComputeSum_FullMethodName = "/mpc.MPC/ComputeSum"
)

// MPCClient is the client API for MPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MPCClient interface {
	ComputeSum(ctx context.Context, in *ComputeSumRequest, opts ...grpc.CallOption) (*ComputeSumResponse, error)
}

type mPCClient struct {
	cc grpc.ClientConnInterface
}

func NewMPCClient(cc grpc.ClientConnInterface) MPCClient {
	return &mPCClient{cc}
}

func (c *mPCClient) ComputeSum(ctx context.Context, in *ComputeSumRequest, opts ...grpc.CallOption) (*ComputeSumResponse, error) {
	out := new(ComputeSumResponse)
	err := c.cc.Invoke(ctx, MPC_ComputeSum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MPCServer is the server API for MPC service.
// All implementations must embed UnimplementedMPCServer
// for forward compatibility
type MPCServer interface {
	ComputeSum(context.Context, *ComputeSumRequest) (*ComputeSumResponse, error)
	mustEmbedUnimplementedMPCServer()
}

// UnimplementedMPCServer must be embedded to have forward compatible implementations.
type UnimplementedMPCServer struct {
}

func (UnimplementedMPCServer) ComputeSum(context.Context, *ComputeSumRequest) (*ComputeSumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeSum not implemented")
}
func (UnimplementedMPCServer) mustEmbedUnimplementedMPCServer() {}

// UnsafeMPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MPCServer will
// result in compilation errors.
type UnsafeMPCServer interface {
	mustEmbedUnimplementedMPCServer()
}

func RegisterMPCServer(s grpc.ServiceRegistrar, srv MPCServer) {
	s.RegisterService(&MPC_ServiceDesc, srv)
}

func _MPC_ComputeSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeSumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPCServer).ComputeSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MPC_ComputeSum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPCServer).ComputeSum(ctx, req.(*ComputeSumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MPC_ServiceDesc is the grpc.ServiceDesc for MPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mpc.MPC",
	HandlerType: (*MPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeSum",
			Handler:    _MPC_ComputeSum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mpc.proto",
}
