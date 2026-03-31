package model

import "time"

// User はユーザーのドメインモデル。
type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
