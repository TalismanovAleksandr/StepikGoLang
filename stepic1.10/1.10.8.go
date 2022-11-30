package main

import (
	"fmt"
	"strings"
)

func main() {
	var first, second string
	fmt.Scan(&first, &second)

	for i := 0; i < len(first); i++ {
		if strings.Contains(second, string(first[i])) {
			fmt.Print(string(first[i]), " ")
		}
	}
}