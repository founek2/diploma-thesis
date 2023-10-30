package routes

import (
	"fmt"
	cartApi "modulith/modules/cart"
	invoiceApi "modulith/modules/invoice"
	"modulith/modules/order/database"
	"modulith/modules/order/endpoints"
	"modulith/modules/order/middleware"
	"modulith/shared"
	"strings"

	"net/http"

	"github.com/AzinKhan/functools"
)

func NewRoutes(db *database.Database, itemApi cartApi.ItemApi, cartApi cartApi.CartApi, invoiceApi invoiceApi.InvoiceApi) shared.Routes {
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
}
