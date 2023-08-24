package usecase

import (
	"bcc-university/src/business/entity"

	"github.com/stretchr/testify/mock"
)

type ClassUseCaseMock struct {
	Mock mock.Mock
}

func GetAllClassUseCase() ([]entity.ClassApi, interface{}) {

	return []entity.ClassApi{}, nil

}

func RemoveUserFromClassUseCase(loginUser entity.User, classId uint, userId uint) interface{} {

	return nil

}

func AdmAddUserToClassUseCase(loginUser entity.User, classId uint, userId uint) interface{} {

	return nil

}

func CreateClassUseCase(userInput entity.CreateUpdateClassBind, loginUser entity.User) (entity.CreateUpdateClassApi, interface{}) {

	return entity.CreateUpdateClassApi{}, nil

}

func EditClassUseCase(userInput entity.CreateUpdateClassBind, loginUser entity.User, classId uint) (entity.CreateUpdateClassApi, interface{}) {

	return entity.CreateUpdateClassApi{}, nil

}

func DeleteClassUseCase(loginUser entity.User, classId uint) interface{} {

	return nil

}

func GetClassParticipantUseCase(loginUser entity.User, classId uint) (entity.ClassParticipantApi, interface{}) {

	return entity.ClassParticipantApi{}, nil

}
