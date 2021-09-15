package piscine

func IsNumeric(s string) bool {
	for i, n := range s {
		if '0' <= n && n <= '9' {
			i++
		} else {
			return false
		}
	}
	return true
}
