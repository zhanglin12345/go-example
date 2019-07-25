package main

import "fmt"

func main() {

	arrays := [6]int{7, 9, 9, 7, 9, 2}
	ones := 0
	twos := 0

	for _, value := range arrays {
		ones = (ones ^ value) &^ twos
		twos = (twos ^ value) &^ ones
	}

	fmt.Println(ones)
	// fmt.Println(twos)

}
