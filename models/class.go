package models

import (
	"time"
)

type Class struct {
	ID                uint               `gorm:"column:id;"`
	Name              string             `gorm:"column:name;"`
	ClassCategoryID   uint               `gorm:"column:class_category_id;"`
	ClassCategory     ClassCategory      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description       string             `gorm:"column:description;"`
	InstructorID      uint               `gorm:"column:instructor_id;"`
	Instructor        Instructor         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ClassParticipants []ClassParticipant `gorm:"foreignKey:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Start             time.Time
	Metadata          Metadata `gorm:"embedded"`
}
