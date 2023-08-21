package repository

import (
	"bcc-university/src/business/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *entity.User) error
	FindUserByEmail(email string) (entity.User, error)
	FindUserByCondition(user interface{}, condition string, value interface{}) error
	ELFindUserByCondition(user interface{}, condition string, value interface{}) error
	UpdateUser(user *entity.User, updateData interface{}) error
	AddUserToClass(user *entity.User, class *entity.Class)
	DropUserFromClass(user *entity.User, class *entity.Class)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (userRepository *UserRepository) CreateUser(user *entity.User) error {

	//create user
	err := userRepository.db.Create(user).Error
	if err != nil {

		return err

	}

	return nil

}

func (userRepository *UserRepository) FindUserByCondition(user interface{}, condition string, value interface{}) error {

	err := userRepository.db.Model(&entity.User{}).First(user, condition, value).Error
	if err != nil {

		return err

	}

	return nil

}

func (userRepository *UserRepository) ELFindUserByCondition(user interface{}, condition string, value interface{}) error {

	err := userRepository.db.Model(&entity.User{}).Preload("Student").Preload("Classes").First(user, condition, value).Error
	if err != nil {

		return err

	}

	return nil

}

func (userRepository *UserRepository) FindUserByEmail(email string) (entity.User, error) {

	var user entity.User
	err := userRepository.db.First(&user, "email = ?", email)
	return user, err.Error

}

func (userRepository *UserRepository) UpdateUser(user *entity.User, updateData interface{}) error {

	err := userRepository.db.Model(user).Updates(updateData).Error

	if err != nil {

		return err

	}

	return nil

}

func (userRepository *UserRepository) AddUserToClass(user *entity.User, class *entity.Class) {

	userRepository.db.Model(user).Association("Classes").Append(class)

	class.Participant += 1
	userRepository.db.Save(class)

}

func (userRepository *UserRepository) DropUserFromClass(user *entity.User, class *entity.Class) {

	userRepository.db.Model(user).Association("Classes").Delete(class)

	class.Participant -= 1
	userRepository.db.Save(class)

}
