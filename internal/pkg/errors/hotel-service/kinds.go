package hotel_service

import (
	"hotel-booking-system/internal/pkg/errors"
	"net/http"
)

const (
	RepositoryHotelErr  = errors.Kind("Something wrong with hotel repository")
	RepositoryReviewErr = errors.Kind("Something wrong with review repository")
	RepositoryRoomErr   = errors.Kind("Something wrong with room repository")
)

const (
	HotelNotFoundErr                = errors.Kind("Hotel not found")
	HotelUuidValidationErr          = errors.Kind("Hotel UUID is not valid")
	HotelFailedToPatch              = errors.Kind("Failed to patch hotel")
	HotelDescriptionValidationError = errors.Kind("Description length cannot be more than 1000 characters")
	HotelCountryValidationError     = errors.Kind("Country length cannot be more than 100 characters")
	HotelCityValidationError        = errors.Kind("City length cannot be more than 100 characters")
	HotelAddressValidationError     = errors.Kind("Address length cannot be more than 250 characters")
	HotelNameValidationError        = errors.Kind("Name length cannot be more than 250 characters")
)

const (
	ReviewUserNotFound      = errors.Kind("User not exists")
	ReviewTextValidationErr = errors.Kind("Text cannot be more than 1500 characters")
	ReviewNotFoundErr       = errors.Kind("Review not found")
	ReviewUuidValidationErr = errors.Kind("Review UUID is not valid")
	UserUuidValidationErr   = errors.Kind("User UUID is not valid")
	ReviewFailedToPatch     = errors.Kind("Failed to patch review")
)

const (
	RoomNotFoundErr             = errors.Kind("Room not found")
	RoomUuidValidationErr       = errors.Kind("Room UUID is not valid")
	RoomBedsValidationErr       = errors.Kind("Beds number should be integer > 0")
	RoomAmountValidationErr     = errors.Kind("Amount of rooms number should be integer > 0")
	RoomNightPriceValidationErr = errors.Kind("Room night price should be > 0")
	RoomTypeValidationErr       = errors.Kind("Room type cannot be empty and cannot be more than 250 characters")
	RoomFailedToPatch           = errors.Kind("Room failed to patch")
)

const (
	PhotoUuidValidationErr = errors.Kind("Photo uuid is not valid")
)

func GetHttpError(err error) int {
	result := errors.GetHttpError(err)

	notFoundErrors := []errors.Kind{
		RoomNotFoundErr,
		HotelNotFoundErr,
		ReviewUserNotFound,
		ReviewNotFoundErr,
	}

	badRequestErrors := []errors.Kind{
		RoomNightPriceValidationErr,
		RoomTypeValidationErr,
		RoomBedsValidationErr,
		RoomAmountValidationErr,
		RoomUuidValidationErr,
		HotelUuidValidationErr,
		HotelDescriptionValidationError,
		HotelCountryValidationError,
		HotelCityValidationError,
		HotelAddressValidationError,
		HotelNameValidationError,
		ReviewTextValidationErr,
		ReviewUuidValidationErr,
		UserUuidValidationErr,
	}

	internalError := []errors.Kind{
		RoomFailedToPatch,
		ReviewFailedToPatch,
		HotelFailedToPatch,
		RepositoryHotelErr,
		RepositoryReviewErr,
		RepositoryRoomErr,
	}

	kind := errors.GetKind(err)

	if errors.Contains(notFoundErrors, kind) {
		result = http.StatusNotFound
	}

	if errors.Contains(badRequestErrors, kind) {
		result = http.StatusBadRequest
	}

	if errors.Contains(internalError, kind) {
		result = http.StatusInternalServerError
	}

	return result
}
