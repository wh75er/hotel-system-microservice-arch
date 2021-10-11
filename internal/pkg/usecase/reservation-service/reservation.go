package reservation_service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	users_proto "hotel-booking-system/internal/pkg/delivery/grpc/auth-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	hotel_proto "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	loyalty_proto "hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service/proto"
	payment_proto "hotel-booking-system/internal/pkg/delivery/grpc/payment-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type ReservationUsecase struct {
	ReservationRepository    models.ReservationRepositoryI
	HotelServiceClient       hotel_proto.HotelServiceClient
	PaymentServiceClient     payment_proto.PaymentServiceClient
	UserServiceClient        users_proto.AuthServiceClient
	UserLoyaltyServiceClient loyalty_proto.LoyaltyServiceClient
	Logger                   logs.LoggerInterface
}

func NewReservationUsecase(
	reservationR models.ReservationRepositoryI,
	hotelClient hotel_proto.HotelServiceClient,
	paymentClient payment_proto.PaymentServiceClient,
	userClient users_proto.AuthServiceClient,
	userLoyaltyClient loyalty_proto.LoyaltyServiceClient,
	logger logs.LoggerInterface,
) models.ReservationUsecaseI {
	return &ReservationUsecase{
		reservationR,
		hotelClient,
		paymentClient,
		userClient,
		userLoyaltyClient,
		logger,
	}
}

func (u *ReservationUsecase) AddReservation(r *models.Reservation) (reservationUuid uuid.UUID, e error) {
	var opError errors.Op = "usecase.AddReservation"

	if err := r.ValidateDate(); err != nil {
		e = err
		u.Logger.Error("Usecase error: ", err)
		return
	}

	if r.UserUuid == uuid.Nil || r.RoomUuid == uuid.Nil {
		e = errors.E(opError, errors.ReservationCreateInvalidRequestErr)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	reservationUuid = uuid.New()

	// Check user uuid
	_, err := u.UserServiceClient.GetUser(context.Background(), &commonProto.UUID{Value: r.UserUuid.String()})
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
	// Check room uuid and that's it available
	room, err := u.HotelServiceClient.GetRoom(context.Background(), &commonProto.UUID{Value: r.RoomUuid.String()})
	if err != nil {
		if status.Code(err) < errors.MaxGrpcCodeValue {
			e = errors.E(opError, errors.HotelServiceUnavailable, err)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		serviceKind := errors.Kind(status.Code(err))
		e = errors.E(opError, serviceKind)
		return
	}
	if room.Amount < 1 {
		e = errors.E(opError, errors.RoomUnavailableErr)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	// Try to access discount service and take user's current discount in percentage
	//		If not available drop everything and create reservation with empty paymentUuid
	//		(there is no point to set price without discount)
	var loyaltyServiceFailed bool
	loyalty, err := u.UserLoyaltyServiceClient.GetDiscount(context.Background(), &commonProto.UUID{Value: r.UserUuid.String()})
	if err != nil {
		loyaltyServiceFailed = true
		if status.Code(err) < errors.MaxGrpcCodeValue {
			e = errors.E(opError, errors.UserLoyaltyServiceUnavailable, err)
			u.Logger.Error("Usecase error: ", e)
		} else {
			serviceKind := errors.Kind(status.Code(err))
			e = errors.E(opError, serviceKind)
			u.Logger.Error("Usecase error: ", e)
		}
		e = nil
	}

	// Create payment instance with considered discount price
	//		If not available drop everything and create reservation with empty paymentUuid
	//		(user will get payment instance with special request later)
	if !loyaltyServiceFailed {
		priceWithDiscount := u.calculatePriceWithDiscount(room.NightPrice, loyalty.Discount)
		paymentUuid, err := u.PaymentServiceClient.CreatePayment(
			context.Background(),
			&payment_proto.CreatePaymentRequest{
				UserUuid: &commonProto.UUID{Value: r.UserUuid.String()},
				Value:    priceWithDiscount,
			},
		)
		if err != nil {
			if status.Code(err) < errors.MaxGrpcCodeValue {
				e = errors.E(opError, errors.PaymentServiceUnavailable, err)
				u.Logger.Error("Usecase error: ", e)
			} else {
				serviceKind := errors.Kind(status.Code(err))
				e = errors.E(opError, serviceKind)
				u.Logger.Error("Usecase error: ", e)
			}
			e = nil
		}

		value := commonProto.ProtoToUuid(paymentUuid)
		validPaymentUuid, err := uuid.Parse(value)
		if err != nil {
			e = errors.E(opError, errors.PaymentUuidValidationErr, err)
			u.Logger.Error("Usecase error: ", e)
			e = nil
		}

		r.PaymentUuid = validPaymentUuid
	}

	r.ReservationUuid = reservationUuid
	r.Status = models.ActiveReservationStatus

	// Add reservation to repository
	tx, err := u.ReservationRepository.CreateReservation(r)
	if err != nil {
		e = errors.E(opError, errors.RepositoryReservationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	// Take one available room
	_, err = u.HotelServiceClient.TakeRoom(context.Background(), &commonProto.UUID{Value: r.RoomUuid.String()})
	if err != nil {
		_ = tx.Rollback()
		if status.Code(err) < errors.MaxGrpcCodeValue {
			e = errors.E(opError, errors.HotelServiceUnavailable, err)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		serviceKind := errors.Kind(status.Code(err))
		e = errors.E(opError, serviceKind)
		return
	}

	var commitFailed bool
	err = tx.Commit()
	if err != nil {
		commitFailed = true
		e = errors.E(opError, errors.RepositoryReservationErr, err)
		u.Logger.Error("Usecase error: ", e)
	}

	// If commit failed try to store room back to stock
	//   If dismiss failed need to store request to queue and try it later
	if commitFailed {
		_, _ = u.HotelServiceClient.DismissRoom(context.Background(), &commonProto.UUID{Value: r.RoomUuid.String()})
	}

	return
}

func (u *ReservationUsecase) CancelReservation(reservationUuid string) (e error) {
	var opError errors.Op = "usecase.CancelReservation"

	validReservationUuid, err := uuid.Parse(reservationUuid)
	if err != nil {
		e = errors.E(opError, errors.ReservationUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	r, err := u.ReservationRepository.GetReservation(validReservationUuid)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.ReservationNotFound, err)
			u.Logger.Error("Usecase error: %v", e)
			return
		}
		e = errors.E(opError, errors.RepositoryReservationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	r.Status = models.CanceledReservationStatus
	tx, err := u.ReservationRepository.PatchReservation(&r)
	if err != nil {
		e = errors.E(opError, errors.RepositoryReservationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	_, err = u.HotelServiceClient.DismissRoom(context.Background(), &commonProto.UUID{Value: r.RoomUuid.String()})
	if err != nil {
		_ = tx.Rollback()
		if status.Code(err) < errors.MaxGrpcCodeValue {
			e = errors.E(opError, errors.HotelServiceUnavailable, err)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		serviceKind := errors.Kind(status.Code(err))
		e = errors.E(opError, serviceKind)
		return
	}

	// Try to commit, if commit failed we're *. Need to add some queue to patch db later when it's up
	//   Taking room back from the stock is pointless, there's no syncing between patching and dismissing operations so
	// 	 someone could potentially take the dismissed room
	err = tx.Commit()
	if err != nil {
		e = errors.E(opError, errors.RepositoryReservationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *ReservationUsecase) GetReservation(reservationUuid string) (r *models.Reservation, e error) {
	var opError errors.Op = "usecase.GetReservation"

	validReservationUuid, err := uuid.Parse(reservationUuid)
	if err != nil {
		e = errors.E(opError, errors.ReservationUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	*r, err = u.ReservationRepository.GetReservation(validReservationUuid)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.ReservationNotFound, err)
			u.Logger.Error("Usecase error: %v", e)
			return
		}
		e = errors.E(opError, errors.RepositoryReservationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *ReservationUsecase) GetReservationsByUser(userUuid string) (r []models.Reservation, e error) {
	var opError errors.Op = "usecase.GetReservationsByUser"

	validUserUuid, err := uuid.Parse(userUuid)
	if err != nil {
		e = errors.E(opError, errors.UserUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	r, err = u.ReservationRepository.GetReservationsByUser(validUserUuid)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.ReservationNotFound, err)
			u.Logger.Error("Usecase error: %v", e)
			return
		}
		e = errors.E(opError, errors.RepositoryReservationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *ReservationUsecase) CreatePayment(reservationUuid string) (paymentUuid uuid.UUID, e error) {
	var opError errors.Op = "usecase.CreatePayment"

	validReservationUuid, err := uuid.Parse(reservationUuid)
	if err != nil {
		e = errors.E(opError, errors.ReservationUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	r, err := u.ReservationRepository.GetReservation(validReservationUuid)
	if err != nil {
		if errors.GetKind(err) == errors.RepositoryNoRows {
			e = errors.E(opError, errors.ReservationNotFound, err)
			u.Logger.Error("Usecase error: %v", e)
			return
		}
		e = errors.E(opError, errors.RepositoryReservationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	// If payment already exists return corresponding error
	if r.PaymentUuid != uuid.Nil {
		e = errors.E(opError, errors.ReservationPaymentExists)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	// Get reservation room
	room, err := u.HotelServiceClient.GetRoom(context.Background(), &commonProto.UUID{Value: r.RoomUuid.String()})
	if err != nil {
		if status.Code(err) < errors.MaxGrpcCodeValue {
			e = errors.E(opError, errors.HotelServiceUnavailable, err)
			u.Logger.Error("Usecase error: ", e)
			return
		}
		serviceKind := errors.Kind(status.Code(err))
		e = errors.E(opError, serviceKind)
		return
	}

	fmt.Println("RESERVATION FROM REPOSITORY: ", r)

	// Try to access discount service and take user's current discount in percentage
	//		If not available return, we're not able to calculate proper price
	//		(there is no point to set price without discount)
	loyalty, err := u.UserLoyaltyServiceClient.GetDiscount(context.Background(), &commonProto.UUID{Value: r.UserUuid.String()})
	if err != nil {
		if status.Code(err) < errors.MaxGrpcCodeValue {
			e = errors.E(opError, errors.UserLoyaltyServiceUnavailable, err)
			u.Logger.Error("Usecase error: ", e)
			e = nil
		} else {
			serviceKind := errors.Kind(status.Code(err))
			e = errors.E(opError, serviceKind)
			return
		}
	}

	// Create payment instance with considered discount price
	//		If not available return
	//		(user will get payment instance with this request later)
	priceWithDiscount := u.calculatePriceWithDiscount(room.NightPrice, loyalty.Discount)
	protoPaymentUuid, err := u.PaymentServiceClient.CreatePayment(
		context.Background(),
		&payment_proto.CreatePaymentRequest{
			UserUuid: &commonProto.UUID{Value: r.UserUuid.String()},
			Value:    priceWithDiscount,
		},
	)
	if err != nil {
		if status.Code(err) < errors.MaxGrpcCodeValue {
			e = errors.E(opError, errors.PaymentServiceUnavailable, err)
			u.Logger.Error("Usecase error: ", e)
		} else {
			serviceKind := errors.Kind(status.Code(err))
			e = errors.E(opError, serviceKind)
			u.Logger.Error("Usecase error: ", e)
		}
		e = nil
	}

	paymentUuid, err = uuid.Parse(protoPaymentUuid.Value)
	if err != nil {
		e = errors.E(opError, errors.PaymentUuidValidationErr, err)
		u.Logger.Error("Usecase error: ", e)
		e = nil
	}

	r.PaymentUuid = paymentUuid

	tx, err := u.ReservationRepository.PatchReservation(&r)
	if err != nil {
		e = errors.E(opError, errors.RepositoryReservationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	// If failed to commit we can just return
	//   to that request later. In payment service
	// 	 there will be idle useless payment instance
	//   but nothing critical
	err = tx.Commit()
	if err != nil {
		e = errors.E(opError, errors.RepositoryReservationErr, err)
		u.Logger.Error("Usecase error: ", e)
		return
	}

	return
}

func (u *ReservationUsecase) calculatePriceWithDiscount(price int64, discount int64) int64 {
	return price - (price * discount / 100)
}
