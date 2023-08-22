package rest

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/sdk/library"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (rest *Rest) Registration(c *gin.Context) {

	//binding user request
	var userInput entity.RegistBind
	err := c.ShouldBindJSON(&userInput)
	if err != nil {

		library.FailedResponse(c, http.StatusBadRequest, "send the correct JSON request", err)
		return

	}

	//create new user
	createdUser, errObject := rest.uc.User.RegistrationUseCase(userInput)
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusCreated, "successfully registration new user!", createdUser)

}

func (rest *Rest) Login(c *gin.Context) {

	//binding user request
	var userInput entity.LoginBind

	err := c.ShouldBindJSON(&userInput)
	if err != nil {

		library.FailedResponse(c, http.StatusBadRequest, "send the correct JSON request", err)
		return

	}

	errObject := rest.uc.User.LoginUseCase(userInput, c)
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusOK, "successes login", nil)

}

func (rest *Rest) EditAccount(c *gin.Context) {

	var userInput entity.EditAccountBind

	err := c.ShouldBindJSON(&userInput)
	if err != nil {

		library.FailedResponse(c, http.StatusBadRequest, "failed to bind input", err)
		return

	}

	loginUser, ok := c.Get("user")
	if !ok {

		library.FailedResponse(c, http.StatusInternalServerError, "failed to generate login user", errors.New("you are not authorized"))
		return

	}

	userApi, errObject := rest.uc.User.EditAccountUseCase(userInput, loginUser.(entity.User))
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusAccepted, "successfully edited", userApi)

}

func (rest *Rest) AddUserToClass(c *gin.Context) {

	var userInput entity.AddClassBind

	err := c.ShouldBindJSON(&userInput)
	if err != nil {

		library.FailedResponse(c, http.StatusBadRequest, "failed to bind input", err)
		return

	}

	loginUser, ok := c.Get("user")
	if !ok {

		library.FailedResponse(c, http.StatusInternalServerError, "failed to generate login user", errors.New("you are not authorized"))
		return
		
	}

	errObject := rest.uc.User.AddUserToClassUseCase(loginUser.(entity.User), userInput.ClassCode)
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusCreated, "successfully join class", nil)

}

func (rest *Rest) DropClass(c *gin.Context) {

	classId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {

		library.FailedResponse(c, http.StatusBadRequest, "failed convert id to int", err)
		return

	}

	loginUser, ok := c.Get("user")
	if !ok {

		library.FailedResponse(c, http.StatusInternalServerError, "failed to generate login user", nil)

	}

	errObject := rest.uc.User.DropClassUseCase(loginUser.(entity.User), uint(classId))
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusOK, "successfully drop class", nil)

}
