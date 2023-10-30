package api

import (
	"context"
	cartApi "modulith/modules/cart"
	invoiceApi "modulith/modules/invoice"
	"modulith/modules/order/database"
	"modulith/modules/order/routes"
	"modulith/shared"
)

type api struct {
	db      database.Database
	itemApi cartApi.ItemApi
	cartApi cartApi.CartApi
}

func Initialize(dsn string, itemApi cartApi.ItemApi, cartApi cartApi.CartApi, invoiceApi invoiceApi.InvoiceApi) (*api, shared.Routes) {
	db := database.Initialize(dsn)

	return &api{
		db: db,
	}, routes.NewRoutes(&db, itemApi, cartApi, invoiceApi)
}

type OrderApi interface {
	GetOrderByOrderId(ctx context.Context, orderId string) (*database.Order, error)
}

func (app *api) GetOrderByOrderId(ctx context.Context, orderId string) (*database.Order, error) {
	return app.db.GetOrderByOrderId(ctx, orderId)
}
