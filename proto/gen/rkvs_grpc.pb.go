// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/rkvs.proto

package rkvs

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
	Rkvs_Get_FullMethodName                = "/rkvs.Rkvs/Get"
	Rkvs_ExecuteTransaction_FullMethodName = "/rkvs.Rkvs/ExecuteTransaction"
)

// RkvsClient is the client API for Rkvs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RkvsClient interface {
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	ExecuteTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*Ack, error)
}

type rkvsClient struct {
	cc grpc.ClientConnInterface
}

func NewRkvsClient(cc grpc.ClientConnInterface) RkvsClient {
	return &rkvsClient{cc}
}

func (c *rkvsClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, Rkvs_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rkvsClient) ExecuteTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, Rkvs_ExecuteTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RkvsServer is the server API for Rkvs service.
// All implementations must embed UnimplementedRkvsServer
// for forward compatibility
type RkvsServer interface {
	Get(context.Context, *Key) (*Value, error)
	ExecuteTransaction(context.Context, *TransactionRequest) (*Ack, error)
	mustEmbedUnimplementedRkvsServer()
}

// UnimplementedRkvsServer must be embedded to have forward compatible implementations.
type UnimplementedRkvsServer struct {
}

func (UnimplementedRkvsServer) Get(context.Context, *Key) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRkvsServer) ExecuteTransaction(context.Context, *TransactionRequest) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteTransaction not implemented")
}
func (UnimplementedRkvsServer) mustEmbedUnimplementedRkvsServer() {}

// UnsafeRkvsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RkvsServer will
// result in compilation errors.
type UnsafeRkvsServer interface {
	mustEmbedUnimplementedRkvsServer()
}

func RegisterRkvsServer(s grpc.ServiceRegistrar, srv RkvsServer) {
	s.RegisterService(&Rkvs_ServiceDesc, srv)
}

func _Rkvs_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RkvsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rkvs_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RkvsServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rkvs_ExecuteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RkvsServer).ExecuteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rkvs_ExecuteTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RkvsServer).ExecuteTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rkvs_ServiceDesc is the grpc.ServiceDesc for Rkvs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rkvs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rkvs.Rkvs",
	HandlerType: (*RkvsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Rkvs_Get_Handler,
		},
		{
			MethodName: "ExecuteTransaction",
			Handler:    _Rkvs_ExecuteTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rkvs.proto",
}