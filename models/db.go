package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

//数据库连接
var DB *gorm.DB

func init() {

	//初始化数据库连接
	if db, err := gorm.Open("", ""); err == nil {
		db.LogMode(true)
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(30)
		DB = db
	} else {
		fmt.Println("数据库连接有误")
		log.Fatal(err)
		os.Exit(-1)
	}
}
