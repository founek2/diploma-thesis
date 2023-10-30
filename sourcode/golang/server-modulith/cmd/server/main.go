package main

import (
	"context"
	"log"
	"net/http"
	"os"

	cartModule "modulith/modules/cart"
	invoiceModule "modulith/modules/invoice"
	orderModule "modulith/modules/order"
	paymentModule "modulith/modules/payment"
	"modulith/shared"
	"modulith/shared/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	ctx    context.Context
	Router *mux.Router
	Tracer trace.Tracer
}

func registerRoutes(routes shared.Routes, router *mux.Router, tracer trace.Tracer) {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = shared.Logger(handler, route.Name)
		handler = middleware.AddTracing(route.Method, route.Pattern)(handler, tracer)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}

func (a *App) Initialize() {
	// Configure OpenTelemetry with sensible defaults.
	uptrace.ConfigureOpentelemetry(
		uptrace.WithDSN(os.Getenv("UPTRACE_DSN")),
		uptrace.WithServiceName("monolith"),
		uptrace.WithServiceVersion("1.0.0"),
	)

	a.Router = mux.NewRouter().StrictSlash(true)

	// a.Db = database.Initialize("postgres://monolith:monolith@192.168.10.88:5432/monolith?sslmode=disable")
	a.Tracer = otel.Tracer("server")
	// a.Router = sw.NewRouter(&a.Db, a.Tracer)

	// var itemApi, itemRoutes = itemModule.Initialize("postgres://modulith:modulith@192.168.10.88:5432/modulith?sslmode=disable")
	// registerRoutes(itemRoutes, a.Router, a.Tracer)

	var databaseUri = os.Getenv("DATABASE_URI")

	var cartApi, cartRoutes = cartModule.Initialize(databaseUri)
	registerRoutes(cartRoutes, a.Router, a.Tracer)

	var invoiceApi, invoiceRoutes = invoiceModule.Initialize(databaseUri)
	registerRoutes(invoiceRoutes, a.Router, a.Tracer)

	var _, orderRoutes = orderModule.Initialize(databaseUri, cartApi, cartApi, invoiceApi)
	registerRoutes(orderRoutes, a.Router, a.Tracer)

	var _, paymentRoutes = paymentModule.Initialize(databaseUri, invoiceApi)
	registerRoutes(paymentRoutes, a.Router, a.Tracer)
}

func (a *App) Run(addr string) {
	defer uptrace.Shutdown(a.ctx)

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
