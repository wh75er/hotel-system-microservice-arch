package payment_service

import (
	"github.com/google/uuid"
	"hotel-booking-system/internal/pkg/errors"
	kinds "hotel-booking-system/internal/pkg/errors/payment-service"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

type PaymentUsecase struct {
	PaymentRepository models.PaymentRepositoryI
	Logger            logs.LoggerInterface
}

func NewPaymentUsecase(
	paymentR models.PaymentRepositoryI,
	logger logs.LoggerInterface,
) models.PaymentUsecaseI {
	return &PaymentUsecase{paymentR, logger}
}

func (u *PaymentUsecase) CreatePayment(price int, userUuid string) (paymentUuid uuid.UUID, e error) {
	var opError errors.Op = "usecase.CreatePayment"

	validUserUuid, err := uuid.Parse(userUuid)
	if err != nil {
		e = errors.E(opError, kinds.PaymentUserUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	// TODO: call to user Service(check that this user exists)

	paymentUuid = uuid.New()

	p := models.Payment{
		PaymentUuid: paymentUuid,
		UserUuid:    validUserUuid,
		Status:      models.NewPaymentStatus,
		Price:       price,
		TimeUpdated: time.Now(),
	}

	err = p.ValidatePrice()
	if err != nil {
		e = err
		u.Logger.Error("Usecase error: ", e)
		return
	}

	err = u.PaymentRepository.AddPayment(&p)
	if err != nil {
		e = errors.E(opError, kinds.RepositoryPaymentErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *PaymentUsecase) MakePayment(paymentUuid string) (e error) {
	var opError errors.Op = "usecase.MakePayment"

	validPaymentUuid, err := uuid.Parse(paymentUuid)
	if err != nil {
		e = errors.E(opError, kinds.PaymentUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	p, err := u.PaymentRepository.GetPayment(validPaymentUuid)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = errors.E(opError, kinds.PaymentNotFoundErr, err)
			u.Logger.Error("Usecase error: %v", e)
			return
		}
		e = errors.E(opError, kinds.RepositoryPaymentErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	p.TimeUpdated = time.Now()
	p.Status = models.PaidPaymentStatus

	err = u.PaymentRepository.ChangePaymentStatus(p)
	if err != nil {
		e = errors.E(opError, kinds.RepositoryPaymentErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	// TODO: make call to UserLoyaltyService to update user's discount

	return
}

func (u *PaymentUsecase) GetPayment(paymentUuid string) (p *models.Payment, e error) {
	var opError errors.Op = "usecase.GetPayment"

	validPaymentUuid, err := uuid.Parse(paymentUuid)
	if err != nil {
		e = errors.E(opError, kinds.PaymentUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	p, err = u.PaymentRepository.GetPayment(validPaymentUuid)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = errors.E(opError, kinds.PaymentNotFoundErr, err)
			u.Logger.Error("Usecase error: %v", e)
			return
		}
		e = errors.E(opError, kinds.RepositoryPaymentErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}
