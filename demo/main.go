package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odinZzzzz/autoRoute/demo/handler"
	"io/fs"
	"net/http"
)

//go:embed static
var staticFs embed.FS

func main() {
	r := gin.Default()
	//设置静态资源目录
	fads, _ := fs.Sub(staticFs, "static")
	r.StaticFS("/", http.FS(fads))

	handler.InitHandler()

	aRoute := handler.InitHandler()
	r.Use(aRoute.RouteMid)
	fmt.Println("服务启动成功 接口启动 http://127.0.0.1:8080")
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
