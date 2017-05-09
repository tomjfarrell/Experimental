package main

import (
	"fmt"
)

func main() {
	var m = make(map[byte]bool)
	var n int
	fmt.Scan(&n)
	var s1 string
	var s2 string


	for i := 0; i < n; i++ {
		fmt.Scan(&s1)
		fmt.Scan(&s2)
    if len(s1) > len(s2) {
	    for _,c := range s2 {
		    m[c] = true
		    for _,c := range s1 {
			    if m[c] {
				    fmt.Println("YES")
				    break
			    }
		    }
	    }
    else {
	    for _,c := range s1 {
		    m[c] = true
	    }

		}
		}
		fmt.Println(z)
	}
}