package payment_service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"hotel-booking-system/internal/pkg/delivery/grpc/interceptors"
	paymentServer "hotel-booking-system/internal/pkg/delivery/grpc/payment-service"
	paymentService "hotel-booking-system/internal/pkg/delivery/grpc/payment-service"
	pb "hotel-booking-system/internal/pkg/delivery/grpc/payment-service/proto"
	http2 "hotel-booking-system/internal/pkg/delivery/http"
	payment_service "hotel-booking-system/internal/pkg/delivery/http/payment-service"
	jwtManager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/repository/postgres"
	paymentRepositories "hotel-booking-system/internal/pkg/repository/postgres/payment-service"
	"hotel-booking-system/internal/pkg/usecase"
	paymentUsecases "hotel-booking-system/internal/pkg/usecase/payment-service"
	"net"
	"net/http"
)

type App struct {
	db         *sqlx.DB
	conf       *config
	configName string
	server     *grpc.Server
	httpRouter http.Handler
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

	jwtTokenManager := jwtManager.NewJWTManager(a.conf.Server.JWTSecret, a.conf.Server.TokenDuration.Duration)

	a.server = grpc.NewServer(
		grpc.UnaryInterceptor(
			interceptors.NewServerAdminAuthInterceptor(
				jwtTokenManager,
				paymentService.AccessiblePaymentServicePaths(),
				a.logger,
			).Unary(),
		),
	)

	paymentRepository := paymentRepositories.NewPaymentRepository(a.db, a.logger)

	paymentUsecase := paymentUsecases.NewPaymentUsecase(paymentRepository, a.logger)
	adminCredsUsecase := usecase.NewAdminCredentialsUsecase(a.conf.AdminCredentials)

	paymentS := paymentServer.NewPaymentServer(
		paymentUsecase,
		adminCredsUsecase,
		jwtTokenManager,
		a.logger,
	)

	pb.RegisterPaymentServiceServer(a.server, paymentS)

	// Set http server for yoo-money notifications
	router := http.NewServeMux()
	payment_service.SetPaymentHttpRoutes(router, paymentUsecase, a.logger)

	loggingMiddleware := http2.LoggingMiddleware(a.logger)
	loggedRouter := loggingMiddleware(router)
	a.httpRouter = loggedRouter

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", a.conf.Server.Port))
	if err != nil {
		a.logger.Fatalf("Failed to listen: %v", err)
	}

	m := cmux.New(lis)

	httpl := m.Match(cmux.HTTP1Fast())
	grpcl := m.MatchWithWriters(
		cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpV2 := m.Match(cmux.HTTP2())

	go a.serveGRPC(grpcl)
	go a.serveHTTP(httpl)
	go a.serveHTTP(httpV2)

	err = m.Serve()
	if err != nil {
		a.logger.Fatalf("Failed to serve cmux listener: %v", err)
	}
}

func (a *App) serveGRPC(lis net.Listener) {
	err := a.server.Serve(lis)
	if err != nil {
		a.logger.Fatalf("Failed to serve GRPC listener: %v", err)
	}
}

func (a *App) serveHTTP(lis net.Listener) {
	err := http.Serve(lis, a.httpRouter)
	if err != nil {
		a.logger.Fatalf("Failed to serve HTTP listener: %v", err)
	}
}

func (a *App) setupStorage() {
	a.db = postgres.EstablishDbConnection(a.logger, a.conf.Storage.Url, a.conf.Storage.MaxPoolConn)

	a.logger.Info("Successfully established connection with database")

	postgres.RunMigrations(a.logger, "file://init/migrations/loyalty-service", a.conf.Storage.Url)

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
