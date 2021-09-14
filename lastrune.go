package piscine

func LastRune(s string) rune {
	x := []rune(s)
	return x[len(x)-1]
}
