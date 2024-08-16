package remote

import (
	"github.com/gin-gonic/gin"
	"github.com/odinZzzzz/autoRoute"
)

type gameRemote struct {
	autoRoute.AutoHandler
}

type paramDefine struct {
	A   int
	Gin *gin.Context
}

func (c gameRemote) Login(msg map[string]interface{}) interface{} {
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
