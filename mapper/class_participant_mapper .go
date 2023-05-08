package mapper

import (
	"strconv"
	"swim-class/dto"
	"swim-class/models"
)

// from []Model to []DTO
func ToClassParticipantDTOList(classParticipantModelList []models.ClassParticipant) ([]dto.ClassParticipantDTO, error) {
	var err error
	classParticipantListDTO := make([]dto.ClassParticipantDTO, len(classParticipantModelList))

	for i, itm := range classParticipantModelList {
		classParticipantListDTO[i], err = ToClassParticipantDTO(itm)
	}

	return classParticipantListDTO, err
}

// from []DTO to []Model
func ToClassParticipantModelList(classParticipantDTOList []dto.ClassParticipantDTO) ([]models.ClassParticipant, error) {
	var err error
	classParticipantModelList := make([]models.ClassParticipant, len(classParticipantDTOList))

	for i, itm := range classParticipantDTOList {
		classParticipantModelList[i], err = ToClassParticipantModel(itm)
	}
	return classParticipantModelList, err
}

// from DTO to Model
func ToClassParticipantModel(dto dto.ClassParticipantDTO) (models.ClassParticipant, error) {
	var classParticipantModel models.ClassParticipant

	classParticipantModel.ID = uint(dto.ID)

	if dto.ClassID != 0 {
		classParticipantModel.ClassID = uint(dto.ClassID)
	}

	if dto.UserID != 0 {
		classParticipantModel.UserID = uint(dto.UserID)
	}

	return classParticipantModel, nil
}

// from Model to DTO
func ToClassParticipantDTO(classParticipantModel models.ClassParticipant) (dto.ClassParticipantDTO, error) {
	var classParticipantDTO dto.ClassParticipantDTO

	// parse or convert time to string
	dateFormat := "2006-01-02"
	datetimeFormat := "2006-01-02 15:04:05"

	// int to string
	instructorClassPhoneStr := strconv.Itoa(classParticipantModel.Class.Instructor.Phone)

	classParticipantDTO.ID = int(classParticipantModel.ID)

	classParticipantDTO.ClassID = int(classParticipantModel.ClassID)
	classParticipantDTO.Class.ID = int(classParticipantModel.Class.ID)
	classParticipantDTO.Class.Name = classParticipantModel.Class.Name

	classParticipantDTO.Class.ClassCategoryID = int(classParticipantModel.Class.ClassCategoryID)
	classParticipantDTO.Class.ClassCategory.ID = int(classParticipantModel.Class.ClassCategory.ID)
	classParticipantDTO.Class.ClassCategory.Name = classParticipantModel.Class.ClassCategory.Name
	classParticipantDTO.Class.ClassCategory.Description = classParticipantModel.Class.ClassCategory.Description
	classParticipantDTO.Class.ClassCategory.CreatedAt = classParticipantModel.Class.ClassCategory.Metadata.CreatedAt.Format(datetimeFormat)
	classParticipantDTO.Class.ClassCategory.UpdatedAt = classParticipantModel.Class.ClassCategory.Metadata.UpdatedAt.Format(datetimeFormat)

	classParticipantDTO.Class.Description = classParticipantModel.Class.Description
	classParticipantDTO.Class.Start = classParticipantModel.Class.Start.Format(datetimeFormat)

	classParticipantDTO.Class.InstructorID = int(classParticipantModel.Class.InstructorID)
	classParticipantDTO.Class.Instructor.ID = int(classParticipantModel.Class.Instructor.ID)
	classParticipantDTO.Class.Instructor.Name = classParticipantModel.Class.Instructor.Name
	classParticipantDTO.Class.Instructor.Gender = classParticipantModel.Class.Instructor.Gender
	classParticipantDTO.Class.Instructor.Phone = instructorClassPhoneStr
	classParticipantDTO.Class.Instructor.CreatedAt = classParticipantModel.Class.Instructor.Metadata.CreatedAt.Format(datetimeFormat)
	classParticipantDTO.Class.Instructor.UpdatedAt = classParticipantModel.Class.Instructor.Metadata.UpdatedAt.Format(datetimeFormat)

	classParticipantDTO.Class.CreatedAt = classParticipantModel.Metadata.CreatedAt.Format(datetimeFormat)
	classParticipantDTO.Class.UpdatedAt = classParticipantModel.Metadata.UpdatedAt.Format(datetimeFormat)

	classParticipantDTO.UserID = int(classParticipantModel.UserID)
	classParticipantDTO.User.ID = int(classParticipantModel.User.ID)
	classParticipantDTO.User.Name = classParticipantModel.User.Name
	classParticipantDTO.User.Email = classParticipantModel.User.Email
	classParticipantDTO.User.Password = "********"
	classParticipantDTO.User.Birthday = classParticipantModel.User.Birthday.Format(dateFormat)
	classParticipantDTO.User.CreatedAt = classParticipantModel.Metadata.CreatedAt.Format(datetimeFormat)
	classParticipantDTO.User.UpdatedAt = classParticipantModel.Metadata.UpdatedAt.Format(datetimeFormat)

	classParticipantDTO.CreatedAt = classParticipantModel.Metadata.CreatedAt.Format(datetimeFormat)
	classParticipantDTO.UpdatedAt = classParticipantModel.Metadata.UpdatedAt.Format(datetimeFormat)

	return classParticipantDTO, nil
}
