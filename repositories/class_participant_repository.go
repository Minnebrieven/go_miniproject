package repositories

import (
	"errors"
	"swim-class/models"

	"gorm.io/gorm"
)

type ClassParticipantRepository interface {
	GetAllClassParticipants() ([]models.ClassParticipant, error)
	GetClassParticipant(models.ClassParticipant) (models.ClassParticipant, error)
	CreateClassParticipant(models.ClassParticipant) (models.ClassParticipant, error)
	UpdateClassParticipant(models.ClassParticipant) (models.ClassParticipant, error)
	DeleteClassParticipant(models.ClassParticipant) error
}

type classParticipantRepository struct {
	db *gorm.DB
}

func NewClassParticipantRepository(db *gorm.DB) *classParticipantRepository {
	return &classParticipantRepository{db}
}

func (cr *classParticipantRepository) GetAllClassParticipants() ([]models.ClassParticipant, error) {
	classParticipants := []models.ClassParticipant{}
	err := cr.db.Find(&classParticipants).Error
	if err != nil {
		return nil, err
	}

	return classParticipants, nil
}

func (cr *classParticipantRepository) GetClassParticipant(classParticipant models.ClassParticipant) (models.ClassParticipant, error) {
	err := cr.db.First(&classParticipant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return classParticipant, errors.New("record not found")
	}
	return classParticipant, err
}

func (cr *classParticipantRepository) CreateClassParticipant(classParticipantData models.ClassParticipant) (models.ClassParticipant, error) {
	err := cr.db.Create(&classParticipantData).Error
	return classParticipantData, err
}

func (cr *classParticipantRepository) UpdateClassParticipant(classParticipantData models.ClassParticipant) (models.ClassParticipant, error) {
	err := cr.db.Save(&classParticipantData).Error
	return classParticipantData, err
}

func (cr *classParticipantRepository) DeleteClassParticipant(classParticipantData models.ClassParticipant) error {
	err := cr.db.Delete(&classParticipantData).Error
	return err
}
