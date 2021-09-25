package jwt_manager

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

type JWTManager struct {
	secret        string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	Role models.Role `json:"role"`
}

func NewJWTManager(secret string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secret, tokenDuration}
}

func (m *JWTManager) Generate(r models.Role) (models.Token, error) {
	var opError errors.Op = "jwt-manager.Generate"

	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.tokenDuration).Unix(),
		},
		Role: r,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(m.secret))
	if err != nil {
		return "", errors.E(opError, errors.JWTSigningErr, err)
	}

	tokenType := models.Token(signedToken)

	return tokenType, nil
}

func (m *JWTManager) Verify(tokenString string) (*UserClaims, error) {
	var opError errors.Op = "jwt-manager.Verify"

	token, err := jwt.ParseWithClaims(
		tokenString,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.E(opError, errors.JWTVerificationSigningMethodErr)
			}

			return []byte(m.secret), nil
		},
	)
	if err != nil {
		return nil, errors.E(opError, errors.JWTVerificationErr, err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.E(opError, errors.JWTTokenClaimsErr, err)
	}

	return claims, nil
}

// UntilExp Return Duration until token expiration date
func (m *JWTManager) UntilExp(tokenString string) (time.Duration, error) {
	var opError errors.Op = "jwt-manager.TimeLeft"

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, &UserClaims{})
	if err != nil {
		return 0, errors.E(opError, errors.JWTVerificationErr, err)
	}

	userClaims, ok := token.Claims.(*UserClaims)
	if !ok {
		return 0, errors.E(opError, errors.JWTVerificationErr, fmt.Errorf("failed to extract payload from token"))
	}

	expTime := userClaims.StandardClaims.ExpiresAt
	return time.Until(time.Unix(expTime, 0)), nil
}
