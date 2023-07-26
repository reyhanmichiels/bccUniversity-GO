package usecase

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/business/repository"
	"bcc-university/src/sdk/library"
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
