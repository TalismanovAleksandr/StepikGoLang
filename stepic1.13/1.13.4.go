package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	isGood := false
	if a > b && a > c {
		isGood = math.Pow(a,2) == math.Pow(b,2) +  math.Pow(c,2)
	} else if b > a && b > c {
		isGood = math.Pow(b,2) == math.Pow(a,2) +  math.Pow(c,2)
	} else if c > b && c > a {
		isGood = math.Pow(c,2) == math.Pow(b,2) +  math.Pow(a,2)
	}
	if isGood {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный");
	}
}