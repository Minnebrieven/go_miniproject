package repositories

import (
	"errors"
	"swim-class/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUser(models.User) (models.User, error)
	CreateUser(userData models.User) (models.User, error)
	UpdateUser(models.User) (models.User, error)
	DeleteUser(models.User) error
	Login(models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetAllUsers() ([]models.User, error) {
	users := []models.User{}
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) GetUser(user models.User) (models.User, error) {
	err := ur.db.First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound){
		return user, errors.New("record not found")
	}
	return user, err
}

func (ur *userRepository) CreateUser(userData models.User) (models.User, error) {
	err := ur.db.Create(&userData).Error
	return userData, err
}

func (ur *userRepository) UpdateUser(userData models.User) (models.User, error) {
	err := ur.db.Save(&userData).Error
	return userData, err
}

func (ur *userRepository) DeleteUser(userData models.User) error {
	err := ur.db.Delete(&userData).Error
	return err
}

func (ur *userRepository) Login(user models.User) (models.User, error) {
	err := ur.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	return user, err
}
