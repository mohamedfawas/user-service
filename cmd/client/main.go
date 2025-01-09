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
	//establish connection to grpc server on 50051 (using insecure credentials)
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() // connection is closed when function exits

	// create a new client for the userservice
	client := pb.NewUserServiceClient(conn)

	// context with 1 second for grpc calls
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // context is cancelled once function exits

	// request to create a new user with specified name and email
	createResp, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("Created user: %v", createResp.User)

	// request to retrieve the user details using the ID of the created user
	getResp, err := client.GetUser(ctx, &pb.GetUserRequest{
		Id: createResp.User.Id,
	})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("Retrieved user: %v", getResp.User)
}
