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
	"microservices/modules/order/database"
	"microservices/modules/order/middleware"
	sharedGetters "microservices/shared/getters"
	"net/http"

	"github.com/gorilla/mux"
)

func CancelOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var db = middleware.GetDb(ctx)
	var params = mux.Vars(r)

	var order, err = db.GetOrderByOrderId(ctx, params["orderId"])
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	order.Status = "cancelled"
	_, err = db.UpdateOrder(ctx, order)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonResponse(order, w)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var ctx = r.Context()
	var db = middleware.GetDb(ctx)
	var cartApi = middleware.GetCartApi(ctx)
	var itemApi = middleware.GetItemApi(ctx)
	var invoiceApi = middleware.GetInvoiceApi(ctx)
	var userId = sharedGetters.GetUserId(r)

	cart, err := cartApi.UpsertShoppingCart(ctx, userId)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	itemsIds, err := cartApi.GetItemIdsInShoppingCart(ctx, cart)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	items, err := itemApi.GetItemsByIds(ctx, itemsIds)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	var sum = 0.
	for _, item := range items {
		item.Price += sum
	}
	var order = &database.Order{
		Title:       "Order",
		Description: "description",
		Quantity:    len(items),
		Price:       sum,
		Status:      "created",
	}

	order, err = db.CreateOrder(ctx, order)
	if err != nil {
		failUnexpectedError(err, w, r)
	}

	_, err = cartApi.RemoveShoppingCart(ctx, cart)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	// ---- CPU usage

	invoice, err := invoiceApi.CreateInvoice(ctx, order.Id, order.Price)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	println("Got invoice", invoice.InvoiceId.String())
	// ------

	order.Status = "processed"
	_, err = db.UpdateOrder(ctx, order)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	var orderResponse = struct {
		database.Order
		InvoiceId string `json:"invoiceId,omitempty"`
	}{
		Order:     *order,
		InvoiceId: invoice.InvoiceId.String(),
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonResponse(orderResponse, w)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
