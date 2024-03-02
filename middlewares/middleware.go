package middlewares

import (
	"aria-cloud/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckLogin(c *gin.Context) {

	var username, token string
	if c.Request.Method == "GET" {
		username = c.Query("username")
		token = c.Query("token")
	} else if c.Request.Method == "POST" {
		username = c.PostForm("username")
		token = c.PostForm("token")
	}

	if username == "" {
		fmt.Println("need username")
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}

	if token == "" {
		fmt.Println("need token")
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	//fmt.Println(username)
	//fmt.Println(token)
	if len(username) < 3 || !controllers.IsTokenVaild(username, token) {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}

	//	验证完成
	c.Set("username", username)
	c.Set("token", token)
	c.Next()
}
