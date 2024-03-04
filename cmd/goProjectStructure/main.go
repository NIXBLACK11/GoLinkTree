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
	mux.HandleFunc("/login", middlewares.AuthorizationMiddleware(routes.LoginHandler))

	// Add route for the user page
	mux.HandleFunc("/{username}", routes.UserPage)

	//Add route for the user to enter details
	mux.HandleFunc("/{username}/setDetails", routes.SetDetails)

	// Start the server on port 8080
	log.Printf("Server listening on port %d", port)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
