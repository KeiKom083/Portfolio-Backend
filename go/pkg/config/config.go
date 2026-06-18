package config

import (
	"fmt"
	"os"
)

// Config はアプリケーション全体の設定を保持する。
type Config struct {
	Port        string
	DatabaseURL string
}

// Load は環境変数から設定を読み込む。
func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
	}, nil
}
