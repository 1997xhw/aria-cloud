package services

import (
	"aria-cloud/models"
	"aria-cloud/util"
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	pwd_salt = "*#890"
)

func LoginHandler(username, password string) (bool, error) {
	encPassword := util.Sha1([]byte(password + pwd_salt))
	user, err := models.AuthenticateUser(username, encPassword)
	if !user || err != nil {
		return false, err
	}
	return true, nil
}

func Register(username, password string) (bool, error) {
	encPassword := util.Sha1([]byte(password + pwd_salt))
	resp, err := models.AuthRegister(username, encPassword)
	if resp {
		return true, nil
	} else {
		return false, err
	}

}

// IsTokenVaild 验证token
func IsTokenVaild(username string, token string) error {
	if len(token) != 40 {
		fmt.Println("token is wrong!")
		ne := errors.New("token is wrong!")
		return ne
	}
	// 1. 判断token时效性
	hexTimestamp := token[len(token)-8:]
	timestamp, err := strconv.ParseInt(hexTimestamp, 16, 64)
	if err != nil {
		fmt.Println("Error converting hex to int:", err)
		ne := errors.New(fmt.Sprintf("Error converting hex to int:", err))
		return ne
	}
	// 将Unix时间戳转换为time.Time
	tokenTime := time.Unix(timestamp, 0)
	// 检查token时间是否超过1小时
	if time.Since(tokenTime).Hours() > 2 {
		fmt.Printf("token已过期！！")
		ne := errors.New("token已过期！！")
		return ne
	}

	// 2. 从数据表tbl_user_token查询username对应的token信息
	dbToken, err := GetTokenByUsername(username)
	if err != nil {
		return err
	}
	fmt.Println("dbtoken: ", dbToken)
	fmt.Println("token: ", token)
	// 3. 对比两个token是否一致
	if dbToken != token {

		fmt.Println("token不一致！！！")
		ne := errors.New("token不一致！！！")
		return ne
	}

	return nil
}

func UpadteUserToken(username, token string) bool {
	return models.UpdateUserToken(username, token)
}

func GetTokenByUsername(username string) (string, error) {
	return models.GetTokenByUsername(username)
}
