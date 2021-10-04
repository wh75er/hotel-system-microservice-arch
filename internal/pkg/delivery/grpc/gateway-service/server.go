package gateway_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	proto1 "hotel-booking-system/internal/pkg/delivery/grpc/auth-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	"hotel-booking-system/internal/pkg/delivery/grpc/gateway-service/proto"
	proto2 "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	proto3 "hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service/proto"
	proto4 "hotel-booking-system/internal/pkg/delivery/grpc/payment-service/proto"
	proto5 "hotel-booking-system/internal/pkg/delivery/grpc/reservation-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type GatewayServer struct {
	proto.UnimplementedGatewayServiceServer
	AdminCredsUsecase        models.CredentialsUsecaseI
	UserServiceClient        proto1.AuthServiceClient
	HotelServiceClient       proto2.HotelServiceClient
	UserLoyaltyServiceClient proto3.LoyaltyServiceClient
	PaymentServiceClient     proto4.PaymentServiceClient
	ReservationServiceClient proto5.ReservationServiceClient
	Logger                   logs.LoggerInterface
}

func NewGatewayServer(
	aCredsU models.CredentialsUsecaseI,
	userClient proto1.AuthServiceClient,
	hotelClient proto2.HotelServiceClient,
	userLoyaltyClient proto3.LoyaltyServiceClient,
	paymentClient proto4.PaymentServiceClient,
	reservationClient proto5.ReservationServiceClient,
	logger logs.LoggerInterface,
) proto.GatewayServiceServer {
	return &GatewayServer{
		AdminCredsUsecase:        aCredsU,
		UserServiceClient:        userClient,
		HotelServiceClient:       hotelClient,
		UserLoyaltyServiceClient: userLoyaltyClient,
		PaymentServiceClient:     paymentClient,
		ReservationServiceClient: reservationClient,
		Logger:                   logger,
	}
}

func (s *GatewayServer) AddReservation(ctx context.Context, pr *proto5.Reservation) (*commonProto.UUID, error) {
	uuid, err := s.ReservationServiceClient.AddReservation(context.Background(), pr)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return uuid, nil
}

func (s *GatewayServer) CancelReservation(ctx context.Context, pu *commonProto.UUID) (*commonProto.Empty, error) {
	plug, err := s.ReservationServiceClient.CancelReservation(context.Background(), pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return plug, nil
}

func (s *GatewayServer) GetReservation(ctx context.Context, pu *commonProto.UUID) (*proto5.Reservation, error) {
	r, err := s.ReservationServiceClient.GetReservation(context.Background(), pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return r, nil
}

func (s *GatewayServer) GetReservationsByUser(ctx context.Context, pu *commonProto.UUID) (*proto5.Reservations, error) {
	r, err := s.ReservationServiceClient.GetReservationsByUser(context.Background(), pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return r, nil
}

func (s *GatewayServer) CreatePayment(ctx context.Context, pu *commonProto.UUID) (*commonProto.UUID, error) {
	uuid, err := s.ReservationServiceClient.CreatePayment(context.Background(), pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return uuid, nil
}

func (s *GatewayServer) AddUser(ctx context.Context, pUser *proto1.User) (*commonProto.Empty, error) {
	plug, err := s.UserServiceClient.AddUser(context.Background(), pUser)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return plug, nil
}

func (s *GatewayServer) Login(ctx context.Context, pUser *proto1.User) (*commonProto.Token, error) {
	token, err := s.UserServiceClient.Login(context.Background(), pUser)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return token, nil
}

func (s *GatewayServer) AddHotel(ctx context.Context, ph *proto2.Hotel) (*commonProto.Empty, error) {
	plug, err := s.HotelServiceClient.AddHotel(context.Background(), ph)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return plug, nil
}

func (s *GatewayServer) GetHotel(ctx context.Context, pu *commonProto.UUID) (*proto2.Hotel, error) {
	h, err := s.HotelServiceClient.GetHotel(context.Background(), pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return h, nil
}

func (s *GatewayServer) GetHotels(ctx context.Context, empty *commonProto.Empty) (*proto2.HotelsResponse, error) {
	h, err := s.HotelServiceClient.GetHotels(context.Background(), empty)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return h, nil
}

func (s *GatewayServer) PatchHotel(ctx context.Context, ph *proto2.Hotel) (*commonProto.Empty, error) {
	plug, err := s.HotelServiceClient.PatchHotel(context.Background(), ph)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return plug, nil
}

func (s *GatewayServer) DeleteHotel(ctx context.Context, pu *commonProto.UUID) (*commonProto.Empty, error) {
	plug, err := s.HotelServiceClient.DeleteHotel(context.Background(), pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return plug, nil
}

func (s *GatewayServer) AddRoom(ctx context.Context, pr *proto2.Room) (*commonProto.Empty, error) {
	plug, err := s.HotelServiceClient.AddRoom(context.Background(), pr)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return plug, nil
}

func (s *GatewayServer) GetRooms(ctx context.Context, pu *commonProto.UUID) (*proto2.RoomsResponse, error) {
	r, err := s.HotelServiceClient.GetRooms(context.Background(), pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return r, nil
}

func (s *GatewayServer) GetRoom(ctx context.Context, pu *commonProto.UUID) (*proto2.Room, error) {
	r, err := s.HotelServiceClient.GetRoom(context.Background(), pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return r, nil
}

func (s *GatewayServer) PatchRoom(ctx context.Context, pr *proto2.Room) (*commonProto.Empty, error) {
	plug, err := s.HotelServiceClient.PatchRoom(context.Background(), pr)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return plug, nil
}

func (s *GatewayServer) DeleteRoom(ctx context.Context, pu *commonProto.UUID) (*commonProto.Empty, error) {
	plug, err := s.HotelServiceClient.DeleteRoom(context.Background(), pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return plug, nil
}

func (s *GatewayServer) GetDiscount(ctx context.Context, pu *commonProto.UUID) (*proto3.Loyalty, error) {
	l, err := s.UserLoyaltyServiceClient.GetDiscount(context.Background(), pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return l, nil
}

func (s *GatewayServer) GetPayment(ctx context.Context, pu *commonProto.UUID) (*proto4.Payment, error) {
	p, err := s.PaymentServiceClient.GetPayment(context.Background(), pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return p, nil
}
