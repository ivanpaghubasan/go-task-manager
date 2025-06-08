package service

import (
	"context"
	"go-task-manager/internal/model"
	"go-task-manager/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type userService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{repo}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	return s.repo.User.Create(ctx, user)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.User.GetUserByEmail(ctx, email)
}
