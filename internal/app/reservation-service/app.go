package reservation_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	auth_service "hotel-booking-system/internal/pkg/delivery/grpc/auth-service"
	users_proto "hotel-booking-system/internal/pkg/delivery/grpc/auth-service/proto"
	hotel_service "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service"
	hotel_proto "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/interceptors"
	loyalty_service "hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service"
	loyalty_proto "hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service/proto"
	paymentService "hotel-booking-system/internal/pkg/delivery/grpc/payment-service"
	payment_proto "hotel-booking-system/internal/pkg/delivery/grpc/payment-service/proto"
	reservationService "hotel-booking-system/internal/pkg/delivery/grpc/reservation-service"
	pb "hotel-booking-system/internal/pkg/delivery/grpc/reservation-service/proto"
	jwtManager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/repository/postgres"
	reservationRepositories "hotel-booking-system/internal/pkg/repository/postgres/reservation-service"
	"hotel-booking-system/internal/pkg/usecase"
	reservationUsecases "hotel-booking-system/internal/pkg/usecase/reservation-service"
	"net"
)

type App struct {
	db                *sqlx.DB
	conf              *config
	configName        string
	server            *grpc.Server
	HotelClient       hotel_proto.HotelServiceClient
	PaymentClient     payment_proto.PaymentServiceClient
	UserLoyaltyClient loyalty_proto.LoyaltyServiceClient
	UsersClient       users_proto.AuthServiceClient
	logger            logs.LoggerInterface
}

func New() *App {
	return &App{
		nil,
		newConfig(),
		"",
		&grpc.Server{},
		nil,
		nil,
		nil,
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
				reservationService.AccessibleReservationServicePaths(),
				a.logger,
			).Unary(),
		),
	)

	reservationRepository := reservationRepositories.NewReservationRepository(a.db, a.logger)

	reservationUsecase := reservationUsecases.NewReservationUsecase(
		reservationRepository,
		a.HotelClient,
		a.PaymentClient,
		a.UsersClient,
		a.UserLoyaltyClient,
		a.logger,
	)
	adminCredsUsecase := usecase.NewAdminCredentialsUsecase(a.conf.AdminCredentials)

	reservationS := reservationService.NewReservationServer(
		reservationUsecase,
		adminCredsUsecase,
		jwtTokenManager,
		a.logger,
	)

	pb.RegisterReservationServiceServer(a.server, reservationS)

	a.logger.Infof("Starting server on port: %v", a.conf.Server.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.conf.Server.Port))
	if err != nil {
		a.logger.Fatalf("Failed to listen: %v", err)
	}

	err = a.server.Serve(lis)
	if err != nil {
		a.logger.Fatalf("Failed to serve GRPC listener: %v", err)
	}
}

func (a *App) setupStorage() {
	a.db = postgres.EstablishDbConnection(a.logger, a.conf.Storage.Url, a.conf.Storage.MaxPoolConn)

	a.logger.Info("Successfully established connection with database")

	postgres.RunMigrations(a.logger, "file://init/migrations/payment-service", a.conf.Storage.Url)

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

	if err := a.conf.setUserServiceFromEnv(); err != nil {
		a.logger.Fatal(err)
	}

	if err := a.conf.setHotelServiceFromEnv(); err != nil {
		a.logger.Fatal(err)
	}

	if err := a.conf.setPaymentServiceFromEnv(); err != nil {
		a.logger.Fatal(err)
	}

	a.logger.Infof("Loaded JWT Key: %v***", a.conf.Server.JWTSecret[:2])
	a.logger.Infof("Loaded Admin Id: %v", a.conf.AdminCredentials.Id)
	a.logger.Infof("Loaded Admin Secret: %v***", a.conf.AdminCredentials.Secret[:2])
	a.logger.Infof("Loaded Loyalty service data: %v", a.conf.UserLoyaltyService)
	a.logger.Infof("Loaded Payment service data: %v", a.conf.PaymentService)
	a.logger.Infof("Loaded User service data: %v", a.conf.UserService)
	a.logger.Infof("Loaded Hotel service data: %v", a.conf.HotelService)
}

func (a *App) establishClientConnectWithAllDependentServices() func() {
	jwtTokenManager := jwtManager.NewJWTManager("", 0)

	userLoyaltyServiceConnCloseFunction := a.setupUserLoyaltyServiceConnection(jwtTokenManager)
	usersServiceConnCloseFunction := a.setupUserServiceConnection(jwtTokenManager)
	hotelServiceConnCloseFunction := a.setupHotelServiceConnection(jwtTokenManager)
	paymentServiceConnCloseFunction := a.setupPaymentServiceConnection(jwtTokenManager)

	return func() {
		userLoyaltyServiceConnCloseFunction()
		usersServiceConnCloseFunction()
		hotelServiceConnCloseFunction()
		paymentServiceConnCloseFunction()
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

func (a *App) setupHotelServiceConnection(jwtTokenManager *jwtManager.JWTManager) func() {
	authInterceptor := interceptors.NewClientAuthInterceptor(
		a.conf.HotelService.Credentials,
		jwtTokenManager,
		interceptors.MethodsRoleMapToSet(hotel_service.AccessibleHotelServicePaths()),
		logrus.New(),
	)

	conn, err := grpc.Dial(
		fmt.Sprintf(a.conf.HotelService.Url),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(authInterceptor.Unary()),
	)
	if err != nil {
		a.logger.Fatalf("Failed to make hotel Service grpc client: %v", err)
	}

	// Create specific client out of connection
	client := hotel_proto.NewHotelServiceClient(conn)

	// Add client GetToken API to auth interceptor
	authInterceptor.GrpcServiceClient = client

	a.HotelClient = client

	return func() { conn.Close() }
}

func (a *App) setupPaymentServiceConnection(jwtTokenManager *jwtManager.JWTManager) func() {
	authInterceptor := interceptors.NewClientAuthInterceptor(
		a.conf.PaymentService.Credentials,
		jwtTokenManager,
		interceptors.MethodsRoleMapToSet(paymentService.AccessiblePaymentServicePaths()),
		logrus.New(),
	)

	conn, err := grpc.Dial(
		fmt.Sprintf(a.conf.PaymentService.Url),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(authInterceptor.Unary()),
	)
	if err != nil {
		a.logger.Fatalf("Failed to make hotel Service grpc client: %v", err)
	}

	// Create specific client out of connection
	client := payment_proto.NewPaymentServiceClient(conn)

	// Add client GetToken API to auth interceptor
	authInterceptor.GrpcServiceClient = client

	a.PaymentClient = client

	return func() { conn.Close() }
}

func (a *App) setupUserServiceConnection(jwtTokenManager *jwtManager.JWTManager) func() {
	authInterceptor := interceptors.NewClientAuthInterceptor(
		a.conf.UserService.Credentials,
		jwtTokenManager,
		interceptors.MethodsRoleMapToSet(auth_service.AccessibleAuthServicePaths()),
		logrus.New(),
	)

	conn, err := grpc.Dial(
		fmt.Sprintf(a.conf.UserService.Url),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(authInterceptor.Unary()),
	)
	if err != nil {
		a.logger.Fatalf("Failed to make User Service grpc client: %v", err)
	}

	// Create specific client out of connection
	client := users_proto.NewAuthServiceClient(conn)

	// Add client GetToken API to auth interceptor
	authInterceptor.GrpcServiceClient = client

	a.UsersClient = client

	return func() { conn.Close() }
}
