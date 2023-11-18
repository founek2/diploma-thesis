package api

import (
	"context"
	"microservices-no-db/modules/clients"
	"microservices-no-db/modules/order/database"
	"microservices-no-db/modules/order/routes"
	"microservices-no-db/shared"
)

type api struct {
	db      database.Database
	itemApi clients.ItemApi
	cartApi clients.CartApi
}

func Initialize(dsn string, itemApi clients.ItemApi, cartApi clients.CartApi, invoiceApi clients.InvoiceApi) (clients.OrderApi, shared.Routes) {
	db := database.Initialize(dsn)

	return &api{
		db: db,
	}, routes.NewRoutes(&db, itemApi, cartApi, invoiceApi)
}

func (app *api) GetOrderByOrderId(ctx context.Context, orderId string) (*database.Order, error) {
	return app.db.GetOrderByOrderId(ctx, orderId)
}
