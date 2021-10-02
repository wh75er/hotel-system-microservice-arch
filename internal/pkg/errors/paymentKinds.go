package errors

const (
	RepositoryPaymentErr = Kind("Something wrong with payment repository")
)

const (
	PaymentExistsErr             = Kind("Payment exists")
	PaymentNotFoundErr           = Kind("Payment not found")
	PaymentUserUuidValidationErr = Kind("User UUID is not valid")
	PaymentUuidValidationErr     = Kind("Payment UUID is not valid")
	PaymentPriceValidationErr    = Kind("Price cannot be less or equal zero")
)
