package rest

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/sdk/library"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rest *rest) ClaimStudentNumber(c *gin.Context) {

	loginUser, _ := c.Get("user")

	//generate student number
	student, errObject := rest.uc.Student.ClaimStudentNumberUseCase(loginUser.(entity.User))
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusOK, "Succesed claim student number", student)

}
