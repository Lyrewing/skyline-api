package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type config struct {
	SiteName      string
	Env           string
	URL           string
	Prefix        string
	ConnectString string
}

var AppSettings config

//构造函数
func init() {
	var bytes []byte
	var err error

	//读取文件
	if bytes, err = ioutil.ReadFile("./config.json"); err != nil {
		fmt.Println("ReadFile:", err.Error())
		os.Exit(-1)
	}

	//去除注释部分
	bytes = []byte(regexp.MustCompile(`/\*.*\*/`).ReplaceAllString(string(bytes), ""))

	//反序列化文件为map
	if err = json.Unmarshal(bytes, &AppSettings); err != nil {
		fmt.Println("反序化失败")
		os.Exit(-1)
	}
	if connectionStr := os.Getenv("MYSQL_DB_CONNECTION"); connectionStr != "" {
		AppSettings.ConnectString = connectionStr
	}
	if statsdUrl:=os.Getenv("STATSD_URL");statsdUrl!=""{
		AppSettings.URL = statsdUrl
	}



}
