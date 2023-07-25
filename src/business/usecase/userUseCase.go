package usecase

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/business/repository"
	"bcc-university/src/sdk/library"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	CreateUser(inputUser entity.CreateUser) (entity.UserRegistResponse, interface{})
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (userUseCase *userUseCase) CreateUser(inputUser entity.CreateUser) (entity.UserRegistResponse, interface{}) {

	//hash pasword
	password, err := bcrypt.GenerateFromPassword([]byte(inputUser.Password), 10)
	if err != nil {

		errorObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate hash password",
			Err:     err,
		}

		return entity.UserRegistResponse{}, errorObject

	}

	//create user
	createdUser := entity.User{
		Name:     inputUser.Name,
		Username: inputUser.Username,
		Email:    inputUser.Email,
		Password: string(password),
	}

	errorObject := userUseCase.userRepository.CreateUser(createdUser)
	if errorObject != nil {
		return entity.UserRegistResponse{}, errorObject
	}

	userResponse := entity.UserRegistResponse{
		Name:     createdUser.Name,
		Email:    createdUser.Email,
		Username: createdUser.Username,
	}

	return userResponse, nil

}
