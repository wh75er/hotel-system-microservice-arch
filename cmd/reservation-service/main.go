package main

import (
	reservationService "hotel-booking-system/internal/app/reservation-service"
)

func main() {
	a := reservationService.New()
	a.Run("release.toml")
}
