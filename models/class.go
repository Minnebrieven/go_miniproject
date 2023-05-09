package models

import (
	"time"
)

type Class struct {
	ID                uint               `gorm:"column:id;"`
	Name              string             `gorm:"column:name;"`
	ClassCategoryID   uint               `gorm:"column:class_category_id;"`
	ClassCategory     ClassCategory      `gorm:"constraint:OnUpdate:CASCADE;"`
	Description       string             `gorm:"column:description;"`
	InstructorID      uint               `gorm:"column:instructor_id;"`
	Instructor        Instructor         `gorm:"constraint:OnUpdate:CASCADE;"`
	ClassParticipants []ClassParticipant `gorm:"foreignKey:id;constraint:OnUpdate:CASCADE;"`
	Start             time.Time
	Metadata          Metadata `gorm:"embedded"`
}
