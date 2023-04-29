package models

import "time"

type Class struct {
	ID              uint
	Name            string
	ClassCategoryID uint
	Description     string
	InstructorID    uint
	ClassParticipants []ClassParticipant
	Start           time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
