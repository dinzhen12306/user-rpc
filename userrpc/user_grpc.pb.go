// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: user.proto

package userrpc

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
	StreamGreeter_AccordingToConditionGet_FullMethodName  = "/stream.StreamGreeter/AccordingToConditionGet"
	StreamGreeter_AccordingToConditionGets_FullMethodName = "/stream.StreamGreeter/AccordingToConditionGets"
	StreamGreeter_InsertUser_FullMethodName               = "/stream.StreamGreeter/InsertUser"
	StreamGreeter_UpdateUser_FullMethodName               = "/stream.StreamGreeter/UpdateUser"
	StreamGreeter_DeletedUser_FullMethodName              = "/stream.StreamGreeter/DeletedUser"
)

// StreamGreeterClient is the client API for StreamGreeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamGreeterClient interface {
	AccordingToConditionGet(ctx context.Context, in *AccordingToConditionGetReq, opts ...grpc.CallOption) (*AccordingToConditionGetResp, error)
	AccordingToConditionGets(ctx context.Context, in *AccordingToConditionGetsReq, opts ...grpc.CallOption) (*AccordingToConditionGetsResp, error)
	InsertUser(ctx context.Context, in *InsertUserReq, opts ...grpc.CallOption) (*InsertUserResp, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
	DeletedUser(ctx context.Context, in *DeletedUserReq, opts ...grpc.CallOption) (*DeletedUserResp, error)
}

type streamGreeterClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamGreeterClient(cc grpc.ClientConnInterface) StreamGreeterClient {
	return &streamGreeterClient{cc}
}

func (c *streamGreeterClient) AccordingToConditionGet(ctx context.Context, in *AccordingToConditionGetReq, opts ...grpc.CallOption) (*AccordingToConditionGetResp, error) {
	out := new(AccordingToConditionGetResp)
	err := c.cc.Invoke(ctx, StreamGreeter_AccordingToConditionGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamGreeterClient) AccordingToConditionGets(ctx context.Context, in *AccordingToConditionGetsReq, opts ...grpc.CallOption) (*AccordingToConditionGetsResp, error) {
	out := new(AccordingToConditionGetsResp)
	err := c.cc.Invoke(ctx, StreamGreeter_AccordingToConditionGets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamGreeterClient) InsertUser(ctx context.Context, in *InsertUserReq, opts ...grpc.CallOption) (*InsertUserResp, error) {
	out := new(InsertUserResp)
	err := c.cc.Invoke(ctx, StreamGreeter_InsertUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamGreeterClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	out := new(UpdateUserResp)
	err := c.cc.Invoke(ctx, StreamGreeter_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamGreeterClient) DeletedUser(ctx context.Context, in *DeletedUserReq, opts ...grpc.CallOption) (*DeletedUserResp, error) {
	out := new(DeletedUserResp)
	err := c.cc.Invoke(ctx, StreamGreeter_DeletedUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamGreeterServer is the server API for StreamGreeter service.
// All implementations must embed UnimplementedStreamGreeterServer
// for forward compatibility
type StreamGreeterServer interface {
	AccordingToConditionGet(context.Context, *AccordingToConditionGetReq) (*AccordingToConditionGetResp, error)
	AccordingToConditionGets(context.Context, *AccordingToConditionGetsReq) (*AccordingToConditionGetsResp, error)
	InsertUser(context.Context, *InsertUserReq) (*InsertUserResp, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error)
	DeletedUser(context.Context, *DeletedUserReq) (*DeletedUserResp, error)
	mustEmbedUnimplementedStreamGreeterServer()
}

// UnimplementedStreamGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedStreamGreeterServer struct {
}

func (UnimplementedStreamGreeterServer) AccordingToConditionGet(context.Context, *AccordingToConditionGetReq) (*AccordingToConditionGetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccordingToConditionGet not implemented")
}
func (UnimplementedStreamGreeterServer) AccordingToConditionGets(context.Context, *AccordingToConditionGetsReq) (*AccordingToConditionGetsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccordingToConditionGets not implemented")
}
func (UnimplementedStreamGreeterServer) InsertUser(context.Context, *InsertUserReq) (*InsertUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertUser not implemented")
}
func (UnimplementedStreamGreeterServer) UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedStreamGreeterServer) DeletedUser(context.Context, *DeletedUserReq) (*DeletedUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletedUser not implemented")
}
func (UnimplementedStreamGreeterServer) mustEmbedUnimplementedStreamGreeterServer() {}

// UnsafeStreamGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamGreeterServer will
// result in compilation errors.
type UnsafeStreamGreeterServer interface {
	mustEmbedUnimplementedStreamGreeterServer()
}

func RegisterStreamGreeterServer(s grpc.ServiceRegistrar, srv StreamGreeterServer) {
	s.RegisterService(&StreamGreeter_ServiceDesc, srv)
}

func _StreamGreeter_AccordingToConditionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccordingToConditionGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamGreeterServer).AccordingToConditionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamGreeter_AccordingToConditionGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamGreeterServer).AccordingToConditionGet(ctx, req.(*AccordingToConditionGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamGreeter_AccordingToConditionGets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccordingToConditionGetsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamGreeterServer).AccordingToConditionGets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamGreeter_AccordingToConditionGets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamGreeterServer).AccordingToConditionGets(ctx, req.(*AccordingToConditionGetsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamGreeter_InsertUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamGreeterServer).InsertUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamGreeter_InsertUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamGreeterServer).InsertUser(ctx, req.(*InsertUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamGreeter_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamGreeterServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamGreeter_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamGreeterServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamGreeter_DeletedUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletedUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamGreeterServer).DeletedUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamGreeter_DeletedUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamGreeterServer).DeletedUser(ctx, req.(*DeletedUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamGreeter_ServiceDesc is the grpc.ServiceDesc for StreamGreeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamGreeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.StreamGreeter",
	HandlerType: (*StreamGreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccordingToConditionGet",
			Handler:    _StreamGreeter_AccordingToConditionGet_Handler,
		},
		{
			MethodName: "AccordingToConditionGets",
			Handler:    _StreamGreeter_AccordingToConditionGets_Handler,
		},
		{
			MethodName: "InsertUser",
			Handler:    _StreamGreeter_InsertUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _StreamGreeter_UpdateUser_Handler,
		},
		{
			MethodName: "DeletedUser",
			Handler:    _StreamGreeter_DeletedUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}