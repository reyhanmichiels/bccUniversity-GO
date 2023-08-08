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
	UpdateUser(updatedUser *entity.User, updateData interface{}) error
	AddUserToClass(user *entity.User, class *entity.Class)
	DropUserFromClass(user *entity.User, class *entity.Class)
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

func (userRepository *userRepository) UpdateUser(updatedUser *entity.User, updateData interface{}) error {

	err := userRepository.db.Model(updatedUser).Updates(updateData).Error

	if err != nil {
		return err
	}

	return nil

}

func (userRepository *userRepository) AddUserToClass(user *entity.User, class *entity.Class) {

	userRepository.db.Model(user).Association("Classes").Append(class)

	class.Participant += 1
	userRepository.db.Save(class)

}

func (userRepository *userRepository) DropUserFromClass(user *entity.User, class *entity.Class) {

	userRepository.db.Model(user).Association("Classes").Delete(class)

	class.Participant -= 1
	userRepository.db.Save(class)

}
