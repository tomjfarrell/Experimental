package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Input directory size(int): ")
	fmt.Scan(&n)


	m := make(map[string]int)
	var name string
	var phone int

	for i := 0; i < n; i++ {
		fmt.Printf("N = %v | i = %v\n", n, i)
		fmt.Print("Input record(name phone_number): ")
		fmt.Scan(&name, &phone)
		fmt.Printf("Name received: %v | Phone received: %v\n", name, phone)
		m[name] = phone
		fmt.Print(m)
	}

  for fmt.Scan(&name)

}
