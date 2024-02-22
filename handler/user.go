package handler

import (
	"aria-cloud/db"
	"aria-cloud/util"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	pwd_salt = "*#890"
)

// SignupHandler 处理用户注册请求
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		file, err := os.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(file)
		if err != nil {
			panic(err)
			return
		}
		return
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	if len(username) < 3 || len(passwd) < 5 {
		_, err := w.Write([]byte("Invalid parameter"))
		if err != nil {
			return
		}
		return
	}
	sha1 := util.Sha1([]byte(passwd + pwd_salt))
	suc := db.UserSignup(username, sha1)

	// 创建用户专属文件夹
	userFileDir := "./data/" + username
	exists, err := util.PathExists(userFileDir)
	if !exists {
		err := os.MkdirAll(userFileDir, 0755)
		if err != nil {
			fmt.Println("创建文件夹时发生错误:", err)
			return
		}
		fmt.Println("文件夹创建成功:", userFileDir)
	} else {
		if err != nil {
			fmt.Println("创建用户专属文件夹失败!!!")
			w.Write([]byte("创建用户专属文件夹失败"))
			return
		} else {
			fmt.Println("已存在用户文件夹!!!")
		}
	}

	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}

// SignInHandler 登陆接口
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		file, err := os.ReadFile("./static/view/signin.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(file)
		if err != nil {
			panic(err)
			return
		}
		return
	}
	// 1. 校验用户名以及密码
	r.ParseForm()
	usename := r.Form.Get("username")
	password := r.Form.Get("password")
	encPassword := util.Sha1([]byte(password + pwd_salt))

	pwdChecked := db.UserSignin(usename, encPassword)
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}

	// 2. 生成访问凭证（token）
	token := GenToken(usename)
	udRes := db.UpdateToken(usename, token)
	if !udRes {
		w.Write([]byte("TONKEN FAILED"))
		return
	}
	// 3. 登陆成功后重定向到首页
	//w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "http://" + r.Host + "/static/view/home.html",
			Username: usename,
			Token:    token,
		},
	}
	//fmt.Println(resp.JSONBytes())
	_, err := w.Write(resp.JSONBytes())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 解析请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	token := r.Form.Get("token")
	// 2. 验证token是否有效
	if !IsTokenVaild(username, token) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 3. 查询用户信息
	userInfo, err := db.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 4. 组装并且响应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: userInfo,
	}
	_, err = w.Write(resp.JSONBytes())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func GenToken(username string) string {
	//40位 md5(username+timestamp+token_salt)+timestamp[:8]
	timestamp := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + timestamp + "_tokonsalt"))
	return tokenPrefix + timestamp[:8]
}

// IsTokenVaild 验证token
func IsTokenVaild(username string, token string) bool {
	if len(token) != 40 {
		fmt.Println("token is wrong!")
		return false
	}
	// 1. 判断token时效性
	hexTimestamp := token[len(token)-8:]
	timestamp, err := strconv.ParseInt(hexTimestamp, 16, 64)
	if err != nil {
		fmt.Println("Error converting hex to int:", err)
		return false
	}
	// 将Unix时间戳转换为time.Time
	tokenTime := time.Unix(timestamp, 0)
	// 检查token时间是否超过1小时
	if time.Since(tokenTime).Hours() > 2 {
		fmt.Printf("token已过期！！")
		return false
	}

	// 2. 从数据表tbl_user_token查询username对应的token信息
	dbToken, _ := db.GetTokenByUsername(username)
	// 3. 对比两个token是否一致
	if dbToken != token {
		fmt.Println("token不一致！！！")
		return false
	}

	return true
}
