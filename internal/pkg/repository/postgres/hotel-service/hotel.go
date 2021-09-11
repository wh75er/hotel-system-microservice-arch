package hotel_service

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/models"
)

type HotelRepository struct {
	Db *sqlx.DB
	Logger *logrus.Logger
}

func NewHotelRepository(db *sqlx.DB, logger *logrus.Logger) models.HotelRepositoryI {
	return &HotelRepository{db, logger }
}

func (r *HotelRepository) GetHotel(hotelUuid uuid.UUID) (h models.Hotel, e error) {
	var opError errors.Op = "postgres.GetHotel"

	e = r.Db.Get(&h, "SELECT name, hotelUuid, photos, description, country, city, " +
		"address, isReady FROM hotels WHERE hotelUuid = $1", hotelUuid)
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr)
		r.Logger.Error("Database error: ", e)
	}

	return
}

func (r *HotelRepository) GetHotels() (h []models.Hotel, e error) {
	var opError errors.Op = "postgres.GetHotels"

	e = r.Db.Select(&h, "SELECT name, hotelUuid, photos, description, country, city, address, " +
		"isReady FROM hotels")
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr)
		r.Logger.Error("Database error: ", e)
	}

	return
}

func (r *HotelRepository) AddHotel(h *models.Hotel) (e error) {
	var opError errors.Op = "postgres.AddHotel"

	_, e = r.Db.Exec("INSERT INTO " +
		"hotels(name, hotelUuid, photos, description, country, city, address, isReady, creationDate) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		h.Name, h.HotelUuid, h.Photos, h.Description, h.Country, h.City, h.Address, h.IsReady, h.CreationDate)
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, e)
		r.Logger.Error("Database error: ", e)
	}

	return
}

func (r *HotelRepository) PatchHotel(h *models.Hotel) (e error) {
	var opError errors.Op = "postgres.PatchHotel"

	_, e = r.Db.Exec("UPDATE hotels SET name = $1, photos = $2, description = $3, country = $4, city = $5, " +
		"address = $6, isReady = $7 WHERE hotelUuid = $8",
		h.Name, h.Photos, h.Description, h.Country, h.City, h.Address, h.IsReady, h.HotelUuid)
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, e)
		r.Logger.Error("Database error: ", e)
	}

	return
}

func (r *HotelRepository) DeleteHotel(hotelUuid uuid.UUID) (e error) {
	var opError errors.Op = "postgres.DeleteHotel"

	_, e = r.Db.Exec("DELETE FROM hotels WHERE hotelUuid = $1", hotelUuid)
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, e)
		r.Logger.Error("Database error: ", e)
	}

	return
}
