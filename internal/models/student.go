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
	Email string `json:"email"`
	Name  string `json:"name"`
	Class string `json:"class"`
}
