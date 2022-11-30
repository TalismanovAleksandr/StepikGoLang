package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scan(&input)
	firstSum := 0
	secondSum := 0
	for index, value := range input {
		strVal := string(value)
		intValue, _ := strconv.Atoi(strVal)
		if index == 0 || index == 1 || index == 2 {
			firstSum += intValue
		} else {
			secondSum += intValue
		}
	}
	if firstSum == secondSum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
