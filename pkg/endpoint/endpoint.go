package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mohamedfawas/user-service/pkg/service"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

type CreateUserRequest struct {
	Name  string
	Email string
}

type CreateUserResponse struct {
	User *service.User
	Err  error
}

type GetUserRequest struct {
	ID string
}

type GetUserResponse struct {
	User *service.User
	Err  error
}

func MakeCreateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		user, err := svc.CreateUser(ctx, req.Name, req.Email)
		return CreateUserResponse{User: user, Err: err}, nil
	}
}

func MakeGetUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		user, err := svc.GetUser(ctx, req.ID)
		return GetUserResponse{User: user, Err: err}, nil
	}
}

func NewEndpoints(svc service.UserService) Endpoints {
	return Endpoints{
		CreateUser: MakeCreateUserEndpoint(svc),
		GetUser:    MakeGetUserEndpoint(svc),
	}
}
