package middleware

import (
	"context"
	"microservices/modules/cart/database"
	"microservices/modules/clients"
	"net/http"
)

func AddDb(next http.Handler, db *database.Database) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", db)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func AddItemApi(next http.Handler, itemApi clients.ItemApi) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "itemApi", itemApi)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetDb(ctx context.Context) *database.Database {
	return ctx.Value("db").(*database.Database)
}
func GetItemApi(ctx context.Context) clients.ItemApi {
	return ctx.Value("itemApi").(clients.ItemApi)
}
