package routes

import (
	"fmt"
	"modulith/modules/invoice/database"
	"modulith/modules/invoice/endpoints"
	"modulith/modules/invoice/middleware"
	"modulith/shared"
	"strings"

	"net/http"

	"github.com/AzinKhan/functools"
)

func NewRoutes(db *database.Database) shared.Routes {
	return functools.Map(func(route shared.Route) shared.Route {
		route.HandlerFunc = middleware.AddDb(route.HandlerFunc, db)
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
		Name:        "CreateInvoice",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/api/v1/invoice",
		HandlerFunc: endpoints.CreateInvoice,
	},
	shared.Route{
		Name:        "GetInvoiceById",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/api/v1/invoice/{invoiceId}",
		HandlerFunc: endpoints.GetInvoiceById,
	},
	shared.Route{
		Name:        "GetInvoiceById",
		Method:      strings.ToUpper("Patch"),
		Pattern:     "/api/v1/invoice/{invoiceId}",
		HandlerFunc: endpoints.UpdateInvoiceById,
	},
	shared.Route{
		Name:        "GetInvoicePdfById",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/api/v1/invoice/{invoiceId}/pdf",
		HandlerFunc: endpoints.GetInvoicePdfById,
	},
}
