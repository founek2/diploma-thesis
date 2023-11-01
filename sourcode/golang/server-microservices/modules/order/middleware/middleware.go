package middleware

import (
	"context"
	"microservices/modules/clients"
	"microservices/modules/order/database"
	"net/http"
)

func AddDb(next http.Handler, db *database.Database) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", db)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func AddItemAndCartApi(next http.Handler, itemApi clients.ItemApi, cartApi clients.CartApi, invoiceApi clients.InvoiceApi) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "itemApi", itemApi)
		ctx = context.WithValue(ctx, "cartApi", cartApi)
		ctx = context.WithValue(ctx, "invoiceApi", invoiceApi)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetDb(ctx context.Context) *database.Database {
	return ctx.Value("db").(*database.Database)
}
func GetItemApi(ctx context.Context) clients.ItemApi {
	return ctx.Value("itemApi").(clients.ItemApi)
}
func GetCartApi(ctx context.Context) clients.CartApi {
	return ctx.Value("cartApi").(clients.CartApi)
}
func GetInvoiceApi(ctx context.Context) clients.InvoiceApi {
	return ctx.Value("invoiceApi").(clients.InvoiceApi)
}
