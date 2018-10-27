package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RefreshToken(context *gin.Context) {
	fmt.Println("cookie")
	context.Next()
}
