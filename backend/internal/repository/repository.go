package repository

import (
	"context"
	"go-task-manager-backend/internal/model"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	QueryTimeoutDuration = time.Second * 5
)

type Repository struct {
	User interface {
		Create(ctx context.Context, user *model.User) error
		GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	}
	Task interface {
		Create(ctx context.Context, task *model.Task) (*model.Task, error)
		GetTask(ctx context.Context, taskID int) (*model.Task, error)
	}
}

func InitRepo(db *sqlx.DB) Repository {
	return Repository{
		User: &UserRepo{db},
		Task: &TaskRepo{db},
	}
}
