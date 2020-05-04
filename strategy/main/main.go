package main

import (
	"designMode/strategy"
	"fmt"
)

func main() {
	money := 100.0
	cc := strategy.NewCashContext("打八折")
	money = cc.GetMoney(money)
	fmt.Println("100打8折实际金额为", money)

	money = 199
	cc = strategy.NewCashContext("满一百返20")
	money = cc.GetMoney(money)
	fmt.Println("199满一百返20实际金额为", money)

	money = 199
	cc = strategy.NewCashContext("无折扣")
	money = cc.GetMoney(money)
	fmt.Println("199没有折扣实际金额为", money)

}