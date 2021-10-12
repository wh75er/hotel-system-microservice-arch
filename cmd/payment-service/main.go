package main

import (
	payment_service "hotel-booking-system/internal/app/payment-service"
)

func main() {
	a := payment_service.New()
	a.Run("develop.toml")
}
