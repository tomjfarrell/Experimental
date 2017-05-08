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
	for i, j := 1, 1; i < 5 && j < 5; j++ {
		total := 0
		total = m[i][j] + m[i-1][j-1] + m[i-1][j] + m[i-1][j+1] + m[i+1][j-1] + m[i+1][j] + m[i+1][j+1]
		fmt.Printf("hourglass total for m[%v][%v]= %v\n", i, j, total)
		if total > max {
			max = total
		}
		if j == 4 {
			i++
			j = 1
		}
	}
	fmt.Println(max)
}


