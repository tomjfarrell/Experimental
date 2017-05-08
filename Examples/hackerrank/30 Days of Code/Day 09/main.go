package main

import (
	"fmt"
)

func factorial(n int) {
	answer := 1
	if n == 0 {
		fmt.Println(0)
	} else {
		for i := 1; i <= n; i++ {
			answer = answer * i
		}
		fmt.Println(answer)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	factorial(n)
}