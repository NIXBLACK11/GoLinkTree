package routes

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func UserPage(w http.ResponseWriter, r *http.Request) {
	if(r.Method=="GET") {
		// Get the username from the URL
		vars := mux.Vars(r)
		username := vars["username"]

		// Do something with the username
		fmt.Fprintf(w, "Hello, %s!", username)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}