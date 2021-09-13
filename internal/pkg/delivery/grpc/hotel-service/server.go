package hotel_service

import (
	"context"
	"hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/models"
)

type HotelServer struct {
	proto.UnimplementedHotelServiceServer
	hotelUsecase models.HotelUsecaseI
	reviewUsecase models.ReviewUsecaseI
	roomUsecase models.RoomUsecaseI
}

func (s *HotelServer) AddHotel(context.Context, *proto.Hotel) (*proto.Empty, error) {
}

func (s *HotelServer) GetHotel(context.Context, *proto.UUID) (*proto.Hotel, error) {
}

func (s *HotelServer) GetHotels(context.Context, *proto.Empty) (*proto.HotelsResponse, error) {
}

func (s *HotelServer) PatchHotel(context.Context, *proto.Hotel) (*proto.Empty, error) {
}

func (s *HotelServer) DeleteHotel(context.Context, *proto.UUID) (*proto.Empty, error) {
}

func (s *HotelServer) AddReview(context.Context, *proto.Review) (*proto.Empty, error) {
}

func (s *HotelServer) GetReview(context.Context, *proto.UUID) (*proto.Review, error) {
}

func (s *HotelServer) GetReviews(context.Context, *proto.UUID) (*proto.ReviewsResponse, error) {
}

func (s *HotelServer) PatchReview(context.Context, *proto.Review) (*proto.Empty, error) {
}

func (s *HotelServer) DeleteReview(context.Context, *proto.UUID) (*proto.Empty, error) {
}

func (s *HotelServer) AddRoom(context.Context, *proto.Room) (*proto.Empty, error) {
}

func (s *HotelServer) GetRooms(context.Context, *proto.UUID) (*proto.RoomsResponse, error) {
}

func (s *HotelServer) PatchRoom(context.Context, *proto.Room) (*proto.Empty, error) {
}

func (s *HotelServer) DeleteRoom(context.Context, *proto.UUID) (*proto.Empty, error) {
}

