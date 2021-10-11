package main

import (
	gatewayService "hotel-booking-system/internal/app/gateway-service"
)

func main() {
	a := gatewayService.New()
	a.Run("develop.toml")
}
