package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/aria?charset=utf8")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	db.SetMaxIdleConns(1000)
	err = db.Ping()
	if err != nil {
		log.Fatalln("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
	fmt.Println("------------------------success connect to airia!!!------------------------")

}

func DBConn() *sql.DB {
	return db
}

func ParseRows(rows *sql.Rows) []map[string]interface{} {
	cols, _ := rows.Columns()
	scanArgs := make([]interface{}, len(cols))
	values := make([]interface{}, len(cols))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		checkErr(err)

		for i, col := range values {
			if col != nil {
				record[cols[i]] = col
			}
		}
		records = append(records, record)
	}
	return records
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
