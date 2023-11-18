package database

import (
	"context"
	"database/sql"
	"time"

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

	var db = bun.NewDB(sqldb, pgdialect.New())
	bundebug.NewQueryHook(bundebug.WithVerbose(true))
	db.AddQueryHook(bundebug.NewQueryHook())
	db.AddQueryHook(bunotel.NewQueryHook(
		bunotel.WithDBName("modulith-item"),
		bunotel.WithFormattedQueries(true),
	))

	return Database{
		db,
	}
}

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
