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
	sqldb.SetMaxOpenConns(10)

	var db = bun.NewDB(sqldb, pgdialect.New())
	bundebug.NewQueryHook(bundebug.WithVerbose(true))
	db.AddQueryHook(bundebug.NewQueryHook())
	db.AddQueryHook(bunotel.NewQueryHook(
		bunotel.WithDBName("modulith-order"),
		bunotel.WithFormattedQueries(true),
	))

	return Database{
		db,
	}
}

func (db Database) GetInvoiceByInvoiceId(ctx context.Context, invoice_id string) (*Invoice, error) {
	invoice := new(Invoice)
	err := db.NewSelect().Model(invoice).Where("invoice_id = ?", invoice_id).Scan(ctx)
	return invoice, err
}

func (db Database) CreateInvoice(ctx context.Context, invoice *Invoice) (*Invoice, error) {
	_, err := db.NewInsert().Model(invoice).Returning("*").Exec(ctx)
	return invoice, err
}

func (db Database) UpdateInvoice(ctx context.Context, invoice *Invoice) (*Invoice, error) {
	_, err := db.NewUpdate().Model(invoice).WherePK().Exec(ctx)

	return invoice, err
}
