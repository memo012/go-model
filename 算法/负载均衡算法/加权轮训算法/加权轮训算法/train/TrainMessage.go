package train

import (
	"fmt"
	"math/rand"
)

// 加权轮训算法
/*
思路：
	1. 首先计算出服务器的权重之和  初始化一个权重 默认为0
	2. 遍历服务器，并初始权重和当前服务器权重进行比较
		如果大于的话，初始权重减掉当前服务器权重并再次遍历
		如果小于的话，该服务器为目标服务器
	3. 如果所有服务器的权重想等的话，采用轮训算法
*/

var list []TrainMessage

type TrainMessage struct {
	name   string // 服务器名称
	weight int32  // 权重
}

func (t *TrainMessage) SetName(name string) {
	t.name = name
}

func (t *TrainMessage) SetWeight(weight int32) {
	t.weight = weight
}

// 初始化服务器个数
func init() {
	trainMessage := new(TrainMessage)
	list = make([]TrainMessage, 0)
	trainMessage.name = "192.168.233.100"
	trainMessage.weight = 5
	list = append(list, *trainMessage)
	trainMessage.name = "192.168.233.101"
	trainMessage.weight = 3
	list = append(list, *trainMessage)
	trainMessage.name = "192.168.233.102"
	trainMessage.weight = 2
	list = append(list, *trainMessage)
}

// 服务器权重之和
func (t *TrainMessage) GetSum() int32 {
	var sum int32 = 0
	for _, value := range list {
		sum += value.weight
	}
	return sum
}

// 获取权重最大值 即 目标服务器
func (t *TrainMessage) getMax() string {
	var value int = rand.Intn(int(t.GetSum()))
	for _, v := range list {
		if value < int(v.weight) {
			return v.name
		}
		value -= int(v.weight)
	}
	// 到这里 说明 服务器的权重值一样 所以用轮训算法即可
	return ""
}

// 获取目标服务器
func (t *TrainMessage) GetTrainMessage()  {
	target := t.getMax()
	if target == "" {
		fmt.Println("服务器权重算法异常...")
		return
	}
	fmt.Println("目标服务器为:", target)
}


