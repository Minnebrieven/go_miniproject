package dto

type InstructorDTO struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
	MetadataDTO
}
