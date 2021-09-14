package piscine

// func Compare(a, b string) int {
// 	x := []byte(a)
// 	y := []byte(b)
// 	for i := range x {
// 		for j := range y {
// 			if x[i] > y[j] || x[i] == y[j] && len(a) > len(b) {
// 				return 1
// 			} else if x[i] < y[j] || x[i] == y[j] && len(b) > len(a) {
// 				return -1
// 			} else {
// 				return 0
// 			}
// 		}
// 	}
// 	return 100
// }

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}
