package service

import "go-task-manager-api/internal/repository"

type Service struct {
	User *UserService
}

func InitService(repo repository.Repository) Service {
    return Service{
        User: &UserService{repo: repo},
    }
}
