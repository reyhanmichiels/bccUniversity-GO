package repository

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/sdk/library"
	"net/http"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(inputUser entity.User) interface{}
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepository *userRepository) CreateUser(inputUser entity.User) interface{} {

	//create user
	user := userRepository.db.Create(&inputUser)
	if user.Error != nil {

		errorObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to create user",
			Err:     user.Error,
		}

		return errorObject

	}

	return nil

}
