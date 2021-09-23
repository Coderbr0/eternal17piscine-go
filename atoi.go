package piscine

func Atoi(s string) int {
	result := 1
	num := 0

	for i, n := range s {
		conv := int(n) - 48
		if 0 <= conv && conv <= 9 {
			num = num*10 + conv
		} else if conv == -3 && i == 0 {
			result = -1
		} else if conv == -5 && i == 0 {
			result = 1
		} else {
			return 0
		}
	}
	num *= result
	return num
}
