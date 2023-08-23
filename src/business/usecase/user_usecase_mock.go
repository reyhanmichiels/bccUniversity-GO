package usecase

import (
	"bcc-university/src/business/entity"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	Mock mock.Mock
}

func (userUseCaseMock *UserUseCaseMock) RegistrationUseCase(userInput entity.RegistBind) (entity.RegistApi, interface{}) {

	args := userUseCaseMock.Mock.Called(userInput)
	if args[1] != nil {

		return entity.RegistApi{}, args[1]

	}

	return args[0].(entity.RegistApi), nil

}

func (userUseCaseMock *UserUseCaseMock) LoginUseCase(userInput entity.LoginBind, c *gin.Context) interface{} {

	args := userUseCaseMock.Mock.Called(userInput)
	return args[0]

}
func (userUseCaseMock *UserUseCaseMock) EditAccountUseCase(userInput entity.EditAccountBind, loginUser entity.User) (entity.UserApi, interface{}) {

	args := userUseCaseMock.Mock.Called(userInput, loginUser)
	if args[1] != nil {

		return entity.UserApi{}, args[1]

	}

	return args[0].(entity.UserApi), nil
}
func (userUseCaseMock *UserUseCaseMock) AddUserToClassUseCase(loginUser entity.User, classCode string) interface{} {

	args := userUseCaseMock.Mock.Called(loginUser, classCode)
	return args[0]

}
func (userUseCaseMock *UserUseCaseMock) DropClassUseCase(loginUser entity.User, classId uint) interface{} {

	args := userUseCaseMock.Mock.Called(loginUser, classId)
	return args[0]

}
