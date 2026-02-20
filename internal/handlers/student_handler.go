package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"GroupBuilder/internal/database"
	"GroupBuilder/internal/models"
)

// ImportStudentList imports a list of students from a CSV file
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
		// We use parameter binding to prevent SQL injection, adhering to security best practices.
		_, err := tx.Exec(`
            INSERT INTO students (email, name, class_id)
            VALUES (?, ?, (SELECT id FROM classes WHERE name = ?))
            ON CONFLICT(email) DO UPDATE SET
                name = excluded.name,
                class_id = excluded.class_id
        `, student.Email, student.Name, student.Class)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// VerifyStudentLoginLink handles the verification of the login link sent to students
func VerifyStudentLoginLink(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Stub implementation
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Student verified"})
	}
}

// GetAllStudents retrieves all students
func GetAllStudents(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Stub implementation
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]models.Student{})
	}
}

// CreateStudent creates a new student
func CreateStudent(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Stub implementation
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Student created"})
	}
}

// GetStudent retrieves a single student by ID
func GetStudent(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		// Stub implementation
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"id": id, "name": "Stub Student"})
	}
}

// UpdateStudent updates an existing student
func UpdateStudent(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		_, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		// Stub implementation
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Student updated"})
	}
}

// DeleteStudent deletes a student
func DeleteStudent(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		_, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		// Stub implementation
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted"})
	}
}
