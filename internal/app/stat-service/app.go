package stat_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"hotel-booking-system/internal/pkg/delivery/grpc/interceptors"
	statServer "hotel-booking-system/internal/pkg/delivery/grpc/stat-service"
	pb "hotel-booking-system/internal/pkg/delivery/grpc/stat-service/proto"
	jwtManager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/repository/postgres"
	statRepositories "hotel-booking-system/internal/pkg/repository/postgres/stat-service"
	"hotel-booking-system/internal/pkg/usecase"
	statUsecases "hotel-booking-system/internal/pkg/usecase/stat-service"
	"net"
)

type App struct {
	db         *sqlx.DB
	conf       *config
	configName string
	server     *grpc.Server
	logger     logs.LoggerInterface
}

func New() *App {
	return &App{
		nil,
		newConfig(),
		"",
		&grpc.Server{},
		logs.NewLogrus(),
	}
}

func (a *App) Run(configFilename string) {
	a.configName = configFilename
	a.setupApp()
	a.setupStorage()

	jwtTokenManager := jwtManager.NewJWTManager(a.conf.Server.JWTSecret, a.conf.Server.TokenDuration.Duration)

	a.server = grpc.NewServer(
		grpc.UnaryInterceptor(
			interceptors.NewServerAdminAuthInterceptor(
				jwtTokenManager,
				statServer.AccessibleStatServicePaths(),
				a.logger,
			).Unary(),
		),
	)

	statRepository := statRepositories.NewStatRepository(a.db, a.logger)

	statUsecase := statUsecases.NewStatUsecase(statRepository, a.logger)
	adminCredsUsecase := usecase.NewAdminCredentialsUsecase(a.conf.AdminCredentials)

	loyaltyS := statServer.NewStatServer(
		statUsecase,
		adminCredsUsecase,
		jwtTokenManager,
		a.logger,
	)

	pb.RegisterStatServiceServer(a.server, loyaltyS)

	a.logger.Infof("Starting server on port: %v", a.conf.Server.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.conf.Server.Port))
	if err != nil {
		a.logger.Fatalf("Failed to listen: %v", err)
	}

	err = a.server.Serve(lis)
	if err != nil {
		a.logger.Fatalf("Failed to serve listener: %v", err)
	}
}

func (a *App) setupStorage() {
	a.db = postgres.EstablishDbConnection(a.logger, a.conf.Storage.Url, a.conf.Storage.MaxPoolConn)

	a.logger.Info("Successfully established connection with database")

	postgres.RunMigrations(a.logger, "file://init/migrations/stat-service", a.conf.Storage.Url)

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

	if err := a.conf.setJWTKeyFromEnv(); err != nil {
		a.logger.Fatal(err)
	}

	if err := a.conf.setAdminCredsFromEnv(); err != nil {
		a.logger.Fatal(err)
	}

	a.logger.Infof("Loaded JWT Key: %v***", a.conf.Server.JWTSecret[:2])
	a.logger.Infof("Loaded Admin Id: %v", a.conf.AdminCredentials.Id)
	a.logger.Infof("Loaded Admin Secret: %v***", a.conf.AdminCredentials.Secret[:2])
}
