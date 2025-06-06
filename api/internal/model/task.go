package model

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          int            `db:"id"`
	UserID      int            `db:"user_id"`
	Title       string         `db:"title"`
	Description sql.NullString `db:"description"`
	Status      string         `db:"status"`
	StartDate   time.Time      `db:"start_date"`
	EndDate     time.Time      `db:"end_date"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
}
