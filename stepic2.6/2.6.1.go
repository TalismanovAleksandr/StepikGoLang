package stepic2_6

import "fmt"

func divide(a int, b int) (int, error){
	//это заглушка
	return 1, nil
}


func main() {
	// package main уже объявлен, нужные импорты уже есть
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	value, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(value)
	}
}
