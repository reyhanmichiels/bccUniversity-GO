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
