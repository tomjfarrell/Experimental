package main

import (
	"fmt"
	"strings"
)

func main() {
	var words []string
	var n int
	fmt.Scan(&n)
	var x string


	for i := 1; i <= n; i++ {
		fmt.Scanf("%s",&x)
		words = append(words, x)
	}

	for i := 0; i < len(words); i++ {
		word := strings.Split(words[i], "")
		var evens []string
		var odds []string
		for i := 0; i < len(word); i++ {
			if i % 2 == 0 {
				evens = append(evens, word[i])
			} else {
				odds = append(odds, word[i])
			}
		}
		fmt.Printf("%v %v\n", strings.Join(evens,""), strings.Join(odds,""))
	}
}