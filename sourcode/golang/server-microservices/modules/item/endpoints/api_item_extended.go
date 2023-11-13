package endpoints

import (
	"encoding/json"
	"microservices/modules/item/middleware"
	"microservices/shared/endpoints"
	"net/http"
)

func GetItemsByIds(w http.ResponseWriter, r *http.Request) {
	var db = middleware.GetDb(r.Context())

	var ids []int64
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		endpoints.FailUnexpectedError(err, w, r)
		return
	}

	items, err := db.GetItemsByIds(r.Context(), ids)
	if err != nil {
		endpoints.FailUnexpectedError(err, w, r)
	} else {
		endpoints.JsonResponse(items, w)
	}
}
