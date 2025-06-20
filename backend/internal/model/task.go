package model

import (
	"time"
)

type Task struct {
	ID          int        `db:"id"`
	UserID      int        `db:"user_id"`
	Title       string     `db:"title"`
	Description *string    `db:"description"`
	Status      string     `db:"status"`
	StartDate   time.Time  `db:"start_date"`
	EndDate     *time.Time `db:"end_date"`
	DueDate     *time.Time `db:"due_date"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
}
