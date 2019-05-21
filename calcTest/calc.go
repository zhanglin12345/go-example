package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/fatih/color"
)

var times int
var initMoney float64 = 200000 //initial money
const loseOdd = 40             //30%
const lossRatio float64 = 3    //5%
const interestMax int = 5      //0%~10%
const tradeTimes = 100         //trade 100 times

func do(money float64) float64 {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	win := random.Intn(1000)
	if win < loseOdd*10 {
		money *= 1 - lossRatio/100
	} else {
		result := random.Intn(interestMax * 10)
		interest := float64(result) / 1000
		money *= 1 + interest
		times++
	}
	return money
}

func main() {
	money := initMoney
	for i := 0; i <= tradeTimes; i++ {
		tmp := money
		money = do(money)
		if tmp > money {
			color.Green(strconv.FormatFloat(money, 'f', 6, 64))
		} else {
			color.Red(strconv.FormatFloat(money, 'f', 6, 64))
		}
	}
	fmt.Printf("win %.2f%% and money is %.2f%%", float64(times)/tradeTimes*100, money/initMoney*100)
}
