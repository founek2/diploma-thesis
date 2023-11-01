package endpoints

import (
	"microservices/modules/cart/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func GetShoppingCartForUser(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	var db = middleware.GetDb(ctx)
	var params = mux.Vars(r)

	var cart, err = db.GetCartByUserId(ctx, params["userId"])
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonResponse(cart, w)
}

func GetItemIdsInShoppingCart(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	var db = middleware.GetDb(ctx)
	var params = mux.Vars(r)

	cart, err := db.GetCartByCartId(ctx, params["cartId"])
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	itemsIds, err := db.GetItemsIdsInCart(ctx, cart)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonResponse(itemsIds, w)
}

func DeleteShoppingCartByCartId(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	var db = middleware.GetDb(ctx)
	var params = mux.Vars(r)

	cart, err := db.GetCartByCartId(ctx, params["cartId"])
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	cart, err = db.RemoveCart(ctx, cart)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonResponse(cart, w)
}
