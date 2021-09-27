package auth_service

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"hotel-booking-system/internal/pkg/models"
	"os"
	"time"
)

const (
	jwtSecretEnv   = "JWT_KEY"
	adminIdEnv     = "ADMIN_ID"
	adminSecretEnv = "ADMIN_SECRET"
	configDst      = "configs/auth-service/"
)

type duration struct {
	time.Duration
}

type config struct {
	Server           Server
	Storage          Storage
	AdminCredentials models.Credentials
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
		Server{
			Port: 3000,
		},
		Storage{
			"postgresql://postgres:postgres@localhost:5432/postgres",
			30,
		},
		models.Credentials{},
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

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
