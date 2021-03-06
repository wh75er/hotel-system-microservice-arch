package hotel_service

import (
	"context"
	"github.com/aglyzov/go-patch"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	stat_proto "hotel-booking-system/internal/pkg/delivery/grpc/stat-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

type RoomUsecase struct {
	HotelRepository models.HotelRepositoryI
	RoomRepository  models.RoomRepositoryI
	StatServiceClient stat_proto.StatServiceClient
	Logger          logs.LoggerInterface
}

func NewRoomUsecase(
	hotelR models.HotelRepositoryI,
	roomR models.RoomRepositoryI,
	statServiceClient stat_proto.StatServiceClient,
	logger logs.LoggerInterface,
) models.RoomUsecaseI {
	return &RoomUsecase{
		hotelR,
		roomR,
		statServiceClient,
		logger,
	}
}

func (u *RoomUsecase) GetRooms(hotelUuid string) (r []models.Room, e error) {
	var opError errors.Op = "usecase.GetRooms"

	validHotelUuid, e := uuid.Parse(hotelUuid)
	if e != nil {
		e = errors.E(opError, errors.HotelUuidValidationErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	_, e = u.HotelRepository.GetHotel(validHotelUuid)
	if e != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.HotelNotFoundErr, e)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		e = errors.E(opError, errors.RepositoryHotelErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	r, e = u.RoomRepository.GetRooms(validHotelUuid)
	if e != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = nil
			return
		}
		e = errors.E(opError, errors.RepositoryRoomErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *RoomUsecase) GetRoom(roomUuid string) (r *models.Room, e error) {
	var opError errors.Op = "usecase.GetRoom"

	validRoomUuid, err := uuid.Parse(roomUuid)
	if e != nil {
		e = errors.E(opError, errors.RoomUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	r = &models.Room{}

	*r, err = u.RoomRepository.GetRoom(validRoomUuid)
	if e != nil {
		e = errors.E(opError, errors.RoomNotFoundErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *RoomUsecase) AddRoom(r *models.Room) (e error) {
	var opError errors.Op = "usecase.AddRoom"

	e = u.validateRoom(opError, r)
	if e != nil {
		u.Logger.Error("Usecase error: ", e)
		return
	}

	_, e = u.HotelRepository.GetHotel(r.HotelUuid)
	if e != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.HotelNotFoundErr, e)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		e = errors.E(opError, errors.RepositoryHotelErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	r.RoomUuid = uuid.New()
	r.CreationDate = time.Now()

	// add room to statistics

	_, err := u.StatServiceClient.UpdateRoomsAmount(context.Background(), &stat_proto.Delta{ Value: int64(r.Amount) })
	if err != nil {
		if status.Code(err) < errors.MaxGrpcCodeValue {
			e = errors.E(opError, errors.StatServiceUnavailable, err)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		serviceKind := errors.Kind(status.Code(err))
		e = errors.E(opError, serviceKind)
		return
	}

	e = u.RoomRepository.AddRoom(r)
	if e != nil {
		e = errors.E(opError, errors.RepositoryRoomErr)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *RoomUsecase) PatchRoom(r *models.Room) (e error) {
	var opError errors.Op = "usecase.PatchRoom"

	e = u.validateRoom(opError, r)
	if e != nil {
		u.Logger.Error("Usecase error: ", e)
		return e
	}

	currentRoom, e := u.RoomRepository.GetRoom(r.RoomUuid)
	if e != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.RoomNotFoundErr, e)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		e = errors.E(opError, errors.RepositoryRoomErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	// reset unchangable fields
	r.CreationDate = time.Time{}
	r.HotelUuid = uuid.UUID{}

	// change room amount in statistics

	newAmount := r.Amount - currentRoom.Amount

	_, err := u.StatServiceClient.UpdateRoomsAmount(context.Background(), &stat_proto.Delta{ Value: int64(newAmount) })
	if err != nil {
		if status.Code(err) < errors.MaxGrpcCodeValue {
			e = errors.E(opError, errors.StatServiceUnavailable, err)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		serviceKind := errors.Kind(status.Code(err))
		e = errors.E(opError, serviceKind)
		return
	}

	_, e = patch.Struct(&currentRoom, r)
	if e != nil {
		e = errors.E(opError, errors.RoomFailedToPatch, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	e = u.RoomRepository.PatchRoom(&currentRoom)
	if e != nil {
		e = errors.E(opError, errors.RepositoryRoomErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *RoomUsecase) TakeRoom(roomUuid string) (e error) {
	err := u.changeRoomAmount(roomUuid, -1)
	if err != nil {
		e = err
		return err
	}

	return
}

func (u *RoomUsecase) DismissRoom(roomUuid string) (e error) {
	err := u.changeRoomAmount(roomUuid, +1)
	if err != nil {
		e = err
		return err
	}

	return
}

func (u *RoomUsecase) changeRoomAmount(roomUuid string, lambda int) (e error) {
	var opError errors.Op = "usecase.changeRoomAmount"

	validRoomUuid, err := uuid.Parse(roomUuid)
	if err != nil {
		e = errors.E(opError, errors.RoomUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	r, err := u.RoomRepository.GetRoom(validRoomUuid)
	if err != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.RoomNotFoundErr, err)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		e = errors.E(opError, errors.RepositoryRoomErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	if r.Amount == 0 && lambda < 0 {
		e = errors.E(opError, errors.RoomUnavailableErr)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	r.Amount += lambda

	err = u.RoomRepository.PatchRoom(&r)
	if err != nil {
		e = errors.E(opError, errors.RepositoryRoomErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *RoomUsecase) DeleteRoom(roomUuid string) (e error) {
	var opError errors.Op = "usecase.DeleteRoom"

	validRoomUuid, e := uuid.Parse(roomUuid)
	if e != nil {
		e = errors.E(opError, errors.RoomUuidValidationErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	e = u.HotelRepository.DeleteHotel(validRoomUuid)
	if e != nil {
		e = errors.E(opError, errors.RepositoryRoomErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *RoomUsecase) validateRoom(opError errors.Op, r *models.Room) (e error) {
	if len(r.RoomType) > 250 || len(r.RoomType) == 0 {
		e = errors.E(opError, errors.RoomTypeValidationErr, e)
		return
	}

	if r.Amount <= 0  {
		e = errors.E(opError, errors.RoomAmountValidationErr, e)
		return
	}

	if r.Beds <= 0 {
		e = errors.E(opError, errors.RoomBedsValidationErr, e)
		return
	}

	if r.NightPrice <= 0 {
		e = errors.E(opError, errors.RoomNightPriceValidationErr, e)
		return
	}

	return
}
