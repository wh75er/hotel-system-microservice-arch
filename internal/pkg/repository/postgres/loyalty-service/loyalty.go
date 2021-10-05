package loyalty_service

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type LoyaltyRepository struct {
	Db     *sqlx.DB
	logger logs.LoggerInterface
}

func NewLoyaltyRepository(db *sqlx.DB, logger logs.LoggerInterface) models.LoyaltyRepositoryI {
	return &LoyaltyRepository{db, logger}
}

func (r *LoyaltyRepository) GetLoyalty(userUid uuid.UUID) (l models.Loyalty, e error) {
	var opError errors.Op = "postgres.GetLoyalty"

	err := r.Db.Get(&l, "SELECT userUuid, status, discount, contributionAmount FROM loyalty WHERE userUuid = $1", userUid)
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

func (r *LoyaltyRepository) AddLoyalty(l *models.Loyalty) (e error) {
	var opError errors.Op = "postgres.AddLoyalty"

	_, err := r.Db.Exec("INSERT INTO "+
		"loyalty(userUuid, status, discount, contributionAmount) VALUES ($1, $2, $3, $4)",
		l.UserUuid, l.Status, l.Discount, l.ContributionAmount)
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

func (r *LoyaltyRepository) UpdateLoyalty(l *models.Loyalty) (e error) {
	var opError errors.Op = "postgres.UpdateLoyalty"

	_, err := r.Db.Exec(
		"UPDATE loyalty SET status = $1, discount = $2, contributionAmount = $3 WHERE userUuid = $8",
		l.Status, l.Discount, l.ContributionAmount,
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
