package model

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User はユーザーのドメインモデル。
type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// NewUser はパスワードをハッシュ化してユーザーを生成する。
func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}
	return &User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
	}, nil
}
