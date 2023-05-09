package repositories

import (
	"errors"
	"swim-class/helpers"
	"swim-class/models"

	"github.com/go-sql-driver/mysql"
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

var mysqlErr *mysql.MySQLError

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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, errors.New("record not found")
	}
	return user, err
}

func (ur *userRepository) CreateUser(userData models.User) (models.User, error) {
	//hashing password
	hashedPassword, err := helpers.HashPassword(userData.Password)
	if err != nil {
		return userData, err
	}
	userData.Password = hashedPassword

	err = ur.db.Create(&userData).Error
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return userData, errors.New("duplicate key found")
	}
	return userData, err
}

func (ur *userRepository) UpdateUser(userData models.User) (models.User, error) {
	err := ur.db.Save(&userData).Error
	return userData, err
}

func (ur *userRepository) DeleteUser(userData models.User) error {
	err := ur.db.First(&userData).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	}
	err = ur.db.Delete(&userData).Error
	return err
}

func (ur *userRepository) Login(user models.User) (models.User, error) {
	userPassword := user.Password
	err := ur.db.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return user, err
	}

	isMatch := helpers.MatchingPasswordHash(userPassword, user.Password)
	if !isMatch {
		return models.User{}, errors.New("login failed")
	}
	return user, err
}
