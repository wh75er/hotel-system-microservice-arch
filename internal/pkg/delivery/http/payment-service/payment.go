package payment_service

import (
	"hotel-booking-system/internal/pkg/errors"
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
		return
	}

	err := req.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	paymentuUid := req.Form.Get("label")
	err = d.paymentUsecase.MakePayment(paymentuUid)
	if err != nil {
		w.WriteHeader(errors.GetHttpError(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
