package controllers

import (
	"fmt"
	"net/http"
    "time"
    "os"

	"sah/helpers"
	"sah/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IndexGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/login")
    }
}

func LoginGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", gin.H{})
    }
}

func LoginPostHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        email := c.PostForm("email")
        password := c.PostForm("password")

        if helpers.EmptyEmailOrPass(email, password) {
            c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
            return
        }

        hash, err := models.GetUserPass(email)
        if err != nil {
            panic(err.Error())
        }

        if !helpers.CheckPassword(hash, password) {
            c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect email or password"})
            return
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "user": email,
            //"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
            "exp": time.Now().Add(time.Second * 30).Unix(),
        })

        tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
        if err != nil {
            fmt.Println("Error signing token.")
            c.HTML(http.StatusInternalServerError, "login.html", gin.H{})
            return
        }
 
        c.SetSameSite(http.SameSiteLaxMode)
        c.SetCookie("token", tokenString, 3600 * 24 * 30, "", "", false, true)

        c.Redirect(http.StatusMovedPermanently, "/dashboard")
    }
}

func LogoutGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        c.SetCookie("token", "", -1, "", "", false, true)
        c.Redirect(http.StatusMovedPermanently, "/login")
    }
}

func RegisterGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        c.HTML(http.StatusOK, "register.html", gin.H{})
    }
}

func RegisterPostHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        email := c.PostForm("email")
        name := c.PostForm("name")
        nhs := c.PostForm("nhs")
        password := c.PostForm("password")

        if helpers.EmptyRegisterParams(email, name, nhs, password) {
            c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "Parameters can't be empty"})
            return
        }

        hash, err := helpers.HashPassword(password)
        if err != nil {
            fmt.Println("Unable to hash password.")
        }

        if !models.ValidRegister(nhs) {
            c.HTML(http.StatusConflict, "register.html", gin.H{"content": "Username or NHS already exists"})
            return
        }

        err = models.RegisterUser(email, name, nhs, hash) 
        if err != nil {
            fmt.Println("Error registering user.")
        }

        c.Redirect(http.StatusMovedPermanently, "/login")
    }
}

func DashboardGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        user, _ := c.Get("user")
        c.HTML(http.StatusOK, "dashboard.html", gin.H{"content": user.(models.User).Name})
    }
}

func NewAppointmentGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        c.HTML(http.StatusOK, "new_appointment.html", gin.H{})
    }
}

func NewAppointmentPostHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        nhs := c.PostForm("nhs")
        date := c.PostForm("date")
        medicalSpecialty := c.PostForm("medicalSpecialty")

        err := models.NewAppointment(nhs, date, medicalSpecialty)
        if err != nil {
            panic(err.Error())
        }

        c.Redirect(http.StatusCreated, "/dashboard")
    }
}

func AppointmentsGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        // TODO
    }
}
