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
			http.Error(w, "Error occurred in user authentication", http.StatusBadRequest)
		} else {
			if exists {
				token, err := jwt.CreateToken(user.Username)
				if err != nil {
					http.Error(w, "Error in user authentication", http.StatusBadRequest)
				}

				// Create a map to hold the response data
				response := map[string]string{"token": "Bearer " + token}

				// Marshal the response into JSON
				responseJSON, err := json.Marshal(response)
				if err != nil {
					http.Error(w, "Failed to create response", http.StatusInternalServerError)
					return
				}

				// Set the Content-Type header
				w.Header().Set("Content-Type", "application/json")

				// Write the JSON response to the client
				w.WriteHeader(http.StatusOK)
				w.Write(responseJSON)
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
