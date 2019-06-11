package main

import (
	"fmt"
	"math"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// testBasic()

	value := 1
	result := findClosest([]int{100, 99, 85, 33, 32, 16, 10, 5, 3, 2}, value)
	fmt.Printf("the array is %v and sum is %v closed to %v \n", result, sum(result), value)
}

func printFalse(a []int, value int, b []int) {
	result := findClosest(a, value)
	if !equal(result, b) {
		fmt.Printf("the array is %v and sum is %v closed to %v \n", result, sum(result), value)
	}
}

func findClosest(list []int, value int) []int {
	diff := math.MaxInt32
	var count int
	var finalArray, tmpArray []int
	for first := range list {
		count++
		slice := list[first:len(list)]

		sum := sum(slice)
		if sum <= value {
			if value-sum < diff {
				finalArray = slice
				break
			}
		}

		remaining := value
	loop:
		for index, element := range slice {
			count++
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
					switch min(diff, tmpDiff, remaining) {
					case diff:
						tmpArray = nil
					case tmpDiff:
						tmpArray = append(tmpArray, element)
						diff = tmpDiff
					case remaining:
						diff = remaining
						if tmpArray == nil {
							tmpArray = append(tmpArray, element)
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

	fmt.Println(count)
	return finalArray
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

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

func min(x, y, z int) int {
	min := check(x < y, x, y).(int)
	return check(min < z, min, z).(int)
}

func check(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func testBasic() {
	list := [6]int{8, 7, 6, 5, 5, 5}

	slice := list[0:]
	printFalse(slice, 1, []int{5})
	printFalse(slice, 2, []int{5})
	printFalse(slice, 3, []int{5})
	printFalse(slice, 4, []int{5})
	printFalse(slice, 5, []int{5})
	printFalse(slice, 6, []int{6})
	printFalse(slice, 7, []int{7})
	printFalse(slice, 8, []int{8})
	printFalse(slice, 9, []int{8})
	printFalse(slice, 10, []int{5, 5})
	printFalse(slice, 11, []int{6, 5})
	printFalse(slice, 12, []int{7, 5})
	printFalse(slice, 13, []int{8, 5})
	printFalse(slice, 14, []int{8, 6})
	printFalse(slice, 15, []int{8, 7})
	printFalse(slice, 16, []int{6, 5, 5})
	printFalse(slice, 17, []int{7, 6, 5})
	printFalse(slice, 18, []int{7, 6, 5})
	printFalse(slice, 20, []int{8, 7, 5})
	printFalse(slice, 21, []int{8, 7, 6})
	printFalse(slice, 22, []int{6, 5, 5, 5})
	printFalse(slice, 23, []int{7, 6, 5, 5})
	printFalse(slice, 24, []int{7, 6, 5, 5})
	printFalse(slice, 25, []int{8, 7, 6, 5})
	printFalse(slice, 26, []int{8, 7, 6, 5})
	printFalse(slice, 27, []int{8, 7, 6, 5})
	printFalse(slice, 28, []int{7, 6, 5, 5, 5})
	printFalse(slice, 29, []int{7, 6, 5, 5, 5})
	printFalse(slice, 30, []int{8, 7, 6, 5, 5})
	printFalse(slice, 31, []int{8, 7, 6, 5, 5})
	printFalse(slice, 32, []int{8, 7, 6, 5, 5})
	printFalse(slice, 33, []int{8, 7, 6, 5, 5})
	printFalse(slice, 34, []int{8, 7, 6, 5, 5, 5})
	printFalse(slice, 35, []int{8, 7, 6, 5, 5, 5})
	printFalse(slice, 36, []int{8, 7, 6, 5, 5, 5})
	printFalse(slice, 37, []int{8, 7, 6, 5, 5, 5})
	printFalse(slice, 100, []int{8, 7, 6, 5, 5, 5})
}
