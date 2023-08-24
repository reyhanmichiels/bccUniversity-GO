package rest

import (
	"bcc-university/src/business/usecase"
	"bcc-university/src/sdk/library"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var classUsecaseMock = usecase.ClassUseCaseMock{
	Mock: mock.Mock{},
}

var classRest = Rest{
	uc: &usecase.UseCase{
		Class: &classUsecaseMock,
	},
}

func TestGetAllClassPath1(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 1 get all class testing %d", i), func(t *testing.T) {

			errObject := library.ErrorObject{
				Code:    http.StatusInternalServerError,
				Message: "test",
				Err:     errors.New("test"),
			}
			functionCall := classUsecaseMock.Mock.On("GetAllClassUseCase").Return(nil, errObject)

			engine := gin.Default()
			engine.GET("/api/v1/classes", classRest.GetAllClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("GET", "/api/v1/classes", nil)
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusInternalServerError, response.Code, "status code should be equal")
			assert.Equal(t, "test", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")

			functionCall.Unset()

		})

	}

}
