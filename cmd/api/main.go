package main

import (
	"database/sql"
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
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}(db)

	// Note: In a complete implementation, we would wrap db into *database.DB if needed.
	// For now, we assume this is correct or will be fixed elsewhere.
	r := routes.SetupRoutes(db)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
