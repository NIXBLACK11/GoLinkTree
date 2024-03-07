package routes

import (
	"net/http"
)

func SetDetails(w http.ResponseWriter, r *http.Request) {
	if(r.Method=="POST") {
		// vars := mux.Vars(r)
		// username := vars["username"]
		
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}