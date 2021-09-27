package auth_service

import "hotel-booking-system/internal/pkg/models"

func AccessibleAuthServicePaths() map[string][]models.Role {
	const authServicePath = "/proto.AuthService/"

	return map[string][]models.Role{
		authServicePath + "GetUser": {models.SERVICE, models.ADMIN},
	}
}
