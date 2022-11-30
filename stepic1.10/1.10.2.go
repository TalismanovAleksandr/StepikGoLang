package main

import "fmt"

func main() {
	var start int
	var end int
	fmt.Scan(&start)
	fmt.Scan(&end)
	var sum int
	for i := start; i <= end; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
