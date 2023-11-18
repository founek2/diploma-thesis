package database

import (
	"context"
	"database/sql"
	"time"

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

	time.Sleep(2 * time.Millisecond)

	return cart, nil
}

func (db Database) GetItemsIdsInCart(ctx context.Context, cart *Cart) ([]int64, error) {
	var items = []RelItemCart{
		RelItemCart{},
		RelItemCart{},
	}

	time.Sleep(2 * time.Millisecond)

	var ids = functools.Map(func(a RelItemCart) int64 {
		return a.ItemSeqId
	}, items)
	return ids, nil
}

func (db Database) AddItemToCart(ctx context.Context, itemId int64, cart *Cart) (*RelItemCart, error) {
	var model = RelItemCart{
		ItemSeqId: itemId,
		CartSeqId: cart.Id,
	}

	time.Sleep(2 * time.Millisecond)

	return &model, nil
}

func (db Database) RemoveItemFromCart(ctx context.Context, itemId int64, cart *Cart) (*RelItemCart, error) {
	var model = RelItemCart{
		ItemSeqId: itemId,
		CartSeqId: cart.Id,
	}
	time.Sleep(2 * time.Millisecond)

	return &model, nil
}

func (db Database) RemoveCart(ctx context.Context, cart *Cart) (*Cart, error) {
	time.Sleep(2 * time.Millisecond)

	return cart, nil
}

// From item

func (db Database) GetItemByItemId(ctx context.Context, item_id string) (*Item, error) {
	item := &Item{}

	time.Sleep(2 * time.Millisecond)

	return item, nil
}

func (db Database) GetItems(ctx context.Context) ([]Item, error) {
	var items = []Item{
		Item{},
		Item{},
		Item{},
		Item{},
		Item{},
	}

	time.Sleep(2 * time.Millisecond)

	return items, nil
}

func (db Database) GetItemsByIds(ctx context.Context, ids []int64) ([]Item, error) {
	var items = []Item{
		Item{},
		Item{},
		Item{},
		Item{},
		Item{},
	}

	time.Sleep(2 * time.Millisecond)

	return items, nil
}
