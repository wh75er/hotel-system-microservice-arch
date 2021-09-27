package hotel_service

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type HotelRepository struct {
	Db     *sqlx.DB
	logger logs.LoggerInterface
}

func NewHotelRepository(db *sqlx.DB, logger logs.LoggerInterface) models.HotelRepositoryI {
	return &HotelRepository{db, logger}
}

func (r *HotelRepository) GetHotel(hotelUuid uuid.UUID) (h models.Hotel, e error) {
	var opError errors.Op = "postgres.GetHotel"

	err := r.Db.QueryRowx("SELECT name, hotelUuid, photos, description, country, city, address, "+
		"isReady FROM hotels WHERE hotelUuid = $1", hotelUuid).Scan(
		&h.Name,
		&h.HotelUuid,
		pq.Array(&h.Photos),
		&h.Description,
		&h.Country,
		&h.City,
		&h.Address,
		&h.IsReady,
	)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	} else if err == sql.ErrNoRows {
		e = errors.E(opError, errors.RepositoryNoRows, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	}

	return
}

func (r *HotelRepository) GetHotels() (h []models.Hotel, e error) {
	var opError errors.Op = "postgres.GetHotels"

	rows, err := r.Db.Queryx("SELECT name, hotelUuid, photos, description, country, city, address, " +
		"isReady FROM hotels")
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

	for rows.Next() {
		var v models.Hotel
		err = rows.Scan(
			&v.Name,
			&v.HotelUuid,
			pq.Array(&v.Photos),
			&v.Description,
			&v.Country,
			&v.City,
			&v.Address,
			&v.IsReady,
		)
		if err != nil {
			e = errors.E(opError, errors.RepositoryQueryErr, err)
			r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
			return
		}
		h = append(h, v)
	}

	return
}

func (r *HotelRepository) AddHotel(h *models.Hotel) (e error) {
	var opError errors.Op = "postgres.AddHotel"

	_, err := r.Db.Exec("INSERT INTO "+
		"hotels(name, hotelUuid, photos, description, country, city, address, isReady, creationDate) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		h.Name, h.HotelUuid, pq.Array(h.Photos), h.Description, h.Country, h.City, h.Address, h.IsReady, h.CreationDate)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	}

	return
}

func (r *HotelRepository) PatchHotel(h *models.Hotel) (e error) {
	var opError errors.Op = "postgres.PatchHotel"

	_, err := r.Db.Exec("UPDATE hotels SET name = $1, photos = $2, description = $3, country = $4, city = $5, "+
		"address = $6, isReady = $7 WHERE hotelUuid = $8",
		h.Name, pq.Array(h.Photos), h.Description, h.Country, h.City, h.Address, h.IsReady, h.HotelUuid)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	}

	return
}

func (r *HotelRepository) DeleteHotel(hotelUuid uuid.UUID) (e error) {
	var opError errors.Op = "postgres.DeleteHotel"

	_, err := r.Db.Exec("DELETE FROM hotels WHERE hotelUuid = $1", hotelUuid)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	}

	return
}
