package usecase

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/business/repository"
	"bcc-university/src/sdk/library"
	"net/http"
)

type ClassUseCase interface {
	GetAllClassUseCase() ([]entity.ClassApi, interface{})
}

type classUseCase struct {
	classRepository repository.ClassRepository
}

func NewClassUseCase(classRepository repository.ClassRepository) ClassUseCase {

	return &classUseCase{
		classRepository: classRepository,
	}

}

func (classUseCase *classUseCase) GetAllClassUseCase() ([]entity.ClassApi, interface{}) {

	var allClass []entity.ClassApi

	err := classUseCase.classRepository.FindAllClass(&allClass)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to get all class",
			Err:     err,
		}
		return nil, errObject

	}

	return allClass, nil

}
