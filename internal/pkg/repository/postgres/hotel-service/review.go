package hotel_service

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type ReviewRepository struct {
	Db     *sqlx.DB
	logger logs.LoggerInterface
}

func NewReviewRepository(db *sqlx.DB, logger logs.LoggerInterface) models.ReviewRepositoryI {
	return &ReviewRepository{db, logger}
}

func (r *ReviewRepository) GetReview(reviewUuid uuid.UUID) (rev models.Review, e error) {
	var opError errors.Op = "postgres.GetReview"

	err := r.Db.QueryRowx("SELECT userUuid, hotelUuid, reviewUuid, text, isAnonymous, photos "+
		"FROM reviews WHERE reviewUuid = $1", reviewUuid).Scan(
		&rev.UserUuid,
		&rev.HotelUuid,
		&rev.ReviewUuid,
		&rev.Text,
		&rev.IsAnonymous,
		pq.Array(&rev.Photos),
	)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	} else if err == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	}

	return
}

func (r *ReviewRepository) GetReviews(hotelUuid uuid.UUID) (rev []models.Review, e error) {
	var opError errors.Op = "postgres.GetReviews"

	rows, err := r.Db.Queryx("SELECT userUuid, hotelUuid, reviewUuid, text, isAnonymous, photos "+
		"FROM reviews WHERE hotelUuid = $1", hotelUuid)
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

	for rows.Next() {
		var v models.Review
		err = rows.Scan(
			&v.UserUuid,
			&v.HotelUuid,
			&v.ReviewUuid,
			&v.Text,
			&v.IsAnonymous,
			pq.Array(&v.Photos),
		)
		if err != nil {
			e = errors.E(opError, errors.RepositoryQueryErr, err)
			r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
			return
		}
		rev = append(rev, v)
	}

	return
}

func (r *ReviewRepository) AddReview(rev *models.Review) (e error) {
	var opError errors.Op = "postgres.AddReview"

	_, err := r.Db.Exec("INSERT INTO "+
		"reviews(userUuid, hotelUuid, reviewUuid, text, isAnonymous, photos) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		rev.UserUuid,
		rev.HotelUuid,
		rev.ReviewUuid,
		rev.Text,
		rev.IsAnonymous,
		pq.Array(rev.Photos),
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

func (r *ReviewRepository) PatchReview(rev *models.Review) (e error) {
	var opError errors.Op = "postgres.PatchReview"

	_, err := r.Db.Exec("UPDATE reviews SET text = $1, photos = $2 WHERE reviewUuid = $3",
		rev.Text,
		pq.Array(rev.Photos),
		rev.ReviewUuid,
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

func (r *ReviewRepository) DeleteReview(reviewUuid uuid.UUID) (e error) {
	var opError errors.Op = "postgres.DeleteReview"

	_, err := r.Db.Exec("DELETE FROM reviews WHERE reviewUuid = $1", reviewUuid)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	}

	return
}
