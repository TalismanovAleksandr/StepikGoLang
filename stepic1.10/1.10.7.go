package main

import "fmt"

func main() {
	var x int
	var p int
	var y int
	result := 0
	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)
	for ; x <= y ; {
		x = x + x * p / 100
		result++
	}
	fmt.Println(result)
}
