package auth_service

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type UserRepository struct {
	Db     *sqlx.DB
	logger logs.LoggerInterface
}

func NewUserRepository(db *sqlx.DB, logger logs.LoggerInterface) models.UserRepositoryI {
	return &UserRepository{db, logger}
}

func (r *UserRepository) AddUser(user *models.User) (e error) {
	var opError errors.Op = "postgres.AddUser"

	_, err := r.Db.Exec("INSERT INTO "+
		"users(userUuid, login, password, role) VALUES ($1, $2, $3, $4)",
		user.UserUuid, user.Login, user.Password, user.Role)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	}

	return
}

func (r *UserRepository) GetUserByUuid(uid uuid.UUID) (user models.User, e error) {
	var opError errors.Op = "postgres.GetUser"

	err := r.Db.Get(&user, "SELECT userUuid, login, password, role FROM users WHERE userUuid = $1", uid)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	}

	return
}

func (r *UserRepository) GetUserByLogin(login string) (user models.User, e error) {
	var opError errors.Op = "postgres.GetUser"

	err := r.Db.Get(&user, "SELECT userUuid, login, password, role FROM users WHERE login = $1", login)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
		return
	}

	return
}
