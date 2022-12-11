package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"sah/models"
)

func AuthRequired(c *gin.Context) {
    tokenString, err := c.Cookie("token")
    if err != nil {
        c.AbortWithStatus(http.StatusUnauthorized)
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
        }

        return []byte(os.Getenv("SECRET_KEY")), nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        exp := claims["exp"].(float64)
        if float64(time.Now().Unix()) > exp {
            c.AbortWithStatus(http.StatusUnauthorized)
        }

        email := claims["user"].(string)
        user := models.GetUser(email)
 
        c.Set("user", user)

        c.Next()

    } else {
        c.AbortWithStatus(http.StatusUnauthorized)
    }
}
