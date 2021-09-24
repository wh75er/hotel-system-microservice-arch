package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"hotel-booking-system/internal/pkg/errors"
	jwt_manager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/models"
	"log"
)

//func AuthInterceptor(
//	ctx context.Context,
//	req interface{},
//	info *grpc.UnaryServerInfo,
//	handler grpc.UnaryHandler,
//) (interface{}, error) {
//	log.Println("--> unary interceptor: ", info.FullMethod)
//	return handler(ctx, req)
//}

type ServerAdminAuthInterceptor struct {
	JwtManager      *jwt_manager.JWTManager
	AccessiblePaths map[string][]models.Role
}

func NewServerAdminAuthInterceptor(jwtManager *jwt_manager.JWTManager) *ServerAdminAuthInterceptor {
	return &ServerAdminAuthInterceptor{
		jwtManager,
		accessibleRoles(),
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

	accessiblePaths, ok := i.AccessiblePaths[method]
	if !ok {
		// everyone can access
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.E(opError, errors.InvalidCredentials)
	}

	values := md["authorization"]
	if len(values) == 0 {
		return errors.E(opError, errors.InvalidCredentials)
	}

	accessToken := values[0]
	claims, err := i.JwtManager.Verify(accessToken)
	if err != nil {
		return errors.E(opError, errors.InvalidCredentials, err)
	}

	for _, role := range accessiblePaths {
		if role == claims.Role {
			return nil
		}
	}

	return errors.E(opError, errors.PermissionDenied)
}

func accessibleRoles() map[string][]models.Role {
	const hotelServicePath = "/proto.HotelService/"

	return map[string][]models.Role{
		hotelServicePath + "AddHotel":     {models.SERVICE},
		hotelServicePath + "GetHotel":     {models.SERVICE},
		hotelServicePath + "GetHotels":    {models.SERVICE},
		hotelServicePath + "PatchHotel":   {models.SERVICE},
		hotelServicePath + "DeleteHotel":  {models.SERVICE},
		hotelServicePath + "AddReview":    {models.SERVICE},
		hotelServicePath + "GetReview":    {models.SERVICE},
		hotelServicePath + "GetReviews":   {models.SERVICE},
		hotelServicePath + "PatchReview":  {models.SERVICE},
		hotelServicePath + "DeleteReview": {models.SERVICE},
		hotelServicePath + "AddRoom":      {models.SERVICE},
		hotelServicePath + "GetRooms":     {models.SERVICE},
		hotelServicePath + "PatchRoom":    {models.SERVICE},
		hotelServicePath + "DeleteRoom":   {models.SERVICE},
	}
}
