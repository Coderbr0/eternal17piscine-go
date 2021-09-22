package piscine

func TrimAtoi(s string) int {
	ar := []rune(s)
	sign := 1
	result := 0
	foundDigit := false

	for _, r := range ar {
		if !foundDigit {
			if r == '-' {
				sign = -1
			} else if r == '+' {
				sign = 1
			}
		}
		if r >= '0' && r <= '9' {
			foundDigit = true
			result = 10*result + int(r-48)
		}
	}
	return sign * result
}
