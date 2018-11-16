package repositories

import (
	"fmt"
	"log"
	"os"
	"skyline-api/conf"
	"skyline-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//数据库连接
var DB *gorm.DB

func init() {

	//初始化数据库连接
	if db, err := gorm.Open("mysql", conf.AppSettings.ConnectString); err == nil {
		db.LogMode(true)
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(30)
		//自动迁移
		db.AutoMigrate(&models.User{})

		DB = db
	} else {
		fmt.Println("数据库连接有误")
		log.Fatal(err)
		os.Exit(-1)
	}
}
