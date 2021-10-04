package reservation_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	auth_service "hotel-booking-system/internal/pkg/delivery/grpc/auth-service"
	users_proto "hotel-booking-system/internal/pkg/delivery/grpc/auth-service/proto"
	gateway_service "hotel-booking-system/internal/pkg/delivery/grpc/gateway-service"
	"hotel-booking-system/internal/pkg/delivery/grpc/gateway-service/proto"
	hotel_service "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service"
	hotel_proto "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/interceptors"
	loyalty_service "hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service"
	loyalty_proto "hotel-booking-system/internal/pkg/delivery/grpc/loyalty-service/proto"
	paymentService "hotel-booking-system/internal/pkg/delivery/grpc/payment-service"
	payment_proto "hotel-booking-system/internal/pkg/delivery/grpc/payment-service/proto"
	reservationService "hotel-booking-system/internal/pkg/delivery/grpc/reservation-service"
	reservation_proto "hotel-booking-system/internal/pkg/delivery/grpc/reservation-service/proto"
	jwtManager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/repository/postgres"
	"hotel-booking-system/internal/pkg/usecase"
	"net"
)

type App struct {
	db                *sqlx.DB
	conf              *config
	configName        string
	server            *grpc.Server
	ReservationClient reservation_proto.ReservationServiceClient
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

	jwtTokenManager := jwtManager.NewJWTManager("", 0)

	a.server = grpc.NewServer(
		grpc.UnaryInterceptor(
			interceptors.NewAuthServiceInterceptor(
				jwtTokenManager,
				a.UsersClient,
				reservationService.AccessibleReservationServicePaths(),
				a.logger,
			).Unary(),
		),
	)

	adminCredsUsecase := usecase.NewAdminCredentialsUsecase(a.conf.AdminCredentials)

	gatewayS := gateway_service.NewGatewayServer(
		adminCredsUsecase,
		a.UsersClient,
		a.HotelClient,
		a.UserLoyaltyClient,
		a.PaymentClient,
		a.ReservationClient,
		a.logger,
	)

	proto.RegisterGatewayServiceServer(a.server, gatewayS)
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
	if err := a.conf.setReservationServiceFromEnv(); err != nil {
		a.logger.Fatal(err)
	}

	a.logger.Infof("Loaded JWT Key: %v***", a.conf.Server.JWTSecret[:2])
	a.logger.Infof("Loaded Admin Id: %v", a.conf.AdminCredentials.Id)
	a.logger.Infof("Loaded Admin Secret: %v***", a.conf.AdminCredentials.Secret[:2])
}

func (a *App) establishClientConnectWithAllDependentServices() func() {
	jwtTokenManager := jwtManager.NewJWTManager("", 0)

	userLoyaltyServiceConnCloseFunction := a.setupUserLoyaltyServiceConnection(jwtTokenManager)
	usersServiceConnCloseFunction := a.setupUserServiceConnection(jwtTokenManager)
	hotelServiceConnCloseFunction := a.setupHotelServiceConnection(jwtTokenManager)
	paymentServiceConnCloseFunction := a.setupPaymentServiceConnection(jwtTokenManager)
	reservationServiceConnCloseFunction := a.setupReservationServiceConnection(jwtTokenManager)

	return func() {
		userLoyaltyServiceConnCloseFunction()
		usersServiceConnCloseFunction()
		hotelServiceConnCloseFunction()
		paymentServiceConnCloseFunction()
		reservationServiceConnCloseFunction()
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

func (a *App) setupReservationServiceConnection(jwtTokenManager *jwtManager.JWTManager) func() {
	authInterceptor := interceptors.NewClientAuthInterceptor(
		a.conf.ReservationService.Credentials,
		jwtTokenManager,
		interceptors.MethodsRoleMapToSet(reservationService.AccessibleReservationServicePaths()),
		logrus.New(),
	)

	conn, err := grpc.Dial(
		fmt.Sprintf(a.conf.ReservationService.Url),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(authInterceptor.Unary()),
	)
	if err != nil {
		a.logger.Fatalf("Failed to make User Service grpc client: %v", err)
	}

	// Create specific client out of connection
	client := reservation_proto.NewReservationServiceClient(conn)

	// Add client GetToken API to auth interceptor
	authInterceptor.GrpcServiceClient = client

	a.ReservationClient = client

	return func() { conn.Close() }
}
