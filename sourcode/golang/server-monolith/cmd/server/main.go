package main

import (
	"context"
	"log"
	"net/http"
	"os"

	sw "monolith/server"
	"monolith/server/database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
		uptrace.WithDSN(os.Getenv("UPTRACE_DSN")),
		uptrace.WithServiceName("monolith"),
		uptrace.WithServiceVersion("1.0.0"),
	)

	a.Db = database.Initialize(os.Getenv("DATABASE_URI"))
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var a = App{ctx: context.Background()}
	a.Initialize()
	a.Run("0.0.0.0:" + os.Getenv("PORT"))
}
