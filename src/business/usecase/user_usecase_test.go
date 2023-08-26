package usecase

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/business/repository"
	"bcc-university/src/sdk/library"
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
