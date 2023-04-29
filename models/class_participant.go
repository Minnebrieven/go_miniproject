package models

import "time"

type ClassParticipant struct {
	ID        uint
	ClassID   uint
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
