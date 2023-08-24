package usecase

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/sdk/library"

	"github.com/stretchr/testify/mock"
)

type ClassUseCaseMock struct {
	Mock mock.Mock
}

func (classUseCaseMock *ClassUseCaseMock) GetAllClassUseCase() ([]entity.ClassApi, interface{}) {

	args := classUseCaseMock.Mock.Called()

	if args[1] != nil {

		return nil, args[1].(library.ErrorObject)

	}

	return args[0].([]entity.ClassApi), nil

}

func (classUseCaseMock *ClassUseCaseMock) RemoveUserFromClassUseCase(loginUser entity.User, classId uint, userId uint) interface{} {

	return nil

}

func (classUseCaseMock *ClassUseCaseMock) AdmAddUserToClassUseCase(loginUser entity.User, classId uint, userId uint) interface{} {

	return nil

}

func (classUseCaseMock *ClassUseCaseMock) CreateClassUseCase(userInput entity.CreateUpdateClassBind, loginUser entity.User) (entity.CreateUpdateClassApi, interface{}) {

	return entity.CreateUpdateClassApi{}, nil

}

func (classUseCaseMock *ClassUseCaseMock) EditClassUseCase(userInput entity.CreateUpdateClassBind, loginUser entity.User, classId uint) (entity.CreateUpdateClassApi, interface{}) {

	return entity.CreateUpdateClassApi{}, nil

}

func (classUseCaseMock *ClassUseCaseMock) DeleteClassUseCase(loginUser entity.User, classId uint) interface{} {

	return nil

}

func (classUseCaseMock *ClassUseCaseMock) GetClassParticipantUseCase(loginUser entity.User, classId uint) (entity.ClassParticipantApi, interface{}) {

	return entity.ClassParticipantApi{}, nil

}
