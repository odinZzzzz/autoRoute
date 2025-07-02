package autoRoute

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/odinZzzzz/autoRoute/tool"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var AutoHandlerMap map[string]map[string]handlerData = make(map[string]map[string]handlerData)
var RouteOpt = RouteOption{
	Debug:    false,
	UseProto: false,
}
var App *AutoRoute

type AutoRoute struct {
}
type RouteOption struct {
	Debug    bool // 是否开启Debug调试详细日志
	UseProto bool //接口接参使用protoBuf
}
type AutoHandler struct {
	HandlerName string
}

func (a *AutoRoute) Register(group string, handler interface{}) {
	a.registerHandler(group, handler)
}

func (a *AutoRoute) FncCall(r *tool.ReqData) interface{} {
	var res interface{}
	handler := AutoHandlerMap[r.HandlerName]
	if method, ok := handler[r.MethodName]; ok {
		res = method.Func(r.Param)
	} else {
		res = map[string]any{
			"code": 404,
			"msg":  "接口不存在",
		}
	}
	return res
}

type StartOption struct {
	Port int
	Host string

	InitHandler func(route *AutoRoute)
	Option      RouteOption
}

var reqChannel = make(chan *tool.ReqData, 10)

func StartServer(option StartOption) *AutoRoute {

	RouteOpt = option.Option

	if !RouteOpt.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	var g errgroup.Group
	r := gin.New()
	r.Use(gin.Recovery())
	//设置静态资源目录
	//handler.InitHandler()
	addWsListener(r)
	aRoute := AutoRoute{}
	option.InitHandler(&aRoute)
	addr := fmt.Sprintf("%s:%d", option.Host, option.Port)
	//r.Use(aRoute.QueueMid)
	r.Use(aRoute.RouteMid)
	server01 := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	go startQueueWatch()
	LogDebug(fmt.Sprintf("autoRouteServer%s 启动成功", addr))
	LogDebug("\n                _        _____                             _____ _             _           _ \n     /\\        | |      / ____|                           / ____| |           | |         | |\n    /  \\  _   _| |_ ___| (___   ___ _ ____   _____ _ __  | (___ | |_ __ _ _ __| |_ ___  __| |\n   / /\\ \\| | | | __/ _ \\\\___ \\ / _ \\ '__\\ \\ / / _ \\ '__|  \\___ \\| __/ _` | '__| __/ _ \\/ _` |\n  / ____ \\ |_| | || (_) |___) |  __/ |   \\ V /  __/ |     ____) | || (_| | |  | ||  __/ (_| |\n /_/    \\_\\__,_|\\__\\___/_____/ \\___|_|    \\_/ \\___|_|    |_____/ \\__\\__,_|_|   \\__\\___|\\__,_|\n                                                                                             \n                                                                                             ")
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
	App = &aRoute
	return &aRoute
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域请求
	},
}
var demoCall = tool.HandleCall{}

// 启动queue取出协程
func startQueueWatch() {
	demoCall.BindHandle("demo")
	for req := range reqChannel {
		req.ResultChan <- App.FncCall(req)
		//select {
		//case req := <-reqChannel:
		//	App.FncCall(req)
		//}
		//wg := sync.WaitGroup{}
		//wg.Add(1)
		//go demoCall.Step(App.FncCall, &wg)
		//
		//wg.Wait()
	}

}

func addWsListener(r *gin.Engine) {
	LogDebug("Websocket 服务监听成功")
	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
			return
		}
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				err := conn.Close()
				if err != nil {
					LogErr(err)
				}
				break
			}

			App.RouteWSMid(conn, messageType, p)

		}
	})
}
