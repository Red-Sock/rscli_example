// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: api/grpc/rscli_example_api.proto

package rscli_example_api

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
	RscliExampleAPI_Version_FullMethodName = "/rscli_example_api.rscli_exampleAPI/Version"
)

// RscliExampleAPIClient is the client API for RscliExampleAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RscliExampleAPIClient interface {
	Version(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type rscliExampleAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewRscliExampleAPIClient(cc grpc.ClientConnInterface) RscliExampleAPIClient {
	return &rscliExampleAPIClient{cc}
}

func (c *rscliExampleAPIClient) Version(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, RscliExampleAPI_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RscliExampleAPIServer is the server API for RscliExampleAPI service.
// All implementations must embed UnimplementedRscliExampleAPIServer
// for forward compatibility
type RscliExampleAPIServer interface {
	Version(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedRscliExampleAPIServer()
}

// UnimplementedRscliExampleAPIServer must be embedded to have forward compatible implementations.
type UnimplementedRscliExampleAPIServer struct {
}

func (UnimplementedRscliExampleAPIServer) Version(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedRscliExampleAPIServer) mustEmbedUnimplementedRscliExampleAPIServer() {}

// UnsafeRscliExampleAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RscliExampleAPIServer will
// result in compilation errors.
type UnsafeRscliExampleAPIServer interface {
	mustEmbedUnimplementedRscliExampleAPIServer()
}

func RegisterRscliExampleAPIServer(s grpc.ServiceRegistrar, srv RscliExampleAPIServer) {
	s.RegisterService(&RscliExampleAPI_ServiceDesc, srv)
}

func _RscliExampleAPI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RscliExampleAPIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RscliExampleAPI_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RscliExampleAPIServer).Version(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RscliExampleAPI_ServiceDesc is the grpc.ServiceDesc for RscliExampleAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RscliExampleAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rscli_example_api.rscli_exampleAPI",
	HandlerType: (*RscliExampleAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _RscliExampleAPI_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/rscli_example_api.proto",
}