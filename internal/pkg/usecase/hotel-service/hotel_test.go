package hotel_service

import (
	"github.com/google/uuid"
	mock2 "github.com/stretchr/testify/mock"
	"hotel-booking-system/internal/pkg/errors"
	logger "hotel-booking-system/internal/pkg/logs/mocks"
	"hotel-booking-system/internal/pkg/models"
	repositoryMocks "hotel-booking-system/internal/pkg/models/mocks"
	"reflect"
	"testing"
	"time"
)

func TestHotelUsecase_GetHotel(t *testing.T) {
	loggerMock := &logger.LoggerInterface{}

	loggerMock.On("Error", mock2.Anything, mock2.Anything).Return()
	loggerMock.On("Errorf", mock2.Anything, mock2.Anything, mock2.Anything).Return()

	hotelMockRepository := repositoryMocks.HotelRepositoryI{}
	roomMockRepository := repositoryMocks.RoomRepositoryI{}

	hotelUsecase := NewHotelUsecase(&hotelMockRepository, &roomMockRepository, loggerMock)

	hotelUuid := uuid.New()
	nonExistedHotelUuid := uuid.New()
	hotelUuidWithoutRoomsReviews := uuid.New()
	hotelPhotos := []uuid.UUID{uuid.New(), uuid.New()}

	room := models.Room{
		RoomType:     "family",
		Amount:       4,
		Beds:         5,
		HotelUuid:    hotelUuid,
		RoomUuid:     uuid.New(),
		CreationDate: time.Now(),
		Offers:       []string{"diving", "food included"},
		NightPrice:   3000,
	}

	hotel := models.Hotel{
		Name:        "Name",
		HotelUuid:   hotelUuid,
		Photos:      hotelPhotos,
		Description: "Desc",
		Country:     "Country",
		City:        "City",
		Address:     "Address",
		IsReady:     true,
	}

	expectedHotelWithoutRoomsReviews := hotel

	expectedHotelWithRoomsReviews := models.Hotel{
		Name:        "Name",
		HotelUuid:   hotelUuid,
		Photos:      hotelPhotos,
		Description: "Desc",
		Country:     "Country",
		City:        "City",
		Address:     "Address",
		IsReady:     true,
		Rooms:       []models.Room{room},
	}

	hotelMockRepository.On(
		"GetHotel", hotelUuid,
	).Return(hotel, nil)

	roomMockRepository.On(
		"GetRooms", hotelUuid,
	).Return([]models.Room{room}, nil)

	hotelMockRepository.On(
		"GetHotel", nonExistedHotelUuid,
	).Return(models.Hotel{}, errors.E(errors.RepositoryNoRows))

	roomMockRepository.On(
		"GetRooms", nonExistedHotelUuid,
	).Return([]models.Room{}, errors.E(errors.RepositoryNoRows))

	hotelMockRepository.On(
		"GetHotel", hotelUuidWithoutRoomsReviews,
	).Return(hotel, nil)

	var emptyRooms []models.Room
	roomMockRepository.On(
		"GetRooms", hotelUuidWithoutRoomsReviews,
	).Return(emptyRooms, errors.E(errors.RepositoryNoRows))

	hotelRepoErr := uuid.New()

	var emptyHotel models.Hotel
	hotelMockRepository.On(
		"GetHotel", hotelRepoErr,
	).Return(emptyHotel, errors.E(errors.RepositoryDownErr))

	roomMockRepository.On(
		"GetRooms", hotelRepoErr,
	).Return(emptyRooms, errors.E(errors.RepositoryNoRows))

	RoomRepoErr := uuid.New()

	hotelMockRepository.On(
		"GetHotel", RoomRepoErr,
	).Return(hotel, nil)

	roomMockRepository.On(
		"GetRooms", RoomRepoErr,
	).Return(emptyRooms, errors.E(errors.RepositoryDownErr))

	ReviewRepoErr := uuid.New()

	hotelMockRepository.On(
		"GetHotel", ReviewRepoErr,
	).Return(hotel, nil)

	roomMockRepository.On(
		"GetRooms", ReviewRepoErr,
	).Return(emptyRooms, errors.E(errors.RepositoryNoRows))

	tests := []struct {
		name      string
		hotelU    models.HotelUsecaseI
		hotelUuid string
		want      models.Hotel
		wantErr   bool
	}{
		{
			name:      "OK",
			hotelU:    hotelUsecase,
			hotelUuid: hotelUuid.String(),
			want:      expectedHotelWithRoomsReviews,
		},
		{
			name:      "Hotel Not Found",
			hotelU:    hotelUsecase,
			hotelUuid: nonExistedHotelUuid.String(),
			wantErr:   true,
		},
		{
			name:      "Hotel without rooms and reviews",
			hotelU:    hotelUsecase,
			hotelUuid: hotelUuidWithoutRoomsReviews.String(),
			want:      expectedHotelWithoutRoomsReviews,
		},
		{
			name:      "Hotel Uuid Validation Error",
			hotelU:    hotelUsecase,
			hotelUuid: "invalid uuid",
			wantErr:   true,
		},
		{
			name:      "Hotel Repository Error",
			hotelU:    hotelUsecase,
			hotelUuid: hotelRepoErr.String(),
			wantErr:   true,
		},
		{
			name:      "Room Repository Error",
			hotelU:    hotelUsecase,
			hotelUuid: RoomRepoErr.String(),
			wantErr:   true,
		},
		{
			name:      "Review Repository Error",
			hotelU:    hotelUsecase,
			hotelUuid: ReviewRepoErr.String(),
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.hotelU.GetHotel(tt.hotelUuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHotel() error new = %v - %v, wantErr %v", err, errors.SourceDetails(err), tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHotel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHotelUsecase_GetHotels(t *testing.T) {
	loggerMock := &logger.LoggerInterface{}

	loggerMock.On("Error", mock2.Anything, mock2.Anything).Return()
	loggerMock.On("Errorf", mock2.Anything, mock2.Anything, mock2.Anything).Return()

	hotelMockRepository := repositoryMocks.HotelRepositoryI{}
	roomMockRepository := repositoryMocks.RoomRepositoryI{}

	hotelUsecase := NewHotelUsecase(&hotelMockRepository, &roomMockRepository, loggerMock)

	hotelUuid1 := uuid.New()
	hotelPhotos1 := []uuid.UUID{uuid.New(), uuid.New()}

	expectedHotels := []models.Hotel{
		{
			Name:        "Name",
			HotelUuid:   hotelUuid1,
			Photos:      hotelPhotos1,
			Description: "Desc",
			Country:     "Country",
			City:        "City",
			Address:     "Address",
			IsReady:     true,
		},
		{
			Name:        "Name",
			HotelUuid:   hotelUuid1,
			Photos:      hotelPhotos1,
			Description: "Desc",
			Country:     "Country",
			City:        "City",
			Address:     "Address",
			IsReady:     true,
		},
	}

	tests := []struct {
		name    string
		hotelU  models.HotelUsecaseI
		mock    func()
		want    []models.Hotel
		wantErr bool
	}{
		{
			name:   "OK",
			hotelU: hotelUsecase,
			mock: func() {
				hotelMockRepository.On("GetHotels").Return(expectedHotels, nil).Once()
			},
			want: expectedHotels,
		},
		{
			name:   "Hotel Repository Error",
			hotelU: hotelUsecase,
			mock: func() {
				var emptyHotels []models.Hotel
				hotelMockRepository.On("GetHotels").Return(emptyHotels, errors.E(errors.RepositoryDownErr)).Once()
			},
			wantErr: true,
		},
		{
			name:   "No Hotels Found",
			hotelU: hotelUsecase,
			mock: func() {
				var emptyHotels []models.Hotel
				hotelMockRepository.On("GetHotels").Return(emptyHotels, nil).Once()
			},
			want: nil,
		},
		{
			name:   "Hotel Repository No Rows Err",
			hotelU: hotelUsecase,
			mock: func() {
				var emptyHotels []models.Hotel
				hotelMockRepository.On("GetHotels").Return(emptyHotels, errors.E(errors.RepositoryNoRows)).Once()
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.hotelU.GetHotels()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHotels() error new = %v - %v, wantErr %v", err, errors.SourceDetails(err), tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHotels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHotelUsecase_AddHotel(t *testing.T) {
	loggerMock := &logger.LoggerInterface{}

	loggerMock.On("Error", mock2.Anything, mock2.Anything).Return()
	loggerMock.On("Errorf", mock2.Anything, mock2.Anything, mock2.Anything).Return()

	hotelMockRepository := repositoryMocks.HotelRepositoryI{}
	roomMockRepository := repositoryMocks.RoomRepositoryI{}

	hotelUsecase := NewHotelUsecase(&hotelMockRepository, &roomMockRepository, loggerMock)

	hotelUuid1 := uuid.New()
	hotelPhotos1 := []uuid.UUID{uuid.New(), uuid.New()}

	expectedHotels := []models.Hotel{
		{
			Name:        "Name",
			HotelUuid:   hotelUuid1,
			Photos:      hotelPhotos1,
			Description: "Desc",
			Country:     "Country",
			City:        "City",
			Address:     "Address",
			IsReady:     true,
		},
		{
			Name:        "Name",
			HotelUuid:   hotelUuid1,
			Photos:      hotelPhotos1,
			Description: "Desc",
			Country:     "Country",
			City:        "City",
			Address:     "Address",
			IsReady:     true,
		},
	}

	tests := []struct {
		name    string
		hotelU  models.HotelUsecaseI
		mock    func()
		want    []models.Hotel
		wantErr bool
	}{
		{
			name:   "OK",
			hotelU: hotelUsecase,
			mock: func() {
				hotelMockRepository.On("GetHotels").Return(expectedHotels, nil).Once()
			},
			want: expectedHotels,
		},
		{
			name:   "Hotel Repository Error",
			hotelU: hotelUsecase,
			mock: func() {
				var emptyHotels []models.Hotel
				hotelMockRepository.On("GetHotels").Return(emptyHotels, errors.E(errors.RepositoryDownErr)).Once()
			},
			wantErr: true,
		},
		{
			name:   "No Hotels Found",
			hotelU: hotelUsecase,
			mock: func() {
				var emptyHotels []models.Hotel
				hotelMockRepository.On("GetHotels").Return(emptyHotels, nil).Once()
			},
			want: nil,
		},
		{
			name:   "Hotel Repository No Rows Err",
			hotelU: hotelUsecase,
			mock: func() {
				var emptyHotels []models.Hotel
				hotelMockRepository.On("GetHotels").Return(emptyHotels, errors.E(errors.RepositoryNoRows)).Once()
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.hotelU.GetHotels()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHotels() error new = %v - %v, wantErr %v", err, errors.SourceDetails(err), tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHotels() = %v, want %v", got, tt.want)
			}
		})
	}
}
