package autoRoute

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
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
	s_time := time.Now()
	handler := autoHandlerMap[handlerName]
	var res interface{}
	if handler, ok := handler[mthodName]; ok {
		args := []reflect.Value{
			reflect.ValueOf(mergePara(c)),
			//newPerson := reflect.New(paramType)
		}
		//
		paramType := handler.Define
		newData := reflect.New(paramType)
		if RouteOpt.UseProto {
			var bytesData []byte
			// Read the Body content
			if c.Request.Body != nil {
				bytesData, _ = ioutil.ReadAll(c.Request.Body)
			}
			err := proto.Unmarshal(bytesData, newData.Interface().(proto.Message))
			if err != nil {
				fmt.Println(err)
				res = map[string]any{
					"code": 401,
					"msg":  "数列化参数错误",
				}
				return
			}
		}
		//如果存在接口 则尝试执行当前接口的预处理函数
		checkPreRes := a.runHandlerPre(handlerName, args)
		if checkPreRes {
			if RouteOpt.UseProto {
				res = handler.Func([]reflect.Value{newData})
			} else {
				res = handler.Func(args)
			}
		} else {
			res = map[string]any{
				"code": 400,
				"msg":  "接口拒绝访问",
			}
		}

	} else {
		//调用失败则返回404 code
		res = map[string]any{
			"code": 404,
			"msg":  "接口不存在",
		}

	}
	dur := time.Since(s_time)
	if dur > 100*time.Millisecond {
		LogWarn(fmt.Sprintf("接口访问耗时超过100ms /%s/%s 耗时%s 超过 ", handlerName, mthodName, dur))
	}
	LogDebug(fmt.Sprintf("接收到接口/%s/%s 耗时%s ", handlerName, mthodName, dur))
	SendMsg(res, c)

}

type wsParam struct {
	Handle string `json:"handle,omitempty"`
	ReqId  int    `json:"req_id,omitempty"`
	Data   []byte `json:"data,omitempty"`
}

// 拦截所有请求的中间件
func (a *AutoRoute) RouteWSMid(conn *websocket.Conn, msgType int, p []byte) {
	var res interface{}
	resType := websocket.TextMessage
	// 将消息转换为JSON
	var param wsParam // 你可以根据实际情况定义更具体的结构体
	if err := json.Unmarshal(p, &param); err != nil {
		LogWarn(fmt.Sprintf("WS 参数格式化失败 JSON Unmarshal error:", err))
		SendWSMsg(conn, resType, res, param)
		return
	}

	parts := strings.Split(param.Handle, "/")
	if len(parts) < 2 {
		SendWSMsg(conn, resType, res, param)
		return
	}
	handlerName, mthodName := parts[0], parts[1]
	s_time := time.Now()
	handler := autoHandlerMap[handlerName]
	if handler, ok := handler[mthodName]; ok {

		paramType := handler.Define
		newData := reflect.New(paramType)
		if RouteOpt.UseProto {

			err := proto.Unmarshal(param.Data, newData.Interface().(proto.Message))
			if err != nil {
				fmt.Println(err)
				res = map[string]any{
					"code": 401,
					"msg":  "数列化参数错误",
				}
				SendWSMsg(conn, resType, res, param)
				return
			}
		}
		//如果存在接口 则尝试执行当前接口的预处理函数
		checkPreRes := a.runHandlerPre(handlerName, []reflect.Value{reflect.ValueOf(map[string]any{})})
		if checkPreRes {
			if RouteOpt.UseProto {
				res = handler.Func([]reflect.Value{newData})
			} else {
				res = handler.Func([]reflect.Value{reflect.ValueOf(param.Data)})
			}
		} else {
			res = map[string]any{
				"code": 400,
				"msg":  "接口拒绝访问",
			}
		}

	} else {
		//调用失败则返回404 code
		res = map[string]any{
			"code": 404,
			"msg":  "接口不存在",
		}

	}
	dur := time.Since(s_time)
	if dur > 100*time.Millisecond {
		LogWarn(fmt.Sprintf("接口访问耗时超过100ms /%s/%s 耗时%s 超过 ", handlerName, mthodName, dur))
	}
	LogDebug(fmt.Sprintf("接收到接口/%s/%s 耗时%s ", handlerName, mthodName, dur))

	SendWSMsg(conn, websocket.TextMessage, res, param)

}

type wsRepose struct {
	ReqId int `json:"req_id,omitempty"`
	Data  any `json:"data,omitempty"`
}

func SendWSMsg(conn *websocket.Conn, messageType int, res interface{}, param wsParam) {
	resData, err := json.Marshal(wsRepose{
		ReqId: param.ReqId,
		Data:  res,
	})
	if err != nil {
		LogErr(err)
		return
	}
	err = conn.WriteMessage(messageType, resData)
	if err != nil {
		LogErr(err)
		return
	}

}

// 拦截所有请求的中间件 true 则通过预处理 false 则拦截接口返回err
func (a *AutoRoute) runHandlerPre(handlerName string, args []reflect.Value) bool {
	handler := autoHandlerMap[handlerName]
	var res bool = true
	if handler, ok := handler["HandlerPre"]; ok {
		// 调用函数成功则返回正确值
		res = handler.Func(args).(bool)
	}
	return res
}

// 定制化路由  通过 handler 结构体匹配到接口
type BaseHandler struct {
	HandlerName string
}

func SendMsg(msg interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, msg)
	c.Abort()
}
