package main

import "github.com/gin-gonic/gin"

func main() {
	engine:= gin.Default()
	engine.Static("/static", "./http/image/static")
	engine.Run(":3000")
}