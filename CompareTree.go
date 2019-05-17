package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

func check(t *tree.Tree, ch chan int) bool {
	if t.Left != nil {
		if !check(t.Left, ch) {
			return false
		}
	}

	x, ok := <-ch
	if ok {
		fmt.Printf("x: %d  y:%d \n", x, t.Value)
		if x != t.Value {
			return false
		}
	} else {
		fmt.Printf("x: nil  y:%d \n", t.Value)
		return false
	}
	if t.Right != nil {
		if !check(t.Right, ch) {
			return false
		}
	}

	return true
}

func doCheck(t *tree.Tree, ch chan int) bool {
	ok := check(t, ch)
	if !ok {
		return false
	}
	x, ok := <-ch
	fmt.Printf("x: %v   y is nil \n", x)
	return !ok
}

func doWork(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// func main() {
// 	ch1 := make(chan int)
// 	t1 := tree.New(1)
// 	t2 := tree.New(2)

// 	t2.Right = &tree.Tree{nil, 999, nil}

// 	go doWork(t1, ch1)

// 	fmt.Println(doCheck(t2, ch1))
// }
