package services

import (
	"reflect"
	"swim-class/models"
	"swim-class/repositories"
)

type InstructorService interface {
	GetAllInstructorsService() ([]models.Instructor, error)
	GetInstructorService(instructorData models.Instructor) (models.Instructor, error)
	CreateInstructorService(instructorData models.Instructor) error
	EditInstructorService(instructorID int, modifiedInstructorData models.Instructor) (models.Instructor, error)
	DeleteInstructorService(instructorID int) error
}

type instructorService struct {
	instructorRepository repositories.InstructorRepository
}

func NewInstructorService(instructorRepo repositories.InstructorRepository) *instructorService {
	return &instructorService{instructorRepository: instructorRepo}
}

func (us *instructorService) GetAllInstructorsService() ([]models.Instructor, error) {
	instructor, err := us.instructorRepository.GetAllInstructors()
	if err != nil {
		return nil, err
	}
	return instructor, nil
}

func (us *instructorService) GetInstructorService(instructorData models.Instructor) (models.Instructor, error) {
	instructor, err := us.instructorRepository.GetInstructor(instructorData)
	return instructor, err
}

func (us *instructorService) CreateInstructorService(instructorData models.Instructor) error {
	//CreateInstructor will return nil if there's no error
	err := us.instructorRepository.CreateInstructor(instructorData)
	return err
}

func (us *instructorService) EditInstructorService(instructorID int, modifiedInstructorData models.Instructor) (models.Instructor, error) {
	//find record first if not exists return error
	instructor := models.Instructor{ID: uint(instructorID)}
	instructor, err := us.instructorRepository.GetInstructor(instructor)
	if err != nil {
		return instructor, err
	}

	//replace current data with new one
	var instructorPointer *models.Instructor = &instructor
	instructorVal := reflect.ValueOf(instructorPointer).Elem()
	instructorType := instructorVal.Type()

	editVal := reflect.ValueOf(modifiedInstructorData)

	for i := 0; i < instructorVal.NumField(); i++ {
		//skip ID field to be edited
		if instructorType.Field(i).Name == "ID" {
			continue
		}

		//edit every field in instructor with modifiedInstructorData
		instructorVal.Field(i).Set(editVal.Field(i))
	}

	result, err := us.instructorRepository.UpdateInstuctor(instructor)
	if err != nil {
		return result, err
	}
	return result, nil

}

func (us *instructorService) DeleteInstructorService(instructorID int) error {
	instructor := models.Instructor{ID: uint(instructorID)}
	err := us.instructorRepository.DeleteInstructor(instructor)
	return err
}
