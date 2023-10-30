package main

import (
	"context"
	"log"
	"net/http"

	cartModule "modulith/modules/cart"
	invoiceModule "modulith/modules/invoice"
	orderModule "modulith/modules/order"
	paymentModule "modulith/modules/payment"
	"modulith/shared"
	"modulith/shared/middleware"

	"github.com/gorilla/mux"
	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type App struct {
	ctx    context.Context
	Router *mux.Router
	Tracer trace.Tracer
}

// func Middleware(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		middleware.AddTracing(r.)
// 			h.ServeHTTP(w, r)
// 	})
// }

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
		uptrace.WithDSN("http://project2_secret_token@192.168.10.88:14317/2"),
		uptrace.WithServiceName("monolith"),
		uptrace.WithServiceVersion("1.0.0"),
	)

	a.Router = mux.NewRouter().StrictSlash(true)

	// a.Db = database.Initialize("postgres://monolith:monolith@192.168.10.88:5432/monolith?sslmode=disable")
	a.Tracer = otel.Tracer("server")
	// a.Router = sw.NewRouter(&a.Db, a.Tracer)

	// var itemApi, itemRoutes = itemModule.Initialize("postgres://modulith:modulith@192.168.10.88:5432/modulith?sslmode=disable")
	// registerRoutes(itemRoutes, a.Router, a.Tracer)

	var cartApi, cartRoutes = cartModule.Initialize("postgres://modulith:modulith@192.168.10.88:5432/modulith?sslmode=disable")
	registerRoutes(cartRoutes, a.Router, a.Tracer)

	var invoiceApi, invoiceRoutes = invoiceModule.Initialize("postgres://modulith:modulith@192.168.10.88:5432/modulith?sslmode=disable")
	registerRoutes(invoiceRoutes, a.Router, a.Tracer)

	var _, orderRoutes = orderModule.Initialize("postgres://modulith:modulith@192.168.10.88:5432/modulith?sslmode=disable", cartApi, cartApi, invoiceApi)
	registerRoutes(orderRoutes, a.Router, a.Tracer)

	var _, paymentRoutes = paymentModule.Initialize("postgres://modulith:modulith@192.168.10.88:5432/modulith?sslmode=disable", invoiceApi)
	registerRoutes(paymentRoutes, a.Router, a.Tracer)
}

func (a *App) Run(addr string) {
	defer uptrace.Shutdown(a.ctx)

	log.Printf("Server started on http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	var a = App{ctx: context.Background()}
	a.Initialize()
	a.Run("0.0.0.0:8080")
}
