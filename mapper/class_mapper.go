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

	if dto.ClassCategory.ID != 0 {
		classModel.ClassCategoryID = uint(dto.ClassCategory.ID)
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

	if dto.Instructor.ID != 0 {
		classModel.InstructorID = uint(dto.Instructor.ID)
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

	classDTO.ClassCategory.ID = int(classModel.ClassCategory.ID)
	classDTO.ClassCategory.Name = classModel.ClassCategory.Name
	classDTO.ClassCategory.Description = classModel.ClassCategory.Description
	classDTO.ClassCategory.Metadata.CreatedAt = classModel.Metadata.CreatedAt.Format(datetimeFormat)
	classDTO.ClassCategory.Metadata.UpdatedAt = classModel.Metadata.UpdatedAt.Format(datetimeFormat)

	classDTO.Description = classModel.Description
	classDTO.Start = classModel.Start.Format(datetimeFormat)

	classDTO.Instructor.ID = int(classModel.Instructor.ID)
	classDTO.Instructor.Name = classModel.Instructor.Name
	classDTO.Instructor.Gender = classModel.Instructor.Gender
	classDTO.Instructor.Phone = instructorPhoneInt
	classDTO.Instructor.Metadata.CreatedAt = classModel.Metadata.CreatedAt.Format(datetimeFormat)
	classDTO.Instructor.Metadata.UpdatedAt = classModel.Metadata.UpdatedAt.Format(datetimeFormat)

	classDTO.Metadata.CreatedAt = classModel.Metadata.CreatedAt.Format(datetimeFormat)
	classDTO.Metadata.UpdatedAt = classModel.Metadata.UpdatedAt.Format(datetimeFormat)

	return classDTO, nil
}
