package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/fatih/color"
)

var initValue float64 = 10000
var times int = 0

func do() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	win := random.Intn(1000)
	if win < 300 {
		initValue *= 0.95
	} else {
		result := random.Intn(100)
		interest := float64(result) / 1000
		initValue *= 1 + interest
		times++
	}
}

func main() {
	for i := 0; i <= 100; i++ {
		tmp := initValue
		do()
		if tmp > initValue {
			color.Green(strconv.FormatFloat(initValue, 'f', 6, 64))
		} else {
			color.Red(strconv.FormatFloat(initValue, 'f', 6, 64))
		}
	}
	fmt.Println(times)
}
