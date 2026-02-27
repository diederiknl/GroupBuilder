package models

import "time"

type StudentLoginToken struct {
	ID        int64     `json:"id"`
	StudentID int64     `json:"student_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Class string `json:"class"`
}
