package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
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
		log.Fatal(err)
		os.Exit(-1)
	}
}
