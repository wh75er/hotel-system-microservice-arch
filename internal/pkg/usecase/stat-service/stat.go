package stat_service

import (
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type StatUsecase struct {
	StatRepository        models.StatRepositoryI
	Logger                   logs.LoggerInterface
}

func NewStatUsecase(
	statR models.StatRepositoryI,
	logger logs.LoggerInterface,
) models.StatUsecaseI {
	return &StatUsecase{
		statR,
		logger,
	}
}

func (u *StatUsecase) GetStat() (s models.Stat, e error) {
	var opError errors.Op = "usecase.GetStat"

	s, err := u.StatRepository.GetStat()
	if err != nil {
		e = errors.E(opError, errors.RepositoryStatErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *StatUsecase) UpdateRoomsAmount(delta int64) (e error) {
	var opError errors.Op = "usecase.UpdateRoomsAmount"

	s, err := u.StatRepository.GetStat()
	if err != nil {
		e = errors.E(opError, errors.RepositoryStatErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	s.RoomsAmount += delta

	err = u.StatRepository.UpdateStat(s)
	if err != nil {
		e = errors.E(opError, errors.RepositoryStatErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *StatUsecase) UpdateReservationsAmount(delta int64) (e error) {
	var opError errors.Op = "usecase.UpdateReservationsAmount"

	s, err := u.StatRepository.GetStat()
	if err != nil {
		e = errors.E(opError, errors.RepositoryStatErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	s.ReservationsAmount += delta

	err = u.StatRepository.UpdateStat(s)
	if err != nil {
		e = errors.E(opError, errors.RepositoryStatErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}
