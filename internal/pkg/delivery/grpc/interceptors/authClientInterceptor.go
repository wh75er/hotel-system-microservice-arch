package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	hotel_service "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service"
	"hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	jwtManager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
	"log"
	"time"
)

const (
	timeUntilTokenExpUpdateInMinutes = 1
)

type GrpcClientI interface {
	GetToken(ctx context.Context, in *proto.Credentials, opts ...grpc.CallOption) (*proto.Token, error)
}

type ClientAuthInterceptor struct {
	GrpcServiceClient  GrpcClientI
	ServiceCredentials models.Credentials
	JwtManager         *jwtManager.JWTManager
	AuthMethods        map[string]bool
	logger             logs.LoggerInterface
	token              models.Token
}

func NewClientAuthInterceptor(
	creds models.Credentials,
	jwtManager *jwtManager.JWTManager,
	authMethods map[string]bool,
	logger logs.LoggerInterface,
) *ClientAuthInterceptor {
	return &ClientAuthInterceptor{
		ServiceCredentials: creds,
		JwtManager:         jwtManager,
		AuthMethods:        authMethods,
		logger:             logger,
	}
}

func (i *ClientAuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("--> unary interceptor: %s", method)

		if i.AuthMethods[method] {
			err := i.updateTokenIfRequired()
			if err != nil {
				return err
			}

			return invoker(i.attachToken(ctx), method, req, reply, cc, opts...)
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func (i *ClientAuthInterceptor) updateTokenIfRequired() error {
	var timeLeft time.Duration

	if i.token != "" {
		i.logger.Info("Extracting token expire date...")
		t, err := i.JwtManager.UntilExp(string(i.token))
		if err != nil {
			i.logger.Errorf("Failed to extract token expire date: %v - %v", err, errors.SourceDetails(err))
			return err
		}

		timeLeft = t
	}

	i.logger.Infof("Token will expire after %v minutes", timeLeft.Minutes())

	if timeLeft.Minutes() < timeUntilTokenExpUpdateInMinutes {
		i.logger.Info("Token is old, updating...")
		err := i.refreshToken()
		if err != nil {
			i.logger.Errorf("Failed to refresh the token: %v - %v", err, errors.SourceDetails(err))
			return err
		}
		i.logger.Info("Successfully updated the token")
	}

	i.logger.Info("Token is ready to be attached")

	return nil
}

func (i *ClientAuthInterceptor) refreshToken() error {
	newToken, err := i.GrpcServiceClient.GetToken(context.Background(), hotel_service.CredentialsToProto(&i.ServiceCredentials))
	if err != nil {
		return err
	}

	i.token = *hotel_service.ProtoToToken(newToken)

	return nil
}

func (i *ClientAuthInterceptor) attachToken(ctx context.Context) context.Context {
	i.logger.Info("Token is successfully attached")
	return metadata.AppendToOutgoingContext(ctx, "authorization", string(i.token))
}
