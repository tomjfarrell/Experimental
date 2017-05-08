package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)

	var data int64

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		data, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	if data % 2 == 1 {
		fmt.Println("Weird")
	} else if data % 2 == 0 {
		if data >= 2 && data <= 5 {
			fmt.Println("Not Weird")
		} else if data >= 6 && data <= 20 {
			fmt.Println("Weird")
		} else if data > 20 {
			fmt.Println("Not Weird")
		}
	}
}