package middlewares

import (
	"GoLinkTree/jwt"
	"GoLinkTree/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func AuthorizationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		var auth models.Auth
		err := json.NewDecoder(r.Body).Decode(&auth)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		
		// Extract the username from the parameters and token from the json data
		vars := mux.Vars(r)
		username := vars["username"]
		Bearertoken := auth.Token

		token := strings.Split(Bearertoken, " ")[1]

		check, err := jwt.AuthToken(username, token)
		
		log.Println(err)

		if(err!=nil) {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		
		if(!check) {
			http.Error(w, "Invalid username", http.StatusBadRequest)
			return
		}

		fmt.Println("Middleware executed")

		next.ServeHTTP(w, r)
	}
}