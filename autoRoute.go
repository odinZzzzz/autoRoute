package autoRoute

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var autoHandlerMap map[string]map[string]handlerData = make(map[string]map[string]handlerData)
var RouteOpt = RouteOption{
	Debug:    false,
	UseProto: false,
}

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

type StartOption struct {
	Port int
	Host string

	InitHandler func(route *AutoRoute)
	Option      RouteOption
}

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
	aRoute := AutoRoute{}
	option.InitHandler(&aRoute)
	addr := fmt.Sprintf("%s:%d", option.Host, option.Port)
	r.Use(aRoute.RouteMid)
	server01 := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	DebugLog(fmt.Sprintf("autoRouteServer%s 启动成功", addr))
	DebugLog("\n                _        _____                             _____ _             _           _ \n     /\\        | |      / ____|                           / ____| |           | |         | |\n    /  \\  _   _| |_ ___| (___   ___ _ ____   _____ _ __  | (___ | |_ __ _ _ __| |_ ___  __| |\n   / /\\ \\| | | | __/ _ \\\\___ \\ / _ \\ '__\\ \\ / / _ \\ '__|  \\___ \\| __/ _` | '__| __/ _ \\/ _` |\n  / ____ \\ |_| | || (_) |___) |  __/ |   \\ V /  __/ |     ____) | || (_| | |  | ||  __/ (_| |\n /_/    \\_\\__,_|\\__\\___/_____/ \\___|_|    \\_/ \\___|_|    |_____/ \\__\\__,_|_|   \\__\\___|\\__,_|\n                                                                                             \n                                                                                             ")
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	return &aRoute
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域请求
	},
}

func addWsListener(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// 在这里编写你的WebSocket逻辑
		// 例如，接收和发送消息
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				conn.Close()
				break
			}

			err = conn.WriteMessage(messageType, p)
			if err != nil {
				conn.Close()
				break
			}
		}
	})
}
