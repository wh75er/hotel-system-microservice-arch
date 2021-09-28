package loyalty_service

import (
	"hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service/proto"
	"hotel-booking-system/internal/pkg/models"
)

func (s *LoyaltyServer) LoyaltyToProto(l *models.Loyalty) *proto.Loyalty {
	return &proto.Loyalty{
		UserUuid:           &proto.UUID{Value: l.UserUuid.String()},
		Status:             string(l.Status),
		Discount:           int64(l.Discount),
		ContributionAmount: int64(l.ContributionAmount),
	}
}

func ProtoToCredentials(c *proto.Credentials) *models.Credentials {
	return &models.Credentials{
		Id:     c.Id,
		Secret: c.Secret,
	}
}

func CredentialsToProto(c *models.Credentials) *proto.Credentials {
	return &proto.Credentials{
		Id:     c.Id,
		Secret: c.Secret,
	}
}

func ProtoToToken(t *proto.Token) *models.Token {
	_token := models.Token(t.Value)
	return &_token
}

func TokenToProto(t *models.Token) *proto.Token {
	return &proto.Token{
		Value: string(*t),
	}
}
