package commonProto

import "hotel-booking-system/internal/pkg/models"

func ProtoToCredentials(c *Credentials) *models.Credentials {
	return &models.Credentials{
		Id:     c.Id,
		Secret: c.Secret,
	}
}

func CredentialsToProto(c *models.Credentials) *Credentials {
	return &Credentials{
		Id:     c.Id,
		Secret: c.Secret,
	}
}

func ProtoToToken(t *Token) *models.Token {
	_token := models.Token(t.Value)
	return &_token
}

func TokenToProto(t *models.Token) *Token {
	return &Token{
		Value: string(*t),
	}
}
