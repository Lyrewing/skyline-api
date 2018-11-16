package user

import (
	"skyline-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userService *services.UserService

func init() {
	userService = &services.UserService{}
}
func List(context *gin.Context) {
	index, _ := strconv.Atoi(context.Query("pageIndex"))
	size, _ := strconv.Atoi(context.Query("pageSize"))
	_, users := userService.GetUsers(index, size)
	context.JSON(200, users)
}
