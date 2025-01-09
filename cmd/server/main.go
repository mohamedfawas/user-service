package main

import (
	"log"
	"net"

	"github.com/mohamedfawas/user-service/pkg/endpoint"
	"github.com/mohamedfawas/user-service/pkg/service"
	"github.com/mohamedfawas/user-service/pkg/transport"
	pb "github.com/mohamedfawas/user-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// listen for incoming connections on TCP port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	svc := service.NewUserService()
	endpoints := endpoint.NewEndpoints(svc)

	grpcServer := grpc.NewServer()

	// associates a service implementation with the grpcServer
	// so that the server can handle incoming gRPC requests for the UserService
	pb.RegisterUserServiceServer(grpcServer, transport.NewGRPCServer(endpoints))
	// transport.NewGRPCServer(endpoints) : allows the business logic
	// to be invoked when the gRPC server receives requests

	log.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
