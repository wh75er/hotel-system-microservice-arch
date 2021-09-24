package hotel_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	kinds "hotel-booking-system/internal/pkg/errors/hotel-service"
	jwt_manager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type HotelServer struct {
	proto.UnimplementedHotelServiceServer
	HotelUsecase      models.HotelUsecaseI
	ReviewUsecase     models.ReviewUsecaseI
	RoomUsecase       models.RoomUsecaseI
	AdminCredsUsecase models.CredentialsUsecaseI
	TokenManager      *jwt_manager.JWTManager
	Logger            logs.LoggerInterface
}

func NewHotelServer(
	hotelU models.HotelUsecaseI,
	reviewU models.ReviewUsecaseI,
	roomU models.RoomUsecaseI,
	aCredsU models.CredentialsUsecaseI,
	jwtManager *jwt_manager.JWTManager,
	logger logs.LoggerInterface,
) proto.HotelServiceServer {
	return &HotelServer{
		HotelUsecase:      hotelU,
		ReviewUsecase:     reviewU,
		RoomUsecase:       roomU,
		AdminCredsUsecase: aCredsU,
		TokenManager:      jwtManager,
		Logger:            logger,
	}
}

func (s *HotelServer) GetToken(ctx context.Context, pc *proto.Credentials) (*proto.Token, error) {
	c := s.ProtoToCredentials(pc)

	err := s.AdminCredsUsecase.Login(c)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	token, err := s.TokenManager.Generate(models.SERVICE)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	pt := s.TokenToProto(&token)

	return pt, nil
}

func (s *HotelServer) AddHotel(ctx context.Context, ph *proto.Hotel) (*proto.Empty, error) {
	h, err := s.ProtoToHotel(ph)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	err = s.HotelUsecase.AddHotel(h)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return nil, nil
}

func (s *HotelServer) GetHotel(ctx context.Context, u *proto.UUID) (*proto.Hotel, error) {
	h, err := s.HotelUsecase.GetHotel(u.Value)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	ph := s.HotelToProto(&h)

	return ph, nil
}

func (s *HotelServer) GetHotels(ctx context.Context, e *proto.Empty) (*proto.HotelsResponse, error) {
	hotels, err := s.HotelUsecase.GetHotels()
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	ph := s.HotelsToProto(hotels)
	return ph, nil
}

func (s *HotelServer) PatchHotel(ctx context.Context, ph *proto.Hotel) (*proto.Empty, error) {
	h, err := s.ProtoToHotel(ph)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	err = s.HotelUsecase.PatchHotel(h)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return nil, nil
}

func (s *HotelServer) DeleteHotel(ctx context.Context, u *proto.UUID) (*proto.Empty, error) {
	err := s.HotelUsecase.DeleteHotel(u.Value)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return nil, nil
}

func (s *HotelServer) AddReview(ctx context.Context, pr *proto.Review) (*proto.Empty, error) {
	r, err := s.ProtoToReview(pr)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	err = s.ReviewUsecase.AddReview(r)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return nil, nil
}

func (s *HotelServer) GetReview(ctx context.Context, u *proto.UUID) (*proto.Review, error) {
	r, err := s.ReviewUsecase.GetReview(u.Value)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	pr := s.ReviewToProto(&r)

	return pr, nil
}

func (s *HotelServer) GetReviews(ctx context.Context, u *proto.UUID) (*proto.ReviewsResponse, error) {
	r, err := s.ReviewUsecase.GetReviews(u.Value)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	pr := s.ReviewsToProto(r)

	return pr, nil
}

func (s *HotelServer) PatchReview(ctx context.Context, pr *proto.Review) (*proto.Empty, error) {
	r, err := s.ProtoToReview(pr)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	err = s.ReviewUsecase.PatchReview(r)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return nil, nil
}

func (s *HotelServer) DeleteReview(ctx context.Context, u *proto.UUID) (*proto.Empty, error) {
	err := s.ReviewUsecase.DeleteReview(u.Value)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return nil, nil
}

func (s *HotelServer) AddRoom(ctx context.Context, pr *proto.Room) (*proto.Empty, error) {
	r, err := s.ProtoToRoom(pr)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	err = s.RoomUsecase.AddRoom(r)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return nil, nil
}

func (s *HotelServer) GetRooms(ctx context.Context, u *proto.UUID) (*proto.RoomsResponse, error) {
	r, err := s.RoomUsecase.GetRooms(u.Value)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	pr := s.RoomsToProto(r)

	return pr, nil
}

func (s *HotelServer) PatchRoom(ctx context.Context, pr *proto.Room) (*proto.Empty, error) {
	r, err := s.ProtoToRoom(pr)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	err = s.RoomUsecase.PatchRoom(r)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return nil, nil
}

func (s *HotelServer) DeleteRoom(ctx context.Context, u *proto.UUID) (*proto.Empty, error) {
	err := s.RoomUsecase.DeleteRoom(u.Value)
	if err != nil {
		s.Logger.Error("Grpc error: ", err)
		err = status.Error(codes.Code(kinds.GetHttpError(err)), err.Error())
		return nil, err
	}

	return nil, nil
}
