package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Scan(&input)

	if input > 10000 {
		fmt.Printf("%e", input)
	} else if input <= 0 {
		fmt.Printf("число %.2f не подходит", input)
	} else {
		result := math.Pow(input, 2.0)
		fmt.Printf("%.4f", result)
	}
}
