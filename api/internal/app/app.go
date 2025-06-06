package app

import (
	"go-task-manager-api/internal/config"
	"go-task-manager-api/internal/repository"
	"go-task-manager-api/internal/service"

	"github.com/jmoiron/sqlx"
)

type Application struct {
	Config  *config.Config
	Service service.Service
}

func New(db *sqlx.DB, cfg *config.Config) (*Application, error) {
	repository := repository.InitRepo(db)
	service := service.InitService(repository)

	return &Application{
		Config:  cfg,
		Service: service,
	}, nil
}
