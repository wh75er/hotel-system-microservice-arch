package errors

import "net/http"

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

func GetHttpError(err error) int {
	badRequestErrors := []Kind{
		JWTVerificationSigningMethodErr,
		JWTVerificationErr,
		JWTTokenClaimsErr,
		JWTSigningErr,
	}

	UnauthorizedErrors := []Kind{
		InvalidCredentials,
	}

	Forbidden := []Kind{
		ExpiredToken,
		PermissionDenied,
	}

	internalError := []Kind{
		JWTGenerationErr,
		RepositoryDownErr,
		RepositoryQueryErr,
		RepositoryNoRows,
		UnexpectedErr,
	}

	kind := GetKind(err)

	if Contains(badRequestErrors, kind) {
		return http.StatusBadRequest
	}

	if Contains(internalError, kind) {
		return http.StatusInternalServerError
	}

	if Contains(Forbidden, kind) {
		return http.StatusForbidden
	}

	if Contains(UnauthorizedErrors, kind) {
		return http.StatusUnauthorized
	}

	return http.StatusInternalServerError
}
