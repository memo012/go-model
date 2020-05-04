package strategy

type cashSuper interface {
	AcceptMoney(money float64) float64
}
