package main

import (
	"autoRoute/demo/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handler.InitHandler()

	aRoute := handler.InitHandler()
	r.Use(aRoute.RouteMid)

	r.Run(":8080")
}
