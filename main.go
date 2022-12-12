package main

import (
	"github.com/gin-gonic/gin"

    "sah/routes"
    "sah/models"
    "sah/middleware"
)

func main() {
    models.ConnectDB()

    r := gin.Default()
    r.LoadHTMLGlob("templates/*.html")

    public := r.Group("/")
    routes.PublicRoutes(public)

    private := r.Group("/")
    private.Use(middleware.AuthRequired)
    routes.PrivateRoutes(private)

    r.RunTLS(":8080", "./certs/server-cert.pem", "./certs/server-key.pem")
}
