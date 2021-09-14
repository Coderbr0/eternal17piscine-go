package piscine

func Compare(a, b string) int {
	x := []byte(a)
	y := []byte(b)
	for indexa := range x {
		for index := range y {
			if x[indexa] > y[index] || x[indexa] == y[index] && len(a) > len(b) {
				return 1
			} else if x[indexa] < y[index] || x[indexa] == y[index] && len(b) > len(a) {
				return -1
			} else {
				return 0
			}
		}
	}
	return 100
}
