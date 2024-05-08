package main

import (
	"GoLinkTree/middlewares"
	"GoLinkTree/models"
	"GoLinkTree/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize MongoDB
    err := models.InitMongoDB()
    if err != nil {
        log.Fatal("Error connecting to MongoDB:", err)
    }
	
	// Create a new mux router
	mux := mux.NewRouter()
	port := 8080

	// Add route to signup user
	mux.HandleFunc("/signup", routes.Signuphandler)

	// Add route for the login API
	mux.HandleFunc("/login", routes.LoginHandler)

	// Add route for the user page
	mux.HandleFunc("/{username}", routes.UserPage)

	// Add route for the user to enter details
	mux.HandleFunc("/{username}/addDetails", middlewares.AuthorizationMiddleware(routes.AddDetails))

	// Add route for the user to delete details
	mux.HandleFunc("/{username}/removeDetails", middlewares.AuthorizationMiddleware(routes.RemoveDetails))

	// Add route to validate the user in the user page
	mux.HandleFunc("/{username}/validate", middlewares.AuthorizationMiddleware(routes.ValidUser))

	// CORS middleware
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Start the server on port 8080
	log.Printf("Server listening on port %d", port)
	err = http.ListenAndServe(":8080", handlers.CORS(headers, origins, methods)(mux))
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
