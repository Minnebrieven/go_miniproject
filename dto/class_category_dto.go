package dto

type ClassCategoryDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"-"`
	UpdatedAt   string `json:"-"`
}
