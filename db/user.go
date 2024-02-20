package db

import (
	mydb "aria-cloud/db/mysql"
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
