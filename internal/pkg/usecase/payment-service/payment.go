package payment_service

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	users_proto "hotel-booking-system/internal/pkg/delivery/grpc/auth-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	loyalty_proto "hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

type PaymentUsecase struct {
	PaymentRepository        models.PaymentRepositoryI
	UserServiceClient        users_proto.AuthServiceClient
	UserLoyaltyServiceClient loyalty_proto.LoyaltyServiceClient
	Logger                   logs.LoggerInterface
}

func NewPaymentUsecase(
	paymentR models.PaymentRepositoryI,
	userClient users_proto.AuthServiceClient,
	userLoyaltyClient loyalty_proto.LoyaltyServiceClient,
	logger logs.LoggerInterface,
) models.PaymentUsecaseI {
	return &PaymentUsecase{
		paymentR,
		userClient,
		userLoyaltyClient,
		logger,
	}
}

func (u *PaymentUsecase) CreatePayment(price int, userUuid string) (paymentUuid uuid.UUID, e error) {
	var opError errors.Op = "usecase.CreatePayment"

	validUserUuid, err := uuid.Parse(userUuid)
	if err != nil {
		e = errors.E(opError, errors.PaymentUserUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	// call to user Service(check that this user exists)
	_, err = u.UserServiceClient.GetUser(context.Background(), &commonProto.UUID{Value: validUserUuid.String()})
	if err != nil {
		if status.Code(err) < errors.MaxGrpcCodeValue {
			e = errors.E(opError, errors.AuthServiceUnavailable, err)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		serviceKind := errors.Kind(status.Code(err))
		e = errors.E(opError, serviceKind)
		return
	}

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
		e = errors.E(opError, errors.RepositoryPaymentErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *PaymentUsecase) MakePayment(paymentUuid string) (e error) {
	var opError errors.Op = "usecase.MakePayment"

	validPaymentUuid, err := uuid.Parse(paymentUuid)
	if err != nil {
		e = errors.E(opError, errors.PaymentUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	p, err := u.PaymentRepository.GetPayment(validPaymentUuid)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.PaymentNotFoundErr, err)
			u.Logger.Error("Usecase error: %v", e)
			return
		}
		e = errors.E(opError, errors.RepositoryPaymentErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	p.TimeUpdated = time.Now()
	p.Status = models.PaidPaymentStatus

	err = u.PaymentRepository.ChangePaymentStatus(p)
	if err != nil {
		e = errors.E(opError, errors.RepositoryPaymentErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	// make call to UserLoyaltyService to update user's discount
	_, err = u.UserLoyaltyServiceClient.UpdateDiscount(context.Background(), &loyalty_proto.UpdateDiscountRequest{
		UserUid:      &commonProto.UUID{Value: p.UserUuid.String()},
		Contribution: int64(p.Price),
	})
	if err != nil {
		if status.Code(err) < errors.MaxGrpcCodeValue {
			e = errors.E(opError, errors.UserLoyaltyServiceUnavailable, err)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		serviceKind := errors.Kind(status.Code(err))
		if serviceKind == errors.LoyaltyNotFoundErr {
			u.Logger.Warnf("loyalty discount account not found for user[uuid]: %v", p.UserUuid)
		} else {
			u.Logger.Warnf("undesirable behaviour on loyalty service: %v", serviceKind)
		}

		// Not valuable service, don't interrupt, only log
		e = nil
	}

	return
}

func (u *PaymentUsecase) GetPayment(paymentUuid string) (p *models.Payment, e error) {
	var opError errors.Op = "usecase.GetPayment"

	validPaymentUuid, err := uuid.Parse(paymentUuid)
	if err != nil {
		e = errors.E(opError, errors.PaymentUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	p, err = u.PaymentRepository.GetPayment(validPaymentUuid)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.PaymentNotFoundErr, err)
			u.Logger.Error("Usecase error: %v", e)
			return
		}
		e = errors.E(opError, errors.RepositoryPaymentErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}
