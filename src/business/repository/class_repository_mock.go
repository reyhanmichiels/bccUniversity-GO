package repository

import (
	"bcc-university/src/business/entity"

	"github.com/stretchr/testify/mock"
)

type ClassRepositoryMock struct {
	Mock mock.Mock
}

func (classRepositoryMock *ClassRepositoryMock) FindAllClass(allClass interface{}) error {

	return nil

}

func (classRepositoryMock *ClassRepositoryMock) ELFindClassByCondition(class interface{}, condition string, value interface{}) error {

	return nil

}

func (classRepositoryMock *ClassRepositoryMock) FindClassByCondition(class interface{}, condition string, value interface{}) error {

	return nil

}

func (classRepositoryMock *ClassRepositoryMock) CreateClass(class *entity.Class) error {

	return nil

}

func (classRepositoryMock *ClassRepositoryMock) Updateclass(class *entity.Class, updateData interface{}) error {

	return nil

}

func (classRepositoryMock *ClassRepositoryMock) DeleteClass(class *entity.Class) error {

	return nil

}

func (classRepositoryMock *ClassRepositoryMock) FindClassParticipant(class *entity.Class, users interface{}) {

}
