package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/odinZzzzz/autoRoute"
)

type gameHandler struct {
	autoRoute.AutoHandler
}
type paramDefine struct {
	A string `json:"a"`
}

func (c gameHandler) Login(msg map[string]interface{}) interface{} {
	loginParam := autoRoute.FormatParam(msg, paramDefine{})
	return c.Suc(gin.H{
		"uid":      10000001,
		"nickname": "芥末",
		"msg":      loginParam,
		"msg1":     loginParam.A,
	})
}
