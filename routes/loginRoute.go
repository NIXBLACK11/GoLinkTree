package routes

import (
	"GoLinkTree/models"
	"encoding/json"
	"fmt"
	"net/http"
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

		exists, err := models.CheckUserExists(user.Username, user.Password)

		if err != nil {
			http.Error(w, "Error occured in user authentication", http.StatusBadRequest)
		} else {
			if exists {
				w.WriteHeader(http.StatusAccepted)
				fmt.Fprintf(w, "Successfully logged in user")
			}
		}
	} else {
		// Return an error response for unsupported methods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
