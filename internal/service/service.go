package service

import "go-task-manager/internal/repository"

type Service struct {
	User UserService
	Task TaskService
}

func InitService(repo repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo),
		Task: TaskService{repo},
	}
}
