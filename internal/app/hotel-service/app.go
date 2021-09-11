package hotel_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/repository/postgres"
)

type App struct {
	db *sqlx.DB
	conf *config
	configName string
	server *echo.Echo
	logger logs.LoggerInterface
}

func New() *App {
	return &App {
		nil,
		newConfig(),
		"",
		echo.New(),
		logs.NewLogrus(),
	}
}

func (a *App) Run(configFilename string) {
	a.configName = configFilename
	a.setupApp()
	a.setupStorage()

//	a.server.Use(middleware.Logger())
//
//	if err := a.server.Start(":" + strconv.Itoa(a.conf.Server.Port)); err == http.ErrServerClosed {
//		a.logger.Fatal(err)
//	}
}

func (a *App) setupStorage() {
	postgres.EstablishDbConnection(a.logger, a.conf.Storage.Url, a.conf.Storage.MaxPoolConn)

	a.logger.Info("Successfully established connection with database")

	postgres.RunMigrations(a.logger, "file://init/migrations/hotel-service", a.conf.Storage.Url)

	a.logger.Info("Successfully ran migrations")
}

func (a *App) setupApp() {
	// Check if there is a configuration file
	if a.configName != "" {
		if err := a.conf.loadFromToml(a.configName); err != nil {
			a.logger.Fatal("Failed to decode configuration file: ", err.Error())
		}
		a.logger.Infof("Loaded configuration file: %v. Current configuration: %v", a.configName, a.conf)
	} else {
		a.logger.Warnf("Configuration file is not specified, using default configuration: %v", a.conf)
	}
}
