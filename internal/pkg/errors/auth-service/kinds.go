package hotel_service

import (
	"hotel-booking-system/internal/pkg/errors"
	"net/http"
)

const (
	RepositoryUserErr = errors.Kind("Something wrong with user repository")
)

const (
	UserExistsErr                     = errors.Kind("User exists cannot add new user")
	UserNotFoundErr                   = errors.Kind("User not found")
	UserUuidValidationErr             = errors.Kind("User UUID is not valid")
	UserPasswordLengthValidationError = errors.Kind("Password length should be less than 128 characters")
	UserPasswordCharsValidationError  = errors.Kind("The password must consist of latin chars with optional numbers or special characters ./\\?,'[]!@#$%^&*()")
	UserLoginLengthValidationError    = errors.Kind("Login should be more than 6 symbols and less than 24 characters")
	UserLoginCharsValidationError     = errors.Kind("Login must consist of latin characters with optional numbers")
	UserLoginUniqueValidationError    = errors.Kind("Login should be unique")
)

const (
	AuthorizationErr = errors.Kind("Invalid login or password")
)

const (
	FailedToHashPassword = errors.Kind("Failed to hash password error")
)

func GetHttpError(err error) int {
	result := errors.GetHttpError(err)

	notFoundErrors := []errors.Kind{
		UserNotFoundErr,
	}

	UnauthorizedErrors := []errors.Kind{
		AuthorizationErr,
	}

	badRequestErrors := []errors.Kind{
		UserExistsErr,
		UserUuidValidationErr,
		UserPasswordLengthValidationError,
		UserLoginLengthValidationError,
		UserLoginCharsValidationError,
		UserLoginUniqueValidationError,
		UserPasswordCharsValidationError,
	}

	internalError := []errors.Kind{
		FailedToHashPassword,
	}

	kind := errors.GetKind(err)

	if errors.Contains(notFoundErrors, kind) {
		result = http.StatusNotFound
	}

	if errors.Contains(badRequestErrors, kind) {
		result = http.StatusBadRequest
	}

	if errors.Contains(UnauthorizedErrors, kind) {
		return http.StatusUnauthorized
	}

	if errors.Contains(internalError, kind) {
		result = http.StatusInternalServerError
	}

	return result
}
