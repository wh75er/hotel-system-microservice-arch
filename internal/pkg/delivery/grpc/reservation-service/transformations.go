package reservation_service

import (
	"github.com/google/uuid"
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	"hotel-booking-system/internal/pkg/delivery/grpc/reservation-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

func ProtoToReservation(pr *proto.Reservation) (*models.Reservation, error) {
	var opError errors.Op = "reservation-service.ProtoToReservation"

	if pr == nil {
		return &models.Reservation{}, nil
	}

	var validReservationUuid uuid.UUID
	if pr.GetReservationUuid() != nil && len(pr.ReservationUuid.Value) != 0 {
		var err error
		validReservationUuid, err = uuid.Parse(pr.ReservationUuid.Value)
		if err != nil {
			e := errors.E(opError, errors.ReservationUuidValidationErr, err)
			return nil, e
		}
	}

	var validUserUuid uuid.UUID
	if pr.GetUserUuid() != nil && len(pr.UserUuid.Value) != 0 {
		var err error
		validUserUuid, err = uuid.Parse(pr.UserUuid.Value)
		if err != nil {
			e := errors.E(opError, errors.UserUuidValidationErr, err)
			return nil, e
		}
	}

	var validRoomUuid uuid.UUID
	if pr.GetRoomUuid() != nil && len(pr.RoomUuid.Value) != 0 {
		var err error
		validRoomUuid, err = uuid.Parse(pr.RoomUuid.Value)
		if err != nil {
			e := errors.E(opError, errors.RoomUuidValidationErr, err)
			return nil, e
		}
	}

	var validPaymentUuid uuid.UUID
	if pr.GetPaymentUuid() != nil && len(pr.PaymentUuid.Value) != 0 {
		var err error
		validPaymentUuid, err = uuid.Parse(pr.PaymentUuid.Value)
		if err != nil {
			e := errors.E(opError, errors.PaymentUuidValidationErr, err)
			return nil, e
		}

	}

	return &models.Reservation{
		ReservationUuid: validReservationUuid,
		RoomUuid:        validRoomUuid,
		UserUuid:        validUserUuid,
		PaymentUuid:     validPaymentUuid,
		Date:            time.Unix(pr.Date, 0),
		Status:          models.ReservationStatus(pr.Status),
	}, nil
}

func ReservationToProto(r *models.Reservation) *proto.Reservation {
	return &proto.Reservation{
		ReservationUuid: &commonProto.UUID{Value: r.ReservationUuid.String()},
		RoomUuid:        &commonProto.UUID{Value: r.RoomUuid.String()},
		UserUuid:        &commonProto.UUID{Value: r.UserUuid.String()},
		PaymentUuid:     &commonProto.UUID{Value: r.PaymentUuid.String()},
		Status:          string(r.Status),
		Date:            r.Date.Unix(),
	}
}

func ReservationsToProto(r []models.Reservation) *proto.Reservations {
	var refReservations []*proto.Reservation

	for _, v := range r {
		t := ReservationToProto(&v)
		refReservations = append(refReservations, t)
	}

	return &proto.Reservations{
		Value: refReservations,
	}
}
