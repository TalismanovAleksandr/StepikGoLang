package main

import "fmt"

func reverseString(str string) (result string) {
	for _ , v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	str := "Educative-Edpresso"

	fmt.Println(str)
	// invoke reverseString
	fmt.Println(reverseString(str))
}
