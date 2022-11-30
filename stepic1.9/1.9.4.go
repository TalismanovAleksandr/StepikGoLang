package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	if input % 400 == 0 {
		fmt.Println("YES")
	} else if input % 4 == 0 && input % 100 != 0{
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
