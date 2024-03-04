package middlewares

import (
	"GoLinkTree/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func AuthorizationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		var auth models.Auth
		err := json.NewDecoder(r.Body).Decode(&auth)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		

		
		fmt.Println("Middleware executed")

		
		next.ServeHTTP(w, r)
	}
}