package piscine

func NRune(s string, n int) rune {
	x := []rune(s)

	if n < 0 || len(s) < n {
		return 0
	}
	return x[n-1]
}
