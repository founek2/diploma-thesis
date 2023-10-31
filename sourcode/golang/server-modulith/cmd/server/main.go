package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	cartModule "modulith/modules/cart"
	invoiceModule "modulith/modules/invoice"
	invoiceClient "modulith/modules/invoice-client"
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

	a.Router = mux.NewRouter().StrictSlash(true)

	a.Tracer = otel.Tracer("server")

	var databaseUri = os.Getenv("DATABASE_URI")

	if os.Getenv("MODE") == "invoice-server" {
		uptrace.ConfigureOpentelemetry(
			uptrace.WithDSN(os.Getenv("UPTRACE_DSN")),
			uptrace.WithServiceName("modulith-invoice-server"),
			uptrace.WithServiceVersion("1.0.0"),
		)

		var _, invoiceRoutes = invoiceModule.Initialize(databaseUri)
		registerRoutes(invoiceRoutes, a.Router, a.Tracer)
	} else if os.Getenv("MODE") == "invoice-client-server" {
		uptrace.ConfigureOpentelemetry(
			uptrace.WithDSN(os.Getenv("UPTRACE_DSN")),
			uptrace.WithServiceName("modulith-invoice-client-server"),
			uptrace.WithServiceVersion("1.0.0"),
		)

		var cartApi, cartRoutes = cartModule.Initialize(databaseUri)
		registerRoutes(cartRoutes, a.Router, a.Tracer)

		var cfg = invoiceClient.NewConfiguration()
		cfg.BasePath = fmt.Sprintf("%s/api/v1", os.Getenv("INVOICE_HOST"))
		println("basePath", cfg.BasePath)
		var client = invoiceClient.NewAPIClient(cfg)

		var _, orderRoutes = orderModule.Initialize(databaseUri, cartApi, cartApi, client.InvoiceApi)
		registerRoutes(orderRoutes, a.Router, a.Tracer)

		var _, paymentRoutes = paymentModule.Initialize(databaseUri, client.InvoiceApi)
		registerRoutes(paymentRoutes, a.Router, a.Tracer)
	} else {
		uptrace.ConfigureOpentelemetry(
			uptrace.WithDSN(os.Getenv("UPTRACE_DSN")),
			uptrace.WithServiceName("modulith"),
			uptrace.WithServiceVersion("1.0.0"),
		)

		var cartApi, cartRoutes = cartModule.Initialize(databaseUri)
		registerRoutes(cartRoutes, a.Router, a.Tracer)

		var invoiceApi, invoiceRoutes = invoiceModule.Initialize(databaseUri)
		registerRoutes(invoiceRoutes, a.Router, a.Tracer)

		var _, orderRoutes = orderModule.Initialize(databaseUri, cartApi, cartApi, invoiceApi)
		registerRoutes(orderRoutes, a.Router, a.Tracer)

		var _, paymentRoutes = paymentModule.Initialize(databaseUri, invoiceApi)
		registerRoutes(paymentRoutes, a.Router, a.Tracer)
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
