package controllers

import (
    "net/http"

    "sah/helpers"
    "sah/models"

    "github.com/gin-gonic/gin"
)

func IndexGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        c.Redirect(http.StatusOK, "login.html")
    }
}

func LoginGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", gin.H{})
    }
}

func LoginPostHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")

        if helpers.EmptyUserOrPass(username, password) {
            c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
            return
        }

        hash, err := models.GetUserPass(username)
        if err != nil {
            panic(err.Error())
        }

        if !helpers.CheckPassword(hash, password) {
            c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect username or password"})
            return
        }

        c.Redirect(http.StatusMovedPermanently, "/dashboard")
    }
}

func LogoutGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", gin.H{})
    }
}

func RegisterGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        c.HTML(http.StatusOK, "register.html", gin.H{})
    }
}

func RegisterPostHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        username := c.PostForm("username")
        name := c.PostForm("name")
        nhs := c.PostForm("nhs")
        password := c.PostForm("password")

        if helpers.EmptyRegisterParams(username, name, nhs, password) {
            c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "Parameters can't be empty"})
            return
        }

        hash, err := helpers.HashPassword(password)
        if err != nil {
            panic(err.Error())
        }

        if !models.ValidRegister(username, nhs) {
            c.HTML(http.StatusConflict, "register.html", gin.H{"content": "Username or NHS already exists"})
            return
        }

        err = models.RegisterUser(username, name, nhs, hash) 
        if err != nil {
            panic(err.Error())
        }

        c.Redirect(http.StatusMovedPermanently, "/login")
    }
}

func DashboardGetHandler() gin.HandlerFunc {
    return func (c *gin.Context) {
        c.HTML(http.StatusOK, "dashboard.html", gin.H{})
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
