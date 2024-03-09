package routes

import (
	"GoLinkTree/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RemoveDetails(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST" {
		vars := mux.Vars(r)
		username := vars["username"]

		// Decode the request body into a slice of Link structs
		var link models.RemLink
		err := json.NewDecoder(r.Body).Decode(&link)
		if err != nil {
			http.Error(w, "Failed to parse request body2", http.StatusBadRequest)
			return
		}

		// Call the function to insert the link
		success, err := models.DeleteLink(username, link)
		if err != nil {
			http.Error(w, "Failed to delete link", http.StatusInternalServerError)
			return
		}

		if success {
			w.WriteHeader(http.StatusOK)
			response := map[string]string{"message": "Update successful"}
			json.NewEncoder(w).Encode(response)
			return
		} else {
			http.Error(w, "Failed to delete link", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}