package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	m := make(map[string]string)
	var x string
	var entry []string

	for i := 0; i < n; i++ {
		fmt.Scanln(&x)
		entry = strings.Split(x, " ")
		fmt.Println(entry)
		m[entry[0]] = entry[1]
		//m["route"] = 66
		fmt.Println(entry)
	}
}
