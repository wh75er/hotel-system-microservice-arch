package models

import (
	"github.com/google/uuid"
	"hotel-booking-system/internal/pkg/errors"
	"time"
)

type PaymentStatus string

const (
	NewPaymentStatus      = "New"
	PaidPaymentStatus     = "Paid"
	CanceledPaymentStatus = "Canceled"
)

type Payment struct {
	PaymentUuid uuid.UUID
	UserUuid    uuid.UUID
	Status      PaymentStatus
	Price       int
	TimeUpdated time.Time
}

type PaymentRepositoryI interface {
	AddPayment(p *Payment) (e error)
	ChangePaymentStatus(p *Payment) (e error)
	GetPayment(paymentUuid uuid.UUID) (p *Payment, e error)
}

type PaymentUsecaseI interface {
	CreatePayment(price int, userUuid string) (paymentUid uuid.UUID, e error)
	MakePayment(paymentUuid string) (e error)
	GetPayment(paymentUuid string) (p *Payment, e error)
}

func (p *Payment) ValidatePrice() error {
	if p.Price <= 0 {
		return errors.E(errors.PaymentPriceValidationErr)
	}

	return nil
}
