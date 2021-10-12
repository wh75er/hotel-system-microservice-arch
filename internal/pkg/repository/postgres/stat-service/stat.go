package stat_service

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type StatRepository struct {
	Db     *sqlx.DB
	logger logs.LoggerInterface
}

func NewStatRepository(db *sqlx.DB, logger logs.LoggerInterface) models.StatRepositoryI {
	return &StatRepository{db, logger}
}

func (r *StatRepository) GetStat() (s models.Stat, e error) {
	var opError errors.Op = "postgres.GetStat"

	err := r.Db.Get(&s, "SELECT RoomsAmount, ReservationsAmount FROM stats LIMIT 1")
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

func (r *StatRepository) UpdateStat(s models.Stat) (e error) {
	var opError errors.Op = "postgres.UpdateStat"

	_, err := r.Db.Exec(
		"UPDATE stats SET RoomsAmount = $1, ReservationsAmount = $2",
		s.RoomsAmount, s.ReservationsAmount,
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
