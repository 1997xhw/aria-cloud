package mysql

import (
	ini "aria-cloud/lib"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func InitDB() {
	conf := ini.LoadServerConfig()
	var err error
	param := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName)
	DB, err = gorm.Open("mysql", param)
	if err != nil {
		log.Fatal(2, err)
	}
	fmt.Println("this")
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName

	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	fmt.Println("########################################## database init on port ", conf.Host)

}
