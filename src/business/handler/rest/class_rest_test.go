package rest

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/business/usecase"
	"bcc-university/src/sdk/library"
	"bytes"
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

func TestRemoveUserFromClassPath3(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 3 remove user from class testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.DELETE("/api/v1/class/:classId/user/:userId", classRest.RemoveUserFromClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/class/%d/user/%[1]d", i), nil)

			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusInternalServerError, response.Code, "http status code should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "failed to generate login user", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "you are not authorized", jsonResponse["error"], "error should be equal")

		})

	}

}

func TestRemoveUserFromClassPath4(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 4 remove user from class testing %d", i), func(t *testing.T) {

			errObject := library.ErrorObject{
				Code:    http.StatusInternalServerError,
				Message: "test",
				Err:     errors.New("test"),
			}
			functionCall := classUsecaseMock.Mock.On("RemoveUserFromClassUseCase", getLoginUser(), uint(i), uint(i)).Return(errObject)

			engine := gin.Default()
			engine.DELETE("/api/v1/class/:classId/user/:userId", setUserLogin, classRest.RemoveUserFromClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/class/%d/user/%[1]d", i), nil)

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

func TestRemoveUserFromClassPath5(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 5 remove user from class testing %d", i), func(t *testing.T) {

			functionCall := classUsecaseMock.Mock.On("RemoveUserFromClassUseCase", getLoginUser(), uint(i), uint(i)).Return(nil)

			engine := gin.Default()
			engine.DELETE("/api/v1/class/:classId/user/:userId", setUserLogin, classRest.RemoveUserFromClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/class/%d/user/%[1]d", i), nil)

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
			assert.Equal(t, "successfully remove user from class", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "success", jsonResponse["status"], "status should be equal")

			functionCall.Unset()

		})

	}

}

func TestAdmAddUserToClassPath1(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 1 admin add user to class testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/class/:classId/user/:userId", classRest.RemoveUserFromClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/class/test/user/test", nil)
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

func TestAdmAddUserToClassPath2(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 2 admin add user to class testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/class/:classId/user/:userId", classRest.RemoveUserFromClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", fmt.Sprintf("/api/v1/class/test/user/%d", i), nil)
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

func TestAdmAddUserToClassPath3(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 3 admin add user to class testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/class/:classId/user/:userId", classRest.RemoveUserFromClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", fmt.Sprintf("/api/v1/class/%d/user/%[1]d", i), nil)
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusInternalServerError, response.Code, "http status code should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "failed to generate login user", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "you are not authorized", jsonResponse["error"], "error should be equal")

		})

	}

}

func TestAdmAddUserToClassPath4(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 4 admin add user to class testing %d", i), func(t *testing.T) {

			errObject := library.ErrorObject{
				Code:    http.StatusInternalServerError,
				Message: "test",
				Err:     errors.New("test"),
			}
			functionCall := classUsecaseMock.Mock.On("AdmAddUserToClassUseCase", getLoginUser(), uint(i), uint(i)).Return(errObject)

			engine := gin.Default()
			engine.POST("/api/v1/class/:classId/user/:userId", setUserLogin, classRest.AdmAddUserToClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", fmt.Sprintf("/api/v1/class/%d/user/%[1]d", i), nil)
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

func TestAdmAddUserToClassPath5(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 5 admin add user to class testing %d", i), func(t *testing.T) {

			functionCall := classUsecaseMock.Mock.On("AdmAddUserToClassUseCase", getLoginUser(), uint(i), uint(i)).Return(nil)

			engine := gin.Default()
			engine.POST("/api/v1/class/:classId/user/:userId", setUserLogin, classRest.AdmAddUserToClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", fmt.Sprintf("/api/v1/class/%d/user/%[1]d", i), nil)
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusCreated, response.Code, "status code should be equal")
			assert.Equal(t, "successfully add user to class", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "success", jsonResponse["status"], "status should be equal")

			functionCall.Unset()

		})

	}

}

func TestCreateClassPath1(t *testing.T) {

	userInput := []entity.CreateUpdateClassBind{
		{},
		{},
		{},
		{},
		{},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 1 Create Class testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/class", classRest.CreateClass)

			jsonInput, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/class", bytes.NewBuffer(jsonInput))
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
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "failed to bind input", jsonResponse["message"], "message should be equal")

		})

	}

}

func TestCreateClassPath2(t *testing.T) {

	userInput := []entity.CreateUpdateClassBind{
		{
			Name:      "testname1",
			Course_id: 1,
		},
		{
			Name:      "testname2",
			Course_id: 2,
		},
		{
			Name:      "testname3",
			Course_id: 3,
		},
		{
			Name:      "testname4",
			Course_id: 4,
		},
		{
			Name:      "testname5",
			Course_id: 5,
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 2 Create Class testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/class", classRest.CreateClass)

			jsonInput, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/class", bytes.NewBuffer(jsonInput))
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusInternalServerError, response.Code, "http status code should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "failed to generate login user", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "you are not authorized", jsonResponse["error"], "error should be equal")

		})

	}

}

func TestCreateClassPath3(t *testing.T) {

	userInput := []entity.CreateUpdateClassBind{
		{
			Name:      "testname1",
			Course_id: 1,
		},
		{
			Name:      "testname2",
			Course_id: 2,
		},
		{
			Name:      "testname3",
			Course_id: 3,
		},
		{
			Name:      "testname4",
			Course_id: 4,
		},
		{
			Name:      "testname5",
			Course_id: 5,
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 3 Create Class testing %d", i), func(t *testing.T) {

			errObject := library.ErrorObject{
				Code:    http.StatusInternalServerError,
				Message: "test",
				Err:     errors.New("test"),
			}
			functionCall := classUsecaseMock.Mock.On("CreateClassUseCase", v, getLoginUser()).Return(entity.CreateUpdateClassApi{}, errObject)

			engine := gin.Default()
			engine.POST("/api/v1/class", setUserLogin, classRest.CreateClass)

			jsonInput, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/class", bytes.NewBuffer(jsonInput))
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusInternalServerError, response.Code, "http status code should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "test", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "test", jsonResponse["error"], "error should be equal")

			functionCall.Unset()

		})

	}

}

func TestCreateClassPath4(t *testing.T) {

	userInput := []entity.CreateUpdateClassBind{
		{
			Name:      "testname1",
			Course_id: 1,
		},
		{
			Name:      "testname2",
			Course_id: 2,
		},
		{
			Name:      "testname3",
			Course_id: 3,
		},
		{
			Name:      "testname4",
			Course_id: 4,
		},
		{
			Name:      "testname5",
			Course_id: 5,
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 4 Create Class testing %d", i), func(t *testing.T) {

			classApi := entity.CreateUpdateClassApi{
				Name:        v.Name,
				Course_id:   v.Course_id,
				Participant: i,
				ClassCode:   "test",
			}
			classApi.Course.Name = "test"
			classApi.Course.Credit = i
			functionCall := classUsecaseMock.Mock.On("CreateClassUseCase", v, getLoginUser()).Return(classApi, nil)

			engine := gin.Default()
			engine.POST("/api/v1/class", setUserLogin, classRest.CreateClass)

			jsonInput, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/class", bytes.NewBuffer(jsonInput))
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			dataResponse := jsonResponse["data"].(map[string]any)
			courseResponse := dataResponse["Course"].(map[string]any)

			assert.Equal(t, http.StatusCreated, response.Code, "http status code should be equal")
			assert.Equal(t, "success", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "successfully create new class", jsonResponse["message"], "message should be equal")
			assert.Equal(t, classApi.Name, dataResponse["name"], "name should be equal")
			assert.Equal(t, classApi.ClassCode, dataResponse["class_code"], "class_code should be equal")
			assert.Equal(t, classApi.Course_id, uint(dataResponse["course_id"].(float64)), "course_id should be equal")
			assert.Equal(t, classApi.Participant, int(dataResponse["participant"].(float64)), "participant should be equal")
			assert.Equal(t, classApi.Course.Name, courseResponse["name"], "course name should be equal")
			assert.Equal(t, classApi.Course.Credit, int(courseResponse["credit"].(float64)), "course credit should be equal")

			functionCall.Unset()

		})

	}

}

func TestEditClassPath1(t *testing.T) {

	userInput := []entity.CreateUpdateClassBind{
		{},
		{},
		{},
		{},
		{},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 1 Edit Class testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/class/:classId", classRest.EditClass)

			jsonInput, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", fmt.Sprintf("/api/v1/class/%d", i), bytes.NewBuffer(jsonInput))
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
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "failed to bind input", jsonResponse["message"], "message should be equal")

		})

	}

}

func TestEditClassPath2(t *testing.T) {

	userInput := []entity.CreateUpdateClassBind{
		{
			Name:      "testname1",
			Course_id: 1,
		},
		{
			Name:      "testname2",
			Course_id: 2,
		},
		{
			Name:      "testname3",
			Course_id: 3,
		},
		{
			Name:      "testname4",
			Course_id: 4,
		},
		{
			Name:      "testname5",
			Course_id: 5,
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 2 Edit Class testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/class/:classId", classRest.EditClass)

			jsonInput, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", fmt.Sprintf("/api/v1/class/%s", "test"), bytes.NewBuffer(jsonInput))
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

func TestEditClassPath3(t *testing.T) {

	userInput := []entity.CreateUpdateClassBind{
		{
			Name:      "testname1",
			Course_id: 1,
		},
		{
			Name:      "testname2",
			Course_id: 2,
		},
		{
			Name:      "testname3",
			Course_id: 3,
		},
		{
			Name:      "testname4",
			Course_id: 4,
		},
		{
			Name:      "testname5",
			Course_id: 5,
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 3 Edit Class testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/class/:classId", classRest.EditClass)

			jsonInput, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", fmt.Sprintf("/api/v1/class/%d", i), bytes.NewBuffer(jsonInput))
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusInternalServerError, response.Code, "http status code should be equal")
			assert.Equal(t, "failed to generate login user", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "you are not authorized", jsonResponse["error"], "error should be equal")

		})

	}

}

func TestEditClassPath4(t *testing.T) {

	userInput := []entity.CreateUpdateClassBind{
		{
			Name:      "testname1",
			Course_id: 1,
		},
		{
			Name:      "testname2",
			Course_id: 2,
		},
		{
			Name:      "testname3",
			Course_id: 3,
		},
		{
			Name:      "testname4",
			Course_id: 4,
		},
		{
			Name:      "testname5",
			Course_id: 5,
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 4 Edit Class testing %d", i), func(t *testing.T) {

			errObject := library.ErrorObject{
				Code:    http.StatusInternalServerError,
				Message: "test",
				Err:     errors.New("test"),
			}
			functionCall := classUsecaseMock.Mock.On("EditClassUseCase", v, getLoginUser(), uint(i)).Return(nil, errObject)

			engine := gin.Default()
			engine.POST("/api/v1/class/:classId", setUserLogin, classRest.EditClass)

			jsonInput, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", fmt.Sprintf("/api/v1/class/%d", i), bytes.NewBuffer(jsonInput))
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusInternalServerError, response.Code, "http status code should be equal")
			assert.Equal(t, "test", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "test", jsonResponse["error"], "error should be equal")

			functionCall.Unset()

		})

	}

}

func TestEditClassPath5(t *testing.T) {

	userInput := []entity.CreateUpdateClassBind{
		{
			Name:      "testname1",
			Course_id: 1,
		},
		{
			Name:      "testname2",
			Course_id: 2,
		},
		{
			Name:      "testname3",
			Course_id: 3,
		},
		{
			Name:      "testname4",
			Course_id: 4,
		},
		{
			Name:      "testname5",
			Course_id: 5,
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 5 Edit Class testing %d", i), func(t *testing.T) {

			classApi := entity.CreateUpdateClassApi{
				Name:        v.Name,
				Course_id:   v.Course_id,
				Participant: i,
				ClassCode:   "test",
			}
			classApi.Course.Name = "test"
			classApi.Course.Credit = i
			functionCall := classUsecaseMock.Mock.On("EditClassUseCase", v, getLoginUser(), uint(i)).Return(classApi, nil)

			engine := gin.Default()
			engine.POST("/api/v1/class/:classId", setUserLogin, classRest.EditClass)

			jsonInput, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", fmt.Sprintf("/api/v1/class/%d", i), bytes.NewBuffer(jsonInput))
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}
			dataResponse := jsonResponse["data"].(map[string]any)
			courseResponse := dataResponse["Course"].(map[string]any)

			assert.Equal(t, http.StatusOK, response.Code, "http status code should be equal")
			assert.Equal(t, "success", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "successfully edit class", jsonResponse["message"], "message should be equal")
			assert.Equal(t, classApi.Name, dataResponse["name"], "name should be equal")
			assert.Equal(t, classApi.ClassCode, dataResponse["class_code"], "class_code should be equal")
			assert.Equal(t, classApi.Course_id, uint(dataResponse["course_id"].(float64)), "course_id should be equal")
			assert.Equal(t, classApi.Participant, int(dataResponse["participant"].(float64)), "participant should be equal")
			assert.Equal(t, classApi.Course.Name, courseResponse["name"], "course name should be equal")
			assert.Equal(t, classApi.Course.Credit, int(courseResponse["credit"].(float64)), "course credit should be equal")

			functionCall.Unset()

		})

	}

}

func TestDeleteClassPath1(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 1 delete class testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.DELETE("/api/v1/class/:classId", classRest.DeleteClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("DELETE", "/api/v1/class/test", nil)	
			if err  != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse) 
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusBadRequest, response.Code, "http status code should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "failed to convert class id to int", jsonResponse["message"], "message should be equal")

		})

	}

}

func TestDeletePath2(t *testing.T) {

	for i:= 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 2 delete class testing %d",i), func(t *testing.T) {

			engine := gin.Default()
			engine.DELETE("/api/v1/class/:classId", classRest.DeleteClass)

			response := httptest.NewRecorder()
			request, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/class/%d", i), nil)	
			if err  != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse) 
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusInternalServerError, response.Code, "http status code should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "failed to generate login user", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "you are not authorized", jsonResponse["error"], "error should be equal")

		})

	}

}
