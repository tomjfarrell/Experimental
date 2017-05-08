package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var data []string

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	mealCost, _ := strconv.ParseFloat(data[0], 64)
	tipPercent, _ := strconv.ParseFloat(data[1], 64)
	taxPercent, _ := strconv.ParseFloat(data[2], 64)

	tip := mealCost*(tipPercent/100)
	tax := mealCost*(taxPercent/100)

	totalCost := mealCost+tip+tax

	fmt.Printf("The total meal cost is %.f dollars.", totalCost)
}