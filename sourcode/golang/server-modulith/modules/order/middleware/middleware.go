package middleware

import (
	"context"
	cartApi "modulith/modules/cart"
	invoiceApi "modulith/modules/invoice"
	"modulith/modules/order/database"
	"net/http"
)

func AddDb(next http.Handler, db *database.Database) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", db)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func AddItemAndCartApi(next http.Handler, itemApi cartApi.ItemApi, cartApi cartApi.CartApi, invoiceApi invoiceApi.InvoiceApi) http.HandlerFunc {
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
func GetItemApi(ctx context.Context) cartApi.ItemApi {
	return ctx.Value("itemApi").(cartApi.ItemApi)
}
func GetCartApi(ctx context.Context) cartApi.CartApi {
	return ctx.Value("cartApi").(cartApi.CartApi)
}
func GetInvoiceApi(ctx context.Context) invoiceApi.InvoiceApi {
	return ctx.Value("invoiceApi").(invoiceApi.InvoiceApi)
}
