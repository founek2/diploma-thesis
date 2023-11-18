package middleware

import (
	"context"
	invoiceApi "modulith-no-db/modules/invoice"
	"modulith-no-db/modules/payment/database"
	"net/http"
)

func AddDb(next http.Handler, db *database.Database) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", db)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AddInvoiceApi(next http.Handler, invoiceApi invoiceApi.InvoiceApi) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "invoiceApi", invoiceApi)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetDb(ctx context.Context) *database.Database {
	return ctx.Value("db").(*database.Database)
}
func GetInvoiceApi(ctx context.Context) invoiceApi.InvoiceApi {
	return ctx.Value("invoiceApi").(invoiceApi.InvoiceApi)
}
