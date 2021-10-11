package hotel_service

import (
	"github.com/aglyzov/go-patch"
	"github.com/google/uuid"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

type HotelUsecase struct {
	HotelRepository models.HotelRepositoryI
	RoomRepository  models.RoomRepositoryI
	Logger          logs.LoggerInterface
}

func NewHotelUsecase(hotelR models.HotelRepositoryI, roomR models.RoomRepositoryI, logger logs.LoggerInterface) models.HotelUsecaseI {
	return &HotelUsecase{hotelR, roomR, logger}
}

func (u *HotelUsecase) GetHotel(hotelUuid string) (h models.Hotel, e error) {
	var opError errors.Op = "usecase.GetHotel"

	validHotelUuid, e := uuid.Parse(hotelUuid)
	if e != nil {
		e = errors.E(opError, errors.HotelUuidValidationErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	h, e = u.HotelRepository.GetHotel(validHotelUuid)
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

	rooms, e := u.RoomRepository.GetRooms(validHotelUuid)
	if e != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = nil
		} else {
			e = errors.E(opError, errors.RepositoryRoomErr, e)
			u.Logger.Error("Usecase error: ", e)
			return
		}
	}

	h.Rooms = rooms

	return
}

func (u *HotelUsecase) GetHotels() (h []models.Hotel, e error) {
	var opError errors.Op = "usecase.GetHotels"

	h, e = u.HotelRepository.GetHotels()
	if e != nil {
		if errors.GetKind(e) == errors.RepositoryNoRows {
			e = nil
		} else {
			e = errors.E(opError, errors.RepositoryHotelErr, e)
			u.Logger.Error("Usecase error: ", e)
			return
		}
	}

	return
}

func (u *HotelUsecase) AddHotel(h *models.Hotel) (e error) {
	var opError errors.Op = "usecase.AddHotels"

	e = u.validateHotel(opError, h)
	if e != nil {
		u.Logger.Error("Usecase error: ", e)
		return
	}

	h.HotelUuid = uuid.New()
	h.CreationDate = time.Now()

	e = u.HotelRepository.AddHotel(h)
	if e != nil {
		e = errors.E(opError, errors.RepositoryHotelErr, e)
		u.Logger.Error("Usecase error: ", e)
	}

	return
}

func (u *HotelUsecase) PatchHotel(h *models.Hotel) (e error) {
	var opError errors.Op = "usecase.PatchHotel"

	e = u.validateHotel(opError, h)
	if e != nil {
		u.Logger.Error("Usecase error: ", e)
		return
	}

	currentHotel, e := u.HotelRepository.GetHotel(h.HotelUuid)
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

	// Reset unchangable fields
	h.CreationDate = time.Time{}
	h.Rooms = []models.Room{}

	_, e = patch.Struct(&currentHotel, h)
	if e != nil {
		e = errors.E(opError, errors.HotelFailedToPatch, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	e = u.HotelRepository.PatchHotel(&currentHotel)
	if e != nil {
		e = errors.E(opError, errors.RepositoryHotelErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *HotelUsecase) DeleteHotel(hotelUuid string) (e error) {
	var opError errors.Op = "usecase.DeleteHotel"

	validHotelUuid, e := uuid.Parse(hotelUuid)
	if e != nil {
		e = errors.E(opError, errors.HotelUuidValidationErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	e = u.HotelRepository.DeleteHotel(validHotelUuid)
	if e != nil {
		e = errors.E(opError, errors.RepositoryHotelErr, e)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *HotelUsecase) validateHotel(opError errors.Op, h *models.Hotel) (e error) {
	if len(h.Name) > 250 || len(h.Name) == 0 {
		e = errors.E(opError, errors.HotelNameValidationError, e)
		return
	}

	if len(h.Description) > 1000 || len(h.Description) == 0 {
		e = errors.E(opError, errors.HotelDescriptionValidationError, e)
		return
	}

	if len(h.Country) > 100 || len(h.Country) == 0 {
		e = errors.E(opError, errors.HotelCountryValidationError, e)
		return
	}

	if len(h.City) > 100 || len(h.City) == 0 {
		e = errors.E(opError, errors.HotelCityValidationError, e)
		return
	}

	if len(h.Address) > 250 || len(h.Address) == 0 {
		e = errors.E(opError, errors.HotelAddressValidationError, e)
		return
	}

	return
}
