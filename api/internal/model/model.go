package model

import "time"

type User struct {
	ID           int       `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

func (u *User) HashPassword(plainText string) string {
	return ""
}

func (u *User) MatchPassword(plainText string) bool {
	return false
}
