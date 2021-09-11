package models

import (
	"github.com/google/uuid"
	"time"
)

type Review struct {
	UserUuid uuid.UUID
	HotelUuid uuid.UUID
	ReviewUuid uuid.UUID
	Text string
	IsAnonymous bool
	Photos []string
	CreationDate time.Time
}

type ReviewRepositoryI interface {
	GetReview(reviewUuid uuid.UUID) (r Review, e error)
	GetReviews(hotelUuid uuid.UUID) (r []Review, e error)
	AddReview(r *Review)  (e error)
	PatchReview(r *Review) (e error)
	DeleteReview(reviewUuid uuid.UUID) (e error)
}

type ReviewUsecaseI interface {
	GetReview(reviewUuid string) (r Review, e error)
	GetReviews(hotelUuid string) (r []Review, e error)
	AddReview(r *Review)  (e error)
	PatchReview(r *Review) (e error)
	DeleteReview(reviewUuid string) (e error)
}
