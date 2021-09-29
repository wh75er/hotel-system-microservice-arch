package payment_service

import (
	"hotel-booking-system/internal/pkg/errors"
	"net/http"
)

const (
	RepositoryPaymentErr = errors.Kind("Something wrong with payment repository")
)

const (
	PaymentExistsErr             = errors.Kind("Payment exists")
	PaymentNotFoundErr           = errors.Kind("Payment not found")
	UserNotFoundErr 			 = errors.Kind("User with specified UUID does not exists")
	PaymentUserUuidValidationErr = errors.Kind("User UUID is not valid")
	PaymentUuidValidationErr 	 = errors.Kind("Payment UUID is not valid")
	PaymentPriceValidationErr	 = errors.Kind("Price cannot be less or equal zero")
)

func GetHttpError(err error) int {
	result := errors.GetHttpError(err)

	notFoundErrors := []errors.Kind{
		PaymentNotFoundErr,
		UserNotFoundErr,
	}

	badRequestErrors := []errors.Kind{
		PaymentExistsErr,
		PaymentUserUuidValidationErr,
		PaymentUuidValidationErr,
		PaymentPriceValidationErr,
	}

	internalError := []errors.Kind{
		RepositoryPaymentErr,
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
