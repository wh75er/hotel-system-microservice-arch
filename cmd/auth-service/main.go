package main

import (
	"hotel-booking-system/internal/app/auth-service"
)

func main() {
	a := auth_service.New()
	a.Run("develop.toml")
}
