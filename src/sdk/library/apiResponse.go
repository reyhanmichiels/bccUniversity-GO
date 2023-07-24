package library

import "github.com/gin-gonic/gin"

func SuccessedResponse(c *gin.Context, code int, message string, data interface{}) {

	c.JSON(code, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})

}

func FailedResponse(c *gin.Context, code int, errorMessage error) {

	c.JSON(code, gin.H{
		"status": "error",
		"error":  errorMessage.Error(),
	})

}
