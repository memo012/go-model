package chain

import (
	"fmt"
	"math"
)

type GeneralManager struct {
	Next Manager
}
func NewGeneralManager() *GeneralManager {
	return &GeneralManager{Next: nil}
}
// 设置下一个处理器
func (pm *GeneralManager) SetNext(handle Manager) {
	pm.Next = handle
}

func (pm *GeneralManager) HandleRequest(money int) {
	fmt.Println(money , math.MaxInt8, math.MaxInt64)
	if money < math.MaxInt64 {
		fmt.Println("3. GeneralManager handle request ...")
		return
	}
	fmt.Println("3. GeneralManager not handle request ...")
	// 让下一个handler处理
	if pm.Next != nil {
		pm.Next.HandleRequest(money)
	}
}
