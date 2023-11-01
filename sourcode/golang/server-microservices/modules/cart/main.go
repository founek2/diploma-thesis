package api

import (
	"context"
	"microservices/modules/cart/database"
	"microservices/modules/cart/routes"
	"microservices/modules/clients"
	"microservices/shared"
)

type api struct {
	db database.Database
}

func Initialize(dsn string, itemApi clients.ItemApi) (clients.CartApi, shared.Routes) {
	db := database.Initialize(dsn)
	return &api{
		db,
	}, routes.NewRoutes(&db, itemApi)
}
func (app *api) UpsertShoppingCart(ctx context.Context, user_id string) (*database.Cart, error) {
	return app.db.GetCartByUserId(ctx, user_id)
}

func (app *api) AddItemToShoppingCart(ctx context.Context, itemId int64, cart *database.Cart) (*database.RelItemCart, error) {
	return app.db.AddItemToCart(ctx, itemId, cart)
}

func (app *api) RemoveItemFromShoppingCart(ctx context.Context, itemId int64, cart *database.Cart) (*database.RelItemCart, error) {
	return app.db.RemoveItemFromCart(ctx, itemId, cart)
}

func (app *api) GetItemIdsInShoppingCart(ctx context.Context, cart *database.Cart) ([]int64, error) {
	return app.db.GetItemsIdsInCart(ctx, cart)
}

func (app *api) RemoveShoppingCart(ctx context.Context, cart *database.Cart) (*database.Cart, error) {
	return app.db.RemoveCart(ctx, cart)
}
