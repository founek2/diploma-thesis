package api

import (
	"context"
	"microservices/modules/clients"
	"microservices/modules/order/database"
	"microservices/modules/order/routes"
	"microservices/shared"
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
