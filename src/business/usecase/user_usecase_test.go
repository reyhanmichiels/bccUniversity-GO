package usecase

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/business/repository"
	"bcc-university/src/sdk/library"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepositoryMock = repository.UserRepositoryMock{
	Mock: mock.Mock{},
}

var classRepositoryMock = repository.ClassRepositoryMock{
	Mock: mock.Mock{},
}

var userUseCase = NewUserUseCase(&userRepositoryMock, &classRepositoryMock)

func TestRegistrationUseCasePath1(t *testing.T) {

	input := []entity.RegistBind{
		{
			Password: strings.Repeat("t", 73),
		},
		{
			Password: strings.Repeat("t", 73),
		},
		{
			Password: strings.Repeat("t", 73),
		},
		{
			Password: strings.Repeat("t", 73),
		},
		{
			Password: strings.Repeat("t", 73),
		},
	}

	for i, v := range input {

		t.Run(fmt.Sprintf("path 1 registration usecase testing %d", i), func(t *testing.T) {

			_, resultErr := userUseCase.RegistrationUseCase(v)

			resultErrObject := resultErr.(library.ErrorObject)
			assert.Equal(t, http.StatusInternalServerError, resultErrObject.Code, "status code should be equal")
			assert.Equal(t, "failed to generate hash password", resultErrObject.Message, "message should be equal")

		})

	}

}

func TestRegistrationUseCasePath2(t *testing.T) {

	input := []entity.RegistBind{
		{
			Password: "test",
			Name:     "test",
			Username: "test",
			Email:    "test@test.com",
		},
		{
			Password: "test",
			Name:     "test",
			Username: "test",
			Email:    "test@test.com",
		},
		{
			Password: "test",
			Name:     "test",
			Username: "test",
			Email:    "test@test.com",
		},
		{
			Password: "test",
			Name:     "test",
			Username: "test",
			Email:    "test@test.com",
		},
		{
			Password: "test",
			Name:     "test",
			Username: "test",
			Email:    "test@test.com",
		},
	}

	for i, v := range input {

		t.Run(fmt.Sprintf("path 2 registration usecase testing %d", i), func(t *testing.T) {

			functionCall := userRepositoryMock.Mock.On("CreateUser").Return(errors.New("test"))

			_, resultErr := userUseCase.RegistrationUseCase(v)

			resultErrObject := resultErr.(library.ErrorObject)
			assert.Equal(t, http.StatusInternalServerError, resultErrObject.Code, "status code should be equal")
			assert.Equal(t, "test", resultErrObject.Err.Error(), "error code should be equal")
			assert.Equal(t, "failed to create user", resultErrObject.Message, "message code should be equal")

			functionCall.Unset()

		})

	}

}

func TestRegistrationUseCasePath3(t *testing.T) {

	input := []entity.RegistBind{
		{
			Password: "test",
			Name:     "test",
			Username: "test",
			Email:    "test@test.com",
		},
		{
			Password: "test",
			Name:     "test",
			Username: "test",
			Email:    "test@test.com",
		},
		{
			Password: "test",
			Name:     "test",
			Username: "test",
			Email:    "test@test.com",
		},
		{
			Password: "test",
			Name:     "test",
			Username: "test",
			Email:    "test@test.com",
		},
		{
			Password: "test",
			Name:     "test",
			Username: "test",
			Email:    "test@test.com",
		},
	}

	for i, v := range input {

		t.Run(fmt.Sprintf("path 3 registration usecase testing %d", i), func(t *testing.T) {

			functionCall := userRepositoryMock.Mock.On("CreateUser").Return(nil)

			resultData, resultErr := userUseCase.RegistrationUseCase(v)
			assert.Nil(t, resultErr, "errod should be nil")
			assert.Equal(t, "test", resultData.Name, "name should be equal")
			assert.Equal(t, "test", resultData.Username, "username should be equal")
			assert.Equal(t, "test@test.com", resultData.Email, "email should be equal")

			functionCall.Unset()

		})

	}

}
