package services

import (
	"reflect"
	"swim-class/dto"
	"swim-class/mapper"
	"swim-class/models"
	"swim-class/repositories"
)

type InstructorService interface {
	GetAllInstructorsService() ([]dto.InstructorDTO, error)
	GetInstructorService(instructorData dto.InstructorDTO) (dto.InstructorDTO, error)
	CreateInstructorService(instructorData dto.InstructorDTO) (dto.InstructorDTO, error)
	EditInstructorService(instructorID int, modifiedInstructorData dto.InstructorDTO) (dto.InstructorDTO, error)
	DeleteInstructorService(instructorID int) error
}

type instructorService struct {
	instructorRepository repositories.InstructorRepository
}

func NewInstructorService(instructorRepo repositories.InstructorRepository) *instructorService {
	return &instructorService{instructorRepository: instructorRepo}
}

func (us *instructorService) GetAllInstructorsService() ([]dto.InstructorDTO, error) {
	instructor, err := us.instructorRepository.GetAllInstructors()
	if err != nil {
		return nil, err
	}

	instructorDTO, err := mapper.ToInstructorDTOList(instructor)
	if err != nil {
		return nil, err
	}
	return instructorDTO, nil
}

func (us *instructorService) GetInstructorService(instructorData dto.InstructorDTO) (dto.InstructorDTO, error) {
	instructorModel, err := mapper.ToInstructorModel(instructorData)
	if err != nil {
		return instructorData, err
	}

	instructorModel, err = us.instructorRepository.GetInstructor(instructorModel)
	if err != nil {
		return instructorData, err
	}

	instructorData, err = mapper.ToInstructorDTO(instructorModel)
	return instructorData, err
}

func (us *instructorService) CreateInstructorService(instructorData dto.InstructorDTO) (dto.InstructorDTO, error) {
	instructorModel, err := mapper.ToInstructorModel(instructorData)
	if err != nil {
		return instructorData, err
	}

	instructorModel, err = us.instructorRepository.CreateInstructor(instructorModel)
	if err != nil {
		return instructorData, err
	}

	instructorData, err = mapper.ToInstructorDTO(instructorModel)
	return instructorData, err
}

func (us *instructorService) EditInstructorService(instructorID int, modifiedInstructorData dto.InstructorDTO) (dto.InstructorDTO, error) {
	//find record first if not exists return error
	instructorModel := models.Instructor{ID: uint(instructorID)}
	instructorModel, err := us.instructorRepository.GetInstructor(instructorModel)
	if err != nil {
		return modifiedInstructorData, err
	}

	modifiedInstructorModel, err := mapper.ToInstructorModel(modifiedInstructorData)
	if err != nil {
		return modifiedInstructorData, err
	}

	//replace exist data with new one
	var instructorPointer *models.Instructor = &instructorModel
	instructorVal := reflect.ValueOf(instructorPointer).Elem()
	instructorType := instructorVal.Type()

	editVal := reflect.ValueOf(modifiedInstructorModel)

	for i := 0; i < instructorVal.NumField(); i++ {
		//skip ID, CreatedAt, UpdatedAt field to be edited
		switch instructorType.Field(i).Name {
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
			instructorVal.Field(i).Set(editVal.Field(i))
		}
	}

	result, err := us.instructorRepository.UpdateInstuctor(instructorModel)
	if err != nil {
		return modifiedInstructorData, err
	}

	modifiedInstructorData, err = mapper.ToInstructorDTO(result)
	return modifiedInstructorData, err

}

func (us *instructorService) DeleteInstructorService(instructorID int) error {
	instructor := models.Instructor{ID: uint(instructorID)}
	err := us.instructorRepository.DeleteInstructor(instructor)
	return err
}
