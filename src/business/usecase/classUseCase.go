package usecase

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/business/repository"
	"bcc-university/src/sdk/library"
	"errors"
	"net/http"
)

type ClassUseCase interface {
	GetAllClassUseCase() ([]entity.ClassApi, interface{})
	RemoveUserFromClassUseCase(loginUser entity.User, classId uint, userId uint) interface{}
	AdmAddUserToClassUseCase(loginUser entity.User, classId uint, userId uint) interface{}
}

type classUseCase struct {
	classRepository repository.ClassRepository
	userRepository  repository.UserRepository
}

func NewClassUseCase(classRepository repository.ClassRepository, userRepository repository.UserRepository) ClassUseCase {

	return &classUseCase{
		classRepository: classRepository,
		userRepository:  userRepository,
	}

}

func (classUseCase *classUseCase) GetAllClassUseCase() ([]entity.ClassApi, interface{}) {

	var allClass []entity.ClassApi

	err := classUseCase.classRepository.FindAllClass(&allClass)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to get all class",
			Err:     err,
		}
		return nil, errObject

	}

	return allClass, nil

}

func (classUseCase *classUseCase) RemoveUserFromClassUseCase(loginUser entity.User, classId uint, userId uint) interface{} {

	//validate if user is admin
	if loginUser.Role != "admin" {

		errObject := library.ErrorObject{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
			Err:     errors.New("this endpoint only can be called by admin"),
		}
		return errObject

	}

	//validate if class exist
	var class entity.Class

	err := classUseCase.classRepository.FindClassByCondition(&class, "id = ?", classId)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "class doesn't exist",
			Err:     err,
		}
		return errObject

	}

	//validate if user exist
	var user entity.User

	err = classUseCase.userRepository.ELFindUserByCondition(&user, "id = ?", userId)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "user doesn't exist",
			Err:     err,
		}
		return errObject

	}

	//check if user is in the class
	var userInClass bool
	for _, v := range user.Classes {

		if v.ID == classId {

			userInClass = true
			break

		}

	}

	if !userInClass {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "user doesn't join the class",
			Err:     errors.New("this endpoint only can be called if user join the class"),
		}
		return errObject

	}

	//remove user from class
	classUseCase.userRepository.DropUserFromClass(&user, &class)

	return nil

}

func (classUseCase *classUseCase) AdmAddUserToClassUseCase(loginUser entity.User, classId uint, userId uint) interface{} {

	//validate if user is admin
	if loginUser.Role != "admin" {

		errObject := library.ErrorObject{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
			Err:     errors.New("this endpoint only can be called by admin"),
		}
		return errObject

	}

	//validate if class exist
	var class entity.Class

	err := classUseCase.classRepository.FindClassByCondition(&class, "id = ?", classId)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "class doesn't exist",
			Err:     err,
		}
		return errObject

	}

	//validate if user exist
	var user entity.User

	err = classUseCase.userRepository.ELFindUserByCondition(&user, "id = ?", userId)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "user doesn't exist",
			Err:     err,
		}
		return errObject

	}

	//check if user is in the class
	var userInClass bool
	for _, v := range user.Classes {

		if v.ID == classId {

			userInClass = true
			break

		}

	}

	if userInClass {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "user already join the class",
			Err:     errors.New("this endpoint only can be called if user doesn't yet join the class"),
		}
		return errObject

	}

	//add user to class
	classUseCase.userRepository.AddUserToClass(&user, &class)
	return nil

}
