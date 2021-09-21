package piscine

func IsSorted(f func(a, b int) int, a []int) bool {

	sortup := true
	sortdown := true
	for i := 0; i < len(a)-1; i++ {
		if !(f(a[i], a[i+1]) >= 0) {
			sortdown = false
		}
		if !(f(a[i], a[i+1]) <= 0) {
			sortup = false
		}
	}
	return sortup || sortdown
}
