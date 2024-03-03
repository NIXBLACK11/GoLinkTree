package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"GoLinkTree/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse the request body
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		// Check if username and password are correct
		if user.Username == "admin" && user.Password == "password" {
			// Return a success response
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Login successful")
		} else {
			// Return an error response
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid username or password")
		}
	} else {
		// Return an error response for unsupported methods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
