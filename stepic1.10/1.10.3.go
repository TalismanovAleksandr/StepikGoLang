package main

import "fmt"

func main() {
	var countOfDigits int
	var value int
	fmt.Scan(&countOfDigits)

	var sum int
	for i := 0; i < countOfDigits; i++ {
		fmt.Scan(&value)
		if isTwoDigits(value) && value % 8 == 0 {
			sum = sum + value
		}
	}
	fmt.Println(sum)
}

func isTwoDigits(value int) bool {
	return value >= 10 && value < 100
}
