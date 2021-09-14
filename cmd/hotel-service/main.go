package main

import (
	"hotel-booking-system/internal/app/hotel-service"
)

func main() {
	a := hotel_service.New()
	a.Run("develop.toml")
}
