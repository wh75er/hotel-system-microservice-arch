package hotel_service

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	mock2 "github.com/stretchr/testify/mock"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"hotel-booking-system/internal/pkg/errors"
	logger "hotel-booking-system/internal/pkg/logs/mocks"
	"hotel-booking-system/internal/pkg/models"
	"reflect"
	"testing"
	"time"
)

func TestHotelRepository_GetHotel(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	loggerMock := &logger.LoggerInterface{}

	loggerMock.On("Errorf", mock2.Anything, mock2.Anything, mock2.Anything).Return()

	r := NewHotelRepository(db, loggerMock)

	hotelUuid := uuid.New()
	photosUuid := []uuid.UUID{uuid.New(), uuid.New()}
	expectedHotel := models.Hotel{
		Name: "Name",
		HotelUuid: hotelUuid,
		Photos: photosUuid,
		Description: "Desc",
		Country: "Country",
		City: "City",
		Address: "Address",
		IsReady: true,
	}

	tests := []struct {
		name    string
		r       models.HotelRepositoryI
		uuid   	uuid.UUID
		mock    func()
		want    models.Hotel
		wantErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			r:    r,
			uuid: hotelUuid,
			mock: func() {
				rows := sqlxmock.NewRows([]string{
					"name",
					"hotelUuid",
					"photos",
					"description",
					"country",
					"city",
					"address",
					"isReady",
				}).AddRow(
					expectedHotel.Name,
					expectedHotel.HotelUuid,
					pq.Array(expectedHotel.Photos),
					expectedHotel.Description,
					expectedHotel.Country,
					expectedHotel.City,
					expectedHotel.Address,
					expectedHotel.IsReady,
				)
				mock.ExpectQuery(
					"SELECT name, hotelUuid, photos, description, country, city, address",
				).WillReturnRows(rows)
			},
			want: expectedHotel,
		},
		{
			name: "Hotel Not Found",
			r:    r,
			uuid: uuid.New(),
			mock: func() {
				rows := sqlxmock.NewRows([]string{
					"name",
					"hotelUuid",
					"photos",
					"description",
					"country",
					"city",
					"address",
					"isReady",
				})
				mock.ExpectQuery(
					"SELECT name, hotelUuid, photos, description, country, city, address",
				).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.r.GetHotel(tt.uuid)
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

func TestHotelRepository_GetHotels(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	loggerMock := &logger.LoggerInterface{}

	loggerMock.On("Errorf", mock2.Anything, mock2.Anything, mock2.Anything).Return()

	r := NewHotelRepository(db, loggerMock)

	hotelUuid1 := uuid.New()
	hotelUuid2 := uuid.New()
	photosUuid1 := []uuid.UUID{uuid.New(), uuid.New()}
	photosUuid2 := []uuid.UUID{uuid.New(), uuid.New()}
	expectedHotels := []models.Hotel{
		{
			Name: "Name1",
			HotelUuid: hotelUuid1,
			Photos: photosUuid1,
			Description: "Desc1",
			Country: "Country1",
			City: "City1",
			Address: "Address1",
			IsReady: true,
		},
		{
			Name: "Name2",
			HotelUuid: hotelUuid2,
			Photos: photosUuid2,
			Description: "Desc2",
			Country: "Country2",
			City: "City2",
			Address: "Address2",
			IsReady: true,
		},
	}

	tests := []struct {
		name    string
		r       models.HotelRepositoryI
		uuid   	uuid.UUID
		mock    func()
		want    []models.Hotel
		wantErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			r:    r,
			mock: func() {
				rows := sqlxmock.NewRows([]string{
					"name",
					"hotelUuid",
					"photos",
					"description",
					"country",
					"city",
					"address",
					"isReady",
				}).AddRow(
					expectedHotels[0].Name,
					expectedHotels[0].HotelUuid,
					pq.Array(expectedHotels[0].Photos),
					expectedHotels[0].Description,
					expectedHotels[0].Country,
					expectedHotels[0].City,
					expectedHotels[0].Address,
					expectedHotels[0].IsReady,
				).AddRow(
					expectedHotels[1].Name,
					expectedHotels[1].HotelUuid,
					pq.Array(expectedHotels[1].Photos),
					expectedHotels[1].Description,
					expectedHotels[1].Country,
					expectedHotels[1].City,
					expectedHotels[1].Address,
					expectedHotels[1].IsReady,
				)
				mock.ExpectQuery(
					"SELECT name, hotelUuid, photos, description, country, city, address",
				).WillReturnRows(rows)
			},
			want: expectedHotels,
		},
		{
			name: "No Hotels Exists",
			r:    r,
			mock: func() {
				rows := sqlxmock.NewRows([]string{
					"name",
					"hotelUuid",
					"photos",
					"description",
					"country",
					"city",
					"address",
					"isReady",
				})
				mock.ExpectQuery(
					"SELECT name, hotelUuid, photos, description, country, city, address",
				).WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.r.GetHotels()
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

func TestHotelRepository_AddHotel(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	loggerMock := &logger.LoggerInterface{}

	loggerMock.On("Errorf", mock2.Anything, mock2.Anything, mock2.Anything).Return()

	r := NewHotelRepository(db, loggerMock)

	hotelUuid := uuid.New()
	photosUuid := []uuid.UUID{uuid.New(), uuid.New()}
	hotel := &models.Hotel{
		Name: "Name",
		HotelUuid: hotelUuid,
		Photos: photosUuid,
		Description: "Desc",
		Country: "Country",
		City: "City",
		Address: "Address",
		IsReady: true,
		CreationDate: time.Now(),
	}

	tests := []struct {
		name    string
		r       models.HotelRepositoryI
		hotel 	models.Hotel
		mock    func()
		want    []models.Hotel
		wantErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			r:    r,
			hotel: *hotel,
			mock: func() {
				result := sqlxmock.NewResult(1, 1)
				mock.ExpectExec(
					"INSERT INTO hotels",
				).WithArgs(
					hotel.Name,
					hotel.HotelUuid,
					pq.Array(hotel.Photos),
					hotel.Description,
					hotel.Country,
					hotel.City,
					hotel.Address,
					hotel.IsReady,
					hotel.CreationDate,
				).WillReturnResult(result)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := tt.r.AddHotel(hotel)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddHotel() error new = %v - %v, wantErr %v", err, errors.SourceDetails(err), tt.wantErr)
				return
			}
		})
	}
}

func TestHotelRepository_PatchHotel(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	loggerMock := &logger.LoggerInterface{}

	loggerMock.On("Errorf", mock2.Anything, mock2.Anything, mock2.Anything).Return()

	r := NewHotelRepository(db, loggerMock)

	hotelUuidOrigin := uuid.New()
	photosUuidOrigin := []uuid.UUID{uuid.New(), uuid.New()}
	hotelOrigin := models.Hotel{
		Name: "Name",
		HotelUuid: hotelUuidOrigin,
		Photos: photosUuidOrigin,
		Description: "Desc",
		Country: "Country",
		City: "City",
		Address: "Address",
		IsReady: true,
		CreationDate: time.Now(),
	}

	patchedHotel := hotelOrigin
	patchedHotel.Name = "NewName"


	otherHotelUuid := uuid.New()
	if otherHotelUuid == hotelUuidOrigin {
		t.Error("PatchHotels() expected different uuids in otherHotelUUid and hotelUuidOrigin")
		return
	}

	patchedHotelWithAnotherUuid := patchedHotel
	patchedHotelWithAnotherUuid.HotelUuid = otherHotelUuid

	tests := []struct {
		name    string
		r       models.HotelRepositoryI
		hotel 	models.Hotel
		mock    func()
		want    []models.Hotel
		wantErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			r:    r,
			hotel: patchedHotel,
			mock: func() {
				result := sqlxmock.NewResult(1, 1)
				mock.ExpectExec(
					"UPDATE hotels SET",
				).WithArgs(
					patchedHotel.Name,
					pq.Array(patchedHotel.Photos),
					patchedHotel.Description,
					patchedHotel.Country,
					patchedHotel.City,
					patchedHotel.Address,
					patchedHotel.IsReady,
					patchedHotel.HotelUuid,
				).WillReturnResult(result)
			},
			wantErr: false,
		},
		{
			name: "Hotel Not Found",
			r:    r,
			hotel: patchedHotelWithAnotherUuid,
			mock: func() {
				result := sqlxmock.NewResult(1, 0)
				mock.ExpectExec(
					"UPDATE hotels SET",
				).WithArgs(
					patchedHotelWithAnotherUuid.Name,
					pq.Array(patchedHotelWithAnotherUuid.Photos),
					patchedHotelWithAnotherUuid.Description,
					patchedHotelWithAnotherUuid.Country,
					patchedHotelWithAnotherUuid.City,
					patchedHotelWithAnotherUuid.Address,
					patchedHotelWithAnotherUuid.IsReady,
					patchedHotelWithAnotherUuid.HotelUuid,
				).WillReturnResult(result)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := tt.r.PatchHotel(&patchedHotel)
			if (err != nil) != tt.wantErr {
				t.Errorf("PatchHotel() error new = %v - %v, wantErr %v", err, errors.SourceDetails(err), tt.wantErr)
				return
			}
		})
	}
}

func TestHotelRepository_Delete(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	loggerMock := &logger.LoggerInterface{}

	loggerMock.On("Errorf", mock2.Anything, mock2.Anything, mock2.Anything).Return()

	r := NewHotelRepository(db, loggerMock)

	hotelUuid := uuid.New()

	tests := []struct {
		name    string
		r       models.HotelRepositoryI
		uuid   	uuid.UUID
		mock    func()
		want    []models.Hotel
		wantErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			r:    r,
			uuid: hotelUuid,
			mock: func() {
				result := sqlxmock.NewResult(1, 1)
				mock.ExpectExec(
					"DELETE FROM hotels WHERE",
				).WithArgs(hotelUuid).WillReturnResult(result)
			},
		},
		{
			name: "Hotel Not Found",
			r:    r,
			uuid: hotelUuid,
			mock: func() {
				result := sqlxmock.NewResult(1, 0)
				mock.ExpectExec(
					"DELETE FROM hotels WHERE",
				).WithArgs(hotelUuid).WillReturnResult(result)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := tt.r.DeleteHotel(hotelUuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteHotel() error new = %v - %v, wantErr %v", err, errors.SourceDetails(err), tt.wantErr)
				return
			}
		})
	}
}
