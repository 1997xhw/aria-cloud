package controllers

import (
	"aria-cloud/services"
	"aria-cloud/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", nil)
}
func Verify(c *gin.Context) {
	username := c.Query("username")
	token := c.Query("token")
	fmt.Println(username, token)
	err := services.IsTokenVaild(username, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	}
}

func SiginUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReturnData struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func RegiesterHandler(c *gin.Context) {
	var user = User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		//绑定失败
		fmt.Println("数据绑定失败：", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "参数格式不正确",
		})
	} else {
		res, err := services.Register(user.Username, user.Password)
		if res {
			c.JSON(http.StatusOK, gin.H{
				"msg": "注册成功",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				//"code": http.StatusInternalServerError,
				"msg": err.Error(),
			})
		}
	}

}

func LoginHandler(c *gin.Context) {
	jsonData, _ := c.GetRawData()
	var m map[string]interface{}
	err := json.Unmarshal(jsonData, &m)
	if err != nil {
		fmt.Println("数据解析失败：", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg": "参数格式不正确",
			//"status": 500,
		})
	}
	fmt.Println(m["username"])
	fmt.Println(m["password"])
	var username string
	var password string
	if val, ok := m["username"].(string); ok {
		username = val
	} else {
		fmt.Println("用户名不是有效的字符串")
		// 处理错误情况
	}

	if val, ok := m["password"].(string); ok {
		password = val
	} else {
		fmt.Println("密码不是有效的字符串")
		// 处理错误情况
	}
	fmt.Println("==================")
	res, err := services.LoginHandler(username, password)
	if res {
		token := GenToken(username)
		fmt.Println("----------------")
		fmt.Println(token)
		_ = services.UpadteUserToken(username, token)
		c.JSON(http.StatusOK, gin.H{
			"msg": "登陆成功",
			"data": ReturnData{
				Username: username,
				Token:    token,
			},
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}

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

func GenToken(username string) string {
	//40位 md5(username+timestamp+token_salt)+timestamp[:8]
	timestamp := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + timestamp + "_tokonsalt"))
	return tokenPrefix + timestamp[:8]
}
