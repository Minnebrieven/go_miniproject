package repositories

import (
	"errors"
	"swim-class/models"

	"gorm.io/gorm"
)

type ClassRepository interface {
	GetAllClasses() ([]models.Class, error)
	GetClass(models.Class) (models.Class, error)
	CreateClass(models.Class) (models.Class, error)
	UpdateClass(models.Class) (models.Class, error)
	DeleteClass(models.Class) error
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *classRepository {
	return &classRepository{db}
}

func (cr *classRepository) GetAllClasses() ([]models.Class, error) {
	classes := []models.Class{}
	err := cr.db.Preload("ClassCategory").Preload("Instructor").Find(&classes).Error
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (cr *classRepository) GetClass(class models.Class) (models.Class, error) {
	err := cr.db.Preload("ClassCategory").Preload("Instructor").Find(&class).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return class, errors.New("record not found")
	}
	return class, err
}

func (cr *classRepository) CreateClass(classData models.Class) (models.Class, error) {
	err := cr.db.Create(&classData).Error
	return classData, err
}

func (cr *classRepository) UpdateClass(classData models.Class) (models.Class, error) {
	err := cr.db.Save(&classData).Error
	return classData, err
}

func (cr *classRepository) DeleteClass(classData models.Class) error {
	err := cr.db.First(&classData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	}
	err = cr.db.Delete(&classData).Error
	return err
}
