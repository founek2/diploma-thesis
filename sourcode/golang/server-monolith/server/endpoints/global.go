package endpoints

import (
	"encoding/json"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func jsonResponse(v any, w http.ResponseWriter) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(500)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(jsonData)
	}
}

func failUnexpectedError(err error, w http.ResponseWriter, r *http.Request) {
	println(err.Error())

	span := trace.SpanFromContext(r.Context())

	span.SetAttributes(attribute.String("http.error", err.Error()))
	span.SetStatus(codes.Error, "Unexpected error")
	w.WriteHeader(http.StatusInternalServerError)
}
