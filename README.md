# Portfolio-Backend

Go + GraphQL (gqlgen) + PostgreSQL によるクリーンアーキテクチャのバックエンド API。

## アーキテクチャ

```
internal/
├── domain/              # ドメイン層（ビジネスルールの中心）
│   ├── model/           #   エンティティ・値オブジェクト
│   └── repository/      #   リポジトリインターフェース
├── usecase/             # ユースケース層（アプリケーションロジック）
├── infrastructure/      # インフラ層（外部依存の実装）
│   ├── database/        #   DB接続
│   └── persistence/     #   リポジトリ実装
├── interface/           # インターフェース層（外部とのやり取り）
│   └── graphql/         #   GraphQL スキーマ・リゾルバ
└── di/                  # 依存性注入（組み立て）
```

**依存の方向**: `interface → usecase → domain ← infrastructure`

## セットアップ

```bash
# 1. 依存パッケージのインストール
go mod tidy

# 2. gqlgen コード生成
make generate

# 3. 環境変数の設定
cp .env.example .env
# .env を編集して DATABASE_URL を設定

# 4. DBマイグレーション
make migrate

# 5. サーバー起動
make run
```

## エンドポイント

| パス | 説明 |
|------|------|
| `/graphql` | GraphQL API |
| `/playground` | GraphQL Playground (開発用 UI) |
| `/health` | ヘルスチェック |

## 開発コマンド

```bash
make generate  # GraphQL スキーマからコード生成
make run       # 開発サーバー起動
make build     # バイナリビルド
make migrate   # DBマイグレーション
```
