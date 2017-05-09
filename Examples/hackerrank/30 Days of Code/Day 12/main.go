package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var s1 string
	var s2 string


	for i := 0; i < n; i++ {
		var m = make(map[byte]bool)
		var z = "NO"
		fmt.Scan(&s1)
		fmt.Scan(&s2)
		if len(s1) > len(s2) {
			fmt.Println("First word is larger, mapping second.")
			for _, c := range s2 {
				m[byte(c)] = true
				fmt.Printf("Set m[%v] = true\n",c)
				for _, c := range s1 {
					if m[byte(c)] {
						fmt.Printf("m[%v] found in map\n", c)
						z = "YES"
						break
					} else {
						fmt.Printf("m[%v] not found in map (%v).\n", c, m)
					}
				}
			}
		} else {
			fmt.Println("Second word is larger, mapping first.")
	    for _,c := range s1 {
		    m[byte(c)] = true
		    fmt.Printf("Set m[%v] = true\n",c)
		    for _, c := range s2 {
			    if m[byte(c)] {
				    fmt.Printf("m[%v] found in map\n", c)
				    z = "YES"
				    break
			    } else {
				    fmt.Printf("m[%v] not found in map (%v).\n", c, m)
			    }
	    }
		}
		}
		fmt.Println(z)
	}
}