package models

import (
	"time"
)

type User struct {
	ID                uint `gorm:"column:id;"`
	Name              string
	Email             string `gorm:"unique"`
	Password          string
	Birthday          time.Time
	ClassParticipants []ClassParticipant `gorm:"foreignKey:id;"`
	IsAdmin           bool               `gorm:"default:false;"`
	Metadata          Metadata           `gorm:"embedded"`
}
