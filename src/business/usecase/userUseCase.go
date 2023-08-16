package usecase

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/business/repository"
	"bcc-university/src/sdk/library"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegistrationUseCase(userInput entity.RegistBind) (entity.RegistApi, interface{})
	LoginUseCase(userInput entity.LoginBind, c *gin.Context) interface{}
	EditAccountUseCase(userInput entity.EditAccountBind, loginUser entity.User) (entity.UserApi, interface{})
	AddUserToClassUseCase(loginUser entity.User, classCode string) interface{}
	DropClassUseCase(loginUser entity.User, classId uint) interface{}
}

type userUseCase struct {
	userRepository  repository.UserRepository
	classRepository repository.ClassRepository
}

func NewUserUseCase(userRepository repository.UserRepository, classRepository repository.ClassRepository) UserUseCase {
	return &userUseCase{
		userRepository:  userRepository,
		classRepository: classRepository,
	}
}

func (userUseCase *userUseCase) RegistrationUseCase(userInput entity.RegistBind) (entity.RegistApi, interface{}) {

	//hash pasword
	password, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), 10)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate hash password",
			Err:     err,
		}

		return entity.RegistApi{}, errObject

	}

	//create user
	user := entity.User{
		Name:     userInput.Name,
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: string(password),
	}

	err = userUseCase.userRepository.CreateUser(&user)
	if err != nil {

		code := http.StatusInternalServerError
		if strings.Contains(err.Error(), "Duplicate entry") {

			code = http.StatusBadRequest

		}

		errObject := library.ErrorObject{
			Code:    code,
			Message: "failed to create user",
			Err:     err,
		}
		return entity.RegistApi{}, errObject

	}

	userApi := entity.RegistApi{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
	}

	return userApi, nil

}

func (userUseCase *userUseCase) LoginUseCase(userInput entity.LoginBind, c *gin.Context) interface{} {

	//verify credential
	user := struct {
		ID       uint
		Password string
	}{}

	err := userUseCase.userRepository.FindUserByCondition(&user, "email = ?", userInput.Email)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "user not found",
			Err:     err,
		}
		return errObject

	}

	//verify credential
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "invalid password!",
			Err:     err,
		}
		return errObject

	}

	//generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_TOKEN")))
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate JWT Token",
			Err:     err,
		}
		return errObject

	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("jwt-token", tokenString, (3600 * 24), "", "", false, true)

	return nil

}

func (userUseCase *userUseCase) EditAccountUseCase(userInput entity.EditAccountBind, loginUser entity.User) (entity.UserApi, interface{}) {

	err := userUseCase.userRepository.UpdateUser(&loginUser, userInput)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to update user",
			Err:     err,
		}
		return entity.UserApi{}, errObject

	}

	userApi := entity.UserApi{
		Name:     loginUser.Name,
		Email:    loginUser.Email,
		Username: loginUser.Username,
		Role:     loginUser.Role,
	}

	return userApi, nil

}

func (userUseCase *userUseCase) AddUserToClassUseCase(loginUser entity.User, classCode string) interface{} {

	//validate if user a student

	if loginUser.Student.Student_id_number == "" {

		errObject := library.ErrorObject{
			Code:    http.StatusUnauthorized,
			Message: "you are not student",
			Err:     errors.New("this endpoint only can be called by student"),
		}
		return errObject

	}

	//validate if the class exist

	var targetClass entity.Class

	err := userUseCase.classRepository.ELFindClassByCondition(&targetClass, "class_code = ?", classCode)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "you have wrong class code",
			Err:     err,
		}
		return errObject

	}

	var userCredit int
	for _, v := range loginUser.Classes {

		//validate if user already take the class
		if v.ClassCode == classCode {

			errObject := library.ErrorObject{
				Code:    http.StatusConflict,
				Message: "you already take this class",
				Err:     errors.New("can't take the same class"),
			}
			return errObject

		}

		//validate if user doesn't take the same course
		if v.Course_id == targetClass.Course_id {

			errObject := library.ErrorObject{
				Code:    http.StatusConflict,
				Message: "you already take this course",
				Err:     errors.New("can't take the same course"),
			}
			return errObject

		}

		userCredit += v.Course.Credit

	}

	//validate if user has enough credit semester left
	if userCredit+targetClass.Course.Credit > 24 {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "you don't have enough credit semester",
			Err:     errors.New("maximum credit is 24"),
		}
		return errObject

	}

	//add user to class
	userUseCase.userRepository.AddUserToClass(&loginUser, &targetClass)

	return nil

}

func (userUseCase *userUseCase) DropClassUseCase(loginUser entity.User, classId uint) interface{} {

	//validate if the class exist
	var class entity.Class
	err := userUseCase.classRepository.FindClassByCondition(&class, "id = ?", classId)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "class not found",
			Err:     err,
		}
		return errObject

	}

	//check if user doesn't join the class
	var userInClass bool
	for _, v := range loginUser.Classes {

		if v.ID == classId {

			userInClass = true
			break
		}

	}

	if !userInClass {

		errObject := library.ErrorObject{
			Code:    http.StatusConflict,
			Message: "you don't join the class",
			Err:     errors.New("wrong class id"),
		}
		return errObject

	}

	//drop user from class
	userUseCase.userRepository.DropUserFromClass(&loginUser, &class)
	return nil
}
