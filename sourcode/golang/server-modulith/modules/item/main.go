package api

import (
	"context"
	"modulith/modules/item/database"
	"modulith/modules/item/routes"
	"modulith/shared"
)

type api struct {
	db database.Database
}

func Initialize(dsn string) (*api, shared.Routes) {
	db := database.Initialize(dsn)

	return &api{
		db: db,
	}, routes.NewRoutes(&db)
}

type ItemApi interface {
	GetItemByItemId(ctx context.Context, item_id string) (*database.Item, error)
	GetItems(ctx context.Context) ([]database.Item, error)
	GetItemsByIds(ctx context.Context, itemIds []int64) ([]database.Item, error)
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
