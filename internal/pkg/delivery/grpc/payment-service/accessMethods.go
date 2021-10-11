package payment_service

import "hotel-booking-system/internal/pkg/models"

func AccessiblePaymentServicePaths() map[string][]models.Role {
	const paymentServicePath = "/proto.PaymentService/"

	return map[string][]models.Role{
		paymentServicePath + "GetPayment":    {models.SERVICE},
		paymentServicePath + "CreatePayment": {models.SERVICE},
	}
}
