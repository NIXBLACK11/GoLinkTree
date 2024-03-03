package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetDetails(w http.ResponseWriter, r *http.Request) {
	if(r.Method=="POST") {
		vars := mux.Vars(r)
		username := vars["username"]

		fmt.Printf("Username is %s", username)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}