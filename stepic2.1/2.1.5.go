package stepic2_1

func sumInt(elements ...int) (int, int) {
	count := 0
	sum := 0
	for _,v := range elements {

		sum += v
		count++
	}
	return count, sum
}