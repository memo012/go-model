package strategy

// 普通情况 不打折
type cashNormal struct {
}

// 对外暴露接口
func NewCashNormal() *cashNormal {
	return new(cashNormal)
}

func (c *cashNormal) AcceptMoney(money float64) float64 {
	return money
}
