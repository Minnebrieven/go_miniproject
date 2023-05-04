package services

import (
	"reflect"
	"swim-class/dto"
	"swim-class/mapper"
	"swim-class/middlewares"
	"swim-class/models"
	"swim-class/repositories"
)

type (
	UserService interface {
		GetAllUsersService() ([]dto.UserDTO, error)
		GetUserService(userDTO dto.UserDTO) (dto.UserDTO, error)
		CreateUserService(userDTO dto.UserDTO) (dto.UserDTO, error)
		EditUserService(userID int, modifiedUserData dto.UserDTO) (dto.UserDTO, error)
		DeleteUserService(userID int) error
		Login(userDTO dto.UserDTO) (dto.UserDTO, string, error)
	}

	userService struct {
		userRepository repositories.UserRepository
	}
)

func NewUserService(userRepo repositories.UserRepository) *userService {
	return &userService{userRepository: userRepo}
}

func (us *userService) GetAllUsersService() ([]dto.UserDTO, error) {
	user, err := us.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	userDTOList, err := mapper.ToUserDTOList(user)
	if err != nil {
		return nil, err
	}
	return userDTOList, nil
}

func (us *userService) GetUserService(userDTO dto.UserDTO) (dto.UserDTO, error) {
	userModel, err := mapper.ToUserModel(userDTO)
	if err != nil {
		return userDTO, err
	}

	result, err := us.userRepository.GetUser(userModel)
	if err != nil {
		return userDTO, err
	}

	userDTO, err = mapper.ToUserDTO(result)
	if err != nil {
		return userDTO, err
	}
	return userDTO, nil
}

func (us *userService) CreateUserService(userDTO dto.UserDTO) (dto.UserDTO, error) {
	//convert DTO to model
	userModel, err := mapper.ToUserModel(userDTO)
	if err != nil {
		return userDTO, err
	}

	userModel, err = us.userRepository.CreateUser(userModel)
	if err != nil {
		return userDTO, err
	}

	//
	userDTO, err = mapper.ToUserDTO(userModel)
	if err != nil {
		return userDTO, err
	}
	return userDTO, nil
}

func (us *userService) EditUserService(userID int, modifiedUserData dto.UserDTO) (dto.UserDTO, error) {
	//find record first if not exists return error
	userModel := models.User{ID: uint(userID)}
	userModel, err := us.userRepository.GetUser(userModel)
	if err != nil {
		return modifiedUserData, err
	}

	modifiedUserModel, err := mapper.ToUserModel(modifiedUserData)
	if err != nil {
		return modifiedUserData, err
	}

	//replace exist data with new one
	var userPointer *models.User = &userModel
	userVal := reflect.ValueOf(userPointer).Elem()
	userType := userVal.Type()

	editVal := reflect.ValueOf(modifiedUserModel)

	for i := 0; i < userVal.NumField(); i++ {
		//skip ID, CreatedAt, UpdatedAt field to be edited
		switch userType.Field(i).Name {
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
			userVal.Field(i).Set(editVal.Field(i))
		}
	}

	result, err := us.userRepository.UpdateUser(userModel)
	if err != nil {
		return modifiedUserData, err
	}

	modifiedUserData, err = mapper.ToUserDTO(result)
	return modifiedUserData, err

}

func (us *userService) DeleteUserService(userID int) error {
	user := models.User{ID: uint(userID)}
	err := us.userRepository.DeleteUser(user)
	return err
}

func (us *userService) Login(userDTO dto.UserDTO) (dto.UserDTO, string, error) {
	userData, err := mapper.ToUserModel(userDTO)
	if err != nil {
		return userDTO, "", err
	}

	user, err := us.userRepository.Login(userData)
	if err != nil {
		return userDTO, "", err
	}

	userDTO, err = mapper.ToUserDTO(user)
	if err != nil {
		return userDTO, "", err
	}

	token, err := middlewares.CreateToken(int(user.ID), user.Email)
	if err != nil {
		return userDTO, "", err
	}
	return userDTO, token, nil
}
