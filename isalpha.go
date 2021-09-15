package piscine

func IsAlpha(s string) bool {
	a := []byte(s)
	for _, number := range a {
		if number == 32 {
			return false
		}
	}
	return true
}
