package models

import (
	"time"
)

type User struct {
	ID                uint      `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Birthday          time.Time `json:"birthday"`
	ClassParticipants []ClassParticipant
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
