package errors

import "net/http"

const (
	MaxGrpcCodeValue = 17
)

const (
	JWTGenerationErr Kind = iota + MaxGrpcCodeValue
	JWTVerificationSigningMethodErr
	JWTVerificationErr
	JWTSigningErr
	JWTTokenClaimsErr
	InvalidCredentials
	ExpiredToken
	PermissionDenied
	RepositoryDownErr
	RepositoryQueryErr
	RepositoryNoRows
	UnexpectedErr
	AuthServiceUnavailable
	PaymentServiceUnavailable
	UserLoyaltyServiceUnavailable
	HotelServiceUnavailable
	PaymentExistsErr
	PaymentUserUuidValidationErr
	PaymentUuidValidationErr
	PaymentPriceValidationErr
	LoyaltyExistsErr
	LoyaltyUserUuidValidationErr
	RoomNightPriceValidationErr
	RoomTypeValidationErr
	RoomBedsValidationErr
	RoomAmountValidationErr
	RoomUuidValidationErr
	HotelUuidValidationErr
	HotelDescriptionValidationError
	HotelCountryValidationError
	HotelCityValidationError
	HotelAddressValidationError
	HotelNameValidationError
	UserUuidValidationErr
	UserExistsErr
	UserPasswordLengthValidationError
	UserLoginLengthValidationError
	UserLoginCharsValidationError
	UserLoginUniqueValidationError
	UserPasswordCharsValidationError
	AuthorizationErr
	PaymentNotFoundErr
	UserNotFoundErr
	LoyaltyNotFoundErr
	RoomNotFoundErr
	HotelNotFoundErr
	RepositoryPaymentErr
	RepositoryLoyaltyErr
	RoomFailedToPatch
	HotelFailedToPatch
	RepositoryHotelErr
	RepositoryRoomErr
	RepositoryUserErr
	FailedToHashPassword
	RoomUnavailableErr
)

func (k Kind) String() string {
	switch k {
	case JWTGenerationErr:
		return "Failed to generate JWT access token"
	case JWTVerificationSigningMethodErr:
		return "Unexpected token signing method"
	case JWTVerificationErr:
		return "Invalid token"
	case JWTSigningErr:
		return "Failed to sign token"
	case JWTTokenClaimsErr:
		return "Invalid token claims"
	case InvalidCredentials:
		return "Invalid credentials"
	case ExpiredToken:
		return "Expired token"
	case PermissionDenied:
		return "Permission Denied - no permission to access"
	case RepositoryDownErr:
		return "Repository connection problem"
	case RepositoryQueryErr:
		return "Failed to perform query"
	case RepositoryNoRows:
		return "No rows were found"
	case UnexpectedErr:
		return "Unexpected error occurred"
	case AuthServiceUnavailable:
		return "Authorization service is unavailable"
	case PaymentServiceUnavailable:
		return "Payment service is unavailable"
	case UserLoyaltyServiceUnavailable:
		return "User loyalty service is unavailable"
	case HotelServiceUnavailable:
		return "Hotel service unavailable"
	case PaymentExistsErr:
		return "Payment exists"
	case PaymentUserUuidValidationErr:
		return "User UUID is not valid"
	case PaymentUuidValidationErr:
		return "Payment UUID is not valid"
	case PaymentPriceValidationErr:
		return "Price cannot be less or equal zero"
	case LoyaltyExistsErr:
		return "User's loyalty exists"
	case LoyaltyUserUuidValidationErr:
		return "User UUID is not valid"
	case RoomNightPriceValidationErr:
		return "Room night price should be > 0"
	case RoomTypeValidationErr:
		return "Room type cannot be empty and cannot be more than 250 characters"
	case RoomBedsValidationErr:
		return "Beds number should be integer > 0"
	case RoomAmountValidationErr:
		return "Amount of rooms number should be integer > 0\""
	case RoomUuidValidationErr:
		return "Room UUID is not valid"
	case HotelUuidValidationErr:
		return "Hotel UUID is not valid"
	case HotelDescriptionValidationError:
		return "Description length cannot be more than 1000 characters"
	case HotelCountryValidationError:
		return "Country length cannot be more than 100 characters"
	case HotelCityValidationError:
		return "City length cannot be more than 100 characters"
	case HotelAddressValidationError:
		return "Address length cannot be more than 250 characters"
	case HotelNameValidationError:
		return "Name length cannot be more than 250 characters"
	case UserUuidValidationErr:
		return "User UUID is not valid"
	case UserExistsErr:
		return "User exists cannot add new user"
	case UserPasswordLengthValidationError:
		return "Password length should be less than 128 characters"
	case UserLoginLengthValidationError:
		return "Login should be more than 6 symbols and less than 24 characters"
	case UserLoginCharsValidationError:
		return "Login must consist of latin characters with optional numbers"
	case UserLoginUniqueValidationError:
		return "Login should be unique"
	case UserPasswordCharsValidationError:
		return "The password must consist of latin chars with optional numbers or special characters ./\\\\?,'[]!@#$%^&*()"
	case AuthorizationErr:
		return "Invalid login or password"
	case PaymentNotFoundErr:
		return "Payment not found"
	case UserNotFoundErr:
		return "User not found"
	case LoyaltyNotFoundErr:
		return "Loyalty for current user not found"
	case RoomNotFoundErr:
		return "Room not found"
	case HotelNotFoundErr:
		return "Hotel not found"
	case RepositoryPaymentErr:
		return "Something wrong with payment repository"
	case RepositoryLoyaltyErr:
		return "Something wrong with loyalty repository"
	case RoomFailedToPatch:
		return "Room failed to patch"
	case HotelFailedToPatch:
		return "Failed to patch hotel"
	case RepositoryHotelErr:
		return "Something wrong with hotel repository"
	case RepositoryRoomErr:
		return "Something wrong with room repository"
	case RepositoryUserErr:
		return "Something wrong with user repository"
	case FailedToHashPassword:
		return "Failed to hash password error"
	case RoomUnavailableErr:
		return "Room is unavailable"
	}

	return "Unimplemented error"
}

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
		RoomUnavailableErr,
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

	return http.StatusInternalServerError
}