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

var userUseCaseMock usecase.UserUseCaseMock = usecase.UserUseCaseMock{
	Mock: mock.Mock{},
}

var userRest = &Rest{
	uc: &usecase.UseCase{
		User: &userUseCaseMock,
	},
}

func TestRegistration(t *testing.T) {

	userInput := []entity.RegistBind{
		{
			Name:     "testName1",
			Username: "test1",
			Email:    "testemail1@test.com",
			Password: "testpass1",
		},
		{
			Name:     "testName2",
			Username: "test2",
			Email:    "testemail2@test.com",
			Password: "testpass2",
		},
		{
			Name:     "testName3",
			Username: "test3",
			Email:    "testemail3@test.com",
			Password: "testpass3",
		},
		{
			Name:     "testName4",
			Username: "test4",
			Email:    "testemail4@test.com",
			Password: "testpass4",
		},
		{
			Name:     "testName5",
			Username: "test5",
			Email:    "testemail5@test.com",
			Password: "testpass5",
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("registration testing %d", i+1), func(t *testing.T) {

			functionReturn := entity.RegistApi{
				Name:     v.Name,
				Email:    v.Email,
				Username: v.Username,
			}

			callFunction := userUseCaseMock.Mock.On("RegistrationUseCase", v).Return(functionReturn, nil)

			engine := gin.Default()
			engine.POST("/api/v1/regist", userRest.Registration)

			jsonData, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/regist", bytes.NewBuffer(jsonData))
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			userResponse := jsonResponse["data"].(map[string]any)

			assert.Equal(t, http.StatusCreated, response.Code, "should give http status code created")
			assert.Equal(t, v.Name, userResponse["name"], "name should be equal")
			assert.Equal(t, v.Username, userResponse["username"], "username should be equal")
			assert.Equal(t, v.Email, userResponse["email"], "email should be equal")
			assert.Equal(t, "successfully registration new user!", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "success", jsonResponse["status"], "status should be equal")

			callFunction.Unset()

		})

	}

}

func TestLogin(t *testing.T) {

	userInput := []entity.LoginBind{
		{
			Email:    "test1@test.com",
			Password: "testpass1",
		},
		{
			Email:    "test2@test.com",
			Password: "testpass2",
		},
		{
			Email:    "test3@test.com",
			Password: "testpass3",
		},
		{
			Email:    "test4@test.com",
			Password: "testpass4",
		},
		{
			Email:    "test5@test.com",
			Password: "testpass5",
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("login testing %d", i+1), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/login", userRest.Login)

			callFunction := userUseCaseMock.Mock.On("LoginUseCase", v).Return(nil)

			jsonData, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonData))
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusOK, response.Code, "http status code should be equal")
			assert.Equal(t, "success", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "successes login", jsonResponse["message"], "message should be equal")
			assert.Equal(t, nil, jsonResponse["data"], "data should be equal")

			callFunction.Unset()

		})

	}

}

// func TestEditAccount(t *testing.T) {

// 	userInput := []entity.EditAccountBind{
// 		{},
// 		{
// 			Username: "test2",
// 		},
// 		{
// 			Username: "test3",
// 		},
// 		{
// 			Username: "test 4",
// 		},
// 	}

// 	for i, v := range userInput {

// 		t.Run(fmt.Sprintf("edit account testing %d", i+1), func(t *testing.T) {

// 			engine := gin.Default()
// 			engine.POST("/api/v1/user", userRest.EditAccount)
// 			callFunction := userUseCaseMock.Mock.On("LoginUseCase", v).Return(nil)

// 			loginUser := entity.User{}

// 			if i == 0 {

// 				jsonData, err := json.Marshal(v)
// 				if err != nil {

// 					t.Fatal(err.Error())

// 				}

// 				response := httptest.NewRecorder()
// 				request, err := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonData))
// 				if err != nil {

// 					t.Fatal(err.Error())

// 				}
// 				engine.ServeHTTP(response, request)

// 				var jsonResponse map[string]any
// 				err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
// 				if err != nil {

// 					t.Fatal(err.Error())

// 				}

// 				assert.Equal(t, http.StatusBadRequest, response.Code, "status code should be equal")
// 				assert.Equal(t, "failed to bind input", jsonResponse["message"], "message should be equal")
// 				assert.Equal(t, "error", jsonResponse["status"], "status should be equal")

// 			}

// 			if i == 1 {

// 				jsonData, err := json.Marshal(v)
// 				if err != nil {

// 					t.Fatal(err.Error())

// 				}

// 				response := httptest.NewRecorder()
// 				request, err := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonData))
// 				if err != nil {

// 					t.Fatal(err.Error())

// 				}
// 				engine.ServeHTTP(response, request)

// 				var jsonResponse map[string]any
// 				err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
// 				if err != nil {

// 					t.Fatal(err.Error())

// 				}

// 				assert.Equal(t, http.StatusInternalServerError, response.Code, "status code should be equal")
// 				assert.Equal(t, "failed to generate login user", jsonResponse["message"], "message should be equal")
// 				assert.Equal(t, "error", jsonResponse["status"], "status should be equal")

// 			}

// 			if i > 1 {

// 				callFunction.Unset()

// 			}

// 		})

// 	}

// }

func TestEditAccountPath1(t *testing.T) {

	userInput := []entity.EditAccountBind{
		{},
		{},
		{},
		{},
		{},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 1 edit account testing %d", i+1), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/user", userRest.EditAccount)

			jsonData, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(jsonData))
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusBadRequest, response.Code, "status code should be equal")
			assert.Equal(t, "failed to bind input", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")

		})

	}

}

func TestEditAccountPath2(t *testing.T) {

	userInput := []entity.EditAccountBind{
		{
			Username: "testuser1",
		},
		{
			Username: "testuser2",
		},
		{
			Username: "testuser3",
		},
		{
			Username: "testuser4",
		},
		{
			Username: "testuser5",
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 2 edit account testing %d", i+1), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/user", userRest.EditAccount)

			jsonData, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(jsonData))
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

func TestEditAccountPath3(t *testing.T) {

	userInput := []entity.EditAccountBind{
		{
			Username: "testuser1",
		},
		{
			Username: "testuser2",
		},
		{
			Username: "testuser3",
		},
		{
			Username: "testuser4",
		},
		{
			Username: "testuser5",
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 3 edit account testing %d", i+1), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/user", setUserLogin, userRest.EditAccount)

			errObject := library.ErrorObject{
				Code:    http.StatusInternalServerError,
				Message: "test",
				Err:     errors.New("test"),
			}
			functionCall := userUseCaseMock.Mock.On("EditAccountUseCase", v, getLoginUser()).Return(entity.UserApi{}, errObject)

			jsonData, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(jsonData))
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

func TestEditAccountPath4(t *testing.T) {

	userInput := []entity.EditAccountBind{
		{
			Username: "testuser1",
		},
		{
			Username: "testuser2",
		},
		{
			Username: "testuser3",
		},
		{
			Username: "testuser4",
		},
		{
			Username: "testuser5",
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 4 edit account testing %d", i+1), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/user", setUserLogin, userRest.EditAccount)

			userApi := entity.UserApi{
				Name:     getLoginUser().Name,
				Username: v.Username,
				Email:    getLoginUser().Email,
				Role:     getLoginUser().Role,
			}
			functionCall := userUseCaseMock.Mock.On("EditAccountUseCase", v, getLoginUser()).Return(userApi, nil)

			jsonData, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(jsonData))
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			userResponse := jsonResponse["data"].(map[string]any)

			assert.Equal(t, http.StatusAccepted, response.Code, "status code should be equal")
			assert.Equal(t, "successfully edited", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "success", jsonResponse["status"], "status should be equal")
			assert.Equal(t, userApi.Name, userResponse["name"], "name should be equal")
			assert.Equal(t, userApi.Username, userResponse["username"], "username should be equal")
			assert.Equal(t, userApi.Email, userResponse["email"], "email should be equal")
			assert.Equal(t, userApi.Role, userResponse["role"], "role should be equal")

			functionCall.Unset()

		})

	}

}

func TestAddUserToClassPath1(t *testing.T) {

	userInput := []entity.AddClassBind{
		{},
		{},
		{},
		{},
		{},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 1 add user to class testing %d", i+1), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/user/class", userRest.AddUserToClass)

			jsonData, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/user/class", bytes.NewBuffer(jsonData))
			if err != nil {

				t.Fatal(err.Error())

			}
			engine.ServeHTTP(response, request)

			var jsonResponse map[string]any
			err = json.Unmarshal(response.Body.Bytes(), &jsonResponse)
			if err != nil {

				t.Fatal(err.Error())

			}

			assert.Equal(t, http.StatusBadRequest, response.Code, "status code should be equal")
			assert.Equal(t, "failed to bind input", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")

		})

	}

}

func TestAddUserToClassPath2(t *testing.T) {

	userInput := []entity.AddClassBind{
		{
			ClassCode: "test1",
		},
		{
			ClassCode: "test2",
		},
		{
			ClassCode: "test3",
		},
		{
			ClassCode: "test4",
		},
		{
			ClassCode: "test5",
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 2 add user to class testing %d", i+1), func(t *testing.T) {

			engine := gin.Default()
			engine.POST("/api/v1/user/class", userRest.AddUserToClass)

			jsonData, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/user/class", bytes.NewBuffer(jsonData))
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

func TestAddUserToClassPath3(t *testing.T) {

	userInput := []entity.AddClassBind{
		{
			ClassCode: "test1",
		},
		{
			ClassCode: "test2",
		},
		{
			ClassCode: "test3",
		},
		{
			ClassCode: "test4",
		},
		{
			ClassCode: "test5",
		},
	}

	for i, v := range userInput {

		t.Run(fmt.Sprintf("path 3 add user to class testing %d", i+1), func(t *testing.T) {

			errObject := library.ErrorObject{
				Code:    http.StatusInternalServerError,
				Message: "test",
				Err:     errors.New("test"),
			}
			functionCall := userUseCaseMock.Mock.On("AddUserToClassUseCase", getLoginUser(), v.ClassCode).Return(errObject)

			engine := gin.Default()
			engine.POST("/api/v1/user/class", setUserLogin, userRest.AddUserToClass)

			jsonData, err := json.Marshal(v)
			if err != nil {

				t.Fatal(err.Error())

			}

			response := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/api/v1/user/class", bytes.NewBuffer(jsonData))
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
			assert.Equal(t, "error", jsonResponse["status"], "status should be equal")
			assert.Equal(t, "test", jsonResponse["message"], "message should be equal")
			assert.Equal(t, "test", jsonResponse["error"], "error should be equal")

			functionCall.Unset()

		})

	}

}

func setUserLogin(c *gin.Context) {

	c.Set("user", entity.User{
		Name:     "test",
		Username: "test",
		Email:    "test@test.com",
		Password: "testtest",
		Role:     "user",
		Student:  entity.Student{},
		Classes:  []entity.Class{},
	})

}

func getLoginUser() entity.User {

	return entity.User{
		Name:     "test",
		Username: "test",
		Email:    "test@test.com",
		Password: "testtest",
		Role:     "user",
		Student:  entity.Student{},
		Classes:  []entity.Class{},
	}

}
