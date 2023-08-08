package repository

import (
	"bcc-university/src/business/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	FindUserByEmail(email string) (entity.User, error)
	FindUserByCondition(user interface{}, condition string, value interface{}) error
	UpdateUser(user *entity.User, updateData interface{}) error
	AddUserToClass(user *entity.User, class *entity.Class)
	DropUserFromClass(user *entity.User, class *entity.Class)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepository *userRepository) CreateUser(user *entity.User) error {

	//create user
	err := userRepository.db.Create(user).Error
	if err != nil {

		return err

	}

	return nil

}

func (userRepository *userRepository) FindUserByCondition(user interface{}, condition string, value interface{}) error {

	err := userRepository.db.Model(&entity.User{}).First(user, condition, value).Error
	if err != nil {

		return err

	}

	return nil

}

func (userRepository *userRepository) FindUserByEmail(email string) (entity.User, error) {

	var user entity.User
	err := userRepository.db.First(&user, "email = ?", email)
	return user, err.Error

}

func (userRepository *userRepository) UpdateUser(user *entity.User, updateData interface{}) error {

	err := userRepository.db.Model(user).Updates(updateData).Error

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
