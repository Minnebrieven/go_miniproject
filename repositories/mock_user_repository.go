package repositories

import (
	"swim-class/models"

	mock "github.com/stretchr/testify/mock"
)

type MockUserRepository interface {
	GetAllUsers() ([]models.User, error)
}

type mockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{}
}

func (mur *mockUserRepository) GetAllUsers() ([]models.User, error) {
	ret := mur.Called()
	return ret.Get(0).([]models.User), ret.Error(1)
}

func (mur *mockUserRepository) GetUser(user models.User) (models.User, error) {
	ret := mur.Called(user)
	return ret.Get(0).(models.User), ret.Error(1)
}

func (mur *mockUserRepository) CreateUser(userData models.User) (models.User, error) {
	ret := mur.Called(userData)
	return ret.Get(0).(models.User), ret.Error(1)
}

func (mur *mockUserRepository) UpdateUser(userData models.User) (models.User, error) {
	ret := mur.Called(userData)
	return ret.Get(0).(models.User), ret.Error(1)
}

func (mur *mockUserRepository) DeleteUser(userData models.User) error {
	ret := mur.Called(userData)
	return ret.Error(0)
}

func (mur *mockUserRepository) Login(user models.User) (models.User, error) {
	ret := mur.Called(user)
	return ret.Get(0).(models.User), ret.Error(1)
}
