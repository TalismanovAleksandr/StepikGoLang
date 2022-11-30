package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scan(&input)
	max := 0
	for _, value := range input {
		intValue, _ := strconv.Atoi(string(value))
		if intValue > max {
			max = intValue
		}
	}
	fmt.Println(max)
}