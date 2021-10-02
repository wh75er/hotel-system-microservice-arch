package payment_service

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
	userServiceUrlEnv                = "USER_SERVICE_URL"
	userServiceAdminIdEnv            = "USER_SERVICE_ADMIN_ID"
	userServiceAdminSecretEnv        = "USER_SERVICE_ADMIN_SECRET"
	configDst                        = "configs/payment-service/"
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

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
