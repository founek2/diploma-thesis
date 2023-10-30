package middleware

import (
	"context"
	"fmt"
	"monolith/server/database"
	"net/http"

	"go.opentelemetry.io/otel/attribute"

	"go.opentelemetry.io/otel/trace"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func AddTracingAndDb(method string, path string) func(next http.Handler, tracer trace.Tracer, db *database.Database) http.Handler {
	return func(next http.Handler, tracer trace.Tracer, db *database.Database) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			ctx, main := tracer.Start(ctx, fmt.Sprintf("%s %s", method, path))
			defer main.End()

			ctx = context.WithValue(ctx, "tracer", tracer)
			ctx = context.WithValue(ctx, "span", main)
			main.SetAttributes(
				attribute.String("http.method", r.Method),
				attribute.String("http.route", r.RequestURI),
				attribute.String("http.url", r.URL.String()),
			)

			lrw := NewLoggingResponseWriter(w)
			next.ServeHTTP(lrw, r.WithContext(ctx))

			main.SetAttributes(attribute.Int("http.status_code", lrw.statusCode))
		})
	}
}

func GetDb(ctx context.Context) *database.Database {
	return ctx.Value("db").(*database.Database)
}
