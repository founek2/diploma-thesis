package main

import (
	"context"
	"log"
	"net/http"

	sw "monolith/server"
	"monolith/server/database"

	"github.com/gorilla/mux"
	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	ctx    context.Context
	Router *mux.Router
	Db     database.Database
	Tracer trace.Tracer
}

func (a *App) Initialize() {
	// Configure OpenTelemetry with sensible defaults.
	uptrace.ConfigureOpentelemetry(
		uptrace.WithDSN("http://project2_secret_token@192.168.10.88:14317/2"),
		uptrace.WithServiceName("monolith"),
		uptrace.WithServiceVersion("1.0.0"),
	)

	a.Db = database.Initialize("postgres://postgres:postgres@localhost:5432/monolith?sslmode=disable")
	a.Tracer = otel.Tracer("server")
	a.Router = sw.NewRouter(&a.Db, a.Tracer)
}

func (a *App) Run(addr string) {
	defer uptrace.Shutdown(a.ctx)
	defer a.Db.Close()

	log.Printf("Server started on http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	var a = App{ctx: context.Background()}
	a.Initialize()
	a.Run("0.0.0.0:8080")
}
