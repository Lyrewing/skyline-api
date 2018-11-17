package user

import (
	"log"
	"skyline-api/models"
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
	total, users := userService.GetUsers(index, size)
	context.JSON(200, gin.H{"pageindex": index, "pagesize": size, "total": total, "data": users})
}

func Add(context *gin.Context) {
	var user models.User
	if err := context.BindJSON(&user); err != nil {
		log.Fatalln("格式转换失败")
		context.JSON(400, gin.H{"msg": "请求数据格式有误"})
		return
	}
	userService.Add(&user)
	context.JSON(200, gin.H{"msg": "添加成功"})
}
