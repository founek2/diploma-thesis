package api

import (
	"context"
	cartApi "modulith-no-db/modules/cart"
	invoiceApi "modulith-no-db/modules/invoice"
	"modulith-no-db/modules/order/database"
	"modulith-no-db/modules/order/routes"
	"modulith-no-db/shared"
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
