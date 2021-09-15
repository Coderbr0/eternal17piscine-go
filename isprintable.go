package piscine

func IsPrintable(s string) bool {
	a := []byte(s)

	for i, char := range a {
		if 32 <= char && char <= 127 {
			i++
		} else {
			return false
		}
	}
	return true
}
