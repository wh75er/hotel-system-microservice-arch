package hotel_service

import (
	"github.com/aglyzov/go-patch"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"hotel-booking-system/internal/pkg/errors"
	kinds "hotel-booking-system/internal/pkg/errors/hotel-service"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

type ReviewUsecase struct {
	HotelRepository models.HotelRepositoryI
	ReviewRepository models.ReviewRepositoryI
	Logger *logrus.Logger
}

func NewReviewUsecase(hotelR models.HotelRepositoryI, reviewR models.ReviewRepositoryI, logger *logrus.Logger) models.ReviewUsecaseI {
	return &ReviewUsecase{hotelR, reviewR, logger}
}

func (u *ReviewUsecase) GetReview(reviewUuid string) (r models.Review, e error) {
	var opError errors.Op = "usecase.GetReview"

	validReviewUuid, e := uuid.Parse(reviewUuid)
	if e != nil {
		e = errors.E(opError, kinds.ReviewUuidValidationErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	r, e = u.ReviewRepository.GetReview(validReviewUuid)
	if e != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = nil
			return
		}
		e = errors.E(opError, kinds.RepositoryReviewErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *ReviewUsecase) GetReviews(hotelUuid string) (r []models.Review, e error) {
	var opError errors.Op = "usecase.GetReviews"

	validHotelUuid, e := uuid.Parse(hotelUuid)
	if e != nil {
		e = errors.E(opError, kinds.HotelUuidValidationErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	_, e = u.HotelRepository.GetHotel(validHotelUuid)
	if e != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = errors.E(opError, kinds.HotelNotFoundErr, e)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		e = errors.E(opError, kinds.RepositoryHotelErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	r, e = u.ReviewRepository.GetReviews(validHotelUuid)
	if e != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = nil
			return
		}
		e = errors.E(opError, kinds.RepositoryReviewErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *ReviewUsecase) AddReview(r *models.Review) (e error) {
	var opError errors.Op = "usecase.AddReviews"

	e = u.validateReview(opError, r)
	if e != nil {
		u.Logger.Error("Usecase error: ", e)
		return
	}

	// Check if r.userUuid user exists

	_, e = u.HotelRepository.GetHotel(r.HotelUuid)
	if e != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = errors.E(opError, kinds.HotelNotFoundErr, e)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		e = errors.E(opError, kinds.RepositoryHotelErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	r.ReviewUuid = uuid.New()
	r.CreationDate = time.Now().UTC()

	e = u.ReviewRepository.AddReview(r)
	if e != nil {
		e = errors.E(opError, kinds.RepositoryReviewErr)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *ReviewUsecase) PatchReview(r *models.Review) (e error) {
	var opError errors.Op = "usecase.PatchReviews"

	e = u.validateReview(opError, r)
	if e != nil {
		u.Logger.Error("Usecase error: ", e)
		return
	}

	currentReview, e := u.ReviewRepository.GetReview(r.ReviewUuid)
	if e != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = errors.E(opError, kinds.ReviewNotFoundErr, e)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		e = errors.E(opError, kinds.RepositoryReviewErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	// reset unchangable fields
	r.CreationDate = time.Time{}
	r.UserUuid = uuid.UUID{}
	r.HotelUuid = uuid.UUID{}

	_, e = patch.Struct(&currentReview, r)
	if e != nil {
		e = errors.E(opError, kinds.ReviewFailedToPatch, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	e = u.ReviewRepository.PatchReview(&currentReview)
	if e != nil {
		e = errors.E(opError, kinds.RepositoryReviewErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *ReviewUsecase) DeleteReview(reviewUuid string) (e error) {
	var opError errors.Op = "usecase.DeleteHotel"

	validReviewUuid, e := uuid.Parse(reviewUuid)
	if e != nil {
		e = errors.E(opError, kinds.ReviewUuidValidationErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	e = u.HotelRepository.DeleteHotel(validReviewUuid)
	if e != nil {
		e = errors.E(opError, kinds.RepositoryReviewErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *ReviewUsecase) validateReview(opError errors.Op, r *models.Review) (e error) {
	if len(r.Text) > 1500 {
		e = errors.E(opError, kinds.ReviewTextValidationErr, e)
		return
	}

	return
}
