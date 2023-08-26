package usecase

import (
	"bcc-university/src/business/repository"

	"github.com/stretchr/testify/mock"
)

var userRepositoryMock = repository.UserRepositoryMock{
	Mock: mock.Mock{},
}

var classRepositoryMock = repository.ClassRepositoryMock{
	Mock: mock.Mock{},
}

var userUseCase = NewUserUseCase(&userRepositoryMock, &classRepositoryMock)