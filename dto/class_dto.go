package dto

type ClassDTO struct {
	ID              int              `json:"id"`
	Name            string           `json:"name"`
	ClassCategoryID int              `json:"class_category_id"`
	ClassCategory   ClassCategoryDTO `json:"class_category"`
	Description     string           `json:"description"`
	Start           string           `json:"start"`
	InstructorID    int              `json:"instructor_id"`
	Instructor      InstructorDTO    `json:"instructor"`
	MetadataDTO
}
