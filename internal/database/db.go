package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./groupbuilder.db")
	if err != nil {
		return nil, err
	}

	// Create tables if they don't exist
	_, err = db.Exec(`-- Projects table
CREATE TABLE IF NOT EXISTS projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    year INTEGER NOT NULL,
    period INTEGER NOT NULL,
    UNIQUE(year, period)
);

-- Classes table
CREATE TABLE IF NOT EXISTS classes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

-- Groups table
CREATE TABLE IF NOT EXISTS groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    class_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    FOREIGN KEY (class_id) REFERENCES classes(id),
    FOREIGN KEY (project_id) REFERENCES projects(id),
    UNIQUE(name, project_id)
);

-- Students table
CREATE TABLE IF NOT EXISTS students (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    class_id INTEGER NOT NULL,
    group_id INTEGER,
    FOREIGN KEY (class_id) REFERENCES classes(id),
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

-- Teachers table
CREATE TABLE IF NOT EXISTS teachers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);

-- Student feedback table
CREATE TABLE IF NOT EXISTS student_feedback (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    student_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    satisfaction_level INTEGER NOT NULL,
    comments TEXT,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

-- Student preferences table
CREATE TABLE IF NOT EXISTS student_preferences (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    student_id INTEGER NOT NULL,
    preferred_student_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    preference_type TEXT NOT NULL,  -- 'PREFER' or 'AVOID'
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (preferred_student_id) REFERENCES students(id),
    FOREIGN KEY (project_id) REFERENCES projects(id),
    UNIQUE(student_id, preferred_student_id, project_id)
);

-- Teacher pins table
CREATE TABLE IF NOT EXISTS teacher_pins (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    teacher_id INTEGER NOT NULL,
    student_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    pin_type TEXT NOT NULL,  -- 'PIN' or 'UNPIN'
    FOREIGN KEY (teacher_id) REFERENCES teachers(id),
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (group_id) REFERENCES groups(id),
    FOREIGN KEY (project_id) REFERENCES projects(id),
    UNIQUE(student_id, project_id)
);    `)

	if err != nil {
		return nil, err
	}

	return db, nil
}
