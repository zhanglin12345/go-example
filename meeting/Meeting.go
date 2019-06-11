package main

import (
	"fmt"
	"math"
)

func sum(value []int) int {
	var sum int
	for _, element := range value {
		sum += element
	}
	return sum
}

func isLastElement(index int, slice []int) bool {
	return index == len(slice)-1
}

func main() {
	list := [6]int{8, 7, 6, 5, 5, 5}
	value := 16
	diff := int(^uint(0) >> 1)

	var finalArray, tmpArray []int
	for first := range list {
		slice := list[first:len(list)]

		sum := sum(slice)
		if sum <= value {
			if value-sum < diff {
				finalArray = slice
			}
		}

		remaining := value
	loop:
		for index, element := range slice {
			switch {
			case remaining-element > 0:
				remaining -= element
				tmpArray = append(tmpArray, element)

				if isLastElement(index, slice) {
					if remaining > diff {
						tmpArray = nil
					}
				}
			case remaining-element < 0:
				if isLastElement(index, slice) {
					tmpDiff := abs(remaining - element)
					if remaining > tmpDiff {
						if tmpDiff >= diff {
							tmpArray = nil
						} else {
							diff = tmpDiff
							tmpArray = append(tmpArray, element)
						}
					} else {
						if remaining > diff {
							tmpArray = nil
						} else {
							diff = remaining
						}
					}
				}
			default:
				diff = 0
				tmpArray = append(tmpArray, element)
				break loop
			}
		}

		if diff == 0 {
			finalArray = tmpArray
			break
		} else if tmpArray != nil {
			finalArray = tmpArray
			tmpArray = nil
		}
	}

	fmt.Printf("the array is %v and sum is %v", finalArray, sum(finalArray))
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}
