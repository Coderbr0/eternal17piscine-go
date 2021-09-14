package piscine

func Compare(a, b string) int {
	x := []byte(a)
	y := []byte(b)

	for i := range a {
		for j := range b {
			if x[i] < y[j] {
				return -1
			} else if x[i] > y[j] {
				return 1
			} else if x[i] == y[j] && len(x) != len(y) {
				x[i] = x[i+1]
				y[j] = y[j+1]
			}
		}
	}
	return 0
}
