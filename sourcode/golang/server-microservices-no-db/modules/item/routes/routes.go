package routes

import (
	"fmt"
	"microservices-no-db/modules/item/database"
	"microservices-no-db/modules/item/endpoints"
	"microservices-no-db/modules/item/middleware"
	"microservices-no-db/shared"
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
