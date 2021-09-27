package hotel_service

import (
	"github.com/google/uuid"
	"hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	kinds "hotel-booking-system/internal/pkg/errors/hotel-service"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

func (s *HotelServer) ReviewToProto(r *models.Review) *proto.Review {
	return &proto.Review{
		UserUuid:    r.UserUuid.String(),
		HotelUuid:   r.HotelUuid.String(),
		ReviewUuid:  r.ReviewUuid.String(),
		Text:        r.Text,
		IsAnonymous: r.IsAnonymous,
		Photos: func() []string {
			var res []string
			for _, v := range r.Photos {
				res = append(res, v.String())
			}
			return res
		}(),
		CreationDate: r.CreationDate.Unix(),
	}
}

func (s *HotelServer) RoomToProto(r *models.Room) *proto.Room {
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

func (s *HotelServer) HotelToProto(h *models.Hotel) *proto.Hotel {
	return &proto.Hotel{
		Name:      h.Name,
		HotelUuid: h.HotelUuid.String(),
		Photos: func() []string {
			var r []string
			for _, v := range h.Photos {
				r = append(r, v.String())
			}
			return r
		}(),
		Description:  h.Description,
		Country:      h.Country,
		City:         h.City,
		Address:      h.Address,
		IsReady:      h.IsReady,
		CreationDate: h.CreationDate.Unix(),
		Reviews: func() []*proto.Review {
			var r []*proto.Review
			for _, v := range h.Reviews {
				r = append(r, s.ReviewToProto(&v))
			}

			return r
		}(),
		Rooms: func() []*proto.Room {
			var r []*proto.Room
			for _, v := range h.Rooms {
				r = append(r, s.RoomToProto(&v))
			}

			return r
		}(),
	}
}

func (s *HotelServer) ProtoToRoom(pr *proto.Room) (r *models.Room, e error) {
	var opError errors.Op = "hotel-service.ProtoToRoom"

	validRoomUuid, err := uuid.Parse(pr.RoomUuid)
	if err != nil {
		e = errors.E(opError, kinds.RoomUuidValidationErr, err)
		s.Logger.Error("Grpc error: ", e)
		return
	}

	validHotelUuid, err := uuid.Parse(pr.HotelUuid)
	if err != nil {
		e = errors.E(opError, kinds.HotelUuidValidationErr, err)
		s.Logger.Error("Grpc error: ", e)
		return
	}

	r = &models.Room{
		RoomType:     pr.RoomType,
		Amount:       int(pr.Amount),
		Beds:         int(pr.Beds),
		HotelUuid:    validHotelUuid,
		RoomUuid:     validRoomUuid,
		CreationDate: time.Unix(pr.CreationDate, 0),
		Offers:       pr.Offers,
		NightPrice:   int(pr.NightPrice),
	}

	return
}

func (s *HotelServer) ProtoToReview(pr *proto.Review) (r *models.Review, e error) {
	var opError errors.Op = "hotel-service.ProtoToReview"

	validReviewUuid, err := uuid.Parse(pr.ReviewUuid)
	if err != nil {
		e = errors.E(opError, kinds.ReviewUuidValidationErr, err)
		s.Logger.Error("Grpc error: ", e)
		return
	}

	validUserUuid, err := uuid.Parse(pr.UserUuid)
	if err != nil {
		e = errors.E(opError, kinds.UserUuidValidationErr, err)
		s.Logger.Error("Grpc error: ", e)
		return
	}

	validHotelUuid, err := uuid.Parse(pr.HotelUuid)
	if err != nil {
		e = errors.E(opError, kinds.HotelUuidValidationErr, err)
		s.Logger.Error("Grpc error: ", e)
		return
	}

	var photos []uuid.UUID
	for _, v := range pr.Photos {
		validPhotoUuid, err := uuid.Parse(v)
		if err != nil {
			e = errors.E(opError, kinds.PhotoUuidValidationErr, err)
			s.Logger.Error("Grpc error: ", e)
			return
		}
		photos = append(photos, validPhotoUuid)
	}

	r = &models.Review{
		UserUuid:     validUserUuid,
		HotelUuid:    validHotelUuid,
		ReviewUuid:   validReviewUuid,
		Text:         pr.Text,
		IsAnonymous:  pr.IsAnonymous,
		Photos:       photos,
		CreationDate: time.Unix(pr.CreationDate, 0),
	}

	return
}

func (s *HotelServer) ProtoToHotel(pr *proto.Hotel) (r *models.Hotel, e error) {
	var opError errors.Op = "hotel-service.ProtoToHotel"

	validHotelUuid, err := uuid.Parse(pr.HotelUuid)
	if err != nil {
		e = errors.E(opError, kinds.HotelUuidValidationErr, err)
		s.Logger.Error("Grpc error: ", e)
		return
	}

	var photos []uuid.UUID
	for _, v := range pr.Photos {
		validPhotoUuid, err := uuid.Parse(v)
		if err != nil {
			e = errors.E(opError, kinds.PhotoUuidValidationErr, err)
			s.Logger.Error("Grpc error: ", e)
			return
		}
		photos = append(photos, validPhotoUuid)
	}

	var reviews []models.Review
	for _, v := range pr.Reviews {
		validReview, err := s.ProtoToReview(v)
		if err != nil {
			e = err
			return
		}

		reviews = append(reviews, *validReview)
	}

	var rooms []models.Room
	for _, v := range pr.Rooms {
		validRoom, err := s.ProtoToRoom(v)
		if err != nil {
			e = err
			return
		}

		rooms = append(rooms, *validRoom)
	}

	r = &models.Hotel{
		Name:         pr.Name,
		HotelUuid:    validHotelUuid,
		Photos:       photos,
		Description:  pr.Description,
		Country:      pr.Country,
		City:         pr.City,
		Address:      pr.Address,
		IsReady:      pr.IsReady,
		CreationDate: time.Unix(pr.CreationDate, 0),
		Rooms:        rooms,
		Reviews:      reviews,
	}

	return
}

func (s *HotelServer) HotelsToProto(h []models.Hotel) *proto.HotelsResponse {
	var protoHotels []*proto.Hotel
	for _, v := range h {
		protoHotels = append(protoHotels, s.HotelToProto(&v))
	}

	return &proto.HotelsResponse{
		Hotels: protoHotels,
	}
}

func (s *HotelServer) ReviewsToProto(r []models.Review) *proto.ReviewsResponse {
	var protoReviews []*proto.Review
	for _, v := range r {
		protoReviews = append(protoReviews, s.ReviewToProto(&v))
	}

	return &proto.ReviewsResponse{
		Reviews: protoReviews,
	}
}

func (s *HotelServer) RoomsToProto(r []models.Room) *proto.RoomsResponse {
	var protoRooms []*proto.Room
	for _, v := range r {
		protoRooms = append(protoRooms, s.RoomToProto(&v))
	}

	return &proto.RoomsResponse{
		Rooms: protoRooms,
	}
}

func ProtoToCredentials(c *proto.Credentials) *models.Credentials {
	return &models.Credentials{
		Id:     c.Id,
		Secret: c.Secret,
	}
}

func CredentialsToProto(c *models.Credentials) *proto.Credentials {
	return &proto.Credentials{
		Id:     c.Id,
		Secret: c.Secret,
	}
}

func ProtoToToken(t *proto.Token) *models.Token {
	_token := models.Token(t.Value)
	return &_token
}

func TokenToProto(t *models.Token) *proto.Token {
	return &proto.Token{
		Value: string(*t),
	}
}
