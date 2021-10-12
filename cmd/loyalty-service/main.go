package main

import (
	"hotel-booking-system/internal/app/loyalty-service"
)

func main() {
	a := loyalty_service.New()
	a.Run("release.toml")
}
