package autoRoute

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/odinZzzzz/autoRoute/tool"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"sync"
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
	handlerName, methodName := parts[1], parts[2]
	s_time := time.Now()
	handler := AutoHandlerMap[handlerName]
	var res interface{}
	if handler, ok := handler[methodName]; ok {
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
		// 定义一个等待组，用于等待协程完成
		resultChan := make(chan interface{})
		var wg sync.WaitGroup
		var queData = &tool.ReqData{
			Open:        true,
			HandlerName: handlerName,
			MethodName:  methodName,
			Param:       args,
			ProtoData:   newData,
			ResultChan:  resultChan,
		}
		//todo 加入请求队列
		a.QueueMid(queData)
		//todo 在此等待数据返回 超时就提前返回

		wg.Add(1)
		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()
		// 启动协程
		go func() {
			defer wg.Done()
			defer close(resultChan)
			select {
			case <-ctx.Done(): // 检测上下文是否超时或被取消
				LogDebug("任务协程超时或被取消，返回")
				res = map[string]any{
					"code": 888,
					"msg":  "接口繁忙,请重试",
				}
				queData.Open = false
				return
			case res = <-resultChan: // 收到返回消息
				LogDebug("协程执行完成,返回前端数据")
				return
			}
		}()

		// 等待协程完成
		wg.Wait()
		////如果存在接口 则尝试执行当前接口的预处理函数
		//checkPreRes := a.runHandlerPre(handlerName, args)
		//if checkPreRes {
		//	if RouteOpt.UseProto {
		//		res = handler.Func([]reflect.Value{newData})
		//	} else {
		//		res = handler.Func(args)
		//	}
		//} else {
		//	res = map[string]any{
		//		"code": 400,
		//		"msg":  "接口拒绝访问",
		//	}
		//}

	} else {
		//调用失败则返回404 code
		res = map[string]any{
			"code": 404,
			"msg":  "接口不存在",
		}

	}
	dur := time.Since(s_time)
	if dur > 100*time.Millisecond {
		LogWarn(fmt.Sprintf("接口访问耗时超过100ms /%s/%s 耗时%s 超过 ", handlerName, methodName, dur))
	}
	LogDebug(fmt.Sprintf("接收到接口/%s/%s 耗时%s ", handlerName, methodName, dur))
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
	handlerName, methodName := parts[0], parts[1]
	s_time := time.Now()
	handler := AutoHandlerMap[handlerName]
	if handler, ok := handler[methodName]; ok {

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
		//todo ws请求 加入请求队列
		var queData = &tool.ReqData{
			Open:        true,
			HandlerName: handlerName,
			MethodName:  methodName,
		}
		a.QueueMid(queData)
		//如果存在接口 则尝试执行当前接口的预处理函数
		checkPreRes := a.RunHandlerPre(handlerName, []reflect.Value{reflect.ValueOf(map[string]any{})})
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
		LogWarn(fmt.Sprintf("接口访问耗时超过100ms /%s/%s 耗时%s 超过 ", handlerName, methodName, dur))
	}
	LogDebug(fmt.Sprintf("接收到接口/%s/%s 耗时%s ", handlerName, methodName, dur))

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
func (a *AutoRoute) RunHandlerPre(handlerName string, args []reflect.Value) bool {
	handler := AutoHandlerMap[handlerName]
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

var maxQue = 10

func (a *AutoRoute) QueueMid(req *tool.ReqData) interface{} {
	var result interface{}
	//var que = tool.GetQueue(req.HandlerName)
	// 寻找队列池
	//if que.Size() >= maxQue {
	//	result = gin.H{
	//		"code": 886,
	//		"msg":  "接口队列繁忙,稍后重试",
	//	}
	//	return result
	//}
	//队列池堆积后显示拥挤
	//que.Enqueue(req)
	reqChannel <- req
	return result

}
