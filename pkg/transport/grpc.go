package transport

import (
	"context"

	kitendpoint "github.com/go-kit/kit/endpoint" // Using alias for go-kit endpoint
	"github.com/mohamedfawas/user-service/pkg/endpoint"
	pb "github.com/mohamedfawas/user-service/proto"
)

type grpcServer struct {
	createUser kitendpoint.Endpoint // Using the aliased import
	getUser    kitendpoint.Endpoint // Using the aliased import
	pb.UnimplementedUserServiceServer
}

func NewGRPCServer(endpoints endpoint.Endpoints) pb.UserServiceServer {
	return &grpcServer{
		createUser: endpoints.CreateUser,
		getUser:    endpoints.GetUser,
	}
}
func (s *grpcServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// Convert proto request to endpoint request
	endpointReq := endpoint.CreateUserRequest{
		Name:  req.Name,
		Email: req.Email,
	}

	response, err := s.createUser(ctx, endpointReq)
	if err != nil {
		return nil, err
	}

	// Convert endpoint response to proto response
	endpointResp := response.(endpoint.CreateUserResponse)
	if endpointResp.Err != nil {
		return nil, endpointResp.Err
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:    endpointResp.User.ID,
			Name:  endpointResp.User.Name,
			Email: endpointResp.User.Email,
		},
	}, nil
}

func (s *grpcServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Convert proto request to endpoint request
	endpointReq := endpoint.GetUserRequest{
		ID: req.Id,
	}

	response, err := s.getUser(ctx, endpointReq)
	if err != nil {
		return nil, err
	}

	// Convert endpoint response to proto response
	endpointResp := response.(endpoint.GetUserResponse)
	if endpointResp.Err != nil {
		return nil, endpointResp.Err
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:    endpointResp.User.ID,
			Name:  endpointResp.User.Name,
			Email: endpointResp.User.Email,
		},
	}, nil
}
