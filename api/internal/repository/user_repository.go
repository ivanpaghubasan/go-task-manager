package repository

import (
	"context"
	"errors"
	"fmt"
	"go-task-manager-api/internal/model"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func (repo *UserRepo) Create(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`

	var userID int

	err := repo.db.QueryRowContext(ctx, query, user.Email, user.PasswordHash).Scan(&userID)
	if err != nil {
		if err == context.DeadlineExceeded {
			return errors.New("Query timed out")
		} else {
			return fmt.Errorf("Failed to insert user data: %v", err)
		}
	}

	return nil
}

func (repo *UserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	query := `SELECT id, email, password_hash, created_at FROM users WHERE email = $1`

	var user model.User

	row := repo.db.QueryRowContext(ctx, query, email)

	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, errors.New("Query timed out")
		} else {
			return nil, fmt.Errorf("Failed to get user data: %v", err)
		}
	}

	return &user, nil

}
