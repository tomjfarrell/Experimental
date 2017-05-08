package main

import (
	"fmt"
	"strconv"
	"strings"
)

func binary(x int64) string {
	b := strconv.FormatInt(x, 2)
	return b
}

func lencheck(y string) {
	z := strings.Split(y, "0")
	max := 0
	for i := 0; i < len(z); i++ {
		if len(z[i]) > max {
			max = len(z[i])
		}
	}
	fmt.Println(max)
}

func main() {
	var x int64
	fmt.Scan(&x)

	lencheck(binary(x))
}
