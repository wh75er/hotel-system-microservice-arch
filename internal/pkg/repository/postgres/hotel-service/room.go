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

type RoomRepository struct {
	Db     *sqlx.DB
	logger logs.LoggerInterface
}

func NewRoomRepository(db *sqlx.DB, logger logs.LoggerInterface) models.RoomRepositoryI {
	return &RoomRepository{db, logger}
}

func (r *RoomRepository) GetRoom(roomUuid uuid.UUID) (room models.Room, e error) {
	var opError errors.Op = "postgres.GetRoom"

	err := r.Db.QueryRowx("SELECT roomType, amount, beds, hotelUuid, roomUuid, creationDate, "+
		"offers, nightPrice FROM rooms WHERE roomUuid = $1", roomUuid).Scan(
		&room.RoomType,
		&room.Amount,
		&room.Beds,
		&room.HotelUuid,
		&room.RoomUuid,
		&room.CreationDate,
		pq.Array(&room.Offers),
		&room.NightPrice,
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

func (r *RoomRepository) GetRooms(hotelUuid uuid.UUID) (rooms []models.Room, e error) {
	var opError errors.Op = "postgres.GetRooms"

	rows, err := r.Db.Queryx("SELECT roomType, amount, beds, hotelUuid, roomUuid, creationDate, "+
		"offers, nightPrice FROM rooms WHERE hotelUuid = $1", hotelUuid)
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
		var v models.Room
		err = rows.Scan(
			&v.RoomType,
			&v.Amount,
			&v.Beds,
			&v.HotelUuid,
			&v.RoomUuid,
			&v.CreationDate,
			pq.Array(&v.Offers),
			&v.NightPrice,
		)
		if err != nil {
			e = errors.E(opError, errors.RepositoryQueryErr, err)
			r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
			return
		}
		rooms = append(rooms, v)
	}

	return
}

func (r *RoomRepository) AddRoom(room *models.Room) (e error) {
	var opError errors.Op = "postgres.AddRoom"

	_, err := r.Db.Exec("INSERT INTO "+
		"rooms(roomType, amount, beds, hotelUuid, roomUuid, creationDate, offers, nightPrice) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		room.RoomType,
		room.Amount,
		room.Beds,
		room.HotelUuid,
		room.RoomUuid,
		room.CreationDate,
		pq.Array(room.Offers),
		room.NightPrice,
	)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	}

	return
}

func (r *RoomRepository) PatchRoom(room *models.Room) (e error) {
	var opError errors.Op = "postgres.PatchRoom"

	_, err := r.Db.Exec("UPDATE rooms SET roomType = $1, Amount = $2, Beds = $3, Offers = $4, NightPrice = $5 "+
		"WHERE roomUuid = $6",
		room.RoomType,
		room.Amount,
		room.Beds,
		pq.Array(room.Offers),
		room.NightPrice,
		room.RoomUuid,
	)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	}

	return
}

func (r *RoomRepository) DeleteRoom(roomUuid uuid.UUID) (e error) {
	var opError errors.Op = "postgres.DeleteRoom"

	_, err := r.Db.Exec("DELETE FROM rooms WHERE roomUuid = $1", roomUuid)
	if err == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	} else if err != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, err)
		r.logger.Errorf("Database error: %v - %v", e, errors.SourceDetails(e))
	}

	return
}
