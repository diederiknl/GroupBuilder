package main

import (
	//"context"
	"log"
	"net/http"

	"GroupBuilder/internal/database"
	"GroupBuilder/internal/routes"
)

// I know this key should not be here. No problem for now...
var jwtKey = []byte("neinneinnein")

func main() {
	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Setup routes
	r := routes.SetupRoutes(db)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
