package service

import "go-task-manager-api/internal/repository"

type Service struct {
	User *UserService
	Task *TaskService
}

func InitService(repo repository.Repository) Service {
	return Service{
		User: &UserService{repo},
		Task: &TaskService{repo},
	}
}
