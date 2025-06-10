package service

import (
	"context"
	"go-task-manager-backend/internal/model"
	"go-task-manager-backend/internal/repository"
)

type TaskService struct {
	service repository.Repository
}

func (s *TaskService) CreateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	return s.service.Task.Create(ctx, task)
}

func (s TaskService) GetAllUserTask(ctx context.Context, userID int) ([]*model.Task, error) {
	return nil, nil
}

func (s TaskService) GetTask(ctx context.Context, taskID int) (*model.Task, error) {
	return s.service.Task.GetTask(ctx, taskID)
}
