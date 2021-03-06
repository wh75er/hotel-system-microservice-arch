package hotel_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	"hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/errors"
	jwt_manager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/logs"
	"hotel-booking-system/internal/pkg/models"
)

type HotelServer struct {
	proto.UnimplementedHotelServiceServer
	HotelUsecase      models.HotelUsecaseI
	RoomUsecase       models.RoomUsecaseI
	AdminCredsUsecase models.CredentialsUsecaseI
	TokenManager      *jwt_manager.JWTManager
	Logger            logs.LoggerInterface
}

func NewHotelServer(
	hotelU models.HotelUsecaseI,
	roomU models.RoomUsecaseI,
	aCredsU models.CredentialsUsecaseI,
	jwtManager *jwt_manager.JWTManager,
	logger logs.LoggerInterface,
) proto.HotelServiceServer {
	return &HotelServer{
		HotelUsecase:      hotelU,
		RoomUsecase:       roomU,
		AdminCredsUsecase: aCredsU,
		TokenManager:      jwtManager,
		Logger:            logger,
	}
}

func (s *HotelServer) GetToken(ctx context.Context, pc *commonProto.Credentials) (*commonProto.Token, error) {
	c := commonProto.ProtoToCredentials(pc)

	err := s.AdminCredsUsecase.Login(c)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	token, err := s.TokenManager.Generate(models.SERVICE)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetHttpError(err)), err.Error())
		return nil, err
	}

	pt := commonProto.TokenToProto(&token)

	return pt, nil
}

func (s *HotelServer) AddHotel(ctx context.Context, ph *proto.Hotel) (*commonProto.Empty, error) {
	h, err := ProtoToHotel(ph)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	err = s.HotelUsecase.AddHotel(h)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.Empty{}, nil
}

func (s *HotelServer) GetHotel(ctx context.Context, u *commonProto.UUID) (*proto.Hotel, error) {
	extractedUuid := commonProto.ProtoToUuid(u)
	h, err := s.HotelUsecase.GetHotel(extractedUuid)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	ph := HotelToProto(&h)

	return ph, nil
}

func (s *HotelServer) GetHotels(ctx context.Context, e *commonProto.Empty) (*proto.HotelsResponse, error) {
	hotels, err := s.HotelUsecase.GetHotels()
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	ph := HotelsToProto(hotels)
	return ph, nil
}

func (s *HotelServer) PatchHotel(ctx context.Context, ph *proto.Hotel) (*commonProto.Empty, error) {
	h, err := ProtoToHotel(ph)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	err = s.HotelUsecase.PatchHotel(h)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.Empty{}, nil
}

func (s *HotelServer) DeleteHotel(ctx context.Context, u *commonProto.UUID) (*commonProto.Empty, error) {
	extractedUuid := commonProto.ProtoToUuid(u)
	err := s.HotelUsecase.DeleteHotel(extractedUuid)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.Empty{}, nil
}

func (s *HotelServer) AddRoom(ctx context.Context, pr *proto.Room) (*commonProto.Empty, error) {
	r, err := ProtoToRoom(pr)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	err = s.RoomUsecase.AddRoom(r)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.Empty{}, nil
}

func (s *HotelServer) GetRooms(ctx context.Context, u *commonProto.UUID) (*proto.RoomsResponse, error) {
	extractedUuid := commonProto.ProtoToUuid(u)
	r, err := s.RoomUsecase.GetRooms(extractedUuid)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	pr := RoomsToProto(r)

	return pr, nil
}

func (s *HotelServer) GetRoom(ctx context.Context, u *commonProto.UUID) (*proto.Room, error) {
	extractedUuid := commonProto.ProtoToUuid(u)
	r, err := s.RoomUsecase.GetRoom(extractedUuid)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	pr := RoomToProto(r)

	return pr, nil
}

func (s *HotelServer) TakeRoom(ctx context.Context, roomUuid *commonProto.UUID) (*commonProto.Empty, error) {
	extractedUuid := commonProto.ProtoToUuid(roomUuid)
	err := s.RoomUsecase.TakeRoom(extractedUuid)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.Empty{}, nil
}

func (s *HotelServer) DismissRoom(ctx context.Context, roomUuid *commonProto.UUID) (*commonProto.Empty, error) {
	extractedUuid := commonProto.ProtoToUuid(roomUuid)
	err := s.RoomUsecase.DismissRoom(extractedUuid)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.Empty{}, nil
}

func (s *HotelServer) PatchRoom(ctx context.Context, pr *proto.Room) (*commonProto.Empty, error) {
	r, err := ProtoToRoom(pr)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	err = s.RoomUsecase.PatchRoom(r)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.Empty{}, nil
}

func (s *HotelServer) DeleteRoom(ctx context.Context, u *commonProto.UUID) (*commonProto.Empty, error) {
	extractedUuid := commonProto.ProtoToUuid(u)
	err := s.RoomUsecase.DeleteRoom(extractedUuid)
	if err != nil {
		s.Logger.Errorf("Grpc error: %v - %v {%v}", err, errors.SourceDetails(err), errors.Ops(err))
		err = status.Error(codes.Code(errors.GetKind(err)), err.Error())
		return nil, err
	}

	return &commonProto.Empty{}, nil
}
