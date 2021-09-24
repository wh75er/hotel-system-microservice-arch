package hotel_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	hotel_server "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service"
	pb "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/interceptors"
	jwt_manager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/repository/postgres"
	hotel_repositories "hotel-booking-system/internal/pkg/repository/postgres/hotel-service"
	hotel_usecases "hotel-booking-system/internal/pkg/usecase/hotel-service"
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

	jwtTokenManager := jwt_manager.NewJWTManager(a.conf.Server.JWTSecret, a.conf.Server.TokenDuration.Duration)

	a.server = grpc.NewServer(
		grpc.UnaryInterceptor(
			interceptors.NewServerAdminAuthInterceptor(jwtTokenManager).Unary(),
		),
	)

	hotelRepository := hotel_repositories.NewHotelRepository(a.db, a.logger)
	roomRepository := hotel_repositories.NewRoomRepository(a.db, a.logger)
	reviewRepository := hotel_repositories.NewReviewRepository(a.db, a.logger)

	hotelUsecase := hotel_usecases.NewHotelUsecase(hotelRepository, roomRepository, reviewRepository, a.logger)
	roomUsecase := hotel_usecases.NewRoomUsecase(hotelRepository, roomRepository, a.logger)
	reviewUsecase := hotel_usecases.NewReviewUsecase(hotelRepository, reviewRepository, a.logger)
	adminCredsUsecase := hotel_usecases.NewAdminCredentialsUsecase(a.conf.AdminCredentials)

	hotelServer := hotel_server.NewHotelServer(
		hotelUsecase,
		reviewUsecase,
		roomUsecase,
		adminCredsUsecase,
		jwtTokenManager,
		a.logger,
	)

	pb.RegisterHotelServiceServer(a.server, hotelServer)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", a.conf.Server.Port))
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
