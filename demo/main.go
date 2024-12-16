package main

import (
	"embed"
	"github.com/odinZzzzz/autoRoute"
	"github.com/odinZzzzz/autoRoute/demo/handler"
)

//go:embed res
var staticFs embed.FS

func main() {
	autoRoute.StartServer(autoRoute.StartOption{
		Port:        8080,
		InitHandler: InitHandler,
		Option: autoRoute.RouteOption{
			Debug:    true,
			UseProto: true,
		},
	})

}
func InitHandler(r *autoRoute.AutoRoute) {
	//创建接口并注册接口
	r.Register("game", handler.GameHandler{AutoHandler: autoRoute.AutoHandler{}})
}
