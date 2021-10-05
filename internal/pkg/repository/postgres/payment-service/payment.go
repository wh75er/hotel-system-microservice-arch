package payment_service

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type PaymentRepository struct {
	Db     *sqlx.DB
	logger logs.LoggerInterface
}

func NewPaymentRepository(db *sqlx.DB, logger logs.LoggerInterface) models.PaymentRepositoryI {
	return &PaymentRepository{db, logger}
}

func (r *PaymentRepository) AddPayment(p *models.Payment) (e error) {
	var opError errors.Op = "postgres.AddPayment"

	_, err := r.Db.Exec("INSERT INTO "+
		"payments(userUuid, paymentUuid, status, price, timeUpdated) VALUES ($1, $2, $3, $4)",
		p.UserUuid, p.PaymentUuid, p.Status, p.Price, p.TimeUpdated)
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

func (r *PaymentRepository) ChangePaymentStatus(p *models.Payment) (e error) {
	var opError errors.Op = "postgres.ChangePaymentStatus"

	_, err := r.Db.Exec(
		"UPDATE payments SET status = $1, timeUpdated = $2 WHERE paymentUuid = $3",
		p.Status, p.TimeUpdated, p.PaymentUuid,
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

func (r *PaymentRepository) GetPayment(paymentUuid uuid.UUID) (p models.Payment, e error) {
	var opError errors.Op = "postgres.GetPayment"

	err := r.Db.Get(&p, "SELECT userUuid, paymentUuid, status, price, dateUpdated FROM payments WHERE paymentUuid = $1", paymentUuid)
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
