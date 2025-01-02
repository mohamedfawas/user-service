package main

import (
	"context"
	"log"
	"time"

	pb "github.com/mohamedfawas/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create a user
	createResp, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("Created user: %v", createResp.User)

	// Get the created user
	getResp, err := client.GetUser(ctx, &pb.GetUserRequest{
		Id: createResp.User.Id,
	})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("Retrieved user: %v", getResp.User)
}
