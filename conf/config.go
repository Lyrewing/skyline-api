package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"skyline-api/utils"
)

type config struct {
	SiteName                                string
	Env                                     string
	URL                                     string
	Prefix                                  string
	ConnectString                           string
	Register_Address                        string
	Register_Service                        string
	Register_Tags                           []string
	Register_Port                           int
	Register_DeregisterCriticalServiceAfter string
	Register_Interval                       string
}

var AppSettings config

//构造函数
func init() {
	var bytes []byte
	var err error

	dir := "./config.development.json"
	if env, e := os.LookupEnv("GOLANG_ENVIRONMENT"); e && env == "product" {
		dir = "./config.json"
	}

	//读取文件
	if bytes, err = ioutil.ReadFile(dir); err != nil {
		fmt.Println("ReadFile:", err.Error())
		os.Exit(-1)
	}

	//对Config文件进行解析
	bytes = []byte(utils.ConfigTemplateRender(string(bytes)))

	//反序列化文件为map
	if err = json.Unmarshal(bytes, &AppSettings); err != nil {
		fmt.Println("反序化失败")
		os.Exit(-1)
	}

}
