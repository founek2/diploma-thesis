package getters

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

func GetTracer(ctx context.Context) trace.Tracer {
	return ctx.Value("tracer").(trace.Tracer)
}
func GetUserId(r *http.Request) string {
	userId := r.Header.Get("User-Id")
	if userId == "" {
		panic("Required header User-Id not set")
	}

	return userId
}
