package main

import "designMode/dute_mondel/chain"

func main() {
	handler1 := chain.NewProjectManager()
	handler2 := chain.NewDepManager()
	handler3 := chain.NewGeneralManager()
	handler1.SetNext(handler2)
	handler2.SetNext(handler3)
	handler1.HandleRequest(4000)
}
