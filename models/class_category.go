package models

type ClassCategory struct {
	ID          uint `gorm:"column:id;"`
	Name        string
	Description string
	Classes     []Class  `gorm:"foreignKey:id;"`
	Metadata    Metadata `gorm:"embedded"`
}
