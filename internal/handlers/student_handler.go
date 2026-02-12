package handlers

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"GroupBuilder/internal/database"
	"GroupBuilder/internal/models"

	"github.com/go-chi/chi/v5"
)

// GetAllStudents returns all students
func GetAllStudents(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]models.Student{})
	}
}

// CreateStudent creates a new student
func CreateStudent(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}
}

// GetStudent returns a single student by ID
func GetStudent(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		if _, err := strconv.ParseInt(idStr, 10, 64); err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// UpdateStudent updates an existing student
func UpdateStudent(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		if _, err := strconv.ParseInt(idStr, 10, 64); err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// DeleteStudent deletes a student
func DeleteStudent(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		if _, err := strconv.ParseInt(idStr, 10, 64); err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func ImportStudentList(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the multipart form
		if err := r.ParseMultipartForm(10 << 20); err != nil { // 10 MB max
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
		if err := saveStudents(db, students); err != nil {
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

		if len(record) < 3 {
			continue // Skip incomplete rows
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
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, student := range students {
		// 1. Get or Create Class ID
		var classID int64
		err := tx.QueryRow("SELECT id FROM classes WHERE name = ?", student.Class).Scan(&classID)
		if err == sql.ErrNoRows {
			res, err := tx.Exec("INSERT INTO classes (name) VALUES (?)", student.Class)
			if err != nil {
				return fmt.Errorf("failed to insert class %s: %v", student.Class, err)
			}
			classID, err = res.LastInsertId()
			if err != nil {
				return fmt.Errorf("failed to get class ID: %v", err)
			}
		} else if err != nil {
			return fmt.Errorf("failed to query class: %v", err)
		}

		// 2. Insert or Update Student
		_, err = tx.Exec(`
            INSERT INTO students (email, name, class_id)
            VALUES (?, ?, ?)
            ON CONFLICT(email) DO UPDATE SET
                name = excluded.name,
                class_id = excluded.class_id
        `, student.Email, student.Name, classID)
		if err != nil {
			return fmt.Errorf("failed to upsert student %s: %v", student.Email, err)
		}
	}

	return tx.Commit()
}
