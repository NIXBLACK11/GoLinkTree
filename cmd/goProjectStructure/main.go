package main

import (
	"GoLinkTree/routes"
	"GoLinkTree/middlewares"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new mux router
	mux := mux.NewRouter()
	port := 8080

	// Add route for the login API
	mux.HandleFunc("/login", routes.LoginHandler)

	// Add route for the user page
	mux.HandleFunc("/{username}", middlewares.AuthorizationMiddleware(routes.UserPage))

	//Add route for the user to enter details
	

	// Start the server on port 8080
	log.Printf("Server listening on port %d", port)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
