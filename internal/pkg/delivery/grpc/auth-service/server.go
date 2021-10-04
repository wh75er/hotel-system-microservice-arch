package auth_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hotel-booking-system/internal/pkg/delivery/grpc/auth-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	"hotel-booking-system/internal/pkg/errors"
	jwt_manager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
	UserUsecase       models.UserUsecaseI
	AdminCredsUsecase models.CredentialsUsecaseI
	TokenManager      *jwt_manager.JWTManager
	Logger            logs.LoggerInterface
}

func NewAuthServer(
	userU models.UserUsecaseI,
	aCredsU models.CredentialsUsecaseI,
	jwtManager *jwt_manager.JWTManager,
	logger logs.LoggerInterface,
) proto.AuthServiceServer {
	return &AuthServer{
		UserUsecase:       userU,
		AdminCredsUsecase: aCredsU,
		TokenManager:      jwtManager,
		Logger:            logger,
	}
}

func (s *AuthServer) GetToken(ctx context.Context, pc *commonProto.Credentials) (*commonProto.Token, error) {
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

func (s *AuthServer) AddUser(ctx context.Context, pu *proto.User) (*commonProto.Empty, error) {
	user, err := ProtoToUser(pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	err = s.UserUsecase.AddUser(user)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return &commonProto.Empty{}, nil
}

func (s *AuthServer) GetUser(ctx context.Context, pUuid *commonProto.UUID) (*proto.User, error) {
	user, err := s.UserUsecase.GetUser(pUuid.Value)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return UserToProto(user), nil
}

func (s *AuthServer) Login(ctx context.Context, pu *proto.User) (*commonProto.Token, error) {
	user, err := ProtoToUser(pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	token, err := s.UserUsecase.Login(user)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	tokenType := models.Token(token)

	return commonProto.TokenToProto(&tokenType), nil
}

func (s *AuthServer) CheckAuth(ctx context.Context, pt *commonProto.Token) (*proto.Role, error) {
	token := commonProto.ProtoToToken(pt)

	role, err := s.UserUsecase.CheckAuth(string(*token))
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	return RoleToProto(&role), nil
}
