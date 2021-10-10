package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"hotel-booking-system/internal/pkg/errors"
	"regexp"
)

const (
	PasswordMinLength = 10
	PasswordMaxLength = 128

	LoginMinLength = 6
	LoginMaxLength = 24

	HashCost = 14
)

type User struct {
	UserUuid uuid.UUID `json:"userUuid,"`
	Login    string    `json:"login,"`
	Password string    `json:"password,omitempty"`
	Role     string    `json:"role,omitempty"`
}

type UserUsecaseI interface {
	AddUser(user *User) (e error)
	AddAdmin(user *User) (e error)
	GetUser(uid string) (user *User, e error)
	Login(user *User) (authToken string, e error)
	CheckAuth(jwtToken string) (role Role, e error)
}

type UserRepositoryI interface {
	AddUser(user *User) (e error)
	GetUserByUuid(uid uuid.UUID) (user User, e error)
	GetUserByLogin(login string) (user User, e error)
}

func (u *User) HashPassword() (e error) {
	var opError errors.Op = "models.HashPassword"

	if len(u.Password) < PasswordMinLength || len(u.Password) > PasswordMaxLength {
		e = errors.E(opError, errors.UserPasswordLengthValidationError)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), HashCost)
	if err != nil {
		return errors.E(opError, errors.FailedToHashPassword, err)
	}

	u.Password = string(hashedPass)

	return
}

func (u *User) CompareWithPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func (u *User) ValidatePassword() error {
	var opError errors.Op = "models.validatePassword"

	if len(u.Password) < PasswordMinLength || len(u.Password) > PasswordMaxLength {
		return errors.E(opError, errors.UserPasswordLengthValidationError)
	}

	validatorRegex := regexp.MustCompile(`^[a-zA-Z0-9./?,'\[\]\\!@#$%^&*()]+$`)
	if match := validatorRegex.MatchString(u.Password); match == false {
		return errors.E(opError, errors.UserPasswordCharsValidationError)
	}

	return nil
}

func (u *User) ValidateLogin() error {
	var opError errors.Op = "models.ValidateLogin"

	if len(u.Login) < LoginMinLength || len(u.Login) > LoginMaxLength {
		return errors.E(opError, errors.UserLoginLengthValidationError)
	}

	validatorRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if match := validatorRegex.MatchString(u.Login); match == false {
		return errors.E(opError, errors.UserLoginCharsValidationError)
	}

	return nil
}
