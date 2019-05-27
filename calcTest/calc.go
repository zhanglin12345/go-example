package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/fatih/color"
)

var times int
var initMoney float64 = 10000 //initial money
const loseOdd = 40            //40%
const lossRatio float64 = 5   //3%
const interestMax int = 10    //0%~5%
const tradeTimes = 10         //trade 100 times

func do(money float64) (float64, float64) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	win := random.Intn(1000)
	var odd float64
	if win < loseOdd*10 {
		odd = lossRatio / 100
		money *= 1 - odd
	} else {
		result := random.Intn(interestMax * 10)
		odd = float64(result) / 1000
		money *= 1 + odd
		times++
	}
	return money, odd
}

func main() {
	money := initMoney
	var odd float64
	for i := 0; i <= tradeTimes; i++ {
		tmp := money
		money, odd = do(money)
		if tmp > money {
			color.Green("%v		-%.2f%%", strconv.FormatFloat(money, 'f', 6, 64), odd*100)
		} else {
			color.Red("%v		%.2f%%", strconv.FormatFloat(money, 'f', 6, 64), odd*100)
		}
	}
	fmt.Printf("初始资金:%v, 交易了%v次，其中%v次赔了%v%%，%v次赚了0到%v%% \n", initMoney, tradeTimes, tradeTimes-times, lossRatio, times, interestMax)
	fmt.Printf("最终资金%v, 是原始资金的%.2f%%", money, money/initMoney*100)
}
