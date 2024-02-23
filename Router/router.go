package Router

import (
	"aria-cloud/Controllers"
	"aria-cloud/Middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/login", Controllers.Login)
	router.GET("/siginup", Controllers.SiginUp)

	aria := router.Group("aria")
	aria.Use(Middlewares.CheckLogin)
	{
		aria.GET("/home", Controllers.UserHome)
	}

	return router
}
