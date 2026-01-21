package main

import (
	"log"
	"net/http"

	"GroupBuilder/internal/database"
	"GroupBuilder/internal/routes"
)

var jwtKey = []byte("neinneinnein")

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	r := routes.SetupRoutes(db)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
