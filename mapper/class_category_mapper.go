package mapper

import (
	"swim-class/dto"
	"swim-class/models"
)

// from []Model to []DTO
func ToClassCategoryDTOList(classCategoryModelList []models.ClassCategory) ([]dto.ClassCategoryDTO, error) {
	var err error
	classCategoryListDTO := make([]dto.ClassCategoryDTO, len(classCategoryModelList))

	for i, itm := range classCategoryModelList {
		classCategoryListDTO[i], err = ToClassCategoryDTO(itm)
	}

	return classCategoryListDTO, err
}

// from []DTO to []Model
func ToClassCategoryModelList(classCategoryDTOList []dto.ClassCategoryDTO) ([]models.ClassCategory, error) {
	var err error
	classCategoryModelList := make([]models.ClassCategory, len(classCategoryDTOList))

	for i, itm := range classCategoryDTOList {
		classCategoryModelList[i], err = ToClassCategoryModel(itm)
	}
	return classCategoryModelList, err
}

// from DTO to Model
func ToClassCategoryModel(dto dto.ClassCategoryDTO) (models.ClassCategory, error) {
	var classCategoryModel models.ClassCategory

	classCategoryModel.ID = uint(dto.ID)
	classCategoryModel.Name = dto.Name
	classCategoryModel.Description = dto.Description

	return classCategoryModel, nil
}

// from Model to DTO
func ToClassCategoryDTO(classCategoryModel models.ClassCategory) (dto.ClassCategoryDTO, error) {
	var classDTO dto.ClassCategoryDTO

	// parse or convert time to string
	// dateFormat := "2006-01-02"
	datetimeFormat := "2006-01-02 15:04:05"

	classDTO.ID = int(classCategoryModel.ID)
	classDTO.Name = classCategoryModel.Name
	classDTO.Description = classCategoryModel.Description
	classDTO.CreatedAt = classCategoryModel.Metadata.CreatedAt.Format(datetimeFormat)
	classDTO.UpdatedAt = classCategoryModel.Metadata.UpdatedAt.Format(datetimeFormat)

	return classDTO, nil
}
