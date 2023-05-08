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
	classListDTO := make([]dto.ClassDTO, len(classModelList))

	for i, itm := range classModelList {
		classListDTO[i], err = ToClassDTO(itm)
	}

	return classListDTO, err
}

// from []DTO to []Model
func ToClassModelList(classDTOList []dto.ClassDTO) ([]models.Class, error) {
	var err error
	classModelList := make([]models.Class, len(classDTOList))

	for i, itm := range classDTOList {
		classModelList[i], err = ToClassModel(itm)
	}
	return classModelList, err
}

// from DTO to Model
func ToClassModel(dto dto.ClassDTO) (models.Class, error) {
	var classModel models.Class

	classModel.ID = uint(dto.ID)
	classModel.Name = dto.Name

	if dto.ClassCategoryID != 0 {
		classModel.ClassCategoryID = uint(dto.ClassCategoryID)
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

	if dto.InstructorID != 0 {
		classModel.InstructorID = uint(dto.InstructorID)
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
	instructorPhoneInt := strconv.Itoa(int(classModel.Instructor.Phone))

	classDTO.ID = int(classModel.ID)
	classDTO.Name = classModel.Name

	classDTO.ClassCategoryID = int(classModel.ClassCategoryID)
	classDTO.ClassCategory.ID = int(classModel.ClassCategory.ID)
	classDTO.ClassCategory.Name = classModel.ClassCategory.Name
	classDTO.ClassCategory.Description = classModel.ClassCategory.Description
	classDTO.ClassCategory.CreatedAt = classModel.Metadata.CreatedAt.Format(datetimeFormat)
	classDTO.ClassCategory.UpdatedAt = classModel.Metadata.UpdatedAt.Format(datetimeFormat)

	classDTO.Description = classModel.Description
	classDTO.Start = classModel.Start.Format(datetimeFormat)

	classDTO.InstructorID = int(classModel.InstructorID)
	classDTO.Instructor.ID = int(classModel.Instructor.ID)
	classDTO.Instructor.Name = classModel.Instructor.Name
	classDTO.Instructor.Gender = classModel.Instructor.Gender
	classDTO.Instructor.Phone = instructorPhoneInt
	classDTO.Instructor.CreatedAt = classModel.Metadata.CreatedAt.Format(datetimeFormat)
	classDTO.Instructor.UpdatedAt = classModel.Metadata.UpdatedAt.Format(datetimeFormat)

	classDTO.CreatedAt = classModel.Metadata.CreatedAt.Format(datetimeFormat)
	classDTO.UpdatedAt = classModel.Metadata.UpdatedAt.Format(datetimeFormat)

	return classDTO, nil
}
