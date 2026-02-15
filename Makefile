.PHONY: generate run build migrate docker-up docker-down docker-logs docker-ps help

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

## Docker起動
docker-up:
	docker compose up -d

## Docker停止
docker-down:
	docker compose down

## Dockerログ確認
docker-logs:
	docker compose logs -f postgres

## Dockerコンテナ状態確認
docker-ps:
	docker compose ps

## ヘルプ
help:
	@echo "make generate    - gqlgen コード生成"
	@echo "make run         - 開発サーバー起動"
	@echo "make build       - バイナリビルド"
	@echo "make migrate     - DBマイグレーション"
	@echo "make docker-up   - Docker起動"
	@echo "make docker-down - Docker停止"
	@echo "make docker-logs - Dockerログ確認"
	@echo "make docker-ps   - Dockerコンテナ状態確認"
