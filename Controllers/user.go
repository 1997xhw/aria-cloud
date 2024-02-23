package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", nil)
}

func SiginUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func UserHome(c *gin.Context) {

	username, exits := c.Get("username")
	if !exits {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	token, exits := c.Get("token")
	if !exits {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
		"token":    token})
}
