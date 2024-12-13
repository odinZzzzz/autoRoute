package handler

import (
	"github.com/odinZzzzz/autoRoute"
)

func InitHandler(opt autoRoute.RouteOption) *autoRoute.AutoRoute {
	//初始化配置
	autoRoute.RouteOpt = opt
	//创建接口并注册接口
	route := autoRoute.AutoRoute{}
	route.Register("game", gameHandler{AutoHandler: autoRoute.AutoHandler{}})
	return &route
}
