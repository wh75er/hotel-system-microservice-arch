package main

import (
	"hotel-booking-system/internal/app/stat-service"
)

func main() {
	a := stat_service.New()
	a.Run("develop.toml")
}
