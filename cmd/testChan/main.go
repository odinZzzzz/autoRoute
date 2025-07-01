package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个通道类型，用于传递任务
type Task struct {
	Data int
}

// 定义一个协程处理函数
func worker(id int, taskChan <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done() // 在协程结束时通知 WaitGroup

	for task := range taskChan {
		fmt.Printf("Worker %d is processing task with data: %d\n", id, task.Data)
		time.Sleep(time.Second) // 模拟处理时间
	}
}

func main() {
	// 创建一个任务通道
	taskChan := make(chan Task, 10)

	// 创建一个 WaitGroup，用于等待所有协程完成
	var wg sync.WaitGroup

	// 启动多个协程
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, taskChan, &wg)
	}

	// 向通道发送任务
	for i := 1; i <= 10; i++ {
		taskChan <- Task{Data: i}
	}

	// 关闭通道
	close(taskChan)

	// 等待所有协程完成
	wg.Wait()
	fmt.Println("All tasks are processed.")
}
