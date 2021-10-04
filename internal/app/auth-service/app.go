package auth_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	authServer "hotel-booking-system/internal/pkg/delivery/grpc/auth-service"
	authService "hotel-booking-system/internal/pkg/delivery/grpc/auth-service"
	pb "hotel-booking-system/internal/pkg/delivery/grpc/auth-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/interceptors"
	loyalty_service "hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service"
	loyalty_proto "hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service/proto"
	jwtManager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/repository/postgres"
	userRepositories "hotel-booking-system/internal/pkg/repository/postgres/auth-service"
	"hotel-booking-system/internal/pkg/usecase"
	userUsecases "hotel-booking-system/internal/pkg/usecase/auth-service"
	"net"
)

type App struct {
	db                *sqlx.DB
	conf              *config
	configName        string
	UserLoyaltyClient loyalty_proto.LoyaltyServiceClient
	server            *grpc.Server
	logger            logs.LoggerInterface
}

func New() *App {
	return &App{
		nil,
		newConfig(),
		"",
		nil,
		&grpc.Server{},
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
				authService.AccessibleAuthServicePaths(),
				a.logger,
			).Unary(),
		),
	)

	userRepository := userRepositories.NewUserRepository(a.db, a.logger)

	userUsecase := userUsecases.NewUserUsecase(userRepository, a.UserLoyaltyClient, jwtTokenManager, a.logger)
	adminCredsUsecase := usecase.NewAdminCredentialsUsecase(a.conf.AdminCredentials)

	authS := authServer.NewAuthServer(
		userUsecase,
		adminCredsUsecase,
		jwtTokenManager,
		a.logger,
	)

	pb.RegisterAuthServiceServer(a.server, authS)
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

	postgres.RunMigrations(a.logger, "file://init/migrations/auth-service", a.conf.Storage.Url)

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

	if err := a.conf.setUserLoyaltyServiceFromEnv(); err != nil {
		a.logger.Fatal(err)
	}

	a.logger.Infof("Loaded JWT Key: %v***", a.conf.Server.JWTSecret[:2])
	a.logger.Infof("Loaded Admin Id: %v", a.conf.AdminCredentials.Id)
	a.logger.Infof("Loaded Admin Secret: %v***", a.conf.AdminCredentials.Secret[:2])
}

func (a *App) establishClientConnectWithAllDependentServices() func() {
	jwtTokenManager := jwtManager.NewJWTManager("", 0)

	userLoyaltyServiceConnCloseFunction := a.setupUserLoyaltyServiceConnection(jwtTokenManager)

	return func() {
		userLoyaltyServiceConnCloseFunction()
	}
}

func (a *App) setupUserLoyaltyServiceConnection(jwtTokenManager *jwtManager.JWTManager) func() {
	authInterceptor := interceptors.NewClientAuthInterceptor(
		a.conf.UserLoyaltyService.Credentials,
		jwtTokenManager,
		interceptors.MethodsRoleMapToSet(loyalty_service.AccessibleLoyaltyServicePaths()),
		logrus.New(),
	)

	conn, err := grpc.Dial(
		fmt.Sprintf(a.conf.UserLoyaltyService.Url),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(authInterceptor.Unary()),
	)
	if err != nil {
		a.logger.Fatalf("Failed to make User Loyalty Service grpc client: %v", err)
	}

	// Create specific client out of connection
	client := loyalty_proto.NewLoyaltyServiceClient(conn)

	// Add client GetToken API to auth interceptor
	authInterceptor.GrpcServiceClient = client

	a.UserLoyaltyClient = client

	return func() { conn.Close() }
}
