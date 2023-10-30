package api

import (
	"context"
	"modulith/modules/cart/database"
	"modulith/modules/cart/routes"
	"modulith/shared"

	"github.com/gorilla/mux"
)

type api struct {
	db database.Database
}

func Initialize(dsn string) (*api, shared.Routes) {
	db := database.Initialize(dsn)
	return &api{
		db,
	}, routes.NewRoutes(&db)
}

func (app *api) BindRouter(router *mux.Router) {
	routes.NewRouter(&app.db, router)
}

type CartApi interface {
	UpsertShoppingCart(ctx context.Context, user_id string) (*database.Cart, error)
	AddItemToShoppingCart(ctx context.Context, itemId int64, cart *database.Cart) (*database.RelItemCart, error)
	RemoveItemFromShoppingCart(ctx context.Context, itemId int64, cart *database.Cart) (*database.RelItemCart, error)
	GetItemIdsInShoppingCart(ctx context.Context, cart *database.Cart) ([]int64, error)
	RemoveShoppingCart(ctx context.Context, cart *database.Cart) (*database.Cart, error)
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
