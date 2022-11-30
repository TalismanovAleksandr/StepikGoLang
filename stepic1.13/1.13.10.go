package main

import "fmt"

func main() {
	var input1 int
	var input2 int
	fmt.Scan(&input1)
	fmt.Scan(&input2)

	maxDividedBySeven := -1
	for i := input1; i <= input2; i++ {
		if i % 7 == 0 {
			maxDividedBySeven = i
		}
	}
	if maxDividedBySeven != -1 {
		fmt.Println(maxDividedBySeven)
	} else {
		fmt.Println("NO")
	}
}