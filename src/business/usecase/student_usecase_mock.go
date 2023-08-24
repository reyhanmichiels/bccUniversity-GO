package usecase

import (
	"bcc-university/src/business/entity"

	"github.com/stretchr/testify/mock"
)

type StudentUseCaseMock struct {
	Mock mock.Mock
}

func (studentUseCaseMock *StudentUseCaseMock) ClaimStudentNumberUseCase(loginUser entity.User) (entity.ClaimStudentNumberApi, interface{}) {

	args := studentUseCaseMock.Mock.Called(loginUser)
	if args[1] != nil {

		return entity.ClaimStudentNumberApi{}, args[1]

	}

	return args[0].(entity.ClaimStudentNumberApi), nil

}
