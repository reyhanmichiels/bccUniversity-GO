package repository

import (
	"bcc-university/src/business/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (userRepositoryMock *UserRepositoryMock) CreateUser(user *entity.User) error {

	return nil

}

func (userRepositoryMock *UserRepositoryMock) FindUserByEmail(email string) (entity.User, error) {

	return entity.User{}, nil

}
func (userRepositoryMock *UserRepositoryMock) FindUserByCondition(user interface{}, condition string, value interface{}) error {

	return nil

}

func (userRepositoryMock *UserRepositoryMock) ELFindUserByCondition(user interface{}, condition string, value interface{}) error {

	return nil

}

func (userRepositoryMock *UserRepositoryMock) UpdateUser(user *entity.User, updateData interface{}) error {

	return nil

}

func (userRepositoryMock *UserRepositoryMock) AddUserToClass(user *entity.User, class *entity.Class) {

	

}

func (userRepositoryMock *UserRepositoryMock) DropUserFromClass(user *entity.User, class *entity.Class) {

	

}
