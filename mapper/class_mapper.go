package mapper

import (
	"strconv"
	"swim-class/dto"
	"swim-class/models"
	"time"
)

// from []Model to []DTO
func ToClassDTOList(classModelList []models.Class) ([]dto.ClassDTO, error) {
	var err error
	userListDTO := make([]dto.ClassDTO, len(classModelList))

	for i, itm := range classModelList {
		userListDTO[i], err = ToClassDTO(itm)
	}

	return userListDTO, err
}

// from []DTO to []Model
func ToClassModelList(ClassDTOList []dto.ClassDTO) ([]models.Class, error) {
	var err error
	classModelList := make([]models.Class, len(ClassDTOList))

	for i, itm := range ClassDTOList {
		classModelList[i], err = ToClassModel(itm)
	}
	return classModelList, err
}

// from DTO to Model
func ToClassModel(dto dto.ClassDTO) (models.Class, error) {
	var classModel models.Class

	classModel.ID = uint(dto.ID)
	classModel.Name = dto.Name

	if dto.ClassCategoryID != "" {
		classCategoryIDString, err := strconv.Atoi(dto.ClassCategoryID)
		if err != nil {
			return classModel, err
		}
		classModel.ClassCategoryID = uint(classCategoryIDString)
	}

	classModel.Description = dto.Description

	if dto.Start != "" {
		// parse or convert start string to time
		datetimeFormat := "2006-01-02 15:04:05"
		parsedStrToTime, err := time.Parse(datetimeFormat, dto.Start)
		if err != nil {
			return classModel, err
		}

		classModel.Start = parsedStrToTime
	}

	if dto.InstructorID != "" {
		instructorIDString, err := strconv.Atoi(dto.InstructorID)
		if err != nil {
			return classModel, err
		}
		classModel.InstructorID = uint(instructorIDString)
	}

	return classModel, nil
}

// from Model to DTO
func ToClassDTO(classModel models.Class) (dto.ClassDTO, error) {
	var classDTO dto.ClassDTO

	// parse or convert time to string
	// dateFormat := "2006-01-02"
	datetimeFormat := "2006-01-02 15:04:05"

	// int to string
	classCategoryIDInt := strconv.Itoa(int(classModel.ClassCategoryID))
	instructorIDInt := strconv.Itoa(int(classModel.InstructorID))

	classDTO.ID = int(classModel.ID)
	classDTO.Name = classModel.Name
	classDTO.ClassCategoryID = classCategoryIDInt
	classDTO.Description = classModel.Description
	classDTO.Start = classModel.Start.Format(datetimeFormat)
	classDTO.InstructorID = instructorIDInt
	classDTO.CreatedAt = classModel.CreatedAt.Format(datetimeFormat)
	classDTO.UpdatedAt = classModel.UpdatedAt.Format(datetimeFormat)

	return classDTO, nil
}
