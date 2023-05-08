package dto

type ClassParticipantDTO struct {
	ID       int         `json:"id"`
	Class    ClassDTO    `json:"class"`
	User     UserDTO     `json:"user"`
	Metadata MetadataDTO `json:"metadata"`
}
