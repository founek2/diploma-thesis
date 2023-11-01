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
	"monolith/server/database"
	"monolith/server/getters"
	"monolith/server/middleware"
	"monolith/server/services"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
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
	var userId = getters.GetUserId(r)
	var tracer = getters.GetTracer(ctx)

	cart, err := db.GetCartByUserId(ctx, userId)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	items, err := db.GetItemsForCart(ctx, cart)
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
		return
	}

	_, err = db.RemoveCart(ctx, cart)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	// ---- CPU usage

	println("Generating pdf")
	var pdfPath = "/tmp/invoice_" + uuid.New().String() + ".pdf"

	_, span := tracer.Start(ctx, "Generating invoice")
	span.SetAttributes(
		attribute.String("pdf.file", pdfPath),
	)

	var invoice = &database.Invoice{
		OrderSeqId: order.Id,
		Price:      order.Price,
		Status:     "created",
		PdfLink:    pdfPath,
	}
	services.ExampleFpdf_Circle(pdfPath)
	span.End()

	db.CreateInvoice(ctx, invoice)

	// ------

	order.Status = "processed"
	_, err = db.UpdateOrder(ctx, order)
	if err != nil {
		failUnexpectedError(err, w, r)
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
