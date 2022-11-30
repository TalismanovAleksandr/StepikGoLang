package main

import "fmt"

func main() {
	var value int
	max := 0
	maxCount := 0
	for fmt.Scan(&value); value != 0; fmt.Scan(&value) {
		if value > max {
			maxCount = 0
			max = value
			maxCount++
		} else if value == max {
			maxCount++
		}
	}
	fmt.Println(maxCount)
}
