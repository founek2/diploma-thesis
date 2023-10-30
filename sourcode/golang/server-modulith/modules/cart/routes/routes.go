package routes

import (
	"fmt"
	"modulith/modules/cart/database"
	"modulith/modules/cart/endpoints"
	"modulith/modules/cart/middleware"
	"modulith/shared"
	"strings"

	"net/http"

	"github.com/AzinKhan/functools"
	"github.com/gorilla/mux"
)

func NewRoutes(db *database.Database) shared.Routes {
	return functools.Map(func(route shared.Route) shared.Route {
		route.HandlerFunc = middleware.AddDb(route.HandlerFunc, db)
		return route
	}, routes)
}

func NewRouter(db *database.Database, router *mux.Router) {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = middleware.AddDb(handler, db)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
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
		Name:        "AddItemToCart",
		Method:      strings.ToUpper("Post"),
		Pattern:     "/api/v1/cart/items/{itemId}",
		HandlerFunc: endpoints.AddItemToCart,
	},
	shared.Route{
		Name:        "RemoveItemFromCart",
		Method:      strings.ToUpper("Delete"),
		Pattern:     "/api/v1/cart/items/{itemId}",
		HandlerFunc: endpoints.RemoveItemFromCart,
	},
	shared.Route{
		Name:        "GetItem",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/api/v1/item/{itemId}",
		HandlerFunc: endpoints.GetItem,
	},
	shared.Route{
		Name:        "GetItems",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/api/v1/item",
		HandlerFunc: endpoints.GetItems,
	},
}
