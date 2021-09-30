package auth_service

import (
	"github.com/google/uuid"
	"hotel-booking-system/internal/pkg/delivery/grpc/auth-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	"hotel-booking-system/internal/pkg/errors"
	kinds "hotel-booking-system/internal/pkg/errors/hotel-service"
	"hotel-booking-system/internal/pkg/models"
)

func (s *AuthServer) ProtoToUser(pu *proto.User) (*models.User, error) {
	var opError errors.Op = "auth-service.ProtoToUser"

	validUserUuid, err := uuid.Parse(pu.UserUuid.Value)
	if err != nil {
		e := errors.E(opError, kinds.UserUuidValidationErr, err)
		s.Logger.Error("Grpc error: ", e)
		return nil, e
	}

	return &models.User{
		UserUuid: validUserUuid,
		Login:    pu.Login,
		Password: pu.Password,
		Role:     pu.Role,
	}, nil
}

func (s *AuthServer) UserToProto(user *models.User) *proto.User {
	return &proto.User{
		UserUuid: &commonProto.UUID{Value: user.UserUuid.String()},
		Login:    user.Login,
		Password: user.Password,
		Role:     user.Role,
	}
}

func (s *AuthServer) RoleToProto(role *models.Role) *proto.Role {
	return &proto.Role{
		Value: string(*role),
	}
}
