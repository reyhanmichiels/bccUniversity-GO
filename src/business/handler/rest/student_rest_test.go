package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestClaimStudentNumberPath1(t *testing.T) {

	for i := 1; i <= 5; i++ {

		t.Run(fmt.Sprintf("path 1 clain student number testing %d", i), func(t *testing.T) {

			engine := gin.Default()
			engine.GET("/api/v1/user/student-number", userRest.ClaimStudentNumber)

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
