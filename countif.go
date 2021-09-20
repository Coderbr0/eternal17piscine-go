package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0

	// apply the function to each string in tab
	// if function is true add to counter

	for i := range tab {
		if f(tab[i]) {
			count++
		}
	}
	return count
}
