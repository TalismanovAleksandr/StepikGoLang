package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scan(&input)
	result := 0
	for i := 0; i < len(input); i++ {
		value, _ := strconv.Atoi(string(input[i]))
		result += value
	}
	fmt.Println(result)
}