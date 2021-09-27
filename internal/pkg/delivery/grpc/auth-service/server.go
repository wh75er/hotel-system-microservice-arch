package auth_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hotel-booking-system/internal/pkg/delivery/grpc/auth-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	kinds "hotel-booking-system/internal/pkg/errors/hotel-service"
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

func (s *AuthServer) GetToken(ctx context.Context, pc *proto.Credentials) (*proto.Token, error) {
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

func (s *AuthServer) AddUser(ctx context.Context, pu *proto.User) (*proto.Empty, error) {
	user, err := s.ProtoToUser(pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	err = s.UserUsecase.AddUser(user)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return &proto.Empty{}, nil
}

func (s *AuthServer) GetUser(ctx context.Context, pUuid *proto.UUID) (*proto.User, error) {
	user, err := s.UserUsecase.GetUser(pUuid.Value)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return s.UserToProto(user), nil
}

func (s *AuthServer) Login(ctx context.Context, pu *proto.User) (*proto.Token, error) {
	user, err := s.ProtoToUser(pu)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	token, err := s.UserUsecase.Login(user)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	tokenType := models.Token(token)

	return TokenToProto(&tokenType), nil
}

func (s *AuthServer) CheckAuth(ctx context.Context, pt *proto.Token) (*proto.Role, error) {
	token := ProtoToToken(pt)

	role, err := s.UserUsecase.CheckAuth(string(*token))
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return s.RoleToProto(&role), nil
}
