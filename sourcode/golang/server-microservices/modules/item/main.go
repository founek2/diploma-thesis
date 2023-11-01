package api

import (
	"context"
	"microservices/modules/clients"
	"microservices/modules/item/database"
	"microservices/modules/item/routes"
	"microservices/shared"
)

type api struct {
	db database.Database
}

func Initialize(dsn string) (clients.ItemApi, shared.Routes) {
	db := database.Initialize(dsn)

	return &api{
		db: db,
	}, routes.NewRoutes(&db)
}

func (app *api) GetItemByItemId(ctx context.Context, item_id string) (*database.Item, error) {
	return app.db.GetItemByItemId(ctx, item_id)
}

func (app *api) GetItems(ctx context.Context) ([]database.Item, error) {
	return app.db.GetItems(ctx)
}

func (app *api) GetItemsByIds(ctx context.Context, itemIds []int64) ([]database.Item, error) {
	return app.db.GetItemsByIds(ctx, itemIds)
}
