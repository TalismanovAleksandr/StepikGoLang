package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	result := ""
	for pos, char := range input {
		//fmt.Printf("character %c starts at byte position %d\n", char, pos)
		if pos == len(input) - 1 {
			result += string(char)
		} else {
			result += string(char) + "*"
		}
	}
	fmt.Println(result)
}

