package main

import (
	"fmt"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)
	x := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	for i := 0; i < d; i++ {
		x = append(x[1:n], x[0])
	}
	for i, v := range x {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}