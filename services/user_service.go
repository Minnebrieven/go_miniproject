package services

import (
	"reflect"
	"swim-class/middlewares"
	"swim-class/models"
	"swim-class/repositories"
)

type (
	UserService interface {
		GetAllUsersService() ([]models.User, error)
		GetUserService(userData models.User) (models.User, error)
		CreateUserService(dataUser models.User) error
		EditUserService(userID int, modifiedUserData models.User) (models.User, error)
		DeleteUserService(userID int) error
		Login(userData models.User) (string, error)
	}

	userService struct {
		userRepository repositories.UserRepository
	}
)

func NewUserService(userRepo repositories.UserRepository) *userService {
	return &userService{userRepository: userRepo}
}

func (us *userService) GetAllUsersService() ([]models.User, error) {
	user, err := us.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) GetUserService(userData models.User) (models.User, error) {
	user, err := us.userRepository.GetUser(userData)
	return user, err
}

func (us *userService) CreateUserService(dataUser models.User) error {
	//CreateUser will return nil if there's no error
	err := us.userRepository.CreateUser(dataUser)
	return err
}

func (us *userService) EditUserService(userID int, modifiedUserData models.User) (models.User, error) {
	//find record first if not exists return error
	user := models.User{ID: uint(userID)}
	user, err := us.userRepository.GetUser(user)
	if err != nil {
		return user, err
	}

	//replace exist data with new one
	var userPointer *models.User = &user
	userVal := reflect.ValueOf(userPointer)
	userType := userVal.Type()

	editVal := reflect.ValueOf(modifiedUserData)

	for i := 0; i < userVal.NumField(); i++ {
		//skip ID field to be edited
		if userType.Field(i).Name == "ID" {
			continue
		}

		//edit every field in user with modifiedUserData
		userVal.Field(i).Set(editVal.Field(i))
	}

	result, err := us.userRepository.UpdateUser(user)
	if err != nil {
		return result, err
	}
	return result, nil

}

func (us *userService) DeleteUserService(userID int) error {
	user := models.User{ID: uint(userID)}
	err := us.userRepository.DeleteUser(user)
	return err
}

func (us *userService) Login(userData models.User) (string, error) {
	user, err := us.userRepository.Login(userData)
	if err != nil {
		return "", err
	}

	token, err := middlewares.CreateToken(int(user.ID), user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}
