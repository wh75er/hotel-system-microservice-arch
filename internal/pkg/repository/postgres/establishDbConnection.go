package postgres

import (
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
	"hotel-booking-system/internal/pkg/logs"
)

func EstablishDbConnection(logger logs.LoggerInterface, dbUrl string, maxPoolConn int) *sqlx.DB {
	db, err := sqlx.Open("pgx", dbUrl)
	if err != nil {
		logger.Fatal("Failed to establish connection with db: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal("Failed to ping db ", err.Error())
	}

	db.SetMaxOpenConns(maxPoolConn)

	return db
}