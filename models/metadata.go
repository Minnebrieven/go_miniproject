package models

import (
	"time"

	"gorm.io/gorm"
)

type Metadata struct {
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time
}
