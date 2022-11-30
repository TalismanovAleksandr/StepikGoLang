package stepic2_1

import "fmt"

func minimumFromFour() int {
	var min int
	fmt.Scan(&min)
	for i := 1; i < 4; i++ {
		var value int
		fmt.Scan(&value)
		if (value < min) {
			min = value
		}
	}
	return min
}