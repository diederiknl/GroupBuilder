package models

type Group struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Class     string `json:"class"`
	ProjectID int    `json:"project_id"`
}
