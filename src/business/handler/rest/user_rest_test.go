package rest

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/business/usecase"
	"bytes"
	"encoding/json"
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

		t.Run(fmt.Sprintf("registration testing %d", i + 1), func(t *testing.T) {

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
