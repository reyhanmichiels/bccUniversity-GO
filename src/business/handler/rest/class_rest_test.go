package rest

import (
	"bcc-university/src/business/entity"
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

func TestGetAllClassPath2(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 2 get all class testing %d", i), func(t *testing.T) {

			classResponse := []entity.ClassApi{
				{
					Name:        "testname1",
					Course_id:   1,
					Participant: 1,
					ClassCode:   "testcode1",
				},
				{
					Name:        "testname1",
					Course_id:   1,
					Participant: 1,
					ClassCode:   "testcode1",
				},
			}
			functionCall := classUsecaseMock.Mock.On("GetAllClassUseCase").Return(classResponse, nil)

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

			assert.Equal(t, http.StatusOK, response.Code, "status code should be equal")
			assert.Equal(t, "successfully get all class", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "success", jsonResponse["status"], "status should be equal")

			for i, v := range jsonResponse["data"].([]interface{}) {

				v := v.(map[string]any)

				assert.Equal(t, classResponse[i].Name, v["name"], fmt.Sprintf("name %d should be equal", i+1))
				assert.Equal(t, classResponse[i].Course_id, uint(v["course_id"].(float64)), fmt.Sprintf("course id %d should be equal", i+1))
				assert.Equal(t, classResponse[i].Participant, int(v["participant"].(float64)), fmt.Sprintf("participant %d should be equal", i+1))
				assert.Equal(t, classResponse[i].ClassCode, v["class_code"], fmt.Sprintf("class code %d should be equal", i+1))

			}

			functionCall.Unset()

		})

	}

}

func TestRemoveUserFromClassPath1(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 1 remove user from class testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.DELETE("/api/v1/class/:classId/user/:userId", classRest.RemoveUserFromClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("DELETE", "/api/v1/class/test/user/test", nil)
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusBadRequest, response.Code, "http status code should be equal")
			assert.Equal(t, "failed to convert user id to int", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")

		})

	}

}

func TestRemoveUserFromClassPath2(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 2 remove user from class testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.DELETE("/api/v1/class/:classId/user/:userId", classRest.RemoveUserFromClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("DELETE", "/api/v1/class/test/user/1", nil)
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusBadRequest, response.Code, "http status code should be equal")
			assert.Equal(t, "failed to convert class id to int", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")

		})

	}

}
