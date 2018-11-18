package middleware

import (
	"github.com/gin-gonic/gin"
)

func Trace(context *gin.Context) {

	context.Next()

}
