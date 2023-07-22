package rest

import (
	"bcc-university/src/business/usecase"
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
}

func InjectRest(usecase *usecase.UseCase) Rest {
	r := &rest{
		uc:  usecase,
		gin: gin.Default(),
	}

	r.Route()
	return r
}
