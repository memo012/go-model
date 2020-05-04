package chain

import "fmt"

type DepManager struct {
	Next Manager
}
func NewDepManager() *DepManager {
	return &DepManager{Next: nil}
}
// 设置下一个处理器
func (pm *DepManager) SetNext(handle Manager) {
	pm.Next = handle
}

func (pm *DepManager) HandleRequest(money int) {
	if money < 5000 {
		fmt.Println("2. DepManager handle request ...")
		return
	}
	fmt.Println("3. DepManager not handle request ...")
	// 让下一个handler处理
	if pm.Next != nil {
		pm.Next.HandleRequest(money)
	}
}