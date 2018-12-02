package routers

import (
	"net/http"
	"skyline-api/controllers/user"
	"skyline-api/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	//注册Consul服务
	
}

func Route(router *gin.Engine) {
	apiPrefix := "api"

	//心跳
	router.GET("/health", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
	})
	api := router.Group(apiPrefix, middleware.RefreshToken)
	api.GET("/user/list", middleware.APIStatsD, user.List)
	api.POST("/user/add", middleware.APIStatsD, user.Add)

}
