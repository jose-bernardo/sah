package routes

import (
    "github.com/gin-gonic/gin"
    "sah/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {
    g.GET("/", controllers.IndexGetHandler())
    g.GET("/login", controllers.LoginGetHandler())
    g.POST("/login", controllers.LoginPostHandler())
    g.GET("/register", controllers.RegisterGetHandler())
    g.POST("/register", controllers.RegisterPostHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
    g.GET("/dashboard", controllers.DashboardGetHandler())
    g.GET("/new_appointment", controllers.NewAppointmentGetHandler())
    g.POST("/new_appointment", controllers.NewAppointmentPostHandler())
    g.GET("/appointments", controllers.AppointmentsGetHandler())
    g.GET("/consultations", controllers.ConsultationsGetHandler())
    g.GET("/logout", controllers.LogoutGetHandler())
}
