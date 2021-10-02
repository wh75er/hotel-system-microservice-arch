package errors

const (
	RepositoryLoyaltyErr = Kind("Something wrong with loyalty repository")
)

const (
	LoyaltyExistsErr             = Kind("User's loyalty exists")
	LoyaltyNotFoundErr           = Kind("Loyalty for current user not found")
	LoyaltyUserUuidValidationErr = Kind("User UUID is not valid")
)
