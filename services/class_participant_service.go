package services

import (
	"reflect"
	"swim-class/dto"
	"swim-class/mapper"
	"swim-class/models"
	"swim-class/repositories"
)

type (
	ClassParticipantService interface {
		GetAllClassParticipantsService() ([]dto.ClassParticipantDTO, error)
		GetAllClassParticipantsByUserIDService(userID int) ([]dto.ClassParticipantDTO, error)
		GetClassParticipantService(ClassParticipantDTO dto.ClassParticipantDTO) (dto.ClassParticipantDTO, error)
		CreateClassParticipantService(ClassParticipantDTO dto.ClassParticipantDTO) (dto.ClassParticipantDTO, error)
		EditClassParticipantService(classParticipantID int, modifiedClassData dto.ClassParticipantDTO) (dto.ClassParticipantDTO, error)
		DeleteClassParticipantService(classParticipantID int) error
	}

	classParticipantService struct {
		classParticipantRepository repositories.ClassParticipantRepository
		classRepository            repositories.ClassRepository
		userRepository             repositories.UserRepository
	}
)

func NewClassParticipantService(classParticipantRepo repositories.ClassParticipantRepository, classRepo repositories.ClassRepository, userRepo repositories.UserRepository) *classParticipantService {
	return &classParticipantService{classParticipantRepository: classParticipantRepo, classRepository: classRepo, userRepository: userRepo}
}

func (cps *classParticipantService) GetAllClassParticipantsService() ([]dto.ClassParticipantDTO, error) {
	classParticipants, err := cps.classParticipantRepository.GetAllClassParticipants()
	if err != nil {
		return nil, err
	}

	ClassParticipantDTOList, err := mapper.ToClassParticipantDTOList(classParticipants)
	if err != nil {
		return nil, err
	}
	return ClassParticipantDTOList, nil
}

func (cps *classParticipantService) GetAllClassParticipantsByUserIDService(userID int) ([]dto.ClassParticipantDTO, error) {
	classParticipants, err := cps.classParticipantRepository.GetAllClassParticipantsByUserID(userID)
	if err != nil {
		return nil, err
	}

	ClassParticipantDTOList, err := mapper.ToClassParticipantDTOList(classParticipants)
	if err != nil {
		return nil, err
	}
	return ClassParticipantDTOList, nil
}

func (cps *classParticipantService) GetClassParticipantService(classParticipantDTO dto.ClassParticipantDTO) (dto.ClassParticipantDTO, error) {
	classParticipantModel, err := mapper.ToClassParticipantModel(classParticipantDTO)
	if err != nil {
		return classParticipantDTO, err
	}

	result, err := cps.classParticipantRepository.GetClassParticipant(classParticipantModel)
	if err != nil {
		return classParticipantDTO, err
	}

	classParticipantDTO, err = mapper.ToClassParticipantDTO(result)
	if err != nil {
		return classParticipantDTO, err
	}
	return classParticipantDTO, nil
}

func (cps *classParticipantService) CreateClassParticipantService(classParticipantDTO dto.ClassParticipantDTO) (dto.ClassParticipantDTO, error) {
	//convert DTO to model
	classParticipantModel, err := mapper.ToClassParticipantModel(classParticipantDTO)
	if err != nil {
		return classParticipantDTO, err
	}

	classParticipantModel.Class, err = cps.classRepository.GetClass(models.Class{ID: classParticipantModel.ClassID})
	if err != nil {
		return classParticipantDTO, err
	}

	classParticipantModel.User, err = cps.userRepository.GetUser(models.User{ID: classParticipantModel.UserID})
	if err != nil {
		return classParticipantDTO, err
	}

	classParticipantModel, err = cps.classParticipantRepository.CreateClassParticipant(classParticipantModel)
	if err != nil {
		return classParticipantDTO, err
	}

	classParticipantDTO, err = mapper.ToClassParticipantDTO(classParticipantModel)
	if err != nil {
		return classParticipantDTO, err
	}
	return classParticipantDTO, nil
}

func (cps *classParticipantService) EditClassParticipantService(classParticipantID int, modifiedClassData dto.ClassParticipantDTO) (dto.ClassParticipantDTO, error) {
	//find record first if not exists return error
	classParticipantModel := models.ClassParticipant{ID: uint(classParticipantID)}
	classParticipantModel, err := cps.classParticipantRepository.GetClassParticipant(classParticipantModel)
	if err != nil {
		return modifiedClassData, err
	}

	if modifiedClassData.ClassID != int(classParticipantModel.ClassID) {
		classParticipantModel.Class, err = cps.classRepository.GetClass(models.Class{ID: uint(modifiedClassData.ClassID)})
		if err != nil {
			return modifiedClassData, err
		}
	}
	if modifiedClassData.UserID != int(classParticipantModel.UserID) {
		classParticipantModel.User, err = cps.userRepository.GetUser(models.User{ID: uint(modifiedClassData.UserID)})
		if err != nil {
			return modifiedClassData, err
		}
	}

	modifiedClassParticipantModel, err := mapper.ToClassParticipantModel(modifiedClassData)
	if err != nil {
		return modifiedClassData, err
	}

	//replace exist data with new one
	var classParticipantPointer *models.ClassParticipant = &classParticipantModel
	classParticipantVal := reflect.ValueOf(classParticipantPointer).Elem()
	classParticipantType := classParticipantVal.Type()

	editVal := reflect.ValueOf(modifiedClassParticipantModel)

	for i := 0; i < classParticipantVal.NumField(); i++ {
		//skip ID, CreatedAt, UpdatedAt field to be edited
		switch classParticipantType.Field(i).Name {
		case "ID":
			continue
		case "Class":
			continue
		case "User":
			continue
		case "CreatedAt":
			continue
		case "UpdatedAt":
			continue
		}

		editField := editVal.Field(i)
		isSet := editField.IsValid() && !editField.IsZero()
		if isSet {
			classParticipantVal.Field(i).Set(editVal.Field(i))
		}
	}

	result, err := cps.classParticipantRepository.UpdateClassParticipant(classParticipantModel)
	if err != nil {
		return modifiedClassData, err
	}

	modifiedClassData, err = mapper.ToClassParticipantDTO(result)
	return modifiedClassData, err

}

func (cps *classParticipantService) DeleteClassParticipantService(classParticipantID int) error {
	class := models.ClassParticipant{ID: uint(classParticipantID)}
	err := cps.classParticipantRepository.DeleteClassParticipant(class)
	return err
}
