package middleware

import (
	"log"
	"net/http"
)

// Logger is a middleware that logs the request method and URL
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		next.ServeHTTP(w, r)
	})

}
