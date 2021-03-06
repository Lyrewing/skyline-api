package middleware

import (
	"log"
	"skyline-api/conf"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/alexcesaro/statsd.v2"
)

var stastdClient *statsd.Client

func instance(client *statsd.Client) *statsd.Client {
	var err error
	if client != nil {
		return client
	}
	if client, err = statsd.New(statsd.Address(conf.AppSettings.URL), statsd.Prefix("skyline")); err != nil {
		log.Fatal(err)
	}
	return client
}
func APIStatsD(context *gin.Context) {
	t := time.Now()
	apiName := context.Request.URL.Path
	//继续
	context.Next()
	//时间间隔
	duration := uint(time.Since(t) / 1e6)

	//监控功能
	//创建一个单利的对象
	stastdClient = instance(stastdClient)
	stastdClient.Timing(apiName, duration)

}
