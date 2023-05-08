package repositories

import (
	"errors"
	"swim-class/models"

	"gorm.io/gorm"
)

type ClassParticipantRepository interface {
	GetAllClassParticipants() ([]models.ClassParticipant, error)
	GetAllClassParticipantsByUserID(userID int) ([]models.ClassParticipant, error)
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

func (cpr *classParticipantRepository) GetAllClassParticipants() ([]models.ClassParticipant, error) {
	classParticipants := []models.ClassParticipant{}
	err := cpr.db.Preload("Class").Preload("Class.ClassCategory").Preload("Class.Instructor").Preload("User").Find(&classParticipants).Error
	if err != nil {
		return nil, err
	}

	return classParticipants, nil
}

func (cpr *classParticipantRepository) GetAllClassParticipantsByUserID(userID int) ([]models.ClassParticipant, error) {
	classParticipants := []models.ClassParticipant{}
	err := cpr.db.Preload("Class").Preload("Class.ClassCategory").Preload("Class.Instructor").Preload("User").Find(&classParticipants, "class_participants.user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}

	return classParticipants, nil
}

func (cpr *classParticipantRepository) GetClassParticipant(classParticipant models.ClassParticipant) (models.ClassParticipant, error) {
	err := cpr.db.Preload("Class").Preload("Class.ClassCategory").Preload("Class.Instructor").Preload("User").Find(&classParticipant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return classParticipant, errors.New("record not found")
	}
	return classParticipant, err
}

func (cpr *classParticipantRepository) CreateClassParticipant(classParticipantData models.ClassParticipant) (models.ClassParticipant, error) {
	err := cpr.db.Create(&classParticipantData).Error
	return classParticipantData, err
}

func (cpr *classParticipantRepository) UpdateClassParticipant(classParticipantData models.ClassParticipant) (models.ClassParticipant, error) {
	err := cpr.db.Save(&classParticipantData).Error
	return classParticipantData, err
}

func (cpr *classParticipantRepository) DeleteClassParticipant(classParticipantData models.ClassParticipant) error {
	err := cpr.db.First(&classParticipantData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	}
	err = cpr.db.Delete(&classParticipantData).Error
	return err
}
