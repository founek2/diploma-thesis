package api

import (
	"context"
	"microservices/modules/clients"
	"microservices/modules/payment/database"
	"microservices/modules/payment/routes"
	"microservices/shared"
)

type api struct {
	db         database.Database
	invoiceApi clients.InvoiceApi
}

func Initialize(dsn string, invoiceApi clients.InvoiceApi) (*api, shared.Routes) {
	db := database.Initialize(dsn)

	return &api{
		db: db,
	}, routes.NewRoutes(&db, invoiceApi)
}

type InvoiceApi interface {
	// GetPaymentByPaymentId(ctx context.Context, paymentId string) (*database.Payment, error)
}

func (app *api) GetPaymentByPaymentId(ctx context.Context, paymentId string) (*database.Payment, error) {
	return app.db.GetPaymentByPaymentId(ctx, paymentId)
}
