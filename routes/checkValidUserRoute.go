package routes

import (
	"encoding/json"
	"net/http"
)

func ValidUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Create a map for the JSON response
		responseJSON := map[string]string{
			"message": "successfully validated",
		}

		// Convert the map to JSON
		responseBytes, err := json.Marshal(responseJSON)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set the Content-Type header
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response to the client
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
