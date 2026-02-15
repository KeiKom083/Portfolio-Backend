package di

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/KeiKom083/Portfolio-Backend/internal/infrastructure/database"
	"github.com/KeiKom083/Portfolio-Backend/internal/infrastructure/persistence"
	"github.com/KeiKom083/Portfolio-Backend/internal/interface/graphql/generated"
	"github.com/KeiKom083/Portfolio-Backend/internal/interface/graphql/resolver"
	"github.com/KeiKom083/Portfolio-Backend/internal/usecase"
	"github.com/KeiKom083/Portfolio-Backend/pkg/config"
)

// InitializeServer はすべての依存関係を組み立て、HTTP handler を返す。
func InitializeServer(cfg *config.Config) (http.Handler, func(), error) {
	ctx := context.Background()

	// --- Infrastructure ---
	pool, err := database.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, nil, err
	}

	// --- Repository ---
	userRepo := persistence.NewUserRepository(pool)

	// --- Usecase ---
	userUsecase := usecase.NewUserUsecase(userRepo)

	// --- GraphQL ---
	resolv := resolver.NewResolver(userUsecase)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: resolv,
		}),
	)

	// --- HTTP Router (標準ライブラリ) ---
	mux := http.NewServeMux()
	mux.Handle("/graphql", corsMiddleware(srv))
	mux.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	cleanup := func() {
		pool.Close()
	}

	return mux, cleanup, nil
}

// corsMiddleware は開発用の CORS ミドルウェア。
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
