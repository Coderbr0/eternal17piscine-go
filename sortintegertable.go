package piscine

func SortIntegerTable(table []int) {
	i := 0
	for i < len(table)-1 {
		if table[i] > table[i+1] {
			table[i], table[i+1] = table[i+1], table[i]
			// resetting i until sorting is done
			i = 0
		} else {
			i++
		}
	}
}
