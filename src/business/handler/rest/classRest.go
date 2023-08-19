package rest

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/sdk/library"
	"net/http"
	"strconv"

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

func (rest *rest) RemoveUserFromClass(c *gin.Context) {

	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {

		library.FailedResponse(c, http.StatusBadRequest, "failed to convert user id to int", err)
		return

	}

	classId, err := strconv.ParseUint(c.Param("classId"), 10, 64)
	if err != nil {

		library.FailedResponse(c, http.StatusBadRequest, "failed to convert class id to int", err)
		return

	}

	loginUser, ok := c.Get("user")
	if !ok {

		library.FailedResponse(c, http.StatusInternalServerError, "failed to generate login user", nil)
		return

	}

	errObject := rest.uc.Class.RemoveUserFromClassUseCase(loginUser.(entity.User), uint(classId), uint(userId))
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusOK, "successfully remove user from class", nil)

}

func (rest *rest) AdmAddUserToClass(c *gin.Context) {

	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {

		library.FailedResponse(c, http.StatusBadRequest, "failed to convert user id to int", err)
		return

	}

	classId, err := strconv.ParseUint(c.Param("classId"), 10, 64)
	if err != nil {

		library.FailedResponse(c, http.StatusBadRequest, "failed to convert class id to int", err)
		return

	}

	loginUser, ok := c.Get("user")
	if !ok {

		library.FailedResponse(c, http.StatusInternalServerError, "failed to generate login user", nil)
		return

	}

	errObject := rest.uc.Class.AdmAddUserToClassUseCase(loginUser.(entity.User), uint(classId), uint(userId))
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusCreated, "successfully add user to class", nil)

}

func (rest *rest) CreateClass(c *gin.Context) {

	var userInput entity.CreateUpdateClassBind

	err := c.ShouldBindJSON(&userInput)
	if err != nil {

		library.FailedResponse(c, http.StatusBadRequest, "failed to bind input", err)
		return

	}

	loginUser, ok := c.Get("user")
	if !ok {

		library.FailedResponse(c, http.StatusInternalServerError, "failed to generate login user", nil)
		return

	}

	class, errObject := rest.uc.Class.CreateClassUseCase(userInput, loginUser.(entity.User))
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusCreated, "successfully create new class", class)

}

func (rest *rest) EditClass(c *gin.Context) {

	var userInput entity.CreateUpdateClassBind

	err := c.ShouldBindJSON(&userInput)
	if err != nil {

		library.FailedResponse(c, http.StatusBadRequest, "failed to bind input", err)
		return

	}

	classId, err := strconv.ParseUint(c.Param("classId"), 10, 64)
	if err != nil {

		library.FailedResponse(c, http.StatusBadRequest, "failed to convert class id to int", err)
		return

	}

	loginUser, ok := c.Get("user")
	if !ok {

		library.FailedResponse(c, http.StatusInternalServerError, "failed to generate login user", nil)
		return

	}

	class, errObject := rest.uc.Class.EditClassUseCase(userInput, loginUser.(entity.User), uint(classId))
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusOK, "successfully edit class", class)

}

func (rest *rest) DeleteClass(c *gin.Context) {

	classId, err := strconv.ParseUint(c.Param("classId"), 10, 64)
	if err != nil {

		library.FailedResponse(c, http.StatusConflict, "failed to convert class id to int", err)
		return

	}

	loginUser, ok := c.Get("user")
	if !ok {

		library.FailedResponse(c, http.StatusInternalServerError, "failed to generate login user", nil)
		return

	}

	errObject := rest.uc.Class.DeleteClassUseCase(loginUser.(entity.User), uint(classId))
	if errObject != nil {

		errObject := errObject.(library.ErrorObject)
		library.FailedResponse(c, errObject.Code, errObject.Message, errObject.Err)
		return

	}

	library.SuccessedResponse(c, http.StatusOK, "successfully delete class", nil)

}
