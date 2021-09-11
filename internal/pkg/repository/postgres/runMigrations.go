package postgres

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"hotel-booking-system/internal/pkg/logs"
)

func RunMigrations(logger logs.LoggerInterface, sourceUrl string, dbUrl string) {
	m, err := migrate.New(
		sourceUrl,
		dbUrl)
	if err != nil {
		logger.Fatal("Failed to start migrations: ", err)
	}
	if err := m.Up(); err != migrate.ErrNoChange && err != nil {
		logger.Fatal("Failed to run migrations: ", err)
	}
}
