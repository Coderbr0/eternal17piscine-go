package piscine

func Join(s []string, sep string) string {
	var needed string

	for i := range s {
		needed += s[i]
		if i != len(s)-1 {
			needed += sep
		}
	}
	return needed
}
