package hotel_service

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type ReviewRepository struct {
	Db     *sqlx.DB
	Logger logs.LoggerInterface
}

func NewReviewRepository(db *sqlx.DB, logger logs.LoggerInterface) models.ReviewRepositoryI {
	return &ReviewRepository{db, logger}
}

func (r *ReviewRepository) GetReview(reviewUuid uuid.UUID) (rev models.Review, e error) {
	var opError errors.Op = "postgres.GetReview"

	e = r.Db.Get(&rev, "SELECT userUuid, hotelUuid, reviewUuid, text, isAnonymous, photos "+
		"FROM review WHERE reviewUuid = $1", reviewUuid)
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr)
		r.Logger.Error("Database error: ", e)
	}

	return
}

func (r *ReviewRepository) GetReviews(hotelUuid uuid.UUID) (rev []models.Review, e error) {
	var opError errors.Op = "postgres.GetReviews"

	e = r.Db.Select(&rev, "SELECT userUuid, hotelUuid, reviewUuid, text, isAnonymous, photos "+
		"FROM review WHERE hotelUuid = $1", hotelUuid)
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr)
		r.Logger.Error("Database error: ", e)
	}

	return
}

func (r *ReviewRepository) AddReview(rev *models.Review) (e error) {
	var opError errors.Op = "postgres.AddReview"

	_, e = r.Db.Exec("INSERT INTO "+
		"reviews(userUuid, hotelUuid, reviewUuid, text, isAnonymous, photos) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		rev.UserUuid, rev.HotelUuid, rev.ReviewUuid, rev.Text, rev.IsAnonymous, rev.Photos)
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, e)
		r.Logger.Error("Database error: ", e)
	}

	return
}

func (r *ReviewRepository) PatchReview(rev *models.Review) (e error) {
	var opError errors.Op = "postgres.PatchReview"

	_, e = r.Db.Exec("UPDATE reviews SET text = $1, photos = $2 WHERE reviewUuid = $3", rev.Text, rev.Photos, rev.ReviewUuid)
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, e)
		r.Logger.Error("Database error: ", e)
	}

	return
}

func (r *ReviewRepository) DeleteReview(reviewUuid uuid.UUID) (e error) {
	var opError errors.Op = "postgres.DeleteReview"

	_, e = r.Db.Exec("DELETE FROM reviews WHERE reviewUuid = $1", reviewUuid)
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, e)
		r.Logger.Error("Database error: ", e)
	}

	return
}
