package selects

import (
	"fmt"
	"sync/atomic"
)

/**
负载均衡算法之一 轮训算法
*/

// 服务器存储地址
var list []string = make([]string, 0)
var index int64

func init() {
	list = append(list, "8080")
	list = append(list, "8081")
}

type SelectMessage struct {
}

func (s  *SelectMessage) GetSelectStrategy()  {
	value := index % int64(len(list))
	atomic.AddInt64(&index, 1)
	fmt.Println(list[value])
}

