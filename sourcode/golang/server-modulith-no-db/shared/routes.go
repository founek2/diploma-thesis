/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * This is a example for performance study based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 * Contact: skalicky.martin@iotdomu.cz
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package shared

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// func NewRouter(tracer trace.Tracer) *mux.Router {
// 	router := mux.NewRouter().StrictSlash(true)
// 	for _, route := range routes {
// 		var handler http.Handler
// 		handler = route.HandlerFunc
// 		handler = Logger(handler, route.Name)
// 		handler = middleware.AddTracingAndDb(route.Method, route.Pattern)(handler, tracer)

// 		router.
// 			Methods(route.Method).
// 			Path(route.Pattern).
// 			Name(route.Name).
// 			Handler(handler)
// 	}

// 	return router
// }

// func Index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!")
// }

// var routes = Routes{
// 	Route{
// 		"Index",
// 		"GET",
// 		"/api/v1/",
// 		Index,
// 	},
// }