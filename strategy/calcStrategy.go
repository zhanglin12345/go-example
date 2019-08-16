package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	initPrice, _ := decimal.NewFromString("100")
	initAmount, _ := decimal.NewFromString("100")
	currentPrice, _ := decimal.NewFromString("95")
	totalMoney, _ := decimal.NewFromString("200000")
	finalRate := decimal.NewFromFloat(0.01)

	lose := initPrice.Sub(currentPrice).Mul(initAmount)
	initMoney := initPrice.Mul(initAmount)
	newAddMoney := lose.Div(finalRate).Sub(initMoney)
	newAddAmount := newAddMoney.Div(currentPrice)
	newPrice := newAddMoney.Add(initMoney).Div(initAmount.Add(newAddAmount))
	leftMoney := totalMoney.Sub(initMoney).Sub(newAddMoney)

	if newAddMoney.LessThanOrEqual(totalMoney) {
		fmt.Print("can buy ")
	} else {
		fmt.Print("have no money! cannot buy ")
	}

	fmt.Printf("%v with %v cost %v, left %v", newAddAmount.StringFixed(0), newPrice.StringFixed(2), newAddMoney.StringFixed(0), leftMoney.StringFixed(0))

}
