package routes

import (
	"github.com/gin-gonic/gin"
	"app/api/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("/users", controllers.GetUsers)
		v1.POST("/users", controllers.CreateUser)
	}
	return router
}