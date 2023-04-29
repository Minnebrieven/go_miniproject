package models

import "time"

type ClassCategory struct {
	ID          uint
	Name        string
	Description string
	Classes     []Class
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
