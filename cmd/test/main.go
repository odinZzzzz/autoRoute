package main

import (
	"github.com/gin-gonic/gin"
	"github.com/odinZzzzz/autoRoute"
)

func main() {
	// 定义一个函数
	autoRoute.StartServer(autoRoute.StartOption{
		Port:        8080,
		InitHandler: InitHandler,
		Option: autoRoute.RouteOption{
			Debug:    true,
			UseProto: false,
		},
	})
}
func InitHandler(r *autoRoute.AutoRoute) {
	//创建接口并注册接口
	r.Register("demo", demoHandler{AutoHandler: autoRoute.AutoHandler{}})
}

type demoHandler struct {
	autoRoute.AutoHandler
}
type paramDefine struct {
	A   string
	Gin *gin.Context
}

func (c demoHandler) Test(msg map[string]interface{}) interface{} {
	loginParam := autoRoute.FormatParam(msg, paramDefine{})
	//参数中会注入gin.Context 置空后返回
	loginParam.Gin = nil

	return c.Suc(gin.H{
		"uid":      10000001,
		"nickname": "芥末",
		"msg":      loginParam,
		"msg1":     loginParam.A,
	})
}
