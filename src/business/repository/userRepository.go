package repository

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/sdk/library"
	"net/http"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(inputUser entity.User) interface{}
	FindUserByEmail(email string) (entity.User, error)
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

func (userRepository *userRepository) FindUserByEmail(email string) (entity.User, error) {

	var user entity.User
	err := userRepository.db.First(&user, "email = ?", email)
	return user, err.Error

}
