package reservation_service

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

type ReservationRepository struct {
	Db     *sqlx.DB
	logger logs.LoggerInterface
}

func NewReservationRepository(db *sqlx.DB, logger logs.LoggerInterface) models.ReservationRepositoryI {
	return &ReservationRepository{db, logger}
}

func (r *ReservationRepository) CreateReservation(v *models.Reservation) (e error) {
	var opError errors.Op = "postgres.AddPayment"

	_, err := r.Db.Exec("INSERT INTO "+
		"reservations(ReservationUuid, RoomUuid, UserUuid, PaymentUuid, Date, Status) VALUES ($1, $2, $3, $4, $5, $6)",
		v.ReservationUuid, v.RoomUuid, v.UserUuid, v.PaymentUuid, v.Date, v.Status,
	)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	}

	return
}

func (r *ReservationRepository) PatchReservation(v *models.Reservation) (e error) {
	var opError errors.Op = "postgres.PatchReservation"

	_, err := r.Db.Exec(
		"UPDATE reservations SET status = $1, paymentUuid = $2 WHERE ReservationUuid = $3",
		v.Status, v.PaymentUuid, v.ReservationUuid,
	)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	}

	return
}

func (r *ReservationRepository) GetReservation(reservationUuid uuid.UUID) (v models.Reservation, e error) {
	var opError errors.Op = "postgres.GetReservation"

	err := r.Db.Get(&v, "SELECT ReservationUuid, RoomUuid, UserUuid, PaymentUuid, Status, Date FROM reservations WHERE reservationUuid = $1", reservationUuid)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	}

	return
}

func (r *ReservationRepository) GetReservationsByUser(userUuid uuid.UUID) (v []models.Reservation, e error) {
	var opError errors.Op = "postgres.GetReservationsByUser"

	err := r.Db.Select(&v, "SELECT ReservationUuid, RoomUuid, PaymentUuid, Status, Date FROM reservations WHERE userUuid = $1", userUuid)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	}

	return
}

func (r *ReservationRepository) CountReservedRoomsForDate(roomUuid uuid.UUID, date time.Time) (count int, e error) {
	var opError errors.Op = "postgres.CountReservedRoomsForDate"

	var reservations []models.Reservation

	err := r.Db.Select(&reservations, "SELECT RoomUuid FROM reservations WHERE RoomUuid=$1 AND Date=$2 AND Status='active'", roomUuid, date)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	}

	count = len(reservations)

	return
}
