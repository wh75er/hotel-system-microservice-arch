package interceptors

import "hotel-booking-system/internal/pkg/models"

func MethodsRoleMapToSet(src map[string][]models.Role) map[string]bool {
	set := make(map[string]bool)

	for k, _ := range src {
		set[k] = true
	}

	return set
}
