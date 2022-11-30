package main

import "fmt"

func main() {
	min := 999999
	minCount := 0
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		if (value < min) {
			min = value
			minCount = 1
		} else if (value == min) {
			minCount++
		}
	}
	fmt.Println(minCount)
}