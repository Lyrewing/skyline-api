package user

import (
	"skyline-api/models"
	"skyline-api/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func List(context *gin.Context) {
	users := []models.User{}
	users = append(users, models.User{Age: 25, Name: "fengzhanyuan", Sex: 1, ID: 2})

	randnum := utils.GenerateRandomNumber(10, 200)
	time.Sleep(time.Millisecond * time.Duration(randnum))
	context.JSON(200, users)
}
