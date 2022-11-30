package main

import "fmt"

func main() {
	var value int
	fmt.Scan(&value)
	if value < 0 {
		fmt.Println("Число отрицательное")
	} else if value > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Ноль")
	}
}