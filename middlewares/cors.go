package middlewares

import (
	"net/http"
	"strings"
)

var methods = []string{http.MethodGet, http.MethodPut, http.MethodOptions}
var allowedMethods = strings.Join(methods, ", ")

// CorsMiddleware adds the CORS headers to response
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", allowedMethods)

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
