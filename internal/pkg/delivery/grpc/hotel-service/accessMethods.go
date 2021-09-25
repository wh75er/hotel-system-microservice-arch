package hotel_service

import "hotel-booking-system/internal/pkg/models"

func AccessibleHotelServicePaths() map[string][]models.Role {
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
