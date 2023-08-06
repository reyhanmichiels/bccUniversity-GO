package rest

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/sdk/library"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rest *rest) Registration(c *gin.Context) {

	//binding user request
	var userFromRequest entity.CreateUser
	err := c.ShouldBindJSON(&userFromRequest)
	if err != nil {

		library.FailedResponse(c, http.StatusConflict, "send the correct JSON request", err)
		return

	}

	//create new user
	createdUser, errorObject := rest.uc.User.CreateUser(userFromRequest)
	if errorObject != nil {

		errorObject := errorObject.(library.ErrorObject)
		library.FailedResponse(c, errorObject.Code, errorObject.Message, errorObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusCreated, "successes registration new user!", createdUser)

}

func (rest *rest) Login(c *gin.Context) {

	//binding user request
	var userFromRequest entity.LoginUser
	err := c.ShouldBindJSON(&userFromRequest)
	if err != nil {

		library.FailedResponse(c, http.StatusConflict, "send the correct JSON request", err)
		return

	}

	//verify credential
	loginUser, errObject := rest.uc.User.VerifyCredential(userFromRequest)
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	//generate jwt token
	token, errObject := rest.uc.User.GenerateJWTToken(loginUser)
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	//set token to cookie
	rest.uc.User.SetToken(c, token)

	library.SuccessedResponse(c, http.StatusOK, "successes login", nil)

}

func (rest *rest) EditAccount(c *gin.Context) {

	var inputUser entity.EditProfileBind

	err := c.ShouldBindJSON(&inputUser)
	if err != nil {

		library.FailedResponse(c, http.StatusConflict, "failed to bind input", err)
		return

	}

	loginUser, ok := c.Get("user")
	if !ok {

		library.FailedResponse(c, http.StatusInternalServerError, "failed to generate login user", nil)

	}

	responseUser, errObject := rest.uc.User.EditProfile(inputUser, loginUser.(entity.User))
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusAccepted, "successfully edited", responseUser)

}
