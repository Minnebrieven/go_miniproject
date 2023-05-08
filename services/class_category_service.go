package services

import (
	"reflect"
	"swim-class/dto"
	"swim-class/mapper"
	"swim-class/models"
	"swim-class/repositories"
)

type (
	ClassCategoryService interface {
		GetAllClassCategoriesService() ([]dto.ClassCategoryDTO, error)
		GetClassCategoryService(classCategoryDTO dto.ClassCategoryDTO) (dto.ClassCategoryDTO, error)
		CreateClassCategoryService(classCategoryDTO dto.ClassCategoryDTO) (dto.ClassCategoryDTO, error)
		EditClassCategoryService(classCategoryID int, modifiedClassCategoryData dto.ClassCategoryDTO) (dto.ClassCategoryDTO, error)
		DeleteClassCategoryService(classCategoryID int) error
	}

	classCategoryService struct {
		classCategoryRepository repositories.ClassCategoryRepository
	}
)

func NewClassCategoryService(classCategoryRepo repositories.ClassCategoryRepository) *classCategoryService {
	return &classCategoryService{classCategoryRepository: classCategoryRepo}
}

func (ccs *classCategoryService) GetAllClassCategoriesService() ([]dto.ClassCategoryDTO, error) {
	classCategories, err := ccs.classCategoryRepository.GetAllClassCategories()
	if err != nil {
		return nil, err
	}

	ClassCategoryDTOList, err := mapper.ToClassCategoryDTOList(classCategories)
	if err != nil {
		return nil, err
	}
	return ClassCategoryDTOList, nil
}

func (ccs *classCategoryService) GetClassCategoryService(classCategoryDTO dto.ClassCategoryDTO) (dto.ClassCategoryDTO, error) {
	classCategoryModel, err := mapper.ToClassCategoryModel(classCategoryDTO)
	if err != nil {
		return classCategoryDTO, err
	}

	result, err := ccs.classCategoryRepository.GetClassCategory(classCategoryModel)
	if err != nil {
		return classCategoryDTO, err
	}

	classCategoryDTO, err = mapper.ToClassCategoryDTO(result)
	if err != nil {
		return classCategoryDTO, err
	}
	return classCategoryDTO, nil
}

func (ccs *classCategoryService) CreateClassCategoryService(classCategoryDTO dto.ClassCategoryDTO) (dto.ClassCategoryDTO, error) {
	//convert DTO to model
	classCategoryModel, err := mapper.ToClassCategoryModel(classCategoryDTO)
	if err != nil {
		return classCategoryDTO, err
	}

	classCategoryModel, err = ccs.classCategoryRepository.CreateClassCategory(classCategoryModel)
	if err != nil {
		return classCategoryDTO, err
	}

	//
	classCategoryDTO, err = mapper.ToClassCategoryDTO(classCategoryModel)
	if err != nil {
		return classCategoryDTO, err
	}
	return classCategoryDTO, nil
}

func (ccs *classCategoryService) EditClassCategoryService(classCategoryID int, modifiedClassCategoryData dto.ClassCategoryDTO) (dto.ClassCategoryDTO, error) {
	//find record first if not exists return error
	classCategoryModel := models.ClassCategory{ID: uint(classCategoryID)}
	classCategoryModel, err := ccs.classCategoryRepository.GetClassCategory(classCategoryModel)
	if err != nil {
		return modifiedClassCategoryData, err
	}

	modifiedClassModel, err := mapper.ToClassCategoryModel(modifiedClassCategoryData)
	if err != nil {
		return modifiedClassCategoryData, err
	}

	//replace exist data with new one
	var classCategoryPointer *models.ClassCategory = &classCategoryModel
	classCategoryVal := reflect.ValueOf(classCategoryPointer).Elem()
	classCategoryType := classCategoryVal.Type()

	editVal := reflect.ValueOf(modifiedClassModel)

	for i := 0; i < classCategoryVal.NumField(); i++ {
		//skip ID, CreatedAt, UpdatedAt field to be edited
		switch classCategoryType.Field(i).Name {
		case "ID":
			continue
		case "CreatedAt":
			continue
		case "UpdatedAt":
			continue
		}

		editField := editVal.Field(i)
		isSet := editField.IsValid() && !editField.IsZero()
		if isSet {
			classCategoryVal.Field(i).Set(editVal.Field(i))
		}
	}

	result, err := ccs.classCategoryRepository.UpdateClassCategory(classCategoryModel)
	if err != nil {
		return modifiedClassCategoryData, err
	}

	modifiedClassCategoryData, err = mapper.ToClassCategoryDTO(result)
	return modifiedClassCategoryData, err

}

func (ccs *classCategoryService) DeleteClassCategoryService(classCategoryID int) error {
	classCategory := models.ClassCategory{ID: uint(classCategoryID)}
	err := ccs.classCategoryRepository.DeleteClassCategory(classCategory)
	return err
}
