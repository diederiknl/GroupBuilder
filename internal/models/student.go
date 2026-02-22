package models

import "time"

type Student struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	ClassID int64  `json:"class_id"`
	Class   string `json:"class"`
	GroupID int64  `json:"group_id"`
}

type StudentLoginToken struct {
	ID        int64     `json:"id"`
	StudentID int64     `json:"student_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
