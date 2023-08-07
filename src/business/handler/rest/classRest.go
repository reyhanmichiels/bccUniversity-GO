package rest

import (
	"bcc-university/src/sdk/library"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (rest *rest) GetAllClass(c *gin.Context) {

	allClass, errObject := rest.uc.Class.GetAllClassUseCase()
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusOK, "successfully get all class", allClass)

}
