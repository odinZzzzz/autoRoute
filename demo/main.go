package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odinZzzzz/autoRoute/demo/handler"
)

func main() {
	r := gin.Default()
	handler.InitHandler()

	aRoute := handler.InitHandler()
	r.Use(aRoute.RouteMid)
	fmt.Println("http://127.0.0.1:8080")
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
