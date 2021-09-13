package hotel_service

import (
	"github.com/google/uuid"
	"hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	kinds "hotel-booking-system/internal/pkg/errors/hotel-service"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

func ReviewToProto(r *models.Review) *proto.Review {
	return &proto.Review {
		UserUuid: r.UserUuid.String(),
		HotelUuid: r.HotelUuid.String(),
		ReviewUuid: r.ReviewUuid.String(),
		Text: r.Text,
		IsAnonymous: r.IsAnonymous,
		Photos: r.Photos,
		CreationDate: r.CreationDate.Unix(),
	}
}

func RoomToProto(r *models.Room) *proto.Room {
	return &proto.Room{
		RoomType:     r.RoomType,
		Amount:       int64(r.Amount),
		Beds:         int64(r.Beds),
		HotelUuid:    r.HotelUuid.String(),
		RoomUuid:     r.RoomUuid.String(),
		CreationDate: r.CreationDate.Unix(),
		Offers:       r.Offers,
		NightPrice:   int64(r.NightPrice),
	}
}

func HotelToProto(h *models.Hotel) *proto.Hotel {
	return &proto.Hotel {
		Name: h.Name,
		HotelUuid: h.HotelUuid.String(),
		Photos: func() []string {
			var r []string
			for _, v := range h.Photos {
				r = append(r, v.String())
			}
			return r
		}(),
		Description: h.Description,
		Country: h.Country,
		City: h.City,
		Address: h.Address,
		IsReady: h.IsReady,
		CreationDate: h.CreationDate.Unix(),
		Reviews: func() []*proto.Review {
			var r []*proto.Review
			for _, v := range h.Reviews {
				r = append(r, ReviewToProto(&v))
			}

			return r
		}(),
		Rooms: func() []*proto.Room {
			var r []*proto.Room
			for _, v := range h.Rooms {
				r = append(r, RoomToProto(&v))
			}

			return r
		}(),
	}
}

func ProtoToRoom(pr *proto.Room) (r *models.Room, e error) {
	var opError errors.Op = "hotel-service.ProtoToRoom"

	validRoomUuid, e := uuid.Parse(pr.RoomUuid)
	if e != nil {
		e = errors.E(opError, kinds.RoomUuidValidationErr, e)
		//u.Logger.Error("Usecase error: ", e)
		return
	}

	validHotelUuid, e := uuid.Parse(pr.HotelUuid)
	if e != nil {
		e = errors.E(opError, kinds.HotelUuidValidationErr, e)
		//u.Logger.Error("Usecase error: ", e)
		return
	}

	r = &models.Room {
		RoomType: pr.RoomType,
		Amount: int(pr.Amount),
		Beds: int(pr.Beds),
		HotelUuid: validHotelUuid,
		RoomUuid: validRoomUuid,
		CreationDate: time.Unix(pr.CreationDate, 0),
		Offers: pr.Offers,
		NightPrice: int(pr.NightPrice),
	}

	return
}

func ProtoToReview(pr *proto.Review) (r *models.Review, e error) {
	var opError errors.Op = "hotel-service.ProtoToReview"

	validReviewUuid, e := uuid.Parse(pr.ReviewUuid)
	if e != nil {
		e = errors.E(opError, kinds.ReviewUuidValidationErr, e)
		//u.Logger.Error("Usecase error: ", e)
		return
	}

	validUserUuid, e := uuid.Parse(pr.UserUuid)
	if e != nil {
		e = errors.E(opError, kinds.UserUuidValidationErr, e)
		//u.Logger.Error("Usecase error: ", e)
		return
	}

	validHotelUuid, e := uuid.Parse(pr.HotelUuid)
	if e != nil {
		e = errors.E(opError, kinds.HotelUuidValidationErr, e)
		//u.Logger.Error("Usecase error: ", e)
		return
	}

	r = &models.Review{
		UserUuid: validUserUuid,
		HotelUuid: validHotelUuid,
		ReviewUuid: validReviewUuid,
		Text: pr.Text,
		IsAnonymous: pr.IsAnonymous,
		Photos: pr.Photos,
		CreationDate: time.Unix(pr.CreationDate, 0),
	}

	return
}

func ProtoToHotel(pr *proto.Hotel) (r *models.Hotel, e error) {
	var opError errors.Op = "hotel-service.ProtoToHotel"

	validHotelUuid, e := uuid.Parse(pr.HotelUuid)
	if e != nil {
		e = errors.E(opError, kinds.HotelUuidValidationErr, e)
		//u.Logger.Error("Usecase error: ", e)
		return
	}

	var photos []uuid.UUID
	for _, v := range pr.Photos {
		validPhotoUuid, e := uuid.Parse(v)
		if e != nil {
			e = errors.E(opError, kinds.PhotoUuidValidationErr, e)
			//u.Logger.Error("Usecase error: ", e)
			return
		}
		photos = append(photos, validPhotoUuid)
	}

	var reviews []models.Review
	for _, v := range pr.Reviews {
		validReview, e := ProtoToReview(v)
		if e != nil {
			return
		}

		reviews = append(reviews, *validReview)
	}

	var rooms []models.Room
	for _, v := range pr.Rooms {
		validRoom, e := ProtoToRoom(v)
		if e != nil {
			return
		}

		rooms = append(rooms, *validRoom)
	}

	r = &models.Hotel {
		Name: pr.Name,
		HotelUuid: validHotelUuid,
		Photos: photos,
		Description: pr.Description,
		Country: pr.Country,
		City: pr.City,
		Address: pr.Address,
		IsReady: pr.IsReady,
		CreationDate: time.Unix(pr.CreationDate, 0),
		Rooms: rooms,
		Reviews: reviews,
	}

	return
}