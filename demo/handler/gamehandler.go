package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/odinZzzzz/autoRoute/autoRoute"
)

type gameHandler struct {
	autoRoute.AutoHandler
}

func (c gameHandler) Init() {
	println(c.HandlerName)
}

func (c gameHandler) Login(msg interface{}) interface{} {

	return c.Suc(gin.H{
		"uid":      10000001,
		"nickname": "芥末",
		"msg":      msg,
	})
}
