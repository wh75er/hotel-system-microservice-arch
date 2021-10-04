package models

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type ReservationStatus string

const (
	ActiveReservationStatus = ReservationStatus("active")
	CanceledReservationStatus = ReservationStatus("canceled")
)

type Reservation struct {
	ReservationUuid uuid.UUID
	RoomUuid uuid.UUID
	UserUuid uuid.UUID
	PaymentUuid uuid.UUID
	Date time.Time
	Status ReservationStatus
}

type ReservationRepositoryI interface {
	CreateReservation(v *Reservation) (tx *sqlx.Tx, e error)
	PatchReservation(v *Reservation) (tx *sqlx.Tx,e error)
	GetReservation(reservationUuid uuid.UUID) (v *Reservation, e error)
	GetReservationsByUser(userUuid uuid.UUID) (v []Reservation, e error)
}

type ReservationUsecaseI interface {
	AddReservation(r *Reservation) (reservationUuid uuid.UUID, e error)
	CancelReservation(reservationUuid string) (e error)
	GetReservation(reservationUuid string) (r *Reservation, e error)
	GetReservationsByUser(userUuid string) (r []Reservation, e error)
	CreatePayment(reservationUuid string) (paymentUuid uuid.UUID, e error)
}
