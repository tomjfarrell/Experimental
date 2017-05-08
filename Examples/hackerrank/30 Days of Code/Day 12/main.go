package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var s1 string
	var s2 string

	for i := 0; i < n; i++ {
		z := "NO"
		fmt.Scan(&s1)
		fmt.Scan(&s2)
    if s1 > s2 {
	    x := strings.Split(s2, "")
	    for i := range x {
		    if strings.Contains(s1, x[i]) {
			    z = "YES"
		    }
	    }
    }
		if s2 > s1 {
			x := strings.Split(s1, "")
			for i := range x {
				if strings.Contains(s2, x[i]) {
					z = "YES"
				}
		}
		}
		fmt.Println(z)
	}
}