package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	result := 0
	for ;input != 0; {
		result += input % 10
		input /= 10
	}

	//fmt.Println(result)

	mainResult := 0
	for ;result != 0; {
		mainResult += result % 10
		result /= 10
	}
	fmt.Println(mainResult)
}