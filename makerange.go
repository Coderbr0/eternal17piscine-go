package piscine

func MakeRange(min, max int) []int {
	// creating a empty slice for condition min > max
	var arr []int
	if min >= max {
		return arr
	}

	// declaring a make and incrementally adding to the make
	usingMake := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		usingMake[i] = i + min
	}
	return usingMake
}
