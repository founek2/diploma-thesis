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
	"fmt"
	"monolith/server/database"
	"monolith/server/getters"
	"monolith/server/middleware"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
)

func PayForInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var db = middleware.GetDb(ctx)
	var params = mux.Vars(r)

	var invoice, err = db.GetInvoiceByInvoiceId(ctx, params["invoiceId"])
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	var tracer = getters.GetTracer(ctx)
	_, span := tracer.Start(ctx, "Calling payment system")
	span.SetAttributes(
		attribute.String("payment.price", fmt.Sprintf("%f", invoice.Price)),
	)
	// Simulate communication with 3rd party payment sevice
	time.Sleep(time.Second)
	span.End()

	invoice.Status = "paid"
	_, err = db.UpdateInvoice(ctx, invoice)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	var payment = &database.Payment{
		CreditCardNumber: 4311_2342_2342_1234,
		Amount:           invoice.Price,
	}

	_, err = db.CreatePayment(ctx, payment)
	if err != nil {
		failUnexpectedError(err, w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonResponse(invoice, w)
}