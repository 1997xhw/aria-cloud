package Middlewares

import (
	"aria-cloud/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckLogin(c *gin.Context) {
	//token, err := c.Cookie("Token")
	username := c.Query("username")
	if username == "" {
		fmt.Println("need username")
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}

	token := c.Query("token")
	if token == "" {
		fmt.Println("need token")
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}
	if len(username) < 3 || !handler.IsTokenVaild(username, token) {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}

	//	验证完成
	c.Set("username", username)
	c.Set("token", token)
	c.Next()
}
