package models

type Feedback struct {
	ID        int    `json:"id"`
	StudentID int    `json:"student_id"`
	GroupID   int    `json:"group_id"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
}
