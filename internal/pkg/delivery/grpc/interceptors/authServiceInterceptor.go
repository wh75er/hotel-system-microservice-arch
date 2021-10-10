/*
		This interceptor provides application layer authentication.
	It uses auth service to validate token and check token's role.
*/
package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	proto1 "hotel-booking-system/internal/pkg/delivery/grpc/auth-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	"hotel-booking-system/internal/pkg/errors"
	jwt_manager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
	"log"
)

type AuthServiceInterceptor struct {
	JwtManager        *jwt_manager.JWTManager
	UserServiceClient proto1.AuthServiceClient
	ApplyMethods      map[string][]models.Role
	logger            logs.LoggerInterface
}

func NewAuthServiceInterceptor(
	jwtManager *jwt_manager.JWTManager,
	userClient proto1.AuthServiceClient,
	applyMethods map[string][]models.Role,
	logger logs.LoggerInterface,
) *AuthServiceInterceptor {
	return &AuthServiceInterceptor{
		jwtManager,
		userClient,
		applyMethods,
		logger,
	}
}

func (i *AuthServiceInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("client unary interceptor: ", info.FullMethod)

		err := i.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		}

		return handler(ctx, req)
	}
}

func (i *AuthServiceInterceptor) authorize(ctx context.Context, method string) error {
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
	userRole, err := i.UserServiceClient.CheckAuth(context.Background(), &commonProto.Token{Value: accessToken})
	if err != nil {
		i.logger.Errorf("Authorization error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		stat, _ := status.FromError(err)
		//err = status.Error(codes.Code(errors.GetHttpError(errors.E(errors.Kind(stat.Code())))), stat.Message())
		return errors.E(opError, errors.Kind(stat.Code()))
	}

	parsedRole := models.Role(userRole.Value)

	for _, role := range accessiblePaths {
		if role == parsedRole {
			i.logger.Info("Successfully authorized")
			return nil
		}
	}

	i.logger.Errorf("Permission denied - invalid request role: %v", parsedRole)

	return errors.E(opError, errors.PermissionDenied)
}
