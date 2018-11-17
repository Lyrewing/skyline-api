package routers

import (
	"skyline-api/controllers/user"
	"skyline-api/middleware"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	apiPrefix := "api"
	api := router.Group(apiPrefix, middleware.RefreshToken)
	api.GET("/user/list", middleware.APIStatsD, user.List)
	api.POST("/user/add",middleware.APIStatsD,user.Add)
}
