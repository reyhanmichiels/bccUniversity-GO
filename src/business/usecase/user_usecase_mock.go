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
	return nil
}
func (userUseCaseMock *UserUseCaseMock) EditAccountUseCase(userInput entity.EditAccountBind, loginUser entity.User) (entity.UserApi, interface{}) {
	return entity.UserApi{}, nil
}
func (userUseCaseMock *UserUseCaseMock) AddUserToClassUseCase(loginUser entity.User, classCode string) interface{} {
	return nil
}
func (userUseCaseMock *UserUseCaseMock) DropClassUseCase(loginUser entity.User, classId uint) interface{} {
	return nil
}