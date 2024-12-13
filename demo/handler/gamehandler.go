package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/odinZzzzz/autoRoute"
	"github.com/odinZzzzz/autoRoute/demo/DAO"
	"github.com/odinZzzzz/autoRoute/demo/define/game"
)

type gameHandler struct {
	autoRoute.AutoHandler
}

// HandlerPre 接口预处理函数 如果入参 不携带指定参数就拒绝访问
func (c gameHandler) HandlerPre(msg map[string]interface{}) bool {
	checkRes := true
	if msg["token"] != "12306" {
		checkRes = false
	}
	return checkRes
}

func (c gameHandler) Login(msg *game.LoginDefine) interface{} {

	return c.Suc(gin.H{
		"uid":      10000001,
		"nickname": "芥末",
		"msg1":     msg.A,
	})
}
func (c gameHandler) TestDAO(msg map[string]interface{}) interface{} {
	data := DAO.BaseDAO{Uid: 10600000679}
	data.BaseDAO()
	return c.Suc(gin.H{
		"uid":      10000001,
		"nickname": "芥末",
	})
}
