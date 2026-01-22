package main

import (
	"log"
	"net/http"

	"GroupBuilder/internal/database"
	"GroupBuilder/internal/routes"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
            log.Printf("Error closing database: %v", err)
		}
	}(db)

	r := routes.SetupRoutes(db)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
