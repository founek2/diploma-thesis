package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"

	"github.com/urfave/cli/v2"

	"modulith-no-db/cmd"
	cartMigrations "modulith-no-db/modules/cart/database/migrations"
	invoiceMigrations "modulith-no-db/modules/invoice/database/migrations"
	orderMigrations "modulith-no-db/modules/order/database/migrations"
	paymentMigrations "modulith-no-db/modules/payment/database/migrations"

	"github.com/uptrace/bun"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URI"))))

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(false),
		bundebug.FromEnv(""),
	))

	app := &cli.App{
		Name: "bun",

		Commands: []*cli.Command{
			cmd.NewDBCommand("db-cart", migrate.NewMigrator(db, cartMigrations.Migrations)),
			cmd.NewDBCommand("db-invoice", migrate.NewMigrator(db, invoiceMigrations.Migrations)),
			cmd.NewDBCommand("db-order", migrate.NewMigrator(db, orderMigrations.Migrations)),
			cmd.NewDBCommand("db-payment", migrate.NewMigrator(db, paymentMigrations.Migrations)),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
