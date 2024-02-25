package db

import (
	mydb "aria-cloud/old-aria/db/mysql"
	"database/sql"
	"fmt"
)

// UserSignup 通过用户名以及密码完成user表的注册操作
func UserSignup(username string, passwd string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user(`user_name`,`user_pwd`)values(?,?)")
	if err != nil {
		fmt.Println("Failed to insert, err" + err.Error())
		return false
	}
	defer stmt.Close()

	exec, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Println("Failed to insert, err" + err.Error())
		return false
	}
	if rowsAffected, err := exec.RowsAffected(); err == nil && rowsAffected > 0 {
		fmt.Println("------------------------ success to user sign up!!! ------------------------")
		return true
	}
	return false
}

func UserSignin(username string, encpwd string) bool {
	stmt, err := mydb.DBConn().Prepare("select * from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	defer stmt.Close()

	query, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else if query == nil {
		fmt.Println("username not found: " + username)
		return false
	}

	pRows := mydb.ParseRows(query)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
		return true
	}
	if string(pRows[0]["user_pwd"].([]byte)) != encpwd {
		fmt.Println("------------------------ password is wrong ------------------------")
	}

	return false
}

// UpdateToken 刷新用户登陆的token
func UpdateToken(username string, token string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"replace into tbl_user_token(`user_name`, `user_token`)values (?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func GetTokenByUsername(username string) (string, error) {
	row, err := mydb.DBConn().Prepare(
		`select user_token from tbl_user_token where user_name=? limit 1`)
	if err != nil {
		fmt.Println("获取token时发生错误: ", err.Error())
		return "", err
	}
	defer row.Close()
	var token string
	err = row.QueryRow(username).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			// 处理没有找到结果的情况
			return "", fmt.Errorf("没有找到对应的token")
		}
		// 处理其他可能的错误
		return "", fmt.Errorf("查询token时发生错误: %v", err)
	}
	return token, nil
}

type User struct {
	Username     string
	email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

func GetUserInfo(username string) (User, error) {
	user := User{}
	row, err := mydb.DBConn().Prepare("select user_name,signup_at from tbl_user where user_name=? limit 1")
	if err != nil {
		if err == sql.ErrNoRows {
			// 处理没有找到结果的情况
			return user, fmt.Errorf("没有找到对应的UserInfo")
		}
		return user, fmt.Errorf(err.Error())
	}
	defer row.Close()
	err = row.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		return user, err
	}

	return user, nil
}
