package autoRoute

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

// 拦截所有请求的中间件
func (a *AutoRoute) RouteMid(c *gin.Context) {
	// 切分 URL 路径
	urlPath := c.Request.URL.Path
	parts := strings.Split(urlPath, "/")
	if len(parts) < 3 {
		return
	}
	handlerName, mthodName := parts[1], parts[2]
	// chuli1
	a.log(fmt.Sprintf("接收到接口 接口名：%s 函数名：%s ", handlerName, mthodName))
	handler := autoHandlerMap[handlerName]
	var res interface{}
	if handler, ok := handler[mthodName]; ok {
		args := []reflect.Value{
			reflect.ValueOf(mergePara(c)),
		}
		// 调用函数成功则返回正确值
		res = handler.(func(args []reflect.Value) interface{})(args)
	} else {
		//调用失败则返回404 code
		res = gin.H{
			"code": 404,
			"msg":  "接口不存在",
		}

	}
	SendMsg(res, c)

}

// 定制化路由  通过 handler 结构体匹配到接口
type BaseHandler struct {
	HandlerName string
}

func SendMsg(msg interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
	})
	c.Abort()
}
