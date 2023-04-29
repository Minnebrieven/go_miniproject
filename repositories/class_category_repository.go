package repositories

import (
	"errors"
	"swim-class/models"

	"gorm.io/gorm"
)

type ClassCategoryRepository interface {
	GetAllClassCategories() ([]models.ClassCategory, error)
	GetClassCategory(models.ClassCategory) (models.ClassCategory, error)
	CreateClassCategory(models.ClassCategory) (models.ClassCategory, error)
	UpdateClassCategory(models.ClassCategory) (models.ClassCategory, error)
	DeleteClassCategory(models.ClassCategory) error
}

type classCategoryRepository struct {
	db *gorm.DB
}

func NewClassCategoryRepository(db *gorm.DB) *classCategoryRepository {
	return &classCategoryRepository{db}
}

func (cr *classCategoryRepository) GetAllClassCategories() ([]models.ClassCategory, error) {
	classCategories := []models.ClassCategory{}
	err := cr.db.Find(&classCategories).Error
	if err != nil {
		return nil, err
	}

	return classCategories, nil
}

func (cr *classCategoryRepository) GetClassCategory(classCategory models.ClassCategory) (models.ClassCategory, error) {
	err := cr.db.First(&classCategory).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return classCategory, errors.New("record not found")
	}
	return classCategory, err
}

func (cr *classCategoryRepository) CreateClassCategory(classCategoryData models.ClassCategory) (models.ClassCategory, error) {
	err := cr.db.Create(&classCategoryData).Error
	return classCategoryData, err
}

func (cr *classCategoryRepository) UpdateClassCategory(classCategoryData models.ClassCategory) (models.ClassCategory, error) {
	err := cr.db.Save(&classCategoryData).Error
	return classCategoryData, err
}

func (cr *classCategoryRepository) DeleteClassCategory(classCategoryData models.ClassCategory) error {
	err := cr.db.Delete(&classCategoryData).Error
	return err
}
