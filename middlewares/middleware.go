package middlewares

import (
	"aria-cloud/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
		//c.Redirect(http.StatusFound, "/login")
		//c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "用户名为空",
		})
		return
	}

	if token == "" {
		log.Println("need token")
		//c.Redirect(http.StatusFound, "/login")
		//c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "token为空",
		})
		return
	}
	//fmt.Println(username)
	//fmt.Println(token)
	err := services.IsTokenVaild(username, token)
	if len(username) < 3 || err != nil {
		//c.Redirect(http.StatusFound, "/login")
		//c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//	验证完成
	c.Set("username", username)
	c.Set("token", token)
	c.Next()
}
