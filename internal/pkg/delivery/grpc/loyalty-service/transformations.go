package loyalty_service

import (
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	"hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service/proto"
	"hotel-booking-system/internal/pkg/models"
)

func LoyaltyToProto(l *models.Loyalty) *proto.Loyalty {
	return &proto.Loyalty{
		UserUuid:           &commonProto.UUID{Value: l.UserUuid.String()},
		Status:             string(l.Status),
		Discount:           int64(l.Discount),
		ContributionAmount: int64(l.ContributionAmount),
	}
}
