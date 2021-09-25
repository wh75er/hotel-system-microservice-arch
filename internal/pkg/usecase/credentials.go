package usecase

import (
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/models"
)

type AdminCredentialsUsecase struct {
	Creds models.Credentials
}

func NewAdminCredentialsUsecase(credentials models.Credentials) models.CredentialsUsecaseI {
	return &AdminCredentialsUsecase{
		Creds: credentials,
	}
}

func (admin *AdminCredentialsUsecase) Login(creds *models.Credentials) error {
	var opError errors.Op = "usecase.Login"

	var err error

	if creds.Id != admin.Creds.Id {
		err = errors.E(opError, errors.InvalidCredentials)
		return err
	}

	if creds.Secret != admin.Creds.Secret {
		err = errors.E(opError, errors.InvalidCredentials)
		return err
	}

	return nil
}
