package middleware

import (
	"context"
	"microservices-no-db/modules/clients"
	"microservices-no-db/modules/payment/database"
	"net/http"
)

func AddDb(next http.Handler, db *database.Database) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", db)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AddInvoiceApi(next http.Handler, invoiceApi clients.InvoiceApi) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "invoiceApi", invoiceApi)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetDb(ctx context.Context) *database.Database {
	return ctx.Value("db").(*database.Database)
}
func GetInvoiceApi(ctx context.Context) clients.InvoiceApi {
	return ctx.Value("invoiceApi").(clients.InvoiceApi)
}
