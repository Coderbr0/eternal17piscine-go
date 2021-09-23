package piscine

func BasicJoin(s []string) string {
	var needed string

	for i := range s {
		needed += s[i]
	}
	return needed
}
