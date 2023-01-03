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
        c.Redirect(http.StatusTemporaryRedirect, "/login")
        return
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
            c.Redirect(http.StatusTemporaryRedirect, "/login")
            return
        }

        nhs := claims["user"].(string)
        user := models.GetUser(nhs)
 
        c.Set("user", user)

        c.Next()

    } else {
        c.Redirect(http.StatusTemporaryRedirect, "/login")
        return
    }
}

// TODO verify if user is already logged in
