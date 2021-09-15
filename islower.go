package piscine

func IsLower(s string) bool {
	for i, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			i++
		} else {
			return false
		}
	}
	return true
}
