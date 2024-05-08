package routes

import (
	"GoLinkTree/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func UserPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Get the username from the URL
		vars := mux.Vars(r)
		username := vars["username"]

		links, err := models.ShowUserLinks(username)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to get user details", http.StatusInternalServerError)
			return
		}

		// Marshal the links into JSON
		responseJSON, err := json.Marshal(links)
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
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
