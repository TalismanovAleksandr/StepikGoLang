package stepic3_1

import "fmt"

//mock
func work(a int) int {
	return a
}

func main() {
	var input int
	m := map[int]int{}
	for i := 0; i < 10; i++ {
		fmt.Scan(&input)
		if value, ok := m[input]; ok {
			fmt.Printf("%d ", value)
		} else {
			result := work(input)
			m[input] = result
			fmt.Printf("%d ", result)
		}
	}
}
