package routes

import (
	"encoding/json"
	"net/http"
)

type TestData struct {
	Name  string `json:"name"`
}

func TestRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data TestData

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid input data", http.StatusBadRequest)
			return
		}

		response := map[string]string{
			"message": "Data received successfully",
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
