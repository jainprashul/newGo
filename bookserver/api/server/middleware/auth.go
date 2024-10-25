package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"xpJain.co/bookserver/auth"
)

// Middleware to authenticate the user
func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request has a valid token
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// split the header into two parts
		splits := strings.Split(authHeader, " ")
		// The first part should be the token type
		if len(splits) != 2 {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}
		if splits[0] != "Bearer" {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}
		// The second part should be the token
		token := splits[1]
		
		// Verify the token
		
		claims, err := auth.VerifyToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Println("Username: ", claims.Username)

		// Call the next handler
		next.ServeHTTP(w, r)


	})
}