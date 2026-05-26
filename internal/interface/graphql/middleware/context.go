package middleware

import (
	"context"
	"net/http"
)

type contextKey string

const (
	responseWriterKey contextKey = "responseWriter"
	requestKey        contextKey = "request"
)

// Inject は http.Request と http.ResponseWriter を context に詰めるミドルウェア。
func Inject(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), responseWriterKey, w)
		ctx = context.WithValue(ctx, requestKey, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ResponseWriter は context から http.ResponseWriter を取り出す。
func ResponseWriter(ctx context.Context) http.ResponseWriter {
	w, _ := ctx.Value(responseWriterKey).(http.ResponseWriter)
	return w
}

// Request は context から *http.Request を取り出す。
func Request(ctx context.Context) *http.Request {
	r, _ := ctx.Value(requestKey).(*http.Request)
	return r
}
