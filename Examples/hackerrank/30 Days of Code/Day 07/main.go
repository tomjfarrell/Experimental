package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var output []int
	var x int

	for i := n; i > 0; i-- {
		fmt.Scan(&x)
		output = append([]int{x}, output...)
	}
	for i := 0; i < len(output); i++ {
		fmt.Printf("%v ", output[i])
	}
}