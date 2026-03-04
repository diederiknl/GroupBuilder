package main

import (
	"database/sql"
	"log"
	"net/http"

	"GroupBuilder/internal/auth"
	"GroupBuilder/internal/database"
	"GroupBuilder/internal/routes"
)

func main() {
	if err := auth.InitKey(); err != nil {
		log.Fatalf("Failed to initialize auth key: %v", err)
	}

	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	r := routes.SetupRoutes(db)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
