package usecase

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/business/repository"
	"bcc-university/src/sdk/library"
	"errors"
	"net/http"
	"strings"
)

type ClassUseCase interface {
	GetAllClassUseCase() ([]entity.ClassApi, interface{})
	RemoveUserFromClassUseCase(loginUser entity.User, classId uint, userId uint) interface{}
	AdmAddUserToClassUseCase(loginUser entity.User, classId uint, userId uint) interface{}
	CreateClassUseCase(userInput entity.CreateUpdateClassBind, loginUser entity.User) (entity.CreateUpdateClassApi, interface{})
	EditClassUseCase(userInput entity.CreateUpdateClassBind, loginUser entity.User, classId uint) (entity.CreateUpdateClassApi, interface{})
	DeleteClassUseCase(loginUser entity.User, classId uint) interface{}
}

type classUseCase struct {
	classRepository  repository.ClassRepository
	userRepository   repository.UserRepository
	courseRepository repository.CourseRepository
}

func NewClassUseCase(classRepository repository.ClassRepository, userRepository repository.UserRepository, courseRepository repository.CourseRepository) ClassUseCase {

	return &classUseCase{
		classRepository:  classRepository,
		userRepository:   userRepository,
		courseRepository: courseRepository,
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
			Code:    http.StatusForbidden,
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
			Code:    http.StatusNotFound,
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
			Code:    http.StatusNotFound,
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
			Code:    http.StatusNotFound,
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
			Code:    http.StatusForbidden,
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
			Code:    http.StatusForbidden,
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
			Code:    http.StatusForbidden,
			Message: "user doesn't exist",
			Err:     err,
		}
		return errObject

	}

	//validate if user is student
	if user.Student.Student_id_number == "" {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "user is not student",
			Err:     errors.New("you only can attach student to the class"),
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

func (classUsecase *classUseCase) CreateClassUseCase(userInput entity.CreateUpdateClassBind, loginUser entity.User) (entity.CreateUpdateClassApi, interface{}) {

	//validate if user is admin
	if loginUser.Role != "admin" {

		errObject := library.ErrorObject{
			Code:    http.StatusForbidden,
			Message: "unauthorized",
			Err:     errors.New("this endpoint only can be called by admin"),
		}
		return entity.CreateUpdateClassApi{}, errObject

	}

	//validate if course exist
	course := struct {
		Name   string
		Credit int
	}{}

	err := classUsecase.courseRepository.FindCourseByCondition(&course, "id = ?", userInput.Course_id)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "course doesn't exist",
			Err:     err,
		}
		return entity.CreateUpdateClassApi{}, errObject

	}

	//create class
	class := entity.Class{
		Name:      userInput.Name,
		Course_id: userInput.Course_id,
		ClassCode: library.GenerateClassCode(userInput.Name),
	}

	err = classUsecase.classRepository.CreateClass(&class)
	if err != nil {

		code := http.StatusInternalServerError
		if strings.Contains(err.Error(), "Duplicate entry") {

			code = http.StatusConflict

		}

		errObject := library.ErrorObject{
			Code:    code,
			Message: "failed to create class",
			Err:     err,
		}
		return entity.CreateUpdateClassApi{}, errObject

	}

	classApi := entity.CreateUpdateClassApi{
		Name:      class.Name,
		Course_id: class.Course_id,
		ClassCode: class.ClassCode,
	}

	classApi.Course.Name = course.Name
	classApi.Course.Credit = course.Credit

	return classApi, nil

}

func (classUseCase *classUseCase) EditClassUseCase(userInput entity.CreateUpdateClassBind, loginUser entity.User, classId uint) (entity.CreateUpdateClassApi, interface{}) {

	//validate if user is admin
	if loginUser.Role != "admin" {

		errObject := library.ErrorObject{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
			Err:     errors.New("this endpoint only can be called by admin"),
		}
		return entity.CreateUpdateClassApi{}, errObject

	}

	//validate if class exist
	var class entity.Class
	err := classUseCase.classRepository.FindClassByCondition(&class, "id = ?", classId)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "Class doesn't exist",
			Err:     err,
		}
		return entity.CreateUpdateClassApi{}, errObject

	}

	//validate if course exist
	course := struct {
		Name   string
		Credit int
	}{}

	err = classUseCase.courseRepository.FindCourseByCondition(&course, "id = ?", userInput.Course_id)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "course doesn't exist",
			Err:     err,
		}
		return entity.CreateUpdateClassApi{}, errObject

	}

	//update class
	updateData := struct {
		Name      string
		Course_id uint
		ClassCode string
	}{
		Name:      userInput.Name,
		Course_id: userInput.Course_id,
		ClassCode: library.GenerateClassCode(userInput.Name),
	}

	err = classUseCase.classRepository.Updateclass(&class, updateData)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to update class",
			Err:     err,
		}
		return entity.CreateUpdateClassApi{}, errObject

	}

	classApi := entity.CreateUpdateClassApi{
		Name:      class.Name,
		Course_id: class.Course_id,
		ClassCode: class.ClassCode,
	}

	classApi.Course.Name = course.Name
	classApi.Course.Credit = course.Credit

	return classApi, nil

}

func (classUseCase *classUseCase) DeleteClassUseCase(loginUser entity.User, classId uint) interface{} {

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
			Message: "Class doesn't exist",
			Err:     err,
		}
		return errObject

	}

	//delete class
	err = classUseCase.classRepository.DeleteClass(&class)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "Class doesn't exist",
			Err:     err,
		}
		return errObject

	}

	return nil

}
