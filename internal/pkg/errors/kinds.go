package errors

import (
	"net/http"
)

const (
	MaxGrpcCodeValue = 17
)

const (
	RoomUnavailableCustomStatus = 419
)

const (
	JWTGenerationErr                = Kind("Failed to generate JWT access token")
	JWTVerificationSigningMethodErr = Kind("Unexpected token signing method")
	JWTVerificationErr              = Kind("Invalid token")
	JWTSigningErr                   = Kind("Failed to sign token")
	JWTTokenClaimsErr               = Kind("Invalid token claims")
	InvalidCredentials              = Kind("Invalid credentials")
	ExpiredToken                    = Kind("Expired token")
	PermissionDenied                = Kind("Permission Denied - no permission to access")
	RepositoryDownErr               = Kind("Repository connection problem")
	RepositoryQueryErr              = Kind("Failed to perform query")
	RepositoryNoRows                = Kind("No rows were found")
	UnexpectedErr                   = Kind("Unexpected error occurred")
)

const (
	AuthServiceUnavailable        = Kind("Authorization service is unavailable")
	PaymentServiceUnavailable     = Kind("Payment service is unavailable")
	UserLoyaltyServiceUnavailable = Kind("User loyalty service is unavailable")
	HotelServiceUnavailable       = Kind("Hotel service unavailable")
)

func GetHttpError(err error) int {
	badRequestErrors := []Kind{
		JWTVerificationSigningMethodErr,
		JWTVerificationErr,
		JWTTokenClaimsErr,
		JWTSigningErr,
		PaymentExistsErr,
		PaymentUserUuidValidationErr,
		PaymentUuidValidationErr,
		PaymentPriceValidationErr,
		LoyaltyExistsErr,
		LoyaltyUserUuidValidationErr,
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
		UserUuidValidationErr,
		PhotoUuidValidationErr,
		UserExistsErr,
		UserUuidValidationErr,
		UserPasswordLengthValidationError,
		UserLoginLengthValidationError,
		UserLoginCharsValidationError,
		UserLoginUniqueValidationError,
		UserPasswordCharsValidationError,
	}

	UnauthorizedErrors := []Kind{
		InvalidCredentials,
		AuthorizationErr,
	}

	Forbidden := []Kind{
		ExpiredToken,
		PermissionDenied,
	}

	notFoundErrors := []Kind{
		PaymentNotFoundErr,
		UserNotFoundErr,
		LoyaltyNotFoundErr,
		RoomNotFoundErr,
		HotelNotFoundErr,
		UserNotFoundErr,
	}

	unprocessibleEntity := []Kind{
		AuthServiceUnavailable,
		PaymentServiceUnavailable,
		UserLoyaltyServiceUnavailable,
		HotelServiceUnavailable,
	}

	internalError := []Kind{
		JWTGenerationErr,
		RepositoryDownErr,
		RepositoryQueryErr,
		RepositoryNoRows,
		UnexpectedErr,
		RepositoryPaymentErr,
		RepositoryLoyaltyErr,
		RoomFailedToPatch,
		HotelFailedToPatch,
		RepositoryHotelErr,
		RepositoryReviewErr,
		RepositoryRoomErr,
		RepositoryUserErr,
		FailedToHashPassword,
	}

	kind := GetKind(err)

	if Contains(badRequestErrors, kind) {
		return http.StatusBadRequest
	}

	if Contains(notFoundErrors, kind) {
		return http.StatusNotFound
	}

	if Contains(internalError, kind) {
		return http.StatusInternalServerError
	}

	if Contains(Forbidden, kind) {
		return http.StatusForbidden
	}

	if Contains(unprocessibleEntity, kind) {
		return http.StatusUnprocessableEntity
	}

	if Contains(UnauthorizedErrors, kind) {
		return http.StatusUnauthorized
	}

	if kind == RoomUnavailableErr {
		return RoomUnavailableCustomStatus
	}

	return http.StatusInternalServerError
}
