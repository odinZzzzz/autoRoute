package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/odinZzzzz/autoRoute"
)

type gameHandler struct {
	autoRoute.AutoHandler
}

// 接口预处理函数 如果入参 不携带指定参数就拒绝访问
func (c gameHandler) HandlerPre(msg map[string]interface{}) bool {
	checkRes := true
	if msg["token"] != "12306" {
		checkRes = false
	}
	return checkRes
}

type paramDefine struct {
	A   string
	Gin *gin.Context
}

func (c gameHandler) Login(msg map[string]interface{}) interface{} {
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
