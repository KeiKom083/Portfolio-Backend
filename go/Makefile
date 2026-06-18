.PHONY: generate run build migrate docker-up docker-down docker-logs docker-ps help

## gqlgen でコード生成
generate:
	go run github.com/99designs/gqlgen generate

## 開発サーバー起動
run:
	export $(shell grep -v '^#' .env | xargs) && go run cmd/server/main.go

## バイナリビルド
build:
	go build -o bin/server cmd/server/main.go

## マイグレーション実行 (Docker経由)
migrate:
	docker exec -i portfolio-postgres psql -U postgres -d portfolio -f - < migrations/001_create_users.sql

## Docker起動
docker-up:
	docker compose -f ../docker-compose.yml up -d

## Docker停止
docker-down:
	docker compose -f ../docker-compose.yml down

## Dockerログ確認
docker-logs:
	docker compose -f ../docker-compose.yml logs -f postgres

## Dockerコンテナ状態確認
docker-ps:
	docker compose -f ../docker-compose.yml ps

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
