package router

import (
	"aria-cloud/controllers"
	"aria-cloud/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/login", controllers.Login)
	router.POST("/register", controllers.RegiesterHandler)
	router.POST("/login", controllers.LoginHandler)
	//router.POST("/file/upload", controllers.UploadFile)
	aria := router.Group("/aria")
	aria.Use(middlewares.CheckLogin)
	{
		aria.GET("/home", controllers.UserHome)
		aria.POST("/file/upload", controllers.UploadFile)
	}

	return router
}
