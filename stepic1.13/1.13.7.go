package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var value int
	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&value)
		if (value == 0) {
			count++
		}
	}
	fmt.Println(count)
}