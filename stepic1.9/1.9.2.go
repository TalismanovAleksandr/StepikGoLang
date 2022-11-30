package main

import "fmt"

func main() {
	var value string
	fmt.Scan(&value)
	if value[0] != value[1] && value[0] != value[2] && value[1] != value[2] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}