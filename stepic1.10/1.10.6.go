package main

import "fmt"

func main() {
	var input int
	for fmt.Scan(&input) ;  ; fmt.Scan(&input) {
		if input < 10 {
			continue
		}
		if input > 100 {
			break
		}
		fmt.Println(input)
	}
}