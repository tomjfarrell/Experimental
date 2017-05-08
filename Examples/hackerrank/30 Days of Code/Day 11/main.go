package main

import (
	"fmt"
	//"strings"
)

func main() {
	var x [6]int
	m := make(map[int][6]int)
	var max int

	for i := 0; i < 6; i++ {
		fmt.Printf("i = %v\n", i)
		fmt.Scan(&x[0],&x[1],&x[2],&x[3],&x[4],&x[5])
		fmt.Printf("x = %v\n", x)
		for range x {
			m[i] = x
		}
	}
	for i := 0; i < 6; i ++ {
		fmt.Println(m[i])
	}
	for i := 1 ; i < 5; i++ {
		total := 0
		for j := 1 ; j < 5; j++ {
			total = m[i][j] + m[i-1][j-1] + m[i-1][j] + m[i-1][j+1] + m[i+1][j-1] + m[i+1][j] + m[i+1][j+1]
			fmt.Printf("hourglass total for m[%v][%v]= %v\n", i, j, total)
			if i == 1 && j == 1 {
				fmt.Println("first calculation set to max")
				max = total
			}
			if total > max {
				max = total
			}
		}
	}
	fmt.Println(max)
}
