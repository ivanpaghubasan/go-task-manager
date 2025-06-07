package service

import (
	"context"
	"go-task-manager/internal/model"
	"go-task-manager/internal/repository"
)

type UserService struct {
	repo repository.Repository
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	return s.repo.User.Create(ctx, user)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.User.GetUserByEmail(ctx, email)
}
