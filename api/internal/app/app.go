package app

import (
	"go-task-manager-api/internal/auth"
	"go-task-manager-api/internal/config"
	"go-task-manager-api/internal/repository"
	"go-task-manager-api/internal/service"

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
