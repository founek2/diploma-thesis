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
	"microservices/modules/payment/database"
	"microservices/modules/payment/middleware"
	"microservices/shared/getters"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
)

func GetPaymentById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var db = middleware.GetDb(r.Context())
	var params = mux.Vars(r)

	var invoice, err = db.GetPaymentByPaymentId(ctx, params["paymentId"])

	if err != nil {
		failUnexpectedError(err, w, r)
	} else {
		jsonResponse(invoice, w)
	}
}

func PayForInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var db = middleware.GetDb(ctx)
	var invoiceApi = middleware.GetInvoiceApi(ctx)
	var params = mux.Vars(r)

	var invoice, err = invoiceApi.GetInvoiceByInvoiceId(ctx, params["invoiceId"])
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
	time.Sleep(500 * time.Millisecond)
	span.End()

	invoice.Status = "paid"
	_, err = invoiceApi.UpdateInvoice(ctx, invoice)
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
