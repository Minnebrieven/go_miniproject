package services

import (
	"reflect"
	"swim-class/dto"
	"swim-class/mapper"
	"swim-class/models"
	"swim-class/repositories"
)

type (
	ClassService interface {
		GetAllClassesService() ([]dto.ClassDTO, error)
		GetClassService(ClassDTO dto.ClassDTO) (dto.ClassDTO, error)
		CreateClassService(ClassDTO dto.ClassDTO) (dto.ClassDTO, error)
		EditClassService(classID int, modifiedClassData dto.ClassDTO) (dto.ClassDTO, error)
		DeleteClassService(classID int) error
	}

	classService struct {
		classRepository         repositories.ClassRepository
		classCategoryRepository repositories.ClassCategoryRepository
		instructorRepository    repositories.InstructorRepository
	}
)

func NewClassService(classRepo repositories.ClassRepository, classCategoryRepo repositories.ClassCategoryRepository, instructorRepo repositories.InstructorRepository) *classService {
	return &classService{classRepository: classRepo, classCategoryRepository: classCategoryRepo, instructorRepository: instructorRepo}
}

func (cs *classService) GetAllClassesService() ([]dto.ClassDTO, error) {
	class, err := cs.classRepository.GetAllClasses()
	if err != nil {
		return nil, err
	}

	ClassDTOList, err := mapper.ToClassDTOList(class)
	if err != nil {
		return nil, err
	}
	return ClassDTOList, nil
}


func (cs *classService) GetClassService(classDTO dto.ClassDTO) (dto.ClassDTO, error) {
	classModel, err := mapper.ToClassModel(classDTO)
	if err != nil {
		return classDTO, err
	}

	result, err := cs.classRepository.GetClass(classModel)
	if err != nil {
		return classDTO, err
	}

	classDTO, err = mapper.ToClassDTO(result)
	if err != nil {
		return classDTO, err
	}
	return classDTO, nil
}

func (cs *classService) CreateClassService(classDTO dto.ClassDTO) (dto.ClassDTO, error) {
	//convert DTO to model
	classModel, err := mapper.ToClassModel(classDTO)
	if err != nil {
		return classDTO, err
	}

	classModel.ClassCategory, err = cs.classCategoryRepository.GetClassCategory(models.ClassCategory{ID: classModel.ClassCategoryID})
	if err != nil {
		return classDTO, err
	}

	classModel.Instructor, err = cs.instructorRepository.GetInstructor(models.Instructor{ID: classModel.InstructorID})
	if err != nil {
		return classDTO, err
	}

	classModel, err = cs.classRepository.CreateClass(classModel)
	if err != nil {
		return classDTO, err
	}

	//
	classDTO, err = mapper.ToClassDTO(classModel)
	if err != nil {
		return classDTO, err
	}
	return classDTO, nil
}

func (cs *classService) EditClassService(classID int, modifiedClassData dto.ClassDTO) (dto.ClassDTO, error) {
	//find record first if not exists return error
	classModel := models.Class{ID: uint(classID)}
	classModel, err := cs.classRepository.GetClass(classModel)
	if err != nil {
		return modifiedClassData, err
	}

	if modifiedClassData.ClassCategoryID != int(classModel.ClassCategoryID) {
		classModel.ClassCategory, err = cs.classCategoryRepository.GetClassCategory(models.ClassCategory{ID: uint(modifiedClassData.ClassCategoryID)})
		if err != nil {
			return modifiedClassData, err
		}
	}
	if modifiedClassData.InstructorID != int(classModel.InstructorID) {
		classModel.Instructor, err = cs.instructorRepository.GetInstructor(models.Instructor{ID: uint(modifiedClassData.InstructorID)})
		if err != nil {
			return modifiedClassData, err
		}
	}

	modifiedClassModel, err := mapper.ToClassModel(modifiedClassData)
	if err != nil {
		return modifiedClassData, err
	}

	//replace exist data with new one
	var classPointer *models.Class = &classModel
	classVal := reflect.ValueOf(classPointer).Elem()
	classType := classVal.Type()

	editVal := reflect.ValueOf(modifiedClassModel)

	for i := 0; i < classVal.NumField(); i++ {
		//skip ID, CreatedAt, UpdatedAt field to be edited
		switch classType.Field(i).Name {
		case "ID":
			continue
		case "ClassCategory":
			continue
		case "Instructor":
			continue
		case "CreatedAt":
			continue
		case "UpdatedAt":
			continue
		}

		editField := editVal.Field(i)
		isSet := editField.IsValid() && !editField.IsZero()
		if isSet {
			classVal.Field(i).Set(editVal.Field(i))
		}
	}

	result, err := cs.classRepository.UpdateClass(classModel)
	if err != nil {
		return modifiedClassData, err
	}

	modifiedClassData, err = mapper.ToClassDTO(result)
	return modifiedClassData, err

}

func (cs *classService) DeleteClassService(classID int) error {
	class := models.Class{ID: uint(classID)}
	err := cs.classRepository.DeleteClass(class)
	return err
}
