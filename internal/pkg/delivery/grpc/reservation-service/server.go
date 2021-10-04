package reservation_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	"hotel-booking-system/internal/pkg/delivery/grpc/reservation-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	jwt_manager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type ReservationServer struct {
	proto.UnimplementedReservationServiceServer
	ReservationUsecase models.ReservationUsecaseI
	AdminCredsUsecase  models.CredentialsUsecaseI
	TokenManager       *jwt_manager.JWTManager
	Logger             logs.LoggerInterface
}

func NewReservationServer(
	reservationU models.ReservationUsecaseI,
	aCredsU models.CredentialsUsecaseI,
	jwtManager *jwt_manager.JWTManager,
	logger logs.LoggerInterface,
) proto.ReservationServiceServer {
	return &ReservationServer{
		ReservationUsecase: reservationU,
		AdminCredsUsecase:  aCredsU,
		TokenManager:       jwtManager,
		Logger:             logger,
	}
}

func (s *ReservationServer) GetToken(ctx context.Context, pc *commonProto.Credentials) (*commonProto.Token, error) {
	c := commonProto.ProtoToCredentials(pc)

	err := s.AdminCredsUsecase.Login(c)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	token, err := s.TokenManager.Generate(models.SERVICE)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	pt := commonProto.TokenToProto(&token)

	return pt, nil
}

func (s *ReservationServer) AddReservation(ctx context.Context, pr *proto.Reservation) (*commonProto.UUID, error) {
	r, err := ProtoToReservation(pr)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	reservationUuid, err := s.ReservationUsecase.AddReservation(r)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.UUID{
		Value: reservationUuid.String(),
	}, nil
}

func (s *ReservationServer) CancelReservation(ctx context.Context, pu *commonProto.UUID) (*commonProto.Empty, error) {
	err := s.ReservationUsecase.CancelReservation(pu.Value)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.Empty{}, nil
}

func (s *ReservationServer) GetReservation(ctx context.Context, pu *commonProto.UUID) (*proto.Reservation, error) {
	r, err := s.ReservationUsecase.GetReservation(pu.Value)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return ReservationToProto(r), nil
}

func (s *ReservationServer) GetReservationsByUser(ctx context.Context, pu *commonProto.UUID) (*proto.Reservations, error) {
	r, err := s.ReservationUsecase.GetReservationsByUser(pu.Value)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return ReservationsToProto(r), nil
}

func (s *ReservationServer) CreatePayment(ctx context.Context, pu *commonProto.UUID) (*commonProto.UUID, error) {
	paymentUuid, err := s.ReservationUsecase.CreatePayment(pu.Value)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.UUID{
		Value: paymentUuid.String(),
	}, nil
}
