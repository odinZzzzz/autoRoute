package tool

import "reflect"

type ReqData struct {
	Open        bool
	HandlerName string
	MethodName  string
	Param       []reflect.Value
	ResultChan  chan interface{}
	ProtoData   reflect.Value
}

// Queue 队列结构体
type Queue struct {
	items []*ReqData // 使用切片存储队列元素
}

// Enqueue 入队操作
func (q *Queue) Enqueue(item *ReqData) {
	q.items = append(q.items, item)
}

// Dequeue 出队操作
func (q *Queue) Dequeue() (*ReqData, bool) {
	if len(q.items) == 0 {
		return &ReqData{}, false // 队列为空
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Size 返回 Queue 的大小
func (q *Queue) Size() int {
	return len(q.items)
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// 声明一个 map，用于存储队列
var queueMap = make(map[string]*Queue)

func GetQueueMap() map[string]*Queue {
	return queueMap
}
func GetQueue(key string) *Queue {
	queue, exists := queueMap[key]
	if !exists {
		queue = &Queue{}      // 创建一个新的队列
		queueMap[key] = queue // 将新队列存入 map
	}
	return queue
}
