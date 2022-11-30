package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(value uint) uint {
		strValue := strconv.Itoa(int(value))
		strResult := ""
		for i := 0; i < len(strValue); i++ {
			char := string(strValue[i])
			atoi, _ := strconv.Atoi(char)
			if atoi % 2 == 0 && atoi != 0{
				strResult += char
			}
		}
		atoi2, _ := strconv.Atoi(strResult)
		if atoi2 == 0 {
			return 100
		}
		return uint(atoi2)
	}
	fmt.Println(fn(727178))
}