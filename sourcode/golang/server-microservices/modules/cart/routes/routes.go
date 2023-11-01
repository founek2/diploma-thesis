package routes

import (
	"fmt"
	"microservices/modules/cart/database"
	"microservices/modules/cart/endpoints"
	"microservices/modules/cart/middleware"
	"microservices/modules/clients"
	"microservices/shared"
	"strings"

	"net/http"

	"github.com/AzinKhan/functools"
)

func NewRoutes(db *database.Database, itemApi clients.ItemApi) shared.Routes {
	return functools.Map(func(route shared.Route) shared.Route {
		route.HandlerFunc = middleware.AddDb(route.HandlerFunc, db)
		route.HandlerFunc = middleware.AddItemApi(route.HandlerFunc, itemApi)
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
		Name:        "GetShoppingCartForUser",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/api/v1/cart/user/{userId}",
		HandlerFunc: endpoints.GetShoppingCartForUser,
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
		Name:        "GetItemIdsInShoppingCart",
		Method:      strings.ToUpper("Get"),
		Pattern:     "/api/v1/cart/{cartId}/item/id",
		HandlerFunc: endpoints.GetItemIdsInShoppingCart,
	},
	shared.Route{
		Name:        "DeleteShoppingCartByCartId",
		Method:      strings.ToUpper("Delete"),
		Pattern:     "/api/v1/cart/{cartId}",
		HandlerFunc: endpoints.DeleteShoppingCartByCartId,
	},
}
