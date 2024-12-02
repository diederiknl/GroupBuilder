package handlers

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"

	"github.com/diederiknl/GroupBuilder/internal/database"
	"github.com/diederiknl/GroupBuilder/internal/models"
)

func ImportStudentList(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the multipart form
		err := r.ParseMultipartForm(10 << 20) // 10 MB max
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Get the file from the form
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Check if the file is a CSV
		if header.Header.Get("Content-Type") != "text/csv" {
			http.Error(w, "Please upload a CSV file", http.StatusBadRequest)
			return
		}

		// Process the CSV file
		students, err := processCSV(file)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error processing CSV: %v", err), http.StatusInternalServerError)
			return
		}

		// Save students to database
		err = saveStudents(db, students)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error saving students: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Successfully imported %d students", len(students))
	}
}

func processCSV(file io.Reader) ([]models.Student, error) {
	reader := csv.NewReader(file)
	var students []models.Student

	// Skip the header row
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		student := models.Student{
			Email: record[0],
			Name:  record[1],
			Class: record[2],
		}
		students = append(students, student)
	}

	return students, nil
}

func saveStudents(db *database.DB, students []models.Student) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, student := range students {
		_, err := tx.Exec(`
            INSERT INTO students (email, name, class)
            VALUES (?, ?, ?)
            ON CONFLICT(email) DO UPDATE SET
                name = excluded.name,
                class = excluded.class
        `, student.Email, student.Name, student.Class)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
