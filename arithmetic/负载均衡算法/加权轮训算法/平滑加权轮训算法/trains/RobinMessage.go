package trains

import (
	"fmt"
	"sync"
)
/*
	思路：
		首先需要初始化一个数据结构，包含三个字段，分别为IP，权重，动态值
		然后遍历循环得到动态值最大的服务器，即可目标服务器
		然后让目标服务器减掉权重之和
		再遍历循环即可
 */
type RobinMessage struct {
	list []WeightTrains
}

var once sync.Once

// 对外暴露函数
func NewRobinMessage() *RobinMessage {
	return &RobinMessage{list: make([]WeightTrains, 0)}
}

// 初始化服务器
func (t *RobinMessage) ini() {
	weightTrains := new(WeightTrains)
	weightTrains.SetWeightTrains("A", 5, 0)
	t.list = append(t.list, *weightTrains)
	weightTrains.SetWeightTrains("B", 1, 0)
	t.list = append(t.list, *weightTrains)
	weightTrains.SetWeightTrains("C", 1, 0)
	t.list = append(t.list, *weightTrains)
}

// 计算权重之和
func (t *RobinMessage) sum() int64 {
	var count int64 = 0
	for _, v := range t.list {
		count += v.weight
	}
	return count
}

// 获取目标服务器
func (t *RobinMessage) GetServer() string {
	once.Do(func() {
		// 初始化 服务器值
		t.ini()
	})
	// 计算权重和
	countSum := t.sum()
	// 计算currentWeight += weight
	len := len(t.list)
	for i := 0; i < len; i++ {
		name := t.list[0].name
		weight := t.list[0].weight
		currentWeight := t.list[0].currentWeight
		t.list = append(t.list[:0], t.list[1:]...)
		weightTrains := new(WeightTrains)
		weightTrains.SetWeightTrains(name, weight, currentWeight+weight)
		t.list = append(t.list, *weightTrains)
	}

	for _, v := range t.list {
		fmt.Printf("%v %v %v \n", v.name, v.weight, v.currentWeight)
	}

	// max(currentWeight)
	var maxWeight WeightTrains
	index := 0
	for i, v := range t.list {
		if maxWeight.currentWeight < v.currentWeight {
			maxWeight = v
			index = i
		}
	}
	name := t.list[index].name
	weight := t.list[index].weight
	currentWeight := t.list[index].currentWeight
	t.list = append(t.list[:index], t.list[index+1:]...)
	t.list = append(t.list, WeightTrains{
		name:          name,
		weight:        weight,
		currentWeight: currentWeight - countSum,
	})
	return maxWeight.name
}
