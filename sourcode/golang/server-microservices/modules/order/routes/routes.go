package routes

import (
	"fmt"
	"microservices/modules/clients"
	"microservices/modules/order/database"
	"microservices/modules/order/endpoints"
	"microservices/modules/order/middleware"
	"microservices/shared"
	"strings"

	"net/http"

	"github.com/AzinKhan/functools"
)

func NewRoutes(db *database.Database, itemApi clients.ItemApi, cartApi clients.CartApi, invoiceApi clients.InvoiceApi) shared.Routes {
	return functools.Map(func(route shared.Route) shared.Route {
		route.HandlerFunc = middleware.AddDb(route.HandlerFunc, db)
		route.HandlerFunc = middleware.AddItemAndCartApi(route.HandlerFunc, itemApi, cartApi, invoiceApi)
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
		Name:        "CancelOrder",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/api/v1/order/{orderId}/cancel",
		HandlerFunc: endpoints.CancelOrder,
	},
	shared.Route{
		Name:        "CreateOrder",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/api/v1/order/create",
		HandlerFunc: endpoints.CreateOrder,
	},
	shared.Route{
		Name:        "GetOrder",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/api/v1/order/{orderId}",
		HandlerFunc: endpoints.GetOrder,
	},
}
