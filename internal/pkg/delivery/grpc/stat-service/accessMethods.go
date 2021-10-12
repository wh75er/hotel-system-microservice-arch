package stat_service

import "hotel-booking-system/internal/pkg/models"

func AccessibleStatServicePaths() map[string][]models.Role {
	const statServicePath = "/proto.StatService/"

	return map[string][]models.Role{
		statServicePath + "GetStat":    {models.SERVICE},
		statServicePath + "UpdateRoomsAmount":        {models.SERVICE},
		statServicePath + "UpdateReservationsAmount": {models.SERVICE},
	}
}
