package hotel_service

import (
	"github.com/google/uuid"
	"hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

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
		Name:         h.Name,
		HotelUuid:    h.HotelUuid.String(),
		Description:  h.Description,
		Country:      h.Country,
		City:         h.City,
		Address:      h.Address,
		IsReady:      h.IsReady,
		CreationDate: h.CreationDate.Unix(),
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
		e = errors.E(opError, errors.RoomUuidValidationErr, err)
		s.Logger.Error("Grpc error: ", e)
		return
	}

	validHotelUuid, err := uuid.Parse(pr.HotelUuid)
	if err != nil {
		e = errors.E(opError, errors.HotelUuidValidationErr, err)
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

func (s *HotelServer) ProtoToHotel(pr *proto.Hotel) (r *models.Hotel, e error) {
	var opError errors.Op = "hotel-service.ProtoToHotel"

	validHotelUuid, err := uuid.Parse(pr.HotelUuid)
	if err != nil {
		e = errors.E(opError, errors.HotelUuidValidationErr, err)
		s.Logger.Error("Grpc error: ", e)
		return
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
		Description:  pr.Description,
		Country:      pr.Country,
		City:         pr.City,
		Address:      pr.Address,
		IsReady:      pr.IsReady,
		CreationDate: time.Unix(pr.CreationDate, 0),
		Rooms:        rooms,
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

func (s *HotelServer) RoomsToProto(r []models.Room) *proto.RoomsResponse {
	var protoRooms []*proto.Room
	for _, v := range r {
		protoRooms = append(protoRooms, s.RoomToProto(&v))
	}

	return &proto.RoomsResponse{
		Rooms: protoRooms,
	}
}
