package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/odinZzzzz/autoRoute"
	"github.com/odinZzzzz/autoRoute/demo/define/game"
)

type GameHandler struct {
	autoRoute.AutoHandler
}

// HandlerPre 接口预处理函数 如果入参接入query 和 body json参数
func (c GameHandler) HandlerPre(msg map[string]interface{}) bool {
	checkRes := true
	//if msg["token"] != "12306" {
	//	checkRes = false
	//}
	return checkRes
}

func (c GameHandler) Test(msg *game.LoginDefine) interface{} {

	return c.Suc(gin.H{
		"uid":      10000001,
		"nickname": "芥末",
		"msg1":     msg.String(),
	})
}
