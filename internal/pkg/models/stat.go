package models

type Stat struct {
	RoomsAmount int64
	ReservationsAmount int64
}

type StatUsecaseI interface {
	GetStat() (s Stat, e error)
	UpdateRoomsAmount(delta int64) (e error)
	UpdateReservationsAmount(delta int64) (e error)
}

type StatRepositoryI interface {
	GetStat() (s Stat, e error)
	UpdateStat(s Stat) (e error)
}
