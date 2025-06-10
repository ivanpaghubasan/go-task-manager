package model

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int       `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

func (u *User) HashPassword(plainText string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Encryption failed: %v", err)
	}
	u.PasswordHash = string(hashedPassword)

	return nil
}

func (u *User) VerifyPassword(plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(plainText))
	if err != nil {
		return false
	}
	return true
}
