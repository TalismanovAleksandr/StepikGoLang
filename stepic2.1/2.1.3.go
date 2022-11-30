package stepic2_1

func vote(x int, y int, z int) int {
	oneCount := 0
	if x == 1 {
		oneCount++
	}

	if y == 1 {
		oneCount++
	}

	if z == 1 {
		oneCount++
	}

	if oneCount >= 2 {
		return 1
	}
	return 0
}