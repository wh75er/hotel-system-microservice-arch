package loyalty_service

import (
	"github.com/google/uuid"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type LoyaltyUsecase struct {
	loyaltyRepository models.LoyaltyRepositoryI
	logger            logs.LoggerInterface
}

func NewLoyaltyUsecase(
	loyaltyR models.LoyaltyRepositoryI,
	logger logs.LoggerInterface,
) models.LoyaltyUsecaseI {
	return &LoyaltyUsecase{loyaltyR, logger}
}

func (u *LoyaltyUsecase) GetDiscount(userUid string) (l *models.Loyalty, e error) {
	var opError errors.Op = "auth-usecase.GetDiscount"

	validUserUuid, err := uuid.Parse(userUid)
	if err != nil {
		e = errors.E(opError, errors.LoyaltyUserUuidValidationErr, err)
		u.logger.Error("Usecase error: ", e)
		return
	}

	l = &models.Loyalty{}

	*l, err = u.loyaltyRepository.GetLoyalty(validUserUuid)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.LoyaltyNotFoundErr, err)
			u.logger.Error("Usecase error: %v", e)
			return
		}
		e = errors.E(opError, errors.RepositoryLoyaltyErr, err)
		u.logger.Error("Usecase error: %v", e)
		return
	}

	return
}

func (u *LoyaltyUsecase) AddUser(userUid string) (e error) {
	var opError errors.Op = "auth-usecase.AddUser"

	validUserUuid, err := uuid.Parse(userUid)
	if err != nil {
		e = errors.E(opError, errors.LoyaltyUserUuidValidationErr, err)
		u.logger.Error("Usecase error: ", e)
		return
	}

	found := true
	_, err = u.loyaltyRepository.GetLoyalty(validUserUuid)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			found = false
			err = nil
		} else {
			e = errors.E(opError, errors.RepositoryLoyaltyErr, err)
			u.logger.Error("Usecase error: %v", e)
			return
		}
	}
	if found {
		e = errors.E(opError, errors.LoyaltyExistsErr, err)
		u.logger.Error("Usecase error: %v", e)
		return
	}

	l := &models.Loyalty{
		UserUuid:           validUserUuid,
		Status:             models.Bronze,
		Discount:           0,
		ContributionAmount: 0,
	}

	err = u.loyaltyRepository.AddLoyalty(l)
	if e != nil {
		e = errors.E(opError, errors.RepositoryLoyaltyErr, e)
		u.logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *LoyaltyUsecase) UpdateDiscount(userUid string, contribution float32) (e error) {
	var opError errors.Op = "auth-usecase.UpdateDiscount"

	validUserUuid, err := uuid.Parse(userUid)
	if err != nil {
		e = errors.E(opError, errors.LoyaltyUserUuidValidationErr, err)
		u.logger.Error("Usecase error: ", e)
		return
	}

	l, err := u.loyaltyRepository.GetLoyalty(validUserUuid)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.LoyaltyNotFoundErr, err)
			u.logger.Error("Usecase error: %v", e)
			return
		}
		e = errors.E(opError, errors.RepositoryLoyaltyErr, err)
		u.logger.Error("Usecase error: %v", e)
		return
	}

	l.ContributionAmount += contribution
	l.UpdateStatus()

	err = u.loyaltyRepository.UpdateLoyalty(&l)
	if e != nil {
		e = errors.E(opError, errors.RepositoryLoyaltyErr, e)
		u.logger.Error("Usecase error: ", e)
		return
	}

	return
}
