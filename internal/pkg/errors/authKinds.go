package errors

const (
	RepositoryUserErr = Kind("Something wrong with user repository")
)

const (
	UserExistsErr                     = Kind("User exists cannot add new user")
	UserNotFoundErr                   = Kind("User not found")
	UserUuidValidationErr             = Kind("User UUID is not valid")
	UserPasswordLengthValidationError = Kind("Password length should be less than 128 characters")
	UserPasswordCharsValidationError  = Kind("The password must consist of latin chars with optional numbers or special characters ./\\?,'[]!@#$%^&*()")
	UserLoginLengthValidationError    = Kind("Login should be more than 6 symbols and less than 24 characters")
	UserLoginCharsValidationError     = Kind("Login must consist of latin characters with optional numbers")
	UserLoginUniqueValidationError    = Kind("Login should be unique")
)

const (
	AuthorizationErr = Kind("Invalid login or password")
)

const (
	FailedToHashPassword = Kind("Failed to hash password error")
)
