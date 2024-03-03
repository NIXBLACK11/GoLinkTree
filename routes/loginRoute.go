package routes

import (
	"GoLinkTree/models"
	"GoLinkTree/jwt"
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
				token, err = jwt.CreateToken(user.Username)
				if err != nil {
					http.Error(w, "Error in user authentication")
				}
				w.WriteHeader(http.StatusAccepted)
				fmt.Fprintf(w, "Successfully logged in user")
			} else {
				w.WriteHeader(http.StatusExpectationFailed)
				fmt.Fprintf(w, "User does not exist")
			}
		}
	} else {
		// Return an error response for unsupported methods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
