package loyalty_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	kinds "hotel-booking-system/internal/pkg/errors/loyalty-service"
	jwt_manager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type LoyaltyServer struct {
	proto.UnimplementedLoyaltyServiceServer
	LoyaltyUsecase    models.LoyaltyUsecaseI
	AdminCredsUsecase models.CredentialsUsecaseI
	TokenManager      *jwt_manager.JWTManager
	Logger            logs.LoggerInterface
}

func NewLoyaltyServer(
	loyaltyU models.LoyaltyUsecaseI,
	aCredsU models.CredentialsUsecaseI,
	jwtManager *jwt_manager.JWTManager,
	logger logs.LoggerInterface,
) proto.LoyaltyServiceServer {
	return &LoyaltyServer{
		LoyaltyUsecase:    loyaltyU,
		AdminCredsUsecase: aCredsU,
		TokenManager:      jwtManager,
		Logger:            logger,
	}
}

func (s *LoyaltyServer) GetToken(ctx context.Context, pc *proto.Credentials) (*proto.Token, error) {
	c := ProtoToCredentials(pc)

	err := s.AdminCredsUsecase.Login(c)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	token, err := s.TokenManager.Generate(models.SERVICE)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	pt := TokenToProto(&token)

	return pt, nil
}

func (s *LoyaltyServer) GetDiscount(ctx context.Context, pu *proto.UUID) (*proto.Loyalty, error) {
	l, err := s.LoyaltyUsecase.GetDiscount(pu.Value)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return s.LoyaltyToProto(l), nil
}

func (s *LoyaltyServer) AddUser(ctx context.Context, pu *proto.UUID) (*proto.Empty, error) {
	err := s.LoyaltyUsecase.AddUser(pu.Value)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (s *LoyaltyServer) UpdateDiscount(ctx context.Context, pr *proto.UpdateDiscountRequest) (*proto.Empty, error) {
	err := s.LoyaltyUsecase.UpdateDiscount(pr.UserUid.Value, int(pr.Contribution))
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return &proto.Empty{}, nil
}
