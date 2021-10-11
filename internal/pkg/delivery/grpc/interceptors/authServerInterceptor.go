/*
		This interceptor provides cross service authentication.
	it uses jwt manager with specified secret to validate tokens.
*/
package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"hotel-booking-system/internal/pkg/errors"
	jwt_manager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
	"log"
)

type ServerAdminAuthInterceptor struct {
	JwtManager   *jwt_manager.JWTManager
	ApplyMethods map[string][]models.Role
	logger       logs.LoggerInterface
}

func NewServerAdminAuthInterceptor(
	jwtManager *jwt_manager.JWTManager,
	applyMethods map[string][]models.Role,
	logger logs.LoggerInterface,
) *ServerAdminAuthInterceptor {
	return &ServerAdminAuthInterceptor{
		jwtManager,
		applyMethods,
		logger,
	}
}

func (i *ServerAdminAuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("server unary interceptor: ", info.FullMethod)

		err := i.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		}

		return handler(ctx, req)
	}
}

func (i *ServerAdminAuthInterceptor) authorize(ctx context.Context, method string) error {
	var opError errors.Op = "jwt-manager.authorize"

	accessiblePaths, ok := i.ApplyMethods[method]
	if !ok {
		// everyone can access
		return nil
	}

	i.logger.Infof("Authorization middleware on method: %v", method)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		i.logger.Error("Failed to obtain metadata from method's context")
		return errors.E(opError, errors.InvalidCredentials)
	}

	values := md["authorization"]
	if len(values) == 0 {
		i.logger.Error("Failed to find authorization data in metadata")
		return errors.E(opError, errors.InvalidCredentials)
	}

	accessToken := values[0]
	claims, err := i.JwtManager.Verify(accessToken)
	if err != nil {
		i.logger.Error("Failed to verify token")
		return errors.E(opError, errors.InvalidCredentials, err)
	}

	for _, role := range accessiblePaths {
		if role == claims.Role {
			i.logger.Info("Successfully authorized")
			return nil
		}
	}

	i.logger.Errorf("Permission denied - invalid request role: %v", claims.Role)

	return errors.E(opError, errors.PermissionDenied)
}
