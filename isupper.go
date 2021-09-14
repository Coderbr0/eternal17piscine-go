package piscine

func IsUpper(s string) bool {
	for i := range s {
		if 'A' <= s[i] && s[i] <= 'Z' && i <= len(s) {
			i++
			if i == len(s) {
				return true
			}
		}
	}
	return false
}
