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
	"encoding/json"
	"modulith/modules/invoice/database"
	"modulith/modules/invoice/middleware"
	"modulith/modules/invoice/services"
	"modulith/shared/endpoints"
	"modulith/shared/getters"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
)

func GetInvoiceById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var db = middleware.GetDb(r.Context())
	var params = mux.Vars(r)

	var invoice, err = db.GetInvoiceByInvoiceId(ctx, params["invoiceId"])

	if err != nil {
		endpoints.FailUnexpectedError(err, w, r)
	} else {
		endpoints.JsonResponse(invoice, w)
	}
}

func UpdateInvoiceById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var db = middleware.GetDb(r.Context())
	var params = mux.Vars(r)

	invoice, err := db.GetInvoiceByInvoiceId(ctx, params["invoiceId"])
	if err != nil {
		endpoints.FailUnexpectedError(err, w, r)
		return
	}

	var data database.Invoice
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		endpoints.FailUnexpectedError(err, w, r)
		return
	}

	invoice.Status = data.Status

	invoice, err = db.UpdateInvoice(ctx, invoice)
	if err != nil {
		endpoints.FailUnexpectedError(err, w, r)
		return
	}

	endpoints.JsonResponse(invoice, w)
}

func GetInvoicePdfById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

type Order struct {
	Price float64 `json:"price"`
	Id    int64   `json:"id"`
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var db = middleware.GetDb(r.Context())

	var data Order
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		endpoints.FailUnexpectedError(err, w, r)
		return
	}

	var pdfPath = "/tmp/invoice_" + uuid.New().String() + ".pdf"

	var tracer = getters.GetTracer(ctx)
	_, span := tracer.Start(ctx, "CreateInvoice")
	span.SetAttributes(
		attribute.String("pdf.file", pdfPath),
	)
	defer span.End()

	var invoice = &database.Invoice{
		OrderSeqId: data.Id,
		Price:      data.Price,
		Status:     "created",
		PdfLink:    pdfPath,
	}
	services.ExampleFpdf_Circle(pdfPath)

	_, err = db.CreateInvoice(ctx, invoice)
	if err != nil {
		endpoints.FailUnexpectedError(err, w, r)
		return
	}

	if err != nil {
		endpoints.FailUnexpectedError(err, w, r)
	} else {
		println("Created invoice", invoice.InvoiceId.String())
		endpoints.JsonResponse(invoice, w)
	}
}
