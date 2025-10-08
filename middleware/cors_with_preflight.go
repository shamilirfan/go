package middleware

import "net/http"

func CorsWithPreflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // * meanes any IP Address can access.
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
