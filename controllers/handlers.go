package controllers

import (
	"fmt"
	"net/http"
    "net"
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
        nhs := c.PostForm("nhs")
        password := c.PostForm("password")

        if helpers.EmptyNhsOrPass(nhs, password) {
            c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
            return
        }

        user := models.GetUser(nhs)

        if !helpers.CheckPassword(user.Password, password) {
            c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect nhs or password"})
            return
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "user": user.Nhs,
            "exp": time.Now().Add(time.Hour).Unix(),
        })

        tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
        if err != nil {
            fmt.Println("Error signing token.")
            c.HTML(http.StatusInternalServerError, "login.html", gin.H{})
            return
        }

        c.SetSameSite(http.SameSiteStrictMode)
        c.SetCookie("token", tokenString, 3600, "", "", false, true)
        if err !=nil {
            fmt.Println("Failed to set cookie.")
        }

        c.Redirect(http.StatusMovedPermanently, "/dashboard")
    }
}


func LogoutGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        c.SetCookie("token", "", -1, "", "", false, true) // Delete user cookie
        c.Redirect(http.StatusTemporaryRedirect, "/login")
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

        if !models.ValidRegister(nhs, email) {
            c.HTML(http.StatusConflict, "register.html", gin.H{"content": "Email or NHS already exists"})
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
        medicalSpecialties, err := models.GetMedicalSpecialties()
        if err != nil {
            fmt.Println("Failed to get medical specialties.")
        }

        c.HTML(http.StatusOK, "new_appointment.html", gin.H{"MedicalSpecialties": medicalSpecialties})
    }
}


func NewAppointmentPostHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        user, _  := c.Get("user")
        date := c.PostForm("date")
        medicalSpecialty := c.PostForm("medicalSpecialty")

        err := models.NewAppointment(user.(models.User).Nhs, date, medicalSpecialty)
        if err != nil {
            fmt.Println("Failed to create appointment.")
        }

        // TODO html new_appointment success
        c.Redirect(http.StatusMovedPermanently, "/new_appointment")
    }
}


func AppointmentsGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        user, _ := c.Get("user")
        appointments , err := models.GetUserAppointments(user.(models.User).Nhs)
        if err != nil {
            fmt.Println("Failed to get user appointments from database.")
        }

	c.HTML(http.StatusOK, "appointments.html", gin.H{"Appointments" : appointments})
    }
}


func ConsultationsGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        // Send OTP to authenticate
        conn, err := net.Dial("tcp", "localhost:8081")
        if err != nil {
            fmt.Println("Failed to connect to the OTP client service.")
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }

        otp, err := helpers.GenerateOTP()
        if err != nil {
            fmt.Println("Failed to generate OTP.")
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }

        _, err = conn.Write([]byte(otp))
        if err != nil {
            fmt.Println("Failed to send OTP.")
            c.AbortWithStatus(http.StatusInternalServerError)
            return
        }

        conn.Close()

        user, _ := c.Get("user")

        err = models.SetOTP(user.(models.User).Nhs, otp)
        if err != nil {
            fmt.Println("Unable to insert OTP into the database.")
            return
        }

        c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
    }
}


/*
func ConsultationsPostHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        user, _ := c.Get("user")
        otpAttempt := c.PostForm("otp")
        otp, err := models.GetOTP(user.(models.User).Nhs)
        if err != nil {
            fmt.Println("Unable to retrieve OTP from database.")
            return
        }

        var content string
        if helpers.ValidateOTP(otp.Value, otpAttempt, otp.Created) {
            err := models.DeleteOTP(user.(models.User).Nhs)
            if err != nil {
                fmt.Println("Unable to delete OTP from database.")
                return
            }
            content = "You are now twice as authenticated."
        } else {
            content = "The submitted OTP is incorrect or has already expired."
        }

        c.HTML(http.StatusOK, "consultations.html", gin.H{"content": content})
    }
}
*/
