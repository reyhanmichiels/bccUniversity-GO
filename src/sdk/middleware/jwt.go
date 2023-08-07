package middleware

import (
	"bcc-university/src/business/entity"
	"bcc-university/src/sdk/db/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthJWT(c *gin.Context) {

	//claims token from cookie
	tokenString, err := c.Cookie("jwt-token")
	if err != nil {

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "unauthorized",
			"error":   err.Error(),
		})
		return

	}

	//parse  and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}

		return []byte(os.Getenv("JWT_SECRET_TOKEN")), nil

	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "unauthorized",
			"error":   err.Error(),
		})
		return

	}

	//check expired
	if float64(time.Now().Unix()) > claims["exp"].(float64) {

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "your session has been expired",
			"error":   nil,
		})
		return

	}

	//set login user
	var user entity.User

	sql.SQLDB.Model(&entity.User{}).Preload("Student").First(&user, claims["user_id"])
	c.Set("user", user)

	c.Next()

}
