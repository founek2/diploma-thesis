package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	cartModule "microservices-no-db/modules/cart"
	clients "microservices-no-db/modules/clients"
	invoiceModule "microservices-no-db/modules/invoice"
	itemModule "microservices-no-db/modules/item"
	orderModule "microservices-no-db/modules/order"
	paymentModule "microservices-no-db/modules/payment"
	"microservices-no-db/shared"
	"microservices-no-db/shared/middleware"

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

	a.Router = mux.NewRouter().StrictSlash(true)

	a.Tracer = otel.Tracer("server")

	var databaseUri = os.Getenv("DATABASE_URI")
	var mode = os.Getenv("MODE")

	uptrace.ConfigureOpentelemetry(
		uptrace.WithDSN(os.Getenv("UPTRACE_DSN")),
		uptrace.WithServiceName("microservice-"+mode),
		uptrace.WithServiceVersion("1.0.0"),
	)

	var cfg = clients.NewConfiguration()
	cfg.BasePath = fmt.Sprintf("%s/api/v1", os.Getenv("API_HOST"))
	println("basePath", cfg.BasePath)
	var client = clients.NewAPIClient(cfg)

	switch mode {
	case "cart":
		var _, routes = cartModule.Initialize(databaseUri, client.ItemClient)
		registerRoutes(routes, a.Router, a.Tracer)
	case "invoice":
		var _, routes = invoiceModule.Initialize(databaseUri)
		registerRoutes(routes, a.Router, a.Tracer)
	case "item":
		var _, routes = itemModule.Initialize(databaseUri)
		registerRoutes(routes, a.Router, a.Tracer)
	case "order":
		var _, routes = orderModule.Initialize(databaseUri, client.ItemClient, client.CartClient, client.InvoiceClient)
		registerRoutes(routes, a.Router, a.Tracer)
	case "payment":
		var _, routes = paymentModule.Initialize(databaseUri, client.InvoiceClient)
		registerRoutes(routes, a.Router, a.Tracer)
	default:
		panic("Unsupported mode provided")
	}
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
