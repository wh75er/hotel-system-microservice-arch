package hotel_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	hotelServer "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service"
	pb "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/interceptors"
	stat_service "hotel-booking-system/internal/pkg/delivery/grpc/stat-service"
	stat_proto "hotel-booking-system/internal/pkg/delivery/grpc/stat-service/proto"
	jwtManager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/repository/postgres"
	hotelRepositories "hotel-booking-system/internal/pkg/repository/postgres/hotel-service"
	"hotel-booking-system/internal/pkg/usecase"
	hotelUsecases "hotel-booking-system/internal/pkg/usecase/hotel-service"
	"net"
)

type App struct {
	db         *sqlx.DB
	conf       *config
	configName string
	server     *grpc.Server
	StatClient       stat_proto.StatServiceClient
	logger     logs.LoggerInterface
}

func New() *App {
	return &App{
		nil,
		newConfig(),
		"",
		&grpc.Server{},
		nil,
		logs.NewLogrus(),
	}
}

func (a *App) Run(configFilename string) {
	a.configName = configFilename
	a.setupApp()
	a.setupStorage()
	connectionsCloseFunction := a.establishClientConnectWithAllDependentServices()
	defer func() {
		connectionsCloseFunction()
	}()

	jwtTokenManager := jwtManager.NewJWTManager(a.conf.Server.JWTSecret, a.conf.Server.TokenDuration.Duration)

	a.server = grpc.NewServer(
		grpc.UnaryInterceptor(
			interceptors.NewServerAdminAuthInterceptor(
				jwtTokenManager,
				hotelServer.AccessibleHotelServicePaths(),
				a.logger,
			).Unary(),
		),
	)

	hotelRepository := hotelRepositories.NewHotelRepository(a.db, a.logger)
	roomRepository := hotelRepositories.NewRoomRepository(a.db, a.logger)

	hotelUsecase := hotelUsecases.NewHotelUsecase(hotelRepository, roomRepository, a.logger)
	roomUsecase := hotelUsecases.NewRoomUsecase(hotelRepository, roomRepository, a.StatClient, a.logger)
	adminCredsUsecase := usecase.NewAdminCredentialsUsecase(a.conf.AdminCredentials)

	hotelS := hotelServer.NewHotelServer(
		hotelUsecase,
		roomUsecase,
		adminCredsUsecase,
		jwtTokenManager,
		a.logger,
	)

	pb.RegisterHotelServiceServer(a.server, hotelS)

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

	if err := a.conf.setStatServiceFromEnv(); err != nil {
		a.logger.Fatal(err)
	}

	a.logger.Infof("Loaded JWT Key: %v***", a.conf.Server.JWTSecret[:2])
	a.logger.Infof("Loaded Admin Id: %v", a.conf.AdminCredentials.Id)
	a.logger.Infof("Loaded Admin Secret: %v***", a.conf.AdminCredentials.Secret[:2])
	a.logger.Infof("Loaded Stat service data: %v", a.conf.StatService)
}

func (a *App) establishClientConnectWithAllDependentServices() func() {
	jwtTokenManager := jwtManager.NewJWTManager("", 0)

	statServiceConnCloseFunction := a.setupStatServiceConnection(jwtTokenManager)

	return func() {
		statServiceConnCloseFunction()
	}
}

func (a *App) setupStatServiceConnection(jwtTokenManager *jwtManager.JWTManager) func() {
	authInterceptor := interceptors.NewClientAuthInterceptor(
		a.conf.StatService.Credentials,
		jwtTokenManager,
		interceptors.MethodsRoleMapToSet(stat_service.AccessibleStatServicePaths()),
		logrus.New(),
	)

	conn, err := grpc.Dial(
		fmt.Sprintf(a.conf.StatService.Url),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(authInterceptor.Unary()),
	)
	if err != nil {
		a.logger.Fatalf("Failed to make User Stat Service grpc client: %v", err)
	}

	// Create specific client out of connection
	client := stat_proto.NewStatServiceClient(conn)

	// Add client GetToken API to auth interceptor
	authInterceptor.GrpcServiceClient = client

	a.StatClient = client

	return func() { conn.Close() }
}
