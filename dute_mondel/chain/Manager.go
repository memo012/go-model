package chain

// 主管 经理 总监 总经理
// 	A  	B  	C 	 D
type Manager interface {
	SetNext(handle Manager) // 设置下一个管理者
	HandleRequest(money int)  // 处理
}
