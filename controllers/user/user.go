package user

import (
	"strconv"
	"skyline-api/models"
	"skyline-api/services"

	"github.com/gin-gonic/gin"
)

var userService *services.UserService

func init() {
	userService = &services.UserService{}
}
func List(context *gin.Context) {
	index,_:= strconv.Atoi(context.Query("pageIndex"));
	size,_:= strconv.Atoi(context.Query("pageSize"))
	users := &[]models.User{}
	userService.GetUsers(index, size, users)
	context.JSON(200, users)
}
