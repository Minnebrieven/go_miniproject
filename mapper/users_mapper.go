package mapper

import (
	"swim-class/dto"
	"swim-class/models"
	"time"
)

// from []Model to []DTO
func ToUserDTOList(userModelList []models.User) ([]dto.UserDTO, error) {
	var err error
	userListDTO := make([]dto.UserDTO, len(userModelList))

	for i, itm := range userModelList {
		userListDTO[i], err = ToUserDTO(itm)
	}

	return userListDTO, err
}

// from []DTO to []Model
func ToUserModelList(userDTOList []dto.UserDTO) ([]models.User, error) {
	var err error
	userModelList := make([]models.User, len(userDTOList))

	for i, itm := range userDTOList {
		userModelList[i], err = ToUserModel(itm)
	}
	return userModelList, err
}

// from DTO to Model
func ToUserModel(dto dto.UserDTO) (models.User, error) {
	var userModel models.User

	userModel.ID = uint(dto.ID)
	userModel.Name = dto.Name
	userModel.Email = dto.Email
	userModel.Password = dto.Password
	
	if dto.Birthday != "" {
		// parse or convert birthday string to time
		dateFormat := "2006-01-02"
		parsedStrToTime, err := time.Parse(dateFormat, dto.Birthday)
		if err != nil {
			return userModel, err
		}

		userModel.Birthday = parsedStrToTime
	}

	return userModel, nil
}

// from Model to DTO
func ToUserDTO(userModel models.User) (dto.UserDTO, error) {
	var userDTO dto.UserDTO

	// parse or convert time to string
	dateFormat := "2006-01-02"
	datetimeFormat := "2006-01-02 15:04:05"

	userDTO.ID = int(userModel.ID)
	userDTO.Name = userModel.Name
	userDTO.Email = userModel.Email
	userDTO.Password = userModel.Password
	userDTO.Birthday = userModel.Birthday.Format(dateFormat)
	userDTO.CreatedAt = userModel.CreatedAt.Format(datetimeFormat)
	userDTO.UpdatedAt = userModel.UpdatedAt.Format(datetimeFormat)

	return userDTO, nil
}
