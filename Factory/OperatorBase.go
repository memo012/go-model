package main

type OperatorBase1 struct {
	left, right int
}

// 赋值
func (op *OperatorBase1) SetLeft (left int){
	op.left = left
}
func (op *OperatorBase1) SetRight (right int){
	op.right = right
}
