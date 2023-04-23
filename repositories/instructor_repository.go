package repositories

import (
	"swim-class/models"

	"gorm.io/gorm"
)

type (
	InstructorRepository interface {
		GetAllInstructors() ([]models.Instructor, error)
		GetInstructor(models.Instructor) (models.Instructor, error)
		CreateInstructor(models.Instructor) error
		UpdateInstuctor(models.Instructor) (models.Instructor, error)
		DeleteInstructor(models.Instructor) error
	}

	instructorRepository struct {
		db *gorm.DB
	}
)

func NewInstructorRepository(db *gorm.DB) *instructorRepository {
	return &instructorRepository{db}
}

func (ir *instructorRepository) GetAllInstructors() ([]models.Instructor, error) {
	instructors := []models.Instructor{}
	err := ir.db.Find(&instructors).Error
	if err != nil {
		return nil, err
	}

	return instructors, nil
}

func (ir *instructorRepository) GetInstructor(instructor models.Instructor) (models.Instructor, error) {
	err := ir.db.First(&instructor).Error
	return instructor, err
}

func (ir *instructorRepository) CreateInstructor(instructorData models.Instructor) error {
	return ir.db.Create(&instructorData).Error
}

func (ir *instructorRepository) UpdateInstuctor(instructorData models.Instructor) (models.Instructor, error) {
	err := ir.db.Save(&instructorData).Error
	return instructorData, err
}

func (ir *instructorRepository) DeleteInstructor(instructorData models.Instructor) error {
	err := ir.db.Delete(&instructorData).Error
	return err
}
