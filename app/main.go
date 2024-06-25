package main

import (
 "net/http"

 "github.com/gin-gonic/gin"
)

func main() {
 r := gin.Default()

 r.GET("/hello", func(c *gin.Context) {
  c.String(http.StatusOK, "Hello world")
})

 r.Run(":8080")
}
