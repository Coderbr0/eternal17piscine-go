package piscine

func IsAlpha(s string) bool {
	for i, l := range s {
		if ('a' <= l && l <= 'z') || ('A' <= l && l <= 'Z') || ('0' <= l && l <= '9') {
			i++
		} else {
			return false
		}
	}
	return true
}
