package payment_service

import (
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	"hotel-booking-system/internal/pkg/delivery/grpc/payment-service/proto"
	"hotel-booking-system/internal/pkg/models"
)

func (s *PaymentServer) PaymentToProto(p *models.Payment) *proto.Payment {
	return &proto.Payment{
		PaymentUuid: &commonProto.UUID{Value: p.PaymentUuid.String()},
		UserUuid:    &commonProto.UUID{Value: p.UserUuid.String()},
		Status:      string(p.Status),
		Price:       int64(p.Price),
		TimeUpdated: p.TimeUpdated.Unix(),
	}
}

