package errors

const (
	RepositoryHotelErr  = Kind("Something wrong with hotel repository")
	RepositoryReviewErr = Kind("Something wrong with review repository")
	RepositoryRoomErr   = Kind("Something wrong with room repository")
)

const (
	HotelNotFoundErr                = Kind("Hotel not found")
	HotelUuidValidationErr          = Kind("Hotel UUID is not valid")
	HotelFailedToPatch              = Kind("Failed to patch hotel")
	HotelDescriptionValidationError = Kind("Description length cannot be more than 1000 characters")
	HotelCountryValidationError     = Kind("Country length cannot be more than 100 characters")
	HotelCityValidationError        = Kind("City length cannot be more than 100 characters")
	HotelAddressValidationError     = Kind("Address length cannot be more than 250 characters")
	HotelNameValidationError        = Kind("Name length cannot be more than 250 characters")
)

const (
	RoomNotFoundErr             = Kind("Room not found")
	RoomUnavailableErr			= Kind("This room is not available")
	RoomUuidValidationErr       = Kind("Room UUID is not valid")
	RoomBedsValidationErr       = Kind("Beds number should be integer > 0")
	RoomAmountValidationErr     = Kind("Amount of rooms number should be integer > 0")
	RoomNightPriceValidationErr = Kind("Room night price should be > 0")
	RoomTypeValidationErr       = Kind("Room type cannot be empty and cannot be more than 250 characters")
	RoomFailedToPatch           = Kind("Room failed to patch")
)

const (
	PhotoUuidValidationErr = Kind("Photo uuid is not valid")
)
