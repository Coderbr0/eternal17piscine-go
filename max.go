package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}

	i := 0
	for i < len(a)-1 {
		if a[i] > a[i+1] {
			a[i], a[i+1] = a[i+1], a[i]
			i = 0
		} else {
			i++
		}
	}
	return a[i]
}
