package loyalty_service

import (
	"hotel-booking-system/internal/pkg/errors"
	"net/http"
)

const (
	RepositoryLoyaltyErr = errors.Kind("Something wrong with loyalty repository")
)

const (
	LoyaltyExistsErr             = errors.Kind("User's loyalty exists")
	LoyaltyNotFoundErr           = errors.Kind("Loyalty for current user not found")
	LoyaltyUserUuidValidationErr = errors.Kind("User UUID is not valid")
)

func GetHttpError(err error) int {
	result := errors.GetHttpError(err)

	notFoundErrors := []errors.Kind{
		LoyaltyNotFoundErr,
	}

	badRequestErrors := []errors.Kind{
		LoyaltyExistsErr,
		LoyaltyUserUuidValidationErr,
	}

	internalError := []errors.Kind{
		RepositoryLoyaltyErr,
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
