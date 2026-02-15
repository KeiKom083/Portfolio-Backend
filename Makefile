.PHONY: generate run build migrate help

## gqlgen でコード生成
generate:
	go run github.com/99designs/gqlgen generate

## 開発サーバー起動
run:
	go run cmd/server/main.go

## バイナリビルド
build:
	go build -o bin/server cmd/server/main.go

## マイグレーション実行 (psql)
migrate:
	psql $(DATABASE_URL) -f migrations/001_create_users.sql

## ヘルプ
help:
	@echo "make generate  - gqlgen コード生成"
	@echo "make run       - 開発サーバー起動"
	@echo "make build     - バイナリビルド"
	@echo "make migrate   - DBマイグレーション"
