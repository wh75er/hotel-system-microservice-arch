/*
	Client for development purposes
*/
package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"hotel-booking-system/internal/pkg/delivery/grpc/commonProto"
	hotel_service "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service"
	pb "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/interceptors"
	jwtManager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/models"
	"log"
)

func main() {
	fmt.Println("Starting client...")

	jwtTokenManager := jwtManager.NewJWTManager("", 0)

	// Create interceptor without client getToken API
	authInterceptor := interceptors.NewClientAuthInterceptor(
		models.Credentials{"DEVELOP_ID", "DEVELOP_SECRET"},
		jwtTokenManager,
		interceptors.MethodsRoleMapToSet(hotel_service.AccessibleHotelServicePaths()),
		logrus.New(),
	)

	conn, err := grpc.Dial(
		fmt.Sprintf("localhost:3000"),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(authInterceptor.Unary()),
	)
	if err != nil {
		log.Fatalf("Failed to make grpc client: %v", err)
	}
	defer conn.Close()

	client := pb.NewHotelServiceClient(conn)

	// Add client AddToken API to auth interceptor
	authInterceptor.GrpcServiceClient = client

	/* Add hotel case */
	//_, err = client.AddHotel(context.Background(), &pb.Hotel{
	//	Name: "test",
	//	HotelUuid: "c6df5496-7146-43ba-8e80-0a22bcaaf6bb",
	//	Description: "test",
	//	Country: "test",
	//	City: "test",
	//	Address: "test",
	//	Rooms: nil,
	//})
	//if err != nil {
	//	fmt.Printf("Failed to create hotel: %v, code: %v\n", err, status.Code(err))
	//}

	/* Get all hotels case */
	hotels, err := client.GetHotels(context.Background(), &commonProto.Empty{})
	if err != nil {
		fmt.Printf("Failed to get hotels: %v, code: %v\n", err, status.Code(err))
	}

	fmt.Println("Hotels: ", hotels)

	/* Get hotel by uuid */
	//hotel, err := client.GetHotel(context.Background(), &pb.UUID{Value: "85bf1386-e41b-4cfd-9cff-f4448c057ce2"})
	//if err != nil {
	//	fmt.Printf("Failed to get hotel: %v, code: %v\n", err, status.Code(err))
	//}
	//
	//fmt.Println("Hotel: ", hotel)

	/* Token middleware update check */

	//// Check that interceptor will get token and update it
	//
	//hotels, err := client.GetHotels(context.Background(), &pb.Empty{})
	//if err != nil {
	//	fmt.Printf("Error occurred while obtaining hotels: %v\n", err)
	//}
	//
	//fmt.Printf("Hotels: %v\n", hotels)
	//
	//// Check that interceptor will make request with existing token
	//
	//time.Sleep(2 * time.Second)
	//
	//hotels, err = client.GetHotels(context.Background(), &pb.Empty{})
	//if err != nil {
	//	fmt.Printf("Error occurred while obtaining hotels: %v\n", err)
	//}
	//
	//fmt.Printf("Hotels: %v\n", hotels)
	//
	//// Check that interceptor will update the token
	//
	//time.Sleep(1 * time.Minute)
	//
	//hotels, err = client.GetHotels(context.Background(), &pb.Empty{})
	//if err != nil {
	//	fmt.Printf("Error occurred while obtaining hotels: %v\n", err)
	//}
	//
	//fmt.Printf("Hotels: %v\n", hotels)

	/* Token Generation Case */

	//token, err := client.GetToken(context.Background(), &pb.Credentials{Id: "DEVELOP_ID", Secret: "DEVELOP_SECRET"})
	//if err != nil {
	//	fmt.Printf("Error occurred while getting a token: %v\n", err)
	//}
	//
	//fmt.Printf("Token: %v\n", token)
}
