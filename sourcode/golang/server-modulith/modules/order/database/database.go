package database

import (
	"context"
	"database/sql"

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
	var sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	sqldb.SetMaxOpenConns(2)

	var db = bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook())
	db.AddQueryHook(bunotel.NewQueryHook(
		bunotel.WithDBName("modulith-order"),
		bunotel.WithFormattedQueries(true),
	))

	return Database{
		db,
	}
}

func (db Database) GetOrderByOrderId(ctx context.Context, order_id string) (*Order, error) {
	order := new(Order)
	err := db.NewSelect().Model(order).Where("invoice_id = ?", order_id).Scan(ctx)
	return order, err
}

func (db Database) CreateOrder(ctx context.Context, order *Order) (*Order, error) {
	_, err := db.NewInsert().Model(order).Returning("*").Exec(ctx)
	return order, err
}

func (db Database) UpdateOrder(ctx context.Context, order *Order) (*Order, error) {
	_, err := db.NewUpdate().Model(order).WherePK().Exec(ctx)

	return order, err
}
