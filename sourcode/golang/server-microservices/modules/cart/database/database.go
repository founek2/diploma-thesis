package database

import (
	"context"
	"database/sql"

	"github.com/AzinKhan/functools"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunotel"
)

type Database struct {
	*bun.DB
}

func Initialize(dsn string) Database {
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	var sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	sqldb.SetMaxOpenConns(2)

	var db = bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook())
	db.AddQueryHook(bunotel.NewQueryHook(
		bunotel.WithDBName("modulith-cart"),
		bunotel.WithFormattedQueries(true),
	))

	return Database{
		db,
	}
}

func (db Database) GetCartByUserId(ctx context.Context, user_id string) (*Cart, error) {
	cart := &Cart{
		UserId: user_id,
	}
	err := db.NewSelect().Model(cart).Where("user_id = ?", user_id).Scan(ctx)
	if err == sql.ErrNoRows {
		_, err = db.NewInsert().Model(cart).On("CONFLICT (user_id) DO NOTHING").Returning("*").Exec(ctx)
		if err != nil {
			return cart, err
		}
	}
	if err != nil {
		return cart, err
	}

	return cart, err
}

func (db Database) GetCartByCartId(ctx context.Context, cart_id string) (*Cart, error) {
	cart := &Cart{}
	err := db.NewSelect().Model(cart).Where("cart_id = ?", cart_id).Scan(ctx)

	return cart, err
}

func (db Database) GetItemsIdsInCart(ctx context.Context, cart *Cart) ([]int64, error) {
	var items []RelItemCart
	err := db.NewSelect().Model(&items).Where("cart_seq_id = ?", cart.Id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	var ids = functools.Map(func(a RelItemCart) int64 {
		return a.ItemSeqId
	}, items)
	return ids, err
}

func (db Database) AddItemToCart(ctx context.Context, itemId int64, cart *Cart) (*RelItemCart, error) {
	var model = RelItemCart{
		ItemSeqId: itemId,
		CartSeqId: cart.Id,
	}

	_, err := db.NewInsert().
		Model(&model).
		On("CONFLICT (item_seq_id, cart_seq_id) DO NOTHING").
		Exec(ctx)

	return &model, err
}

func (db Database) RemoveItemFromCart(ctx context.Context, itemId int64, cart *Cart) (*RelItemCart, error) {
	var model = RelItemCart{
		ItemSeqId: itemId,
		CartSeqId: cart.Id,
	}
	_, err := db.NewDelete().
		Model(&model).
		Where("item_seq_id = ?", model.ItemSeqId).
		Where("cart_seq_id = ?", cart.Id).
		Exec(ctx)

	return &model, err
}

func (db Database) RemoveCart(ctx context.Context, cart *Cart) (*Cart, error) {
	_, err := db.NewDelete().
		Model(&RelItemCart{}).
		Where("cart_seq_id = ?", cart.Id).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	_, err = db.NewDelete().
		Model(cart).
		WherePK().
		Exec(ctx)

	return cart, err
}
