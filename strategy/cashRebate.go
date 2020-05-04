package strategy

type cashRebate struct {
	Rebate float64 // 折扣
}

// 对外暴露接口
func NewCashRebate(rebate float64) *cashRebate {
	instance := new(cashRebate)
	instance.Rebate = rebate
	return instance
}

func (c *cashRebate) AcceptMoney(money float64) float64 {
	return money * c.Rebate
}