package main

import "fmt"

// A X B
// X  操作
// A B 操作数

func main() {
	var fac OperatorFactory = &PlusOperatorFactory{}
	op := fac.Create()
	op.SetLeft(12)
	op.SetRight(12)
	fmt.Println(op.Result())
}
