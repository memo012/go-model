package strategy

type cashReturn struct {
	MoneyCondition float64
	MoneyReturn float64
}

func NewCashReturn(moneyCondition, moneyReturn float64) *cashReturn {
	cash := new(cashReturn)
	cash.MoneyReturn = moneyReturn
	cash.MoneyCondition = moneyCondition
	return cash
}

func (c *cashReturn) AcceptMoney(money float64) float64 {
	if money >= c.MoneyCondition {
		moneyMinus := int(money / c.MoneyCondition)
		return money - float64(moneyMinus) * c.MoneyReturn
	}
	return money
}
