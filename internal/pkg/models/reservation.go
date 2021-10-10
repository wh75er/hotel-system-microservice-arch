package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"hotel-booking-system/internal/pkg/errors"
	"time"
)

type ReservationStatus string

const (
	ActiveReservationStatus   = ReservationStatus("active")
	CanceledReservationStatus = ReservationStatus("canceled")
)

type Reservation struct {
	ReservationUuid uuid.UUID
	RoomUuid        uuid.UUID
	UserUuid        uuid.UUID
	PaymentUuid     uuid.UUID
	Date            time.Time
	Status          ReservationStatus
}

type ReservationRepositoryI interface {
	CreateReservation(v *Reservation) (tx *sqlx.Tx, e error)
	PatchReservation(v *Reservation) (tx *sqlx.Tx, e error)
	GetReservation(reservationUuid uuid.UUID) (v Reservation, e error)
	GetReservationsByUser(userUuid uuid.UUID) (v []Reservation, e error)
}

type ReservationUsecaseI interface {
	AddReservation(r *Reservation) (reservationUuid uuid.UUID, e error)
	CancelReservation(reservationUuid string) (e error)
	GetReservation(reservationUuid string) (r *Reservation, e error)
	GetReservationsByUser(userUuid string) (r []Reservation, e error)
	CreatePayment(reservationUuid string) (paymentUuid uuid.UUID, e error)
}

func (r Reservation) ValidateDate() error {
	if r.Date.Unix() == 0 {
		return errors.E(errors.ReservationDateValidationErr)
	}

	today := time.Now()
	if r.Date.Before(today) {
		return errors.E(errors.ReservationDateValidationLateErr)
	}

	monthAgo := today.AddDate(0, 1, 0)
	if r.Date.After(monthAgo) {
		return errors.E(errors.ReservationDateValidationFarErr)
	}

	return nil
}
