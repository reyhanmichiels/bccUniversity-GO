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

var studentUseCaseMock = usecase.StudentUseCaseMock{
	Mock: mock.Mock{},
}

var studentRest = &Rest{
	uc: &usecase.UseCase{
		Student: &studentUseCaseMock,
	},
}

func TestClaimStudentNumberPath1(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 1 claim student number testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.GET("/api/v1/user/student-number", studentRest.ClaimStudentNumber)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("GET", "/api/v1/user/student-number", nil)
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
			assert.Equal(t, "failed to generate login user", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")

		})

	}

}

func TestClaimStudentNumberPath2(t *testing.T) {

	for i := 1; i <= 5; i++ {

		errObject := library.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "test",
			Err:     errors.New("test"),
		}
		functionCall := studentUseCaseMock.Mock.On("ClaimStudentNumberUseCase", getLoginUser()).Return(nil, errObject)

		t.Run(fmt.Sprintf("path 2 claim student number testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.GET("/api/v1/user/student-number", setUserLogin, studentRest.ClaimStudentNumber)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("GET", "/api/v1/user/student-number", nil)
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

func TestClaimStudentNumberPath3(t *testing.T) {

	for i := 1; i <= 5; i++ {

		claimStudentNumberApi := entity.ClaimStudentNumberApi{
			Student_id_number: "test",
		}
		functionCall := studentUseCaseMock.Mock.On("ClaimStudentNumberUseCase", getLoginUser()).Return(claimStudentNumberApi, nil)

		t.Run(fmt.Sprintf("path 3 claim student number testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.GET("/api/v1/user/student-number", setUserLogin, studentRest.ClaimStudentNumber)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("GET", "/api/v1/user/student-number", nil)
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			data := jsonResponse["data"].(map[string]any)

			assert.Equal(t, http.StatusOK, response.Code, "status code should be equal")
			assert.Equal(t, "Succesed claim student number", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "success", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "test", data["student_id_number"], "student id number should be equal")

			functionCall.Unset()

		})

	}

}
