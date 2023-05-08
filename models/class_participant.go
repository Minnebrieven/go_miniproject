package models

type ClassParticipant struct {
	ID       uint
	ClassID  uint `gorm:"column:class_id;"`
	Class    Class
	UserID   uint `gorm:"column:user_id;"`
	User     User
	Metadata Metadata `gorm:"embedded"`
}
