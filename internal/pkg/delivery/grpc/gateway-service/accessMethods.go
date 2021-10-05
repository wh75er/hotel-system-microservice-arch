package gateway_service

import (
	"hotel-booking-system/internal/pkg/models"
)

func AccessibleGatewayServicePaths() map[string][]models.Role {
	const gatewayServicePath = "/proto.GatewayService/"

	return map[string][]models.Role{
		gatewayServicePath + "AddReservation":        {models.USER, models.ADMIN},
		gatewayServicePath + "CancelReservation":     {models.USER, models.ADMIN},
		gatewayServicePath + "GetReservation":        {models.USER, models.ADMIN},
		gatewayServicePath + "GetReservationsByUser": {models.USER, models.ADMIN},
		gatewayServicePath + "CreatePayment":         {models.USER, models.ADMIN},
		gatewayServicePath + "AddHotel":              {models.ADMIN},
		gatewayServicePath + "GetHotel":              {models.USER, models.ADMIN},
		gatewayServicePath + "GetHotels":             {models.USER, models.ADMIN},
		gatewayServicePath + "PatchHotel":            {models.ADMIN},
		gatewayServicePath + "DeleteHotel":           {models.ADMIN},
		gatewayServicePath + "AddRoom":               {models.ADMIN},
		gatewayServicePath + "GetRooms":              {models.USER, models.ADMIN},
		gatewayServicePath + "GetRoom":               {models.USER, models.ADMIN},
		gatewayServicePath + "PatchRoom":             {models.ADMIN},
		gatewayServicePath + "DeleteRoom":            {models.ADMIN},
		gatewayServicePath + "GetDiscount":           {models.USER, models.ADMIN},
		gatewayServicePath + "GetPayment":            {models.USER, models.ADMIN},
	}
}
