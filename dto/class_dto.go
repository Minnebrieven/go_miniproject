package dto

type ClassDTO struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	ClassCategoryID string `json:"class_category_id"`
	Description     string `json:"description"`
	Start           string `json:"start"`
	InstructorID    string `json:"instructor_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
