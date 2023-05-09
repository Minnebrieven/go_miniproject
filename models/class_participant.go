package models

type ClassParticipant struct {
	ID       uint
	ClassID  uint     `gorm:"column:class_id;"`
	Class    Class    `gorm:"foreignKey:ClassID"`
	UserID   uint     `gorm:"column:user_id;"`
	User     User     `gorm:"foreignKey:UserID"`
	Metadata Metadata `gorm:"embedded"`
}
