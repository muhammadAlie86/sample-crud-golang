package middlewares

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Auth Middleware")

		next.ServeHTTP(w, r)
	})
}
