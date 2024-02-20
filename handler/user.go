package handler

import (
	"aria-cloud/db"
	"aria-cloud/util"
	"fmt"
	"net/http"
	"os"
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
	w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
}
func GenToken(username string) string {
	//40位 md5(username+timestamp+token_salt)+timestamp[:8]
	timestamp := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + timestamp + "_tokonsalt"))
	return tokenPrefix + timestamp[:8]
}
