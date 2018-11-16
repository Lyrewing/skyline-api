package main

import (
	"skyline-api/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	routers.Route(app)
	app.Run(":8081")

}
