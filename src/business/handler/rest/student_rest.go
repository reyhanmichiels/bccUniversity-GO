package rest

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/sdk/library"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rest *Rest) ClaimStudentNumber(c *gin.Context) {

	loginUser, ok := c.Get("user")
	if !ok {

		library.FailedResponse(c, http.StatusInternalServerError, "failed to generate login user", nil)

	}

	//generate student number
	student, errObject := rest.uc.Student.ClaimStudentNumberUseCase(loginUser.(entity.User))
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusOK, "Succesed claim student number", student)

}
