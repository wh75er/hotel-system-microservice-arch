// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	commonProto "hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StatServiceClient is the client API for StatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatServiceClient interface {
	GetToken(ctx context.Context, in *commonProto.Credentials, opts ...grpc.CallOption) (*commonProto.Token, error)
	GetStat(ctx context.Context, in *commonProto.Empty, opts ...grpc.CallOption) (*Stat, error)
	UpdateRoomsAmount(ctx context.Context, in *Delta, opts ...grpc.CallOption) (*commonProto.Empty, error)
	UpdateReservationsAmount(ctx context.Context, in *Delta, opts ...grpc.CallOption) (*commonProto.Empty, error)
}

type statServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatServiceClient(cc grpc.ClientConnInterface) StatServiceClient {
	return &statServiceClient{cc}
}

func (c *statServiceClient) GetToken(ctx context.Context, in *commonProto.Credentials, opts ...grpc.CallOption) (*commonProto.Token, error) {
	out := new(commonProto.Token)
	err := c.cc.Invoke(ctx, "/proto.StatService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statServiceClient) GetStat(ctx context.Context, in *commonProto.Empty, opts ...grpc.CallOption) (*Stat, error) {
	out := new(Stat)
	err := c.cc.Invoke(ctx, "/proto.StatService/GetStat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statServiceClient) UpdateRoomsAmount(ctx context.Context, in *Delta, opts ...grpc.CallOption) (*commonProto.Empty, error) {
	out := new(commonProto.Empty)
	err := c.cc.Invoke(ctx, "/proto.StatService/UpdateRoomsAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statServiceClient) UpdateReservationsAmount(ctx context.Context, in *Delta, opts ...grpc.CallOption) (*commonProto.Empty, error) {
	out := new(commonProto.Empty)
	err := c.cc.Invoke(ctx, "/proto.StatService/UpdateReservationsAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatServiceServer is the server API for StatService service.
// All implementations must embed UnimplementedStatServiceServer
// for forward compatibility
type StatServiceServer interface {
	GetToken(context.Context, *commonProto.Credentials) (*commonProto.Token, error)
	GetStat(context.Context, *commonProto.Empty) (*Stat, error)
	UpdateRoomsAmount(context.Context, *Delta) (*commonProto.Empty, error)
	UpdateReservationsAmount(context.Context, *Delta) (*commonProto.Empty, error)
	mustEmbedUnimplementedStatServiceServer()
}

// UnimplementedStatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatServiceServer struct {
}

func (UnimplementedStatServiceServer) GetToken(context.Context, *commonProto.Credentials) (*commonProto.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedStatServiceServer) GetStat(context.Context, *commonProto.Empty) (*Stat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStat not implemented")
}
func (UnimplementedStatServiceServer) UpdateRoomsAmount(context.Context, *Delta) (*commonProto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoomsAmount not implemented")
}
func (UnimplementedStatServiceServer) UpdateReservationsAmount(context.Context, *Delta) (*commonProto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReservationsAmount not implemented")
}
func (UnimplementedStatServiceServer) mustEmbedUnimplementedStatServiceServer() {}

// UnsafeStatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatServiceServer will
// result in compilation errors.
type UnsafeStatServiceServer interface {
	mustEmbedUnimplementedStatServiceServer()
}

func RegisterStatServiceServer(s grpc.ServiceRegistrar, srv StatServiceServer) {
	s.RegisterService(&StatService_ServiceDesc, srv)
}

func _StatService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StatService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatServiceServer).GetToken(ctx, req.(*commonProto.Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatService_GetStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatServiceServer).GetStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StatService/GetStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatServiceServer).GetStat(ctx, req.(*commonProto.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatService_UpdateRoomsAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Delta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatServiceServer).UpdateRoomsAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StatService/UpdateRoomsAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatServiceServer).UpdateRoomsAmount(ctx, req.(*Delta))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatService_UpdateReservationsAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Delta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatServiceServer).UpdateReservationsAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StatService/UpdateReservationsAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatServiceServer).UpdateReservationsAmount(ctx, req.(*Delta))
	}
	return interceptor(ctx, in, info, handler)
}

// StatService_ServiceDesc is the grpc.ServiceDesc for StatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StatService",
	HandlerType: (*StatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _StatService_GetToken_Handler,
		},
		{
			MethodName: "GetStat",
			Handler:    _StatService_GetStat_Handler,
		},
		{
			MethodName: "UpdateRoomsAmount",
			Handler:    _StatService_UpdateRoomsAmount_Handler,
		},
		{
			MethodName: "UpdateReservationsAmount",
			Handler:    _StatService_UpdateReservationsAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/pkg/delivery/grpc/stat-service/proto/scheme.proto",
}