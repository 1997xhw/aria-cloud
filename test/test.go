package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//filehash := "368f6fcb4026f43218140ac0d8fade97bb18127e"
	//fileMeta := meta.GetFileMeta(filehash)
	//marshal, err := json.Marshal(fileMeta)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//fmt.Printf("%s\n", marshal)
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/aria?charset=utf8")
	if err != nil {
		panic(err)
		return
	}
	err = db.Ping()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("------------------------success connect to airia!!!")
}
