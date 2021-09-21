package piscine

func Atoi(s string) int {
	// result is a switch for numbers that include + or - as first index
	result := 1
	num := 0
	for i, n := range s {
		conv := int(n) - 48
		if conv >= 0 && conv <= 9 {
			num = (num * 10) + conv
			// if first index is a negative sign
		} else if conv == -3 && i == 0 {
			result = -1
			// else if positive sign as first index
		} else if conv == -5 && i == 0 {
			result = 1
		} else {
			return 0
		}
	}
	// num will multiply result and keep adding
	num *= result
	return num
}
