package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.LoadHTMLGlob("templates/*.html")

  r.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{})
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
