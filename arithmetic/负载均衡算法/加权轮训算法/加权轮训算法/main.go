package main

import "designMode/arithmetic/负载均衡算法/加权轮训算法/加权轮训算法/train"

func main() {
	trainMessage := new(train.TrainMessage)
	for i := 0; i < 50; i++ {
		trainMessage.GetTrainMessage()
	}

}
