package routes

import (
	"github.com/gin-gonic/gin"
	"app/api/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		// 第二引数に「*gin.Context」を引数に取る「handlers ...gin.HandlerFunc」を渡しておけば、内部でContextを生成してくれる。
		v1.GET("/users", controllers.GetUsers)
		v1.POST("/users", controllers.CreateUser)
		v1.POST("login", controllers.Login)
	}
	return router
}