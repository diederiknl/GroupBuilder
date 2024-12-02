package models

type Project struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Period int    `json:"period"`
}
