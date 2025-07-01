package tool

import "sync"

type HandleCall struct {
	handleList []string
}

func (c *HandleCall) BindHandle(handle string) {
	c.handleList = append(c.handleList, handle)
}
func (c *HandleCall) Step(Fnc func(r *ReqData) interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, handle := range c.handleList {
		for {
			queue := GetQueue(handle)
			reqData, suc := queue.Dequeue()
			if !suc || reqData.Open == false {
				break
			}
			var res = Fnc(reqData)
			reqData.ResultChan <- res

		}

	}
}
