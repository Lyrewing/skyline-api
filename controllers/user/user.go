package user

import (
	"github.com/gin-gonic/gin"
	"skyline-api/models"
	"skyline-api/services"
)

var userService *services.UserService

func init() {
	userService = &services.UserService{}
}
func List(context *gin.Context) {
	users := &[]models.User{}
	userService.GetUsers(1, 10, users)
	context.JSON(200, users)
}
