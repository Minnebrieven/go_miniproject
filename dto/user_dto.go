package dto

type UserDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
	Metadata MetadataDTO `json:"metadata"`
}
