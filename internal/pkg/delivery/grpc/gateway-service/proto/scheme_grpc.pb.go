// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto1 "hotel-booking-system/internal/pkg/delivery/grpc/auth-service/proto"
	commonProto "hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	proto2 "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	proto3 "hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service/proto"
	proto4 "hotel-booking-system/internal/pkg/delivery/grpc/payment-service/proto"
	proto "hotel-booking-system/internal/pkg/delivery/grpc/reservation-service/proto"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayServiceClient interface {
	AddReservation(ctx context.Context, in *proto.Reservation, opts ...grpc.CallOption) (*commonProto.UUID, error)
	CancelReservation(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*commonProto.Empty, error)
	GetReservation(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto.Reservation, error)
	GetReservationsByUser(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto.Reservations, error)
	CreatePayment(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*commonProto.UUID, error)
	AddUser(ctx context.Context, in *proto1.User, opts ...grpc.CallOption) (*commonProto.Empty, error)
	Login(ctx context.Context, in *proto1.User, opts ...grpc.CallOption) (*commonProto.Token, error)
	CheckAuth(ctx context.Context, in *commonProto.Token, opts ...grpc.CallOption) (*proto1.Role, error)
	AddHotel(ctx context.Context, in *proto2.Hotel, opts ...grpc.CallOption) (*commonProto.Empty, error)
	GetHotel(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto2.Hotel, error)
	GetHotels(ctx context.Context, in *commonProto.Empty, opts ...grpc.CallOption) (*proto2.HotelsResponse, error)
	PatchHotel(ctx context.Context, in *proto2.Hotel, opts ...grpc.CallOption) (*commonProto.Empty, error)
	DeleteHotel(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*commonProto.Empty, error)
	AddRoom(ctx context.Context, in *proto2.Room, opts ...grpc.CallOption) (*commonProto.Empty, error)
	GetRooms(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto2.RoomsResponse, error)
	GetRoom(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto2.Room, error)
	PatchRoom(ctx context.Context, in *proto2.Room, opts ...grpc.CallOption) (*commonProto.Empty, error)
	DeleteRoom(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*commonProto.Empty, error)
	GetDiscount(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto3.Loyalty, error)
	GetPayment(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto4.Payment, error)
}

type gatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayServiceClient(cc grpc.ClientConnInterface) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) AddReservation(ctx context.Context, in *proto.Reservation, opts ...grpc.CallOption) (*commonProto.UUID, error) {
	out := new(commonProto.UUID)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/AddReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) CancelReservation(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*commonProto.Empty, error) {
	out := new(commonProto.Empty)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/CancelReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetReservation(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto.Reservation, error) {
	out := new(proto.Reservation)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/GetReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetReservationsByUser(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto.Reservations, error) {
	out := new(proto.Reservations)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/GetReservationsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) CreatePayment(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*commonProto.UUID, error) {
	out := new(commonProto.UUID)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/CreatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) AddUser(ctx context.Context, in *proto1.User, opts ...grpc.CallOption) (*commonProto.Empty, error) {
	out := new(commonProto.Empty)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) Login(ctx context.Context, in *proto1.User, opts ...grpc.CallOption) (*commonProto.Token, error) {
	out := new(commonProto.Token)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) CheckAuth(ctx context.Context, in *commonProto.Token, opts ...grpc.CallOption) (*proto1.Role, error) {
	out := new(proto1.Role)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/CheckAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) AddHotel(ctx context.Context, in *proto2.Hotel, opts ...grpc.CallOption) (*commonProto.Empty, error) {
	out := new(commonProto.Empty)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/AddHotel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetHotel(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto2.Hotel, error) {
	out := new(proto2.Hotel)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/GetHotel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetHotels(ctx context.Context, in *commonProto.Empty, opts ...grpc.CallOption) (*proto2.HotelsResponse, error) {
	out := new(proto2.HotelsResponse)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/GetHotels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) PatchHotel(ctx context.Context, in *proto2.Hotel, opts ...grpc.CallOption) (*commonProto.Empty, error) {
	out := new(commonProto.Empty)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/PatchHotel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) DeleteHotel(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*commonProto.Empty, error) {
	out := new(commonProto.Empty)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/DeleteHotel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) AddRoom(ctx context.Context, in *proto2.Room, opts ...grpc.CallOption) (*commonProto.Empty, error) {
	out := new(commonProto.Empty)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/AddRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetRooms(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto2.RoomsResponse, error) {
	out := new(proto2.RoomsResponse)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/GetRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetRoom(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto2.Room, error) {
	out := new(proto2.Room)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/GetRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) PatchRoom(ctx context.Context, in *proto2.Room, opts ...grpc.CallOption) (*commonProto.Empty, error) {
	out := new(commonProto.Empty)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/PatchRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) DeleteRoom(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*commonProto.Empty, error) {
	out := new(commonProto.Empty)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/DeleteRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetDiscount(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto3.Loyalty, error) {
	out := new(proto3.Loyalty)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/GetDiscount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetPayment(ctx context.Context, in *commonProto.UUID, opts ...grpc.CallOption) (*proto4.Payment, error) {
	out := new(proto4.Payment)
	err := c.cc.Invoke(ctx, "/proto.GatewayService/GetPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
// All implementations must embed UnimplementedGatewayServiceServer
// for forward compatibility
type GatewayServiceServer interface {
	AddReservation(context.Context, *proto.Reservation) (*commonProto.UUID, error)
	CancelReservation(context.Context, *commonProto.UUID) (*commonProto.Empty, error)
	GetReservation(context.Context, *commonProto.UUID) (*proto.Reservation, error)
	GetReservationsByUser(context.Context, *commonProto.UUID) (*proto.Reservations, error)
	CreatePayment(context.Context, *commonProto.UUID) (*commonProto.UUID, error)
	AddUser(context.Context, *proto1.User) (*commonProto.Empty, error)
	Login(context.Context, *proto1.User) (*commonProto.Token, error)
	CheckAuth(context.Context, *commonProto.Token) (*proto1.Role, error)
	AddHotel(context.Context, *proto2.Hotel) (*commonProto.Empty, error)
	GetHotel(context.Context, *commonProto.UUID) (*proto2.Hotel, error)
	GetHotels(context.Context, *commonProto.Empty) (*proto2.HotelsResponse, error)
	PatchHotel(context.Context, *proto2.Hotel) (*commonProto.Empty, error)
	DeleteHotel(context.Context, *commonProto.UUID) (*commonProto.Empty, error)
	AddRoom(context.Context, *proto2.Room) (*commonProto.Empty, error)
	GetRooms(context.Context, *commonProto.UUID) (*proto2.RoomsResponse, error)
	GetRoom(context.Context, *commonProto.UUID) (*proto2.Room, error)
	PatchRoom(context.Context, *proto2.Room) (*commonProto.Empty, error)
	DeleteRoom(context.Context, *commonProto.UUID) (*commonProto.Empty, error)
	GetDiscount(context.Context, *commonProto.UUID) (*proto3.Loyalty, error)
	GetPayment(context.Context, *commonProto.UUID) (*proto4.Payment, error)
	mustEmbedUnimplementedGatewayServiceServer()
}

// UnimplementedGatewayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (UnimplementedGatewayServiceServer) AddReservation(context.Context, *proto.Reservation) (*commonProto.UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReservation not implemented")
}
func (UnimplementedGatewayServiceServer) CancelReservation(context.Context, *commonProto.UUID) (*commonProto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelReservation not implemented")
}
func (UnimplementedGatewayServiceServer) GetReservation(context.Context, *commonProto.UUID) (*proto.Reservation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservation not implemented")
}
func (UnimplementedGatewayServiceServer) GetReservationsByUser(context.Context, *commonProto.UUID) (*proto.Reservations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationsByUser not implemented")
}
func (UnimplementedGatewayServiceServer) CreatePayment(context.Context, *commonProto.UUID) (*commonProto.UUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedGatewayServiceServer) AddUser(context.Context, *proto1.User) (*commonProto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedGatewayServiceServer) Login(context.Context, *proto1.User) (*commonProto.Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGatewayServiceServer) CheckAuth(context.Context, *commonProto.Token) (*proto1.Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuth not implemented")
}
func (UnimplementedGatewayServiceServer) AddHotel(context.Context, *proto2.Hotel) (*commonProto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHotel not implemented")
}
func (UnimplementedGatewayServiceServer) GetHotel(context.Context, *commonProto.UUID) (*proto2.Hotel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotel not implemented")
}
func (UnimplementedGatewayServiceServer) GetHotels(context.Context, *commonProto.Empty) (*proto2.HotelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHotels not implemented")
}
func (UnimplementedGatewayServiceServer) PatchHotel(context.Context, *proto2.Hotel) (*commonProto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchHotel not implemented")
}
func (UnimplementedGatewayServiceServer) DeleteHotel(context.Context, *commonProto.UUID) (*commonProto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHotel not implemented")
}
func (UnimplementedGatewayServiceServer) AddRoom(context.Context, *proto2.Room) (*commonProto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoom not implemented")
}
func (UnimplementedGatewayServiceServer) GetRooms(context.Context, *commonProto.UUID) (*proto2.RoomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRooms not implemented")
}
func (UnimplementedGatewayServiceServer) GetRoom(context.Context, *commonProto.UUID) (*proto2.Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoom not implemented")
}
func (UnimplementedGatewayServiceServer) PatchRoom(context.Context, *proto2.Room) (*commonProto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchRoom not implemented")
}
func (UnimplementedGatewayServiceServer) DeleteRoom(context.Context, *commonProto.UUID) (*commonProto.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
func (UnimplementedGatewayServiceServer) GetDiscount(context.Context, *commonProto.UUID) (*proto3.Loyalty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscount not implemented")
}
func (UnimplementedGatewayServiceServer) GetPayment(context.Context, *commonProto.UUID) (*proto4.Payment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayment not implemented")
}
func (UnimplementedGatewayServiceServer) mustEmbedUnimplementedGatewayServiceServer() {}

// UnsafeGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServiceServer will
// result in compilation errors.
type UnsafeGatewayServiceServer interface {
	mustEmbedUnimplementedGatewayServiceServer()
}

func RegisterGatewayServiceServer(s grpc.ServiceRegistrar, srv GatewayServiceServer) {
	s.RegisterService(&GatewayService_ServiceDesc, srv)
}

func _GatewayService_AddReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Reservation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).AddReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/AddReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).AddReservation(ctx, req.(*proto.Reservation))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_CancelReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).CancelReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/CancelReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).CancelReservation(ctx, req.(*commonProto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/GetReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetReservation(ctx, req.(*commonProto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetReservationsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetReservationsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/GetReservationsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetReservationsByUser(ctx, req.(*commonProto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/CreatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).CreatePayment(ctx, req.(*commonProto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto1.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).AddUser(ctx, req.(*proto1.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto1.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Login(ctx, req.(*proto1.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/CheckAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).CheckAuth(ctx, req.(*commonProto.Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_AddHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto2.Hotel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).AddHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/AddHotel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).AddHotel(ctx, req.(*proto2.Hotel))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/GetHotel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetHotel(ctx, req.(*commonProto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetHotels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetHotels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/GetHotels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetHotels(ctx, req.(*commonProto.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_PatchHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto2.Hotel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).PatchHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/PatchHotel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).PatchHotel(ctx, req.(*proto2.Hotel))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_DeleteHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).DeleteHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/DeleteHotel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).DeleteHotel(ctx, req.(*commonProto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_AddRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto2.Room)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).AddRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/AddRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).AddRoom(ctx, req.(*proto2.Room))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/GetRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetRooms(ctx, req.(*commonProto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/GetRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetRoom(ctx, req.(*commonProto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_PatchRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto2.Room)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).PatchRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/PatchRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).PatchRoom(ctx, req.(*proto2.Room))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/DeleteRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).DeleteRoom(ctx, req.(*commonProto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/GetDiscount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetDiscount(ctx, req.(*commonProto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonProto.UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GatewayService/GetPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetPayment(ctx, req.(*commonProto.UUID))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayService_ServiceDesc is the grpc.ServiceDesc for GatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddReservation",
			Handler:    _GatewayService_AddReservation_Handler,
		},
		{
			MethodName: "CancelReservation",
			Handler:    _GatewayService_CancelReservation_Handler,
		},
		{
			MethodName: "GetReservation",
			Handler:    _GatewayService_GetReservation_Handler,
		},
		{
			MethodName: "GetReservationsByUser",
			Handler:    _GatewayService_GetReservationsByUser_Handler,
		},
		{
			MethodName: "CreatePayment",
			Handler:    _GatewayService_CreatePayment_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _GatewayService_AddUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _GatewayService_Login_Handler,
		},
		{
			MethodName: "CheckAuth",
			Handler:    _GatewayService_CheckAuth_Handler,
		},
		{
			MethodName: "AddHotel",
			Handler:    _GatewayService_AddHotel_Handler,
		},
		{
			MethodName: "GetHotel",
			Handler:    _GatewayService_GetHotel_Handler,
		},
		{
			MethodName: "GetHotels",
			Handler:    _GatewayService_GetHotels_Handler,
		},
		{
			MethodName: "PatchHotel",
			Handler:    _GatewayService_PatchHotel_Handler,
		},
		{
			MethodName: "DeleteHotel",
			Handler:    _GatewayService_DeleteHotel_Handler,
		},
		{
			MethodName: "AddRoom",
			Handler:    _GatewayService_AddRoom_Handler,
		},
		{
			MethodName: "GetRooms",
			Handler:    _GatewayService_GetRooms_Handler,
		},
		{
			MethodName: "GetRoom",
			Handler:    _GatewayService_GetRoom_Handler,
		},
		{
			MethodName: "PatchRoom",
			Handler:    _GatewayService_PatchRoom_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _GatewayService_DeleteRoom_Handler,
		},
		{
			MethodName: "GetDiscount",
			Handler:    _GatewayService_GetDiscount_Handler,
		},
		{
			MethodName: "GetPayment",
			Handler:    _GatewayService_GetPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/pkg/delivery/grpc/gateway-service/proto/scheme.proto",
}
