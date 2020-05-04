package strategy

type cashContext struct {
	strategy cashSuper
}

func NewCashContext(cashType string) *cashContext {
	c := new(cashContext)
	switch cashType {
	case "打八折":
		c.strategy = NewCashRebate(0.8)
	case "满一百返20":
		c.strategy = NewCashReturn(100.0, 20.0)
	default:
		c.strategy = NewCashNormal()
	}
	return c
}

// 策略完成后  可以直接调用策略函数
func (c *cashContext) GetMoney(money float64) float64 {
	return c.strategy.AcceptMoney(money)
}
