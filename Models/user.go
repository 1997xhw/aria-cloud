package Models

import (
	"aria-cloud/Databases/mysql"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	UserName string `gorm:"column:user_name;"`
	UserPwd  string `gorm:"column:user_pwd;"`
}

func (User) TableName() string {
	return "tbl_user"
}
func AuthenticateUser(username, enpwd string) (bool, error) {
	var user User
	fmt.Println(username, enpwd)
	result := mysql.DB.Where("user_name=?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// 用户不存在
			return false, result.Error
		}
		// 数据库错误
		return false, result.Error
	}

	// 检查密码是否匹配
	if user.UserPwd != enpwd {
		// 密码不匹配
		var ErrPasswordMismatch = errors.New("password does not match")
		return false, ErrPasswordMismatch
	}

	// 用户名和密码都匹配
	return true, nil
}

type UserToken struct {
	UserName  string `gorm:"column:user_name;"`
	UserToken string `gorm:"column:user_token"`
}

// TableName 设置UserToken模型对应的表名
func (UserToken) TableName() string {
	return "tbl_user_token"
}

func UpdateUserToken(username, token string) bool {

	result := mysql.DB.Model(&UserToken{}).Where("user_name = ?", username).Updates(UserToken{UserToken: token})

	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	}

	return true
}
