package middlewares

import (
	"GoLinkTree/jwt"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func AuthorizationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// Split the Authorization header to get the token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}
		token := tokenParts[1]

		// Extract the username from the request
		vars := mux.Vars(r)
		username := vars["username"]

		// Validate the token
		check, err := jwt.AuthToken(username, token)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if !check {
			http.Error(w, "Invalid username", http.StatusBadRequest)
			return
		}

		fmt.Println("Middleware executed")

		next.ServeHTTP(w, r)
	}
}
