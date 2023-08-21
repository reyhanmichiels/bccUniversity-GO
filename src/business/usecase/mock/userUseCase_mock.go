package mock

import (
	"bcc-university/src/business/entity"

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
