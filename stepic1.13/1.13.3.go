package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)
	hours := input / 3600
	reminderHours := input % 3600
	minutes := reminderHours / 60

	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
