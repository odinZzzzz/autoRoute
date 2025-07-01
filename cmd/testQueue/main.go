package main

import (
	"fmt"
	"github.com/odinZzzzz/autoRoute/tool"
	"time"
)

func main() {
	// 定义一个函数
	var que = tool.GetQueue("game")
	for i := 0; i < 100; i++ {
		go que.Enqueue(i)
	}
	//for i := 0; i < 101; i++ {
	//	go que.Dequeue()
	//}
	time.Sleep(10 * time.Second)
	for {
		a, b := que.Dequeue()
		if b {
			fmt.Println(a)
		}
	}
}
