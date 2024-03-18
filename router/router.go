package router

import (
	"aria-cloud/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/login", controllers.Login)
	router.GET("/verify", controllers.Verify)

	router.POST("/register", controllers.RegiesterHandler)
	router.POST("/login", controllers.LoginHandler)
	//router.POST("/file/upload", controllers.UploadFile)
	aria := router.Group("/aria")
	//aria.Use(middlewares.CheckLogin)
	aria.Use()
	{
		aria.GET("/home", controllers.UserHome)
		aria.POST("/file/upload", controllers.UploadFile)
		aria.GET("/file/allList", controllers.AllFileList)
		aria.POST("/file/delete", controllers.DeleteFile)
	}

	return router
}
