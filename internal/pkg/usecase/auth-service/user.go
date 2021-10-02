package auth_service

import (
	"github.com/google/uuid"
	"hotel-booking-system/internal/pkg/errors"
	jwt_manager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type UserUsecase struct {
	UserRepository models.UserRepositoryI
	jwtManager     *jwt_manager.JWTManager
	Logger         logs.LoggerInterface
}

func NewUserUsecase(
	userR models.UserRepositoryI,
	jwtManager *jwt_manager.JWTManager,
	logger logs.LoggerInterface,
) models.UserUsecaseI {
	return &UserUsecase{userR, jwtManager, logger}
}

func (u *UserUsecase) GetUser(uid string) (user *models.User, e error) {
	var opError errors.Op = "usecase.GetUser"

	validUserUuid, err := uuid.Parse(uid)
	if err != nil {
		e = errors.E(opError, errors.UserUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	user, err = u.UserRepository.GetUserByUuid(validUserUuid)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.UserNotFoundErr, err)
			u.Logger.Error("Usecase error: %v", e)
			return
		}
		e = errors.E(opError, errors.RepositoryUserErr, err)
		u.Logger.Error("Usecase error: %v", e)
		return
	}

	return
}

func (u *UserUsecase) AddUser(user *models.User) (e error) {
	var opError errors.Op = "usecase.AddUser"

	err := user.ValidateLogin()
	if err != nil {
		e = err
		return
	}

	err = user.ValidatePassword()
	if err != nil {
		e = err
		return
	}

	foundUser, err := u.UserRepository.GetUserByLogin(user.Login)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = nil
			return
		}
		e = errors.E(opError, errors.RepositoryUserErr, err)
		u.Logger.Error("Usecase error: %v", e)
		return
	}

	if foundUser != nil {
		e = errors.E(opError, errors.UserExistsErr, err)
		u.Logger.Error("Usecase error: %v", e)
		return
	}

	user.UserUuid = uuid.New()
	user.Role = string(models.USER)
	err = user.HashPassword()
	if err != nil {
		e = err
		u.Logger.Error("Usecase error: %v", e)
		return
	}

	e = u.UserRepository.AddUser(user)
	if e != nil {
		e = errors.E(opError, errors.RepositoryUserErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *UserUsecase) Login(user *models.User) (authToken string, e error) {
	var opError errors.Op = "usecase.Login"

	err := user.ValidateLogin()
	if err != nil {
		e = err
		return
	}

	err = user.ValidatePassword()
	if err != nil {
		e = err
		return
	}

	foundUser, err := u.UserRepository.GetUserByLogin(user.Login)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.AuthorizationErr, err)
			u.Logger.Error("Usecase error: %v", e)
			return
		}
		e = errors.E(opError, errors.RepositoryUserErr, err)
		u.Logger.Error("Usecase error: %v", e)
		return
	}

	err = foundUser.CompareWithPassword(user.Password)
	if err != nil {
		e = errors.E(opError, errors.AuthorizationErr, err)
		u.Logger.Error("Usecase error: %v", e)
		return
	}

	token, err := u.jwtManager.Generate(foundUser.UserUuid, foundUser.Login, models.Role(foundUser.Role))
	if err != nil {
		e = err
		u.Logger.Error("Usecase error: %v", e)
		return
	}

	authToken = string(token)

	return
}

func (u *UserUsecase) CheckAuth(jwtToken string) (role models.Role, e error) {
	claims, err := u.jwtManager.Verify(jwtToken)
	if err != nil {
		e = err
		u.Logger.Error("Usecase error: %v", e)
		return
	}

	role = claims.Role

	return
}
