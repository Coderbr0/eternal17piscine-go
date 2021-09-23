package piscine

func ActiveBits(n int) int {
	result := 0

	for ; n > 0; n = n / 2 {
		if n%2 == 1 {
			result++
		}
	}
	return result
}
