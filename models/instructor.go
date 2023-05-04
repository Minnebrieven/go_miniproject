package models

import "time"

type Instructor struct {
	ID        uint
	Name      string
	Gender    string
	Phone     int
	Classes   []Class
	CreatedAt time.Time
	UpdatedAt time.Time
}
