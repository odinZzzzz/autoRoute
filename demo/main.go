package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/odinZzzzz/autoRoute/demo/DAO"
	"github.com/odinZzzzz/autoRoute/demo/handler"
	"github.com/odinZzzzz/autoRoute/demo/remote"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

//go:embed static
var staticFs embed.FS

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域请求
	},
}

func main() {
	var g errgroup.Group
	DAO.InitMongo()
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      startServer(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server01.ListenAndServe()
	})

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      startRpcServer(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
func startServer() http.Handler {

	r := gin.New()
	r.Use(gin.Recovery())
	//设置静态资源目录
	//fads, _ := fs.Sub(staticFs, "static")
	//r.StaticFS("/", http.FS(fads))

	//r.SetTrustedProxies([]string{"127.0.0.1"})
	//添加ws监听
	addWsListener(r)
	handler.InitHandler()
	aRoute := handler.InitHandler()

	r.Use(aRoute.RouteMid)

	fmt.Println("服务启动 :8080")
	return r
}
func startRpcServer() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())

	//添加ws监听
	addWsListener(r)
	aRoute := remote.InitRemote()
	r.Use(aRoute.RouteMid)
	return r
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
