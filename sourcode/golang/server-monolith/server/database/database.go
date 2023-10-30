package database

import (
	"context"
	"database/sql"
	"fmt"

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
	sqldb.SetMaxOpenConns(2 * 4)

	var db = bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook())
	db.AddQueryHook(bunotel.NewQueryHook(
		bunotel.WithDBName("monolith"),
		bunotel.WithFormattedQueries(true),
	))

	return Database{
		db,
	}
}

func (db Database) GetItemByItemId(ctx context.Context, item_id string) (*Item, error) {
	item := new(Item)
	err := db.NewSelect().Model(item).Where("item_id = ?", item_id).Scan(ctx)
	return item, err
}

func (db Database) GetItems(ctx context.Context) ([]Item, error) {
	var items []Item
	err := db.NewSelect().Model(&items).Scan(ctx)
	return items, err
}

func (db Database) GetOrderByOrderId(ctx context.Context, order_id string) (*Order, error) {
	order := new(Order)
	err := db.NewSelect().Model(order).Where("invoice_id = ?", order_id).Scan(ctx)
	return order, err
}

func (db Database) GetInvoiceByInvoiceId(ctx context.Context, invoice_id string) (*Invoice, error) {
	invoice := new(Invoice)
	err := db.NewSelect().Model(invoice).Where("invoice_id = ?", invoice_id).Scan(ctx)
	return invoice, err
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

func (db Database) GetItemsForCart(ctx context.Context, cart *Cart) ([]Item, error) {
	var items []Item
	err := db.NewRaw(
		"SELECT * FROM item AS i WHERE EXISTS (SELECT id FROM rel_item_cart AS r where i.id = r.item_seq_id AND cart_seq_id = ?)",
		cart.Id,
	).Scan(ctx, &items)
	return items, err
}

func (db Database) GetPaymentByPaymentId(ctx context.Context, paymentId string) (*Payment, error) {
	payment := new(Payment)
	_, err := db.NewSelect().Model(payment).Where("payment_id = ?", paymentId).Exec(ctx)
	return payment, err
}

func (db Database) CreateOrder(ctx context.Context, order *Order) (*Order, error) {
	_, err := db.NewInsert().Model(order).Returning("*").Exec(ctx)
	return order, err
}

func (db Database) CreateInvoice(ctx context.Context, invoice *Invoice) (*Invoice, error) {
	_, err := db.NewInsert().Model(invoice).Returning("*").Exec(ctx)
	return invoice, err
}

func (db Database) CreatePayment(ctx context.Context, payment *Payment) (*Payment, error) {
	_, err := db.NewInsert().Model(payment).Returning("*").Exec(ctx)
	return payment, err
}

func (db Database) AddItemToCart(ctx context.Context, item *Item, cart *Cart) (*RelItemCart, error) {
	var model = RelItemCart{
		ItemSeqId: item.Id,
		CartSeqId: cart.Id,
	}
	fmt.Printf("%+v\n", model)

	_, err := db.NewInsert().
		Model(&model).
		On("CONFLICT (item_seq_id, cart_seq_id) DO NOTHING").
		Exec(ctx)

	return &model, err
}

func (db Database) RemoveItemFromCart(ctx context.Context, item *Item, cart *Cart) (*RelItemCart, error) {
	var model = RelItemCart{
		ItemSeqId: item.Id,
		CartSeqId: cart.Id,
	}
	_, err := db.NewDelete().
		Model(&model).
		Where("item_seq_id = ?", model.ItemSeqId).
		Where("cart_seq_id = ?", cart.Id).
		Exec(ctx)

	return &model, err
}

func (db Database) UpdateOrder(ctx context.Context, order *Order) (*Order, error) {
	_, err := db.NewUpdate().Model(order).WherePK().Exec(ctx)

	return order, err
}

func (db Database) UpdateInvoice(ctx context.Context, invoice *Invoice) (*Invoice, error) {
	_, err := db.NewUpdate().Model(invoice).WherePK().Exec(ctx)

	return invoice, err
}
