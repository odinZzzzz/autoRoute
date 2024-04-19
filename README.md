<div align="center">
<br/>
<br/>
  <h1 align="center">
    AutoRoute
  </h1>
</div>

#### 项目简介
>  Gin框架的路由中间件
>  极速创建 极速开发
>  致力于接口开发尽可能少的修改文件
>  纵享丝滑




#### 运行项目

``` gameHandler.go
import (
	"autoRoute"
	"github.com/gin-gonic/gin"
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
```
```bash
go mod tidy
go run main.go
```
