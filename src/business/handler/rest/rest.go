package rest

import (
	"bcc-university/src/business/usecase"
	"bcc-university/src/sdk/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Rest interface {
	Run()
}

type rest struct {
	uc  *usecase.UseCase
	gin *gin.Engine
}

func (r *rest) Run() {
	r.gin.Run()
}

func (r *rest) Route() {
	v1 := r.gin.Group("/api/v1")

	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Api Running",
		})
	})

	v1.POST("/regist", r.Registration)

	v1.POST("/login", r.Login)

	v1.GET("/user/student-number", middleware.AuthJWT, r.ClaimStudentNumber)

	v1.POST("/user", middleware.AuthJWT, r.EditAccount)

	v1.GET("/classes", middleware.AuthJWT, r.GetAllClass)

	v1.POST("/user/class", middleware.AuthJWT, r.AddUserToClass)

	v1.DELETE("/user/class/:id", middleware.AuthJWT, r.DropClass)

	v1.DELETE("/class/:classId/user/:userId", middleware.AuthJWT, r.RemoveUserFromClass)

	v1.POST("/class/:classId/user/:userId", middleware.AuthJWT, r.AdmAddUserToClass)

	v1.POST("/class", middleware.AuthJWT, r.CreateClass)

	v1.POST("/class/:classId", middleware.AuthJWT, r.EditClass)

	v1.DELETE("class/:classId", middleware.AuthJWT, r.DeleteClass)

}

func InjectRest(usecase *usecase.UseCase) Rest {
	r := &rest{
		uc:  usecase,
		gin: gin.Default(),
	}

	r.Route()
	return r
}
