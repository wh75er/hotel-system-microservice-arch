package stat_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	"hotel-booking-system/internal/pkg/delivery/grpc/stat-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	jwt_manager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type StatServer struct {
	proto.UnimplementedStatServiceServer
	StatUsecase    models.StatUsecaseI
	AdminCredsUsecase models.CredentialsUsecaseI
	TokenManager      *jwt_manager.JWTManager
	Logger            logs.LoggerInterface
}

func NewStatServer(
	statU models.StatUsecaseI,
	aCredsU models.CredentialsUsecaseI,
	jwtManager *jwt_manager.JWTManager,
	logger logs.LoggerInterface,
) proto.StatServiceServer {
	return &StatServer{
		StatUsecase:    statU,
		AdminCredsUsecase: aCredsU,
		TokenManager:      jwtManager,
		Logger:            logger,
	}
}

func (s *StatServer) GetToken(ctx context.Context, pc *commonProto.Credentials) (*commonProto.Token, error) {
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

func (s *StatServer) GetStat(ctx context.Context, p *commonProto.Empty) (*proto.Stat, error) {
	stat, err := s.StatUsecase.GetStat()
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	pStat := &proto.Stat{}

	pStat.RoomsAmount = stat.RoomsAmount
	pStat.ReservationsAmount = stat.ReservationsAmount

	return pStat, nil
}

func (s *StatServer) UpdateRoomsAmount(ctx context.Context, pd *proto.Delta) (*commonProto.Empty, error) {
	err := s.StatUsecase.UpdateRoomsAmount(pd.GetValue())
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.Empty{}, nil
}

func (s *StatServer) UpdateReservationsAmount(ctx context.Context, pd *proto.Delta) (*commonProto.Empty, error) {
	err := s.StatUsecase.UpdateReservationsAmount(pd.GetValue())
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.Empty{}, nil
}
