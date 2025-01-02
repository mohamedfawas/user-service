package endpoint

import "github.com/mohamedfawas/user-service/pkg/service"

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
