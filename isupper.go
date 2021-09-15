package piscine

func IsUpper(s string) bool {
	for i, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			i++
		} else {
			return false
		}
	}
	return true
}
