package hotel_service

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/models"
)

type RoomRepository struct {
	Db *sqlx.DB
	Logger *logrus.Logger
}

func NewRoomRepository(db *sqlx.DB, logger *logrus.Logger) models.RoomRepositoryI {
	return &RoomRepository{db, logger }
}

func (r *RoomRepository) GetRoom(roomUuid uuid.UUID) (room models.Room, e error) {
	var opError errors.Op = "postgres.GetRoom"

	e = r.Db.Select(&room, "SELECT roomType, amount, beds, hotelUuid, roomUuid, creationDate, " +
		"offers, nightPrice WHERE roomUuid = $1", roomUuid)
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

func (r *RoomRepository) GetRooms(hotelUuid uuid.UUID) (rooms []models.Room, e error) {
	var opError errors.Op = "postgres.GetRooms"

	e = r.Db.Select(&rooms, "SELECT roomType, amount, beds, hotelUuid, roomUuid, creationDate, " +
		"offers, nightPrice WHERE hotelUuid = $1", hotelUuid)
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

func (r *RoomRepository) AddRoom(room *models.Room)  (e error) {
	var opError errors.Op = "postgres.AddRoom"

	_, e = r.Db.Exec("INSERT INTO " +
		"rooms(roomType, amount, beds, hotelUuid, roomUuid, creationDate, offers, nightPrice) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		room.RoomType, room.Amount, room.Beds, room.HotelUuid, room.RoomUuid, room.CreationDate, room.Offers, room.NightPrice)
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, e)
		r.Logger.Error("Database error: ", e)
	}

	return
}

func (r *RoomRepository) PatchRoom(room *models.Room) (e error) {
	var opError errors.Op = "postgres.PatchRoom"

	_, e = r.Db.Exec("UPDATE rooms SET roomType = $1, Amount = $2, Beds = $3, Offers = $4, NightPrice = $5 " +
		"WHERE roomUuid = $6", room.RoomType, room.Amount, room.Beds, room.Offers, room.NightPrice, room.RoomUuid)
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, e)
		r.Logger.Error("Database error: ", e)
	}

	return
}

func (r *RoomRepository) DeleteRoom(roomUuid uuid.UUID) (e error) {
	var opError errors.Op = "postgres.DeleteRoom"

	_, e = r.Db.Exec("DELETE FROM rooms WHERE roomUuid = $1", roomUuid)
	if e == sql.ErrConnDone {
		e = errors.E(opError, errors.RepositoryDownErr, e)
		r.Logger.Error("Database error: ", e)
	} else if e != nil {
		e = errors.E(opError, errors.RepositoryQueryErr, e)
		r.Logger.Error("Database error: ", e)
	}

	return
}
