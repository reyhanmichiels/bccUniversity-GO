package usecase

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/business/repository"
	"bcc-university/src/sdk/library"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	CreateUser(inputUser entity.CreateUser) (entity.UserRegistResponse, interface{})
	VerifyCredential(inputUser entity.LoginUser) (entity.User, interface{})
	GenerateJWTToken(loginUser entity.User) (string, interface{})
	SetToken(c *gin.Context, token string)
	EditProfile(inputUser entity.EditProfileBind, loginUser entity.User) (entity.ResponseUser, interface{})
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

func (userUseCase *userUseCase) VerifyCredential(inputUser entity.LoginUser) (entity.User, interface{}) {

	//find user by email
	user, err := userUseCase.userRepository.FindUserByEmail(inputUser.Email)
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "user not found",
			Err:     err,
		}
		return entity.User{}, errObject

	}

	//verify credential
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password))
	if err != nil {

		errObject := library.ErrorObject{
			Code:    http.StatusUnauthorized,
			Message: "invalid password!",
			Err:     err,
		}
		return entity.User{}, errObject

	}

	return user, nil

}

func (userUseCase *userUseCase) GenerateJWTToken(loginUser entity.User) (string, interface{}) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": loginUser.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_TOKEN")))
	if err != nil {

		errorObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to generate JWT Token",
			Err:     err,
		}
		return "", errorObject

	}

	return tokenString, nil

}

func (userUseCase *userUseCase) SetToken(c *gin.Context, token string) {

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("jwt-token", token, (3600 * 24), "", "", false, true)

}

func (userUseCase *userUseCase) EditProfile(inputUser entity.EditProfileBind, loginUser entity.User) (entity.ResponseUser, interface{}) {

	err := userUseCase.userRepository.UpdateUser(&loginUser, inputUser)
	if err != nil {

		errorObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to update user",
			Err:     err,
		}
		return entity.ResponseUser{}, errorObject

	}

	userResponse := entity.ResponseUser{
		Name:     loginUser.Name,
		Email:    loginUser.Email,
		Username: loginUser.Username,
		Role:     loginUser.Role,
	}

	return userResponse, nil

}

func (userUseCase *userUseCase) AddUserToClassUseCase(loginUser entity.User, classCode string) interface{} {

	var targetClass entity.Class

	err := userUseCase.classRepository.ELFindClassByClassCode(&targetClass, classCode)
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
	err := userUseCase.classRepository.FindClassById(&class, classId)
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
