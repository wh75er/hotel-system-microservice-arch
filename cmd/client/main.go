package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "hotel-booking-system/internal/pkg/delivery/grpc/hotel-service/proto"
	"log"
)

func main() {
	fmt.Println("Starting client...")

	conn, err := grpc.Dial(fmt.Sprintf("localhost:3000"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to make grpc client: %v", err)
	}
	defer conn.Close()

	client := pb.NewHotelServiceClient(conn)

	hotels, err := client.GetHotels(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Printf("Error occurred while obtaining hotels: %v\n", err)
	}

	fmt.Printf("Hotels: %v\n", hotels)
}
