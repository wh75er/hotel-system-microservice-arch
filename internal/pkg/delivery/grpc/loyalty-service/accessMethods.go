package loyalty_service

import "hotel-booking-system/internal/pkg/models"

func AccessibleLoyaltyServicePaths() map[string][]models.Role {
	const authServicePath = "/proto.LoyaltyService/"

	return map[string][]models.Role{
		authServicePath + "GetDiscount":    {models.SERVICE},
		authServicePath + "AddUser":        {models.SERVICE},
		authServicePath + "UpdateDiscount": {models.SERVICE},
	}
}
