package main

import "github.com/gin-gonic/gin"

import "net/http"

func main() {
	engine:= gin.Default();
	engine.LoadHTMLGlob("http/html/templates/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "hello gin!",
		})
	})
	engine.Run(":3000")
}