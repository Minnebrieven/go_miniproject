package mapper

import (
	"strconv"
	"swim-class/dto"
	"swim-class/models"
)

// from []Model to []DTO
func ToInstructorDTOList(instructorModelList []models.Instructor) ([]dto.InstructorDTO, error) {
	var err error
	instructorDTOList := make([]dto.InstructorDTO, len(instructorModelList))

	for i, itm := range instructorModelList {
		instructorDTOList[i], err = ToInstructorDTO(itm)
	}

	return instructorDTOList, err
}

// from []DTO to []Model
func ToInstructorModelList(instructorDTOList []dto.InstructorDTO) ([]models.Instructor, error) {
	var err error
	instructorModelList := make([]models.Instructor, len(instructorDTOList))

	for i, itm := range instructorDTOList {
		instructorModelList[i], err = ToInstructorModel(itm)
	}
	return instructorModelList, err
}

// from DTO to Model
func ToInstructorModel(dto dto.InstructorDTO) (models.Instructor, error) {
	var instructorModel models.Instructor

	instructorModel.ID = uint(dto.ID)
	instructorModel.Name = dto.Name
	instructorModel.Gender = dto.Gender

	// string to int
	if dto.Phone != "" {
		phoneStrToInt, err := strconv.Atoi(dto.Phone)
		if err != nil {
			return instructorModel, err
		}
		instructorModel.Phone = phoneStrToInt
	}

	return instructorModel, nil
}

// from Model to DTO
func ToInstructorDTO(instructorModel models.Instructor) (dto.InstructorDTO, error) {
	var instructorDTO dto.InstructorDTO

	// parse or convert time to string
	datetimeFormat := "2006-01-02 15:04:05"

	// int to string
	phoneIntToStr := strconv.Itoa(instructorModel.Phone)

	instructorDTO.ID = int(instructorModel.ID)
	instructorDTO.Name = instructorModel.Name
	instructorDTO.Gender = instructorModel.Gender
	instructorDTO.Phone = phoneIntToStr
	instructorDTO.Metadata.CreatedAt = instructorModel.Metadata.CreatedAt.Format(datetimeFormat)
	instructorDTO.Metadata.UpdatedAt = instructorModel.Metadata.UpdatedAt.Format(datetimeFormat)

	return instructorDTO, nil
}
