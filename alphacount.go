package piscine

func AlphaCount(s string) int {
	count := 0
	for i := 0; i <= len(s)-1; i++ {
		if ('a' <= s[i] && s[i] <= 'z') || ('A' <= s[i] && s[i] <= 'Z') {
			count++
		}
	}
	return count
}
