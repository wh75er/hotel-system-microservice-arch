package hotel_service

import (
	"github.com/google/uuid"
	"hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

func RoomToProto(r *models.Room) *proto.Room {
	return &proto.Room{
		RoomType:     r.RoomType,
		Amount:       int64(r.Amount),
		Beds:         int64(r.Beds),
		HotelUuid:    r.HotelUuid.String(),
		RoomUuid:     r.RoomUuid.String(),
		CreationDate: r.CreationDate.Unix(),
		Offers:       r.Offers,
		NightPrice:   r.NightPrice,
	}
}

func HotelToProto(h *models.Hotel) *proto.Hotel {
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
				r = append(r, RoomToProto(&v))
			}

			return r
		}(),
	}
}

func ProtoToRoom(pr *proto.Room) (r *models.Room, e error) {
	var opError errors.Op = "hotel-service.ProtoToRoom"

	if pr == nil {
		return &models.Room{}, nil
	}

	var validRoomUuid uuid.UUID
	if len(pr.RoomUuid) != 0 {
		var err error
		validRoomUuid, err = uuid.Parse(pr.RoomUuid)
		if err != nil {
			e = errors.E(opError, errors.RoomUuidValidationErr, err)
			return
		}
	}

	var validHotelUuid uuid.UUID
	if len(pr.HotelUuid) != 0 {
		var err error
		validHotelUuid, err = uuid.Parse(pr.HotelUuid)
		if err != nil {
			e = errors.E(opError, errors.HotelUuidValidationErr, err)
			return
		}
	}

	r = &models.Room{
		RoomType:     pr.RoomType,
		Amount:       int(pr.Amount),
		Beds:         int(pr.Beds),
		HotelUuid:    validHotelUuid,
		RoomUuid:     validRoomUuid,
		CreationDate: time.Unix(pr.CreationDate, 0),
		Offers:       pr.Offers,
		NightPrice:   pr.NightPrice,
	}

	return
}

func ProtoToHotel(pr *proto.Hotel) (r *models.Hotel, e error) {
	var opError errors.Op = "hotel-service.ProtoToHotel"

	if pr == nil {
		return &models.Hotel{}, nil
	}

	var validHotelUuid uuid.UUID
	if len(pr.HotelUuid) != 0 {
		var err error
		validHotelUuid, err = uuid.Parse(pr.HotelUuid)
		if err != nil {
			e = errors.E(opError, errors.HotelUuidValidationErr, err)
			return
		}
	}

	var rooms []models.Room
	for _, v := range pr.Rooms {
		validRoom, err := ProtoToRoom(v)
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

func HotelsToProto(h []models.Hotel) *proto.HotelsResponse {
	var protoHotels []*proto.Hotel
	for _, v := range h {
		protoHotels = append(protoHotels, HotelToProto(&v))
	}

	return &proto.HotelsResponse{
		Hotels: protoHotels,
	}
}

func RoomsToProto(r []models.Room) *proto.RoomsResponse {
	var protoRooms []*proto.Room
	for _, v := range r {
		protoRooms = append(protoRooms, RoomToProto(&v))
	}

	return &proto.RoomsResponse{
		Rooms: protoRooms,
	}
}
