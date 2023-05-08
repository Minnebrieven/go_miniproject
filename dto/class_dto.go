package dto

type ClassDTO struct {
	ID            int              `json:"id"`
	Name          string           `json:"name"`
	ClassCategory ClassCategoryDTO `json:"class_category"`
	Description   string           `json:"description"`
	Start         string           `json:"start"`
	Instructor    InstructorDTO    `json:"instructor"`
	Metadata      MetadataDTO      `json:"metadata"`
}
