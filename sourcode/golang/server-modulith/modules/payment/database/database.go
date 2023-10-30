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
		bunotel.WithDBName("modulith-payment"),
		bunotel.WithFormattedQueries(true),
	))

	return Database{
		db,
	}
}

func (db Database) CreatePayment(ctx context.Context, payment *Payment) (*Payment, error) {
	_, err := db.NewInsert().Model(payment).Returning("*").Exec(ctx)
	return payment, err
}

func (db Database) GetPaymentByPaymentId(ctx context.Context, paymentId string) (*Payment, error) {
	payment := new(Payment)
	_, err := db.NewSelect().Model(payment).Where("payment_id = ?", paymentId).Exec(ctx)
	return payment, err
}
