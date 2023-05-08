package dto

type ClassParticipantDTO struct {
	ID        int      `json:"id"`
	ClassID   int      `json:"class_id"`
	Class     ClassDTO `json:"class"`
	UserID    int      `json:"user_id"`
	User      UserDTO  `json:"user"`
	CreatedAt string   `json:"-"`
	UpdatedAt string   `json:"-"`
}
