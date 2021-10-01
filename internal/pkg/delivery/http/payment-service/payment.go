package payment_service

import (
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
	"net/http"
)

type PaymentHttp struct {
	paymentUsecase models.PaymentUsecaseI
	logger         logs.LoggerInterface
}

func SetPaymentHttpRoutes(router *http.ServeMux, paymentU models.PaymentUsecaseI, logger logs.LoggerInterface) {
	deliveryInstance := PaymentHttp{
		paymentUsecase: paymentU,
		logger:         logger,
	}

	router.HandleFunc("/api/v1/pay", deliveryInstance.makePaymentHandler)
}

func (d *PaymentHttp) makePaymentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	w.WriteHeader(http.StatusOK)
}
