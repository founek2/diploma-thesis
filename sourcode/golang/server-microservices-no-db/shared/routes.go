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
