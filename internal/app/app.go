package app

import (
	"go-task-manager/internal/auth"
	"go-task-manager/internal/config"
	"go-task-manager/internal/repository"
	"go-task-manager/internal/service"

	"github.com/jmoiron/sqlx"
)

type Application struct {
	Config     *config.Config
	Service    *service.Service
	JWTManager *auth.JWTManager
}

func New(db *sqlx.DB, cfg *config.Config) (*Application, error) {
	repository := repository.InitRepo(db)
	service := service.InitService(repository)
	jwtManager := auth.NewJWTManager(cfg.JWTSecret)

	return &Application{
		Config:     cfg,
		Service:    service,
		JWTManager: jwtManager,
	}, nil
}
