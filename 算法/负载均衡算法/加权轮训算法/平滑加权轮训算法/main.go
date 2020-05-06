package main

import (
	trains2 "designMode/算法/负载均衡算法/加权轮训算法/平滑加权轮训算法/trains"
	"fmt"
)

func main() {
	trains := trains2.NewRobinMessage()
	for i := 0; i < 5; i++ {
		fmt.Println(trains.GetServer())
	}
}
