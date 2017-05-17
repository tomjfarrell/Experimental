package main

import (
	"fmt"
)

func check(mag map[string]int, ran map[string]int) {
	ans := "No"

	for key, value := range ran {
		if mag[key] >= value {
			ans = "Yes"
		}
	}

	fmt.Println(ans)
}

func main() {
	var m,n int //m=#wordsmagazine,n=#wordsransom
	var x string //temporary word scan storage
	fmt.Scan(&m, &n)
	mag := make(map[string]int) //map for magazine words mag[word]=count
	ran := make(map[string]int) //map for ransom words ran[word]=count

	for i := 0; i < m; i++ {
		fmt.Scan(&x)
		_, ok := mag[x]
		if ok {
			mag[x]++
		} else {
			mag[x] = 1
		}
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		_, ok := ran[x]
		if ok {
			ran[x]++
		} else {
			ran[x] = 1
		}
	}

	check(mag, ran)
}

//https://www.hackerrank.com/challenges/ctci-ransom-note