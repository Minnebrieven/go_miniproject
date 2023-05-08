package models

type Instructor struct {
	ID       uint `gorm:"column:id;"`
	Name     string
	Gender   string
	Phone    int
	Classes  []Class  `gorm:"foreignKey:id;"`
	Metadata Metadata `gorm:"embedded"`
}
