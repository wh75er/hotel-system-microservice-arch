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
)

func TestReviewRepository_GetReview(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	loggerMock := &logger.LoggerInterface{}

	loggerMock.On("Errorf", mock2.Anything, mock2.Anything, mock2.Anything).Return()

	r := NewReviewRepository(db, loggerMock)

	reviewUuid := uuid.New()
	userUuid := uuid.New()
	hotelUuid := uuid.New()
	photosUuid := []uuid.UUID{uuid.New(), uuid.New()}
	expectedReview := models.Review{
		UserUuid: userUuid,
		HotelUuid: hotelUuid,
		ReviewUuid: reviewUuid,
		Text: "Test text",
		IsAnonymous: true,
		Photos: photosUuid,
	}

	tests := []struct {
		name    string
		r       models.ReviewRepositoryI
		uuid   	uuid.UUID
		mock    func()
		want    models.Review
		wantErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			r:    r,
			uuid: reviewUuid,
			mock: func() {
				rows := sqlxmock.NewRows([]string{
					"userUuid",
					"hotelUuid",
					"reviewUuid",
					"text",
					"isAnonymous",
					"photos",
				}).AddRow(
					expectedReview.UserUuid,
					expectedReview.HotelUuid,
					expectedReview.ReviewUuid,
					expectedReview.Text,
					expectedReview.IsAnonymous,
					pq.Array(expectedReview.Photos),
				)
				mock.ExpectQuery(
					"SELECT userUuid, hotelUuid, reviewUuid, text, isAnonymous, photos FROM reviews ",
				).WillReturnRows(rows)
			},
			want: expectedReview,
		},
		{
			name: "Review Not Found",
			r:    r,
			uuid: uuid.New(),
			mock: func() {
				rows := sqlxmock.NewRows([]string{
					"userUuid",
					"hotelUuid",
					"reviewUuid",
					"text",
					"isAnonymous",
					"photos",
				})
				mock.ExpectQuery(
					"SELECT userUuid, hotelUuid, reviewUuid, text, isAnonymous, photos FROM reviews",
				).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.r.GetReview(tt.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReview() error new = %v - %v, wantErr %v", err, errors.SourceDetails(err), tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReview() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReviewRepository_GetReviews(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	loggerMock := &logger.LoggerInterface{}

	loggerMock.On("Errorf", mock2.Anything, mock2.Anything, mock2.Anything).Return()

	r := NewReviewRepository(db, loggerMock)

	hotelUuid := uuid.New()
	nonExistingHotelUuid := uuid.New()

	expectedReviews := []models.Review{
		{
			UserUuid: uuid.New(),
			HotelUuid: hotelUuid,
			ReviewUuid: uuid.New(),
			Text: "Test text",
			IsAnonymous: true,
			Photos: []uuid.UUID{uuid.New(), uuid.New()},
		},
		{
			UserUuid: uuid.New(),
			HotelUuid: hotelUuid,
			ReviewUuid: uuid.New(),
			Text: "Test text",
			IsAnonymous: true,
			Photos: []uuid.UUID{uuid.New(), uuid.New()},
		},
	}

	tests := []struct {
		name    string
		r       models.ReviewRepositoryI
		uuid   	uuid.UUID
		mock    func()
		want    []models.Review
		wantErr bool
	}{
		{
			//When everything works as expected
			name: "OK",
			r:    r,
			uuid: hotelUuid,
			mock: func() {
				rows := sqlxmock.NewRows([]string{
					"userUuid",
					"hotelUuid",
					"reviewUuid",
					"text",
					"isAnonymous",
					"photos",
				}).AddRow(
					expectedReviews[0].UserUuid,
					expectedReviews[0].HotelUuid,
					expectedReviews[0].ReviewUuid,
					expectedReviews[0].Text,
					expectedReviews[0].IsAnonymous,
					pq.Array(expectedReviews[0].Photos),
				).AddRow(
					expectedReviews[1].UserUuid,
					expectedReviews[1].HotelUuid,
					expectedReviews[1].ReviewUuid,
					expectedReviews[1].Text,
					expectedReviews[1].IsAnonymous,
					pq.Array(expectedReviews[1].Photos),
				)
				mock.ExpectQuery(
					"SELECT userUuid, hotelUuid, reviewUuid, text, isAnonymous, photos FROM reviews",
				).WithArgs(hotelUuid).WillReturnRows(rows)
			},
			want: expectedReviews,
		},
		{
			name: "Hotel Not Found",
			r:    r,
			uuid: nonExistingHotelUuid,
			mock: func() {
				rows := sqlxmock.NewRows([]string{
					"userUuid",
					"hotelUuid",
					"reviewUuid",
					"text",
					"isAnonymous",
					"photos",
				}).AddRow(
					expectedReviews[0].UserUuid,
					expectedReviews[0].HotelUuid,
					expectedReviews[0].ReviewUuid,
					expectedReviews[0].Text,
					expectedReviews[0].IsAnonymous,
					pq.Array(expectedReviews[0].Photos),
				).AddRow(
					expectedReviews[1].UserUuid,
					expectedReviews[1].HotelUuid,
					expectedReviews[1].ReviewUuid,
					expectedReviews[1].Text,
					expectedReviews[1].IsAnonymous,
					pq.Array(expectedReviews[1].Photos),
				)
				emptyRows := sqlxmock.NewRows([]string{
					"userUuid",
					"hotelUuid",
					"reviewUuid",
					"text",
					"isAnonymous",
					"photos",
				})
				mock.ExpectQuery(
					"SELECT userUuid, hotelUuid, reviewUuid, text, isAnonymous, photos FROM reviews",
				).WithArgs(nonExistingHotelUuid).WillReturnRows(emptyRows)
				mock.ExpectQuery(
					"SELECT userUuid, hotelUuid, reviewUuid, text, isAnonymous, photos FROM reviews",
				).WithArgs(hotelUuid).WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.r.GetReviews(tt.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReviews() error new = %v - %v, wantErr %v", err, errors.SourceDetails(err), tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReviews() = %v, want %v", got, tt.want)
			}
		})
	}
}
