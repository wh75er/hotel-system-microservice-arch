package models

import (
	"github.com/google/uuid"
	"time"
)

type Hotel struct {
	Name         string
	HotelUuid    uuid.UUID
	Photos       []uuid.UUID
	Description  string
	Country      string
	City         string
	Address      string
	IsReady      bool
	CreationDate time.Time
	Reviews      []Review
	Rooms        []Room
}

type HotelRepositoryI interface {
	GetHotel(hotelUuid uuid.UUID) (h Hotel, e error)
	GetHotels() (h []Hotel, e error)
	AddHotel(h *Hotel) (e error)
	PatchHotel(h *Hotel) (e error)
	DeleteHotel(hotelUuid uuid.UUID) (e error)
}

type HotelUsecaseI interface {
	GetHotel(hotelUuid string) (h Hotel, e error)
	GetHotels() (h []Hotel, e error)
	AddHotel(h *Hotel) (e error)
	PatchHotel(h *Hotel) (e error)
	DeleteHotel(hotelUuid string) (e error)
}
