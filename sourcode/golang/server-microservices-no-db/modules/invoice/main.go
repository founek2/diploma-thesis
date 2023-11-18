package api

import (
	"context"
	"microservices-no-db/modules/clients"
	"microservices-no-db/modules/invoice/database"
	"microservices-no-db/modules/invoice/routes"
	"microservices-no-db/modules/invoice/services"
	"microservices-no-db/shared"
	"microservices-no-db/shared/getters"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
)

type api struct {
	db      database.Database
	itemApi clients.ItemApi
	cartApi clients.CartApi
}

func Initialize(dsn string) (clients.InvoiceApi, shared.Routes) {
	db := database.Initialize(dsn)

	return &api{
		db: db,
	}, routes.NewRoutes(&db)
}

func (app *api) GetInvoiceByInvoiceId(ctx context.Context, invoiceId string) (*database.Invoice, error) {
	var tracer = getters.GetTracer(ctx)
	_, span := tracer.Start(ctx, "GetInvoiceByInvoiceId")
	span.SetAttributes(
		attribute.String("invoice.invoice_id", invoiceId),
	)
	defer span.End()

	return app.db.GetInvoiceByInvoiceId(ctx, invoiceId)
}
func (app *api) UpdateInvoice(ctx context.Context, invoice *database.Invoice) (*database.Invoice, error) {
	var tracer = getters.GetTracer(ctx)
	_, span := tracer.Start(ctx, "UpdateInvoice")
	span.SetAttributes(
		attribute.String("invoice.invoice_id", invoice.InvoiceId.String()),
	)
	defer span.End()

	return app.db.UpdateInvoice(ctx, invoice)
}
func (app *api) CreateInvoice(ctx context.Context, orderId int64, price float64) (*database.Invoice, error) {
	var pdfPath = "/tmp/invoice_" + uuid.New().String() + ".pdf"

	var tracer = getters.GetTracer(ctx)
	_, span := tracer.Start(ctx, "CreateInvoice")
	span.SetAttributes(
		attribute.String("pdf.file", pdfPath),
	)
	defer span.End()

	var invoice = &database.Invoice{
		OrderSeqId: orderId,
		Price:      price,
		Status:     "created",
		PdfLink:    pdfPath,
	}
	services.ExampleFpdf_Circle(pdfPath)

	return app.db.CreateInvoice(ctx, invoice)
}