package reservation_service

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"hotel-booking-system/internal/pkg/models"
	"os"
	"time"
)

const (
	jwtSecretEnv                     = "JWT_KEY"
	adminIdEnv                       = "ADMIN_ID"
	adminSecretEnv                   = "ADMIN_SECRET"
	userLoyaltyServiceUrlEnv         = "USER_LOYALTY_SERVICE_URL"
	userLoyaltyServiceAdminIdEnv     = "USER_LOYALTY_SERVICE_ADMIN_ID"
	userLoyaltyServiceAdminSecretEnv = "USER_LOYALTY_SERVICE_ADMIN_SECRET"
	hotelServiceUrlEnv               = "HOTEL_SERVICE_URL"
	hotelServiceAdminIdEnv           = "HOTEL_SERVICE_ADMIN_ID"
	hotelServiceAdminSecretEnv       = "HOTEL_SERVICE_ADMIN_SECRET"
	paymentServiceUrlEnv             = "PAYMENT_SERVICE_URL"
	paymentServiceAdminIdEnv         = "PAYMENT_SERVICE_ADMIN_ID"
	paymentServiceAdminSecretEnv     = "PAYMENT_SERVICE_ADMIN_SECRET"
	userServiceUrlEnv                = "USER_SERVICE_URL"
	userServiceAdminIdEnv            = "USER_SERVICE_ADMIN_ID"
	userServiceAdminSecretEnv        = "USER_SERVICE_ADMIN_SECRET"
	statServiceUrlEnv         = "STAT_SERVICE_URL"
	statServiceAdminIdEnv     = "STAT_SERVICE_ADMIN_ID"
	statServiceAdminSecretEnv = "STAT_SERVICE_ADMIN_SECRET"
	configDst                        = "configs/reservation-service/"
)

type DependencyService struct {
	Url         string
	Credentials models.Credentials
}

type duration struct {
	time.Duration
}

type config struct {
	Server             Server
	Storage            Storage
	AdminCredentials   models.Credentials
	UserLoyaltyService DependencyService
	HotelService       DependencyService
	PaymentService     DependencyService
	StatService        DependencyService
	UserService        DependencyService
}

type Server struct {
	Port          int
	JWTSecret     string
	TokenDuration duration
}

type Storage struct {
	Url         string
	MaxPoolConn int
}

func newConfig() *config {
	return &config{
		Server: Server{
			Port: 3000,
		},
		Storage: Storage{
			"postgresql://postgres:postgres@localhost:5432/postgres",
			30,
		},
		AdminCredentials: models.Credentials{},
	}
}

func (c *config) loadFromToml(tomlData string) (e error) {
	_, e = toml.DecodeFile(configDst+tomlData, c)
	return
}

func (c *config) setJWTKeyFromEnv() error {
	key, err := getEnvVariable(jwtSecretEnv)
	if err != nil {
		return err
	}

	c.Server.JWTSecret = key

	return nil
}

func (c *config) setAdminCredsFromEnv() error {
	id, err := getEnvVariable(adminIdEnv)
	if err != nil {
		return err
	}

	c.AdminCredentials.Id = id

	secret, err := getEnvVariable(adminSecretEnv)
	if err != nil {
		return err
	}

	c.AdminCredentials.Secret = secret

	return nil
}

func getEnvVariable(varName string) (string, error) {
	v := os.Getenv(varName)
	if v == "" {
		return "", fmt.Errorf("failed to find environment variable: %v", varName)
	}

	return v, nil
}

func (c *config) setUserLoyaltyServiceFromEnv() error {
	url, err := getEnvVariable(userLoyaltyServiceUrlEnv)
	if err != nil {
		return err
	}

	c.UserLoyaltyService.Url = url

	id, err := getEnvVariable(userLoyaltyServiceAdminIdEnv)
	if err != nil {
		return err
	}

	c.UserLoyaltyService.Credentials.Id = id

	secret, err := getEnvVariable(userLoyaltyServiceAdminSecretEnv)
	if err != nil {
		return err
	}

	c.UserLoyaltyService.Credentials.Secret = secret

	return nil
}

func (c *config) setHotelServiceFromEnv() error {
	url, err := getEnvVariable(hotelServiceUrlEnv)
	if err != nil {
		return err
	}

	c.HotelService.Url = url

	id, err := getEnvVariable(hotelServiceAdminIdEnv)
	if err != nil {
		return err
	}

	c.HotelService.Credentials.Id = id

	secret, err := getEnvVariable(hotelServiceAdminSecretEnv)
	if err != nil {
		return err
	}

	c.HotelService.Credentials.Secret = secret

	return nil
}

func (c *config) setPaymentServiceFromEnv() error {
	url, err := getEnvVariable(paymentServiceUrlEnv)
	if err != nil {
		return err
	}

	c.PaymentService.Url = url

	id, err := getEnvVariable(paymentServiceAdminIdEnv)
	if err != nil {
		return err
	}

	c.PaymentService.Credentials.Id = id

	secret, err := getEnvVariable(paymentServiceAdminSecretEnv)
	if err != nil {
		return err
	}

	c.PaymentService.Credentials.Secret = secret

	return nil
}

func (c *config) setUserServiceFromEnv() error {
	url, err := getEnvVariable(userServiceUrlEnv)
	if err != nil {
		return err
	}

	c.UserService.Url = url

	id, err := getEnvVariable(userServiceAdminIdEnv)
	if err != nil {
		return err
	}

	c.UserService.Credentials.Id = id

	secret, err := getEnvVariable(userServiceAdminSecretEnv)
	if err != nil {
		return err
	}

	c.UserService.Credentials.Secret = secret

	return nil
}

func (c *config) setStatServiceFromEnv() error {
	url, err := getEnvVariable(statServiceUrlEnv)
	if err != nil {
		return err
	}

	c.StatService.Url = url

	id, err := getEnvVariable(statServiceAdminIdEnv)
	if err != nil {
		return err
	}

	c.StatService.Credentials.Id = id

	secret, err := getEnvVariable(statServiceAdminSecretEnv)
	if err != nil {
		return err
	}

	c.StatService.Credentials.Secret = secret

	return nil
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
