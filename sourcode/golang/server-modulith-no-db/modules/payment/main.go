package api

import (
	"context"
	invoiceApi "modulith-no-db/modules/invoice"
	"modulith-no-db/modules/payment/database"
	"modulith-no-db/modules/payment/routes"
	"modulith-no-db/shared"
)

type api struct {
	db         database.Database
	invoiceApi invoiceApi.InvoiceApi
}

func Initialize(dsn string, invoiceApi invoiceApi.InvoiceApi) (*api, shared.Routes) {
	db := database.Initialize(dsn)

	return &api{
		db: db,
	}, routes.NewRoutes(&db, invoiceApi)
}

type InvoiceApi interface {
	GetPaymentByPaymentId(ctx context.Context, paymentId string) (*database.Payment, error)
}

func (app *api) GetPaymentByPaymentId(ctx context.Context, paymentId string) (*database.Payment, error) {
	return app.db.GetPaymentByPaymentId(ctx, paymentId)
}
