package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID    string
	Name  string
	Email string
}

type UserService interface {
	CreateUser(ctx context.Context, name, email string) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
}

type userService struct {
	users map[string]*User
}

func NewUserService() UserService {
	return &userService{
		users: make(map[string]*User),
	}
}

func (s *userService) CreateUser(ctx context.Context, name, email string) (*User, error) {
	user := &User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
	s.users[user.ID] = user
	return user, nil
}

func (s *userService) GetUser(ctx context.Context, id string) (*User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}
