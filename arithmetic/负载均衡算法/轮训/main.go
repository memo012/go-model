package main

import "designMode/arithmetic/负载均衡算法/轮训/selects"

func main() {
	var sel *selects.SelectMessage
	for i := 0; i < 15; i++ {
		sel = new(selects.SelectMessage)
		sel.GetSelectStrategy()
	}
}

