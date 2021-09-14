package models

import (
	"github.com/google/uuid"
	"time"
)

type Room struct {
	RoomType     string
	Amount       int
	Beds         int
	HotelUuid    uuid.UUID
	RoomUuid     uuid.UUID
	CreationDate time.Time
	Offers       []string
	NightPrice   int
}

type RoomRepositoryI interface {
	GetRoom(roomUuid uuid.UUID) (room Room, e error)
	GetRooms(hotelUuid uuid.UUID) (rooms []Room, e error)
	AddRoom(room *Room) (e error)
	PatchRoom(room *Room) (e error)
	DeleteRoom(roomUuid uuid.UUID) (e error)
}

type RoomUsecaseI interface {
	GetRooms(hotelUuid string) (r []Room, e error)
	AddRoom(r *Room) (e error)
	PatchRoom(r *Room) (e error)
	DeleteRoom(roomUuid string) (e error)
}
