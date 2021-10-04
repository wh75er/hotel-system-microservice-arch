package reservation_service

import (
	"github.com/google/uuid"
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	"hotel-booking-system/internal/pkg/delivery/grpc/reservation-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	"hotel-booking-system/internal/pkg/models"
	"time"
)

func (s *ReservationServer) ProtoToReservation(pr *proto.Reservation) (*models.Reservation, error) {
	var opError errors.Op = "reservation-service.ProtoToReservation"

	var validReservationUuid uuid.UUID
	if len(pr.ReservationUuid.Value) != 0 {
		var err error
		validReservationUuid, err = uuid.Parse(pr.ReservationUuid.Value)
		if err != nil {
			e := errors.E(opError, errors.ReservationUuidValidationErr, err)
			s.Logger.Error("Grpc error: ", e)
			return nil, e
		}
	}

	var validUserUuid uuid.UUID
	if len(pr.UserUuid.Value) != 0 {
		var err error
		validUserUuid, err = uuid.Parse(pr.UserUuid.Value)
		if err != nil {
			e := errors.E(opError, errors.UserUuidValidationErr, err)
			s.Logger.Error("Grpc error: ", e)
			return nil, e
		}
	}

	var validRoomUuid uuid.UUID
	if len(pr.UserUuid.Value) != 0 {
		var err error
		validRoomUuid, err = uuid.Parse(pr.RoomUuid.Value)
		if err != nil {
			e := errors.E(opError, errors.RoomUuidValidationErr, err)
			s.Logger.Error("Grpc error: ", e)
			return nil, e
		}
	}

	var validPaymentUuid uuid.UUID
	if len(pr.PaymentUuid.Value) != 0 {
		var err error
		validPaymentUuid, err = uuid.Parse(pr.PaymentUuid.Value)
		if err != nil {
			e := errors.E(opError, errors.PaymentUuidValidationErr, err)
			s.Logger.Error("Grpc error: ", e)
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

func (s *ReservationServer) ReservationToProto(r *models.Reservation) *proto.Reservation {
	return &proto.Reservation{
		ReservationUuid: &commonProto.UUID{Value: r.ReservationUuid.String()},
		RoomUuid:        &commonProto.UUID{Value: r.RoomUuid.String()},
		UserUuid:        &commonProto.UUID{Value: r.UserUuid.String()},
		PaymentUuid:     &commonProto.UUID{Value: r.PaymentUuid.String()},
		Status:          string(r.Status),
		Date:            r.Date.Unix(),
	}
}

func (s *ReservationServer) ReservationsToProto(r []models.Reservation) *proto.Reservations {
	var refReservations []*proto.Reservation

	for _, v := range r {
		t := s.ReservationToProto(&v)
		refReservations = append(refReservations, t)
	}

	return &proto.Reservations{
		Value: refReservations,
	}
}
