package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)


	m := make(map[string]int)
	var name string
	var phone int

	for i := 0; i < n; i++ {
		fmt.Scan(&name, &phone)
		m[name] = phone
	}

	for {
		read, _ := fmt.Scan(&name)
		if read > 0 {
			phone, ok := m[name]
			if !ok {
				fmt.Println("Not found")
			} else {
				fmt.Printf("%v=%v\n", name, phone)
			}
		} else {
			os.Exit(1)
		}
	}
}