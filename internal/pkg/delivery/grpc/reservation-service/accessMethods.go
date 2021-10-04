package reservation_service

import (
	"hotel-booking-system/internal/pkg/models"
)

func AccessibleReservationServicePaths() map[string][]models.Role {
	const reservationServicePath = "/proto.ReservationService/"

	return map[string][]models.Role{
		reservationServicePath + "AddReservation":        {models.SERVICE},
		reservationServicePath + "CancelReservation":     {models.SERVICE},
		reservationServicePath + "GetReservation":        {models.SERVICE},
		reservationServicePath + "GetReservationsByUser": {models.SERVICE},
		reservationServicePath + "CreateReservation":     {models.SERVICE},
	}
}
