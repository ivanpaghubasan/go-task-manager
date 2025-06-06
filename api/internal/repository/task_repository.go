package repository

import (
	"context"
	"fmt"
	"go-task-manager-api/internal/model"

	"github.com/jmoiron/sqlx"
)

type TaskRepo struct {
	Client *sqlx.DB
}

func (repo *TaskRepo) Create(ctx context.Context, task *model.Task) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `INSERT INTO tasks (user_id, title, description, status, start_date, due_date, end_date)
        VALUES (:user_id, :title, :description, :status, :start_date, :due_date, :end_date)`
	rows, err := repo.Client.NamedQueryContext(ctx, query, task)
	if err != nil {
		return nil, fmt.Errorf("faile to insert task: %v", err)
	}

	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("failed to close rows for inserting task: %v", err)
	}

	return task, nil
}

func (repo *TaskRepo) GetTask(ctx context.Context, taskID int) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var task model.Task

	err := repo.Client.GetContext(ctx, &task, "SELECT * FROM tasks WHERE id = $1", taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to get task data: %v", err)
	}

	return &task, nil

}
