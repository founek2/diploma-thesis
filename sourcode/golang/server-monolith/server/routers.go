/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * This is a example for performance study based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 * Contact: skalicky.martin@iotdomu.cz
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"monolith/server/database"
	"monolith/server/endpoints"
	"monolith/server/middleware"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/trace"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(db *database.Database, tracer trace.Tracer) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		handler = middleware.AddTracingAndDb(route.Method, route.Pattern)(handler, tracer, db)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v1/",
		Index,
	},

	Route{
		"GetInvoiceById",
		strings.ToUpper("Get"),
		"/api/v1/invoice/{invoiceId}",
		endpoints.GetInvoiceById,
	},

	Route{
		"GetInvoicePdfById",
		strings.ToUpper("Get"),
		"/api/v1/invoice/{invoiceId}/pdf",
		endpoints.GetInvoicePdfById,
	},

	Route{
		"CancelOrder",
		strings.ToUpper("Post"),
		"/api/v1/order/{orderId}/cancel",
		endpoints.CancelOrder,
	},

	Route{
		"CreateOrder",
		strings.ToUpper("Post"),
		"/api/v1/order/create",
		endpoints.CreateOrder,
	},

	Route{
		"GetOrder",
		strings.ToUpper("Get"),
		"/api/v1/order/{orderId}",
		endpoints.GetOrder,
	},

	Route{
		"GetPaymentById",
		strings.ToUpper("Get"),
		"/api/v1/payment/{paymentId}",
		endpoints.GetPaymentById,
	},

	Route{
		"PayForInvoice",
		strings.ToUpper("Post"),
		"/api/v1/payment/invoice/{invoiceId}",
		endpoints.PayForInvoice,
	},

	Route{
		"AddItemToCart",
		strings.ToUpper("Post"),
		"/api/v1/cart/items/{itemId}",
		endpoints.AddItemToCart,
	},

	Route{
		"RemoveItemFromCart",
		strings.ToUpper("Delete"),
		"/api/v1/cart/items/{itemId}",
		endpoints.RemoveItemFromCart,
	},

	Route{
		"GetItem",
		strings.ToUpper("Get"),
		"/api/v1/item/{itemId}",
		endpoints.GetItem,
	},

	Route{
		"GetItems",
		strings.ToUpper("Get"),
		"/api/v1/item",
		endpoints.GetItems,
	},
}
