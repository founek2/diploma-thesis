package routes

import (
	"fmt"
	"microservices/modules/clients"
	"microservices/modules/payment/database"
	"microservices/modules/payment/endpoints"
	"microservices/modules/payment/middleware"
	"microservices/shared"
	"strings"

	"net/http"

	"github.com/AzinKhan/functools"
)

func NewRoutes(db *database.Database, invoiceApi clients.InvoiceApi) shared.Routes {
	return functools.Map(func(route shared.Route) shared.Route {
		route.HandlerFunc = middleware.AddDb(route.HandlerFunc, db)
		route.HandlerFunc = middleware.AddInvoiceApi(route.HandlerFunc, invoiceApi)
		return route
	}, routes)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = shared.Routes{
	shared.Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/api/v1/",
		HandlerFunc: Index,
	},
	shared.Route{
		Name:        "GetPayment",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/api/v1/payment/{paymentId}",
		HandlerFunc: endpoints.GetPaymentById,
	},
	shared.Route{
		Name:        "PayForInvoice",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/api/v1/payment/invoice/{invoiceId}",
		HandlerFunc: endpoints.PayForInvoice,
	},
}
