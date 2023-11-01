/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * This is a example for performance study based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 * Contact: skalicky.martin@iotdomu.cz
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package endpoints

import (
	"microservices/modules/item/database"
	"microservices/modules/item/middleware"
	"net/http"
	"strconv"
	"strings"

	"github.com/AzinKhan/functools"
	"github.com/gorilla/mux"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var db = middleware.GetDb(r.Context())
	var params = mux.Vars(r)

	var item, err = db.GetItemByItemId(ctx, params["itemId"])

	if err != nil {
		failUnexpectedError(err, w, r)
	} else {
		jsonResponse(item, w)
	}
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	var db = middleware.GetDb(r.Context())

	var queryIDs = r.URL.Query().Get("ids")

	var items []database.Item
	var err error

	if queryIDs != "" {
		ids_str := strings.Split(queryIDs, ",")
		var ids = functools.Map(func(id string) int64 {
			i, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				panic(err)
			}
			return i
		}, ids_str)

		items, err = db.GetItemsByIds(r.Context(), ids)
	} else {
		items, err = db.GetItems(r.Context())
	}
	if err != nil {
		failUnexpectedError(err, w, r)
	} else {
		jsonResponse(items, w)
	}
}