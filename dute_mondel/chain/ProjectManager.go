package chain

import "fmt"

type ProjectManager struct {
	Next Manager
}

func NewProjectManager() *ProjectManager {
	return &ProjectManager{Next: nil}
}

// 设置下一个处理器
func (pm *ProjectManager) SetNext(handle Manager) {
	pm.Next = handle
}

func (pm *ProjectManager) HandleRequest(money int) {
	if money < 1000 {
		fmt.Println("1. ProjectManager handle request ...")
		return
	}
	fmt.Println("1. ProjectManager not handle request ...")
	// 让下一个handler处理
	if pm.Next != nil {
		pm.Next.HandleRequest(money)
	}
}
