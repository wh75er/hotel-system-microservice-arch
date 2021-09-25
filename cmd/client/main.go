/*
	Client for development purposes
*/
package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	hotel_service "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service"
	pb "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"hotel-booking-system/internal/pkg/delivery/grpc/interceptors"
	jwtManager "hotel-booking-system/internal/pkg/jwt-manager"
	"hotel-booking-system/internal/pkg/models"
	"log"
	"time"
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

	// Check that interceptor will get token and update it

	hotels, err := client.GetHotels(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Printf("Error occurred while obtaining hotels: %v\n", err)
	}

	fmt.Printf("Hotels: %v\n", hotels)

	// Check that interceptor will make request with existing token

	time.Sleep(2 * time.Second)

	hotels, err = client.GetHotels(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Printf("Error occurred while obtaining hotels: %v\n", err)
	}

	fmt.Printf("Hotels: %v\n", hotels)

	// Check that interceptor will update the token

	time.Sleep(1 * time.Minute)

	hotels, err = client.GetHotels(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Printf("Error occurred while obtaining hotels: %v\n", err)
	}

	fmt.Printf("Hotels: %v\n", hotels)

	//token, err := client.GetToken(context.Background(), &pb.Credentials{Id: "DEVELOP_ID", Secret: "DEVELOP_SECRET"})
	//if err != nil {
	//	fmt.Printf("Error occurred while getting a token: %v\n", err)
	//}
	//
	//fmt.Printf("Token: %v\n", token)
}
