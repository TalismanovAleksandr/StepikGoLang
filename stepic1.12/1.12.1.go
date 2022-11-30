package main

import "fmt"
//для выполнения задачи из комментария достаем и вставляем в окошло следующее
/*
	var workArray [10]uint8
	var input uint8
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&input)
		workArray[i] = input
	}

	first := -1
	second := -1
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			if first != -1 && second != -1 {
				tmp := workArray[first]
				workArray[first] = workArray[second]
				workArray[second] = tmp

				first = -1
				second = -1
			}
			fmt.Scan(&first)
		} else if i%2 == 1 {
			fmt.Scan(&second)
		}
	}
	tmp := workArray[first]
	workArray[first] = workArray[second]
	workArray[second] = tmp

	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
*/
func main() {
	var workArray [10]uint8
	var input uint8
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&input)
		workArray[i] = input
	}

	first := -1
	second := -1
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			if first != -1 && second != -1 {
				tmp := workArray[first]
				workArray[first] = workArray[second]
				workArray[second] = tmp

				first = -1
				second = -1
			}
			fmt.Scan(&first)
		} else if i%2 == 1 {
			fmt.Scan(&second)
		}
	}
	tmp := workArray[first]
	workArray[first] = workArray[second]
	workArray[second] = tmp

	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], " ")
	}
	//99 137 151 187 117 71 20 187 93 67
}